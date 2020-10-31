data "google_client_config" "default" {}

resource "google_kms_key_ring" "default" {
  name     = "admin-keyring"
  project  = data.google_client_config.default.project
  location = var.region

  depends_on = [google_project_service.cloudkms]
}

resource "google_kms_crypto_key" "default" {
  name            = "default-key"
  key_ring        = google_kms_key_ring.default.self_link
  rotation_period = "1209600s"

}

// firestore DB has been initialised through portal, potentially automatable through local-exec script
// gcloud firestore databases create --region=europe-west2

# resource "google_firestore_index" "default" {
#   collection = "todos"

#   fields {
#     field_path = "id"
#     order      = "ASCENDING"
#   }

#   fields {
#     field_path = "description"
#     order      = "ASCENDING"
#   }

#   fields {
#     field_path = "__name__"
#     order      = "DESCENDING"
#   }
# }
