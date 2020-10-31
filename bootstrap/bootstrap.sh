#!/bin/bash

# Script Root for this script, used to source environment.
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"
# Import the environment variables
source $DIR/environment.sh
set -euo pipefail

# Configuration Variables, tweak these based upon changes to GCP.
PROJECT_ROLES='roles/storage.admin'
REQUIRED_SERVICES='cloudresourcemanager.googleapis.com'

DIRTY=0

if [[ "$GOOGLE_PROJECT" == "" ]]
  then
    echo -e "Admin Project: The Admin Project in the \$GOOGLE_PROJECT variables doesn't appear to be set."
    DIRTY=1
  else
    echo -e "Admin Project: The Admin Project ($GOOGLE_PROJECT) is set."
fi

EXISTING_PROJECT=$(gcloud projects list --filter=name:${GOOGLE_PROJECT} --format=json)
if [[ "$EXISTING_PROJECT" == "[]" ]]
  then
    echo -e "Project: The Admin Project in the \$GOOGLE_PROJECT variables doesn't exist."
    DIRTY=1
  else
    echo -e "Admin Project: The Project ($GOOGLE_PROJECT) exists."
fi

if [[ "$DIRTY" == "1" ]]
  then
    echo -e "Error: One of more guard clauses failed check environment.sh and gcloud configuration matches up"
    exit 1
fi

# Enable the required services
for svc in $REQUIRED_SERVICES
do
  echo -e "Admin Project: Enabling $svc."
  gcloud services enable "$svc"
done

# Create the user if it does not exist

EXISTING_USER=$(gcloud iam service-accounts list --format=json --filter=name:terraform)

if [[ "$EXISTING_USER" == "[]" ]]
  then
    echo -e "Terraform User: The Terraform does not exist yet, creating it and granting permissions."
    gcloud iam service-accounts create terraform --display-name "Terraform admin account"

    echo -e "Terraform User: Downloading JSON credentials."
    gcloud iam service-accounts keys create ${TF_CREDS} --iam-account terraform@${GOOGLE_PROJECT}.iam.gserviceaccount.com
  else
    echo -e "Terraform User: The Admin Project ($GOOGLE_PROJECT) already exists."
fi

# Assign roles for the user, don't need to check first as this works repeatedly
for prole in $PROJECT_ROLES
do
  echo -e "Terraform User: Granting $prole role to the Admin Project ($GOOGLE_PROJECT)."
  gcloud projects add-iam-policy-binding ${GOOGLE_PROJECT} --member serviceAccount:terraform@${GOOGLE_PROJECT}.iam.gserviceaccount.com --role $prole
done

# Create terraform backend
if (gsutil ls ${TF_STATE_BUCKET} &>/dev/null)
  then
    echo -e "Admin Project: The Terraform state bucket ($TF_STATE_BUCKET) for ($GOOGLE_PROJECT) already exists."
  else
    echo -e "Admin Project: Creating Terraform state bucket ($TF_STATE_BUCKET) for ($GOOGLE_PROJECT)."
    gsutil mb -p ${GOOGLE_PROJECT} -c multi_regional -b on ${TF_STATE_BUCKET}
    gsutil versioning set on ${TF_STATE_BUCKET}
fi

# All done
echo -e "Bootstrap: Bootstrapping Complete."
