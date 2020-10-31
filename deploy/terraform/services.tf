resource "google_project_service" "cloudkms" {
  project = data.google_client_config.default.project
  service = "cloudkms.googleapis.com"
}

resource "google_project_service" "appengine" {
  project = data.google_client_config.default.project
  service = "appengine.googleapis.com"
}