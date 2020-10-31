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

resource "google_app_engine_application" "firestore" {
  project       = data.google_client_config.default.project
  location_id   = var.firestore_location
  database_type = "CLOUD_FIRESTORE"

  depends_on = [google_project_service.appengine]
}
