#!/bin/bash
export GOOGLE_PROJECT=archiebw-todo
export TF_VAR_admin_project=${GOOGLE_PROJECT}
export GOOGLE_CREDENTIALS=~/.config/gcloud/${GOOGLE_PROJECT}-service-account.json
export GOOGLE_APPLICATION_CREDENTIALS=~/.config/gcloud/${GOOGLE_PROJECT}-service-account.json
export TF_STATE_BUCKET=gs://${GOOGLE_PROJECT}-tf-state
