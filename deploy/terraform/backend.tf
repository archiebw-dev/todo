terraform {
  backend "gcs" {
    bucket  = "archiebw-todo-tf-state"
    prefix  = "dev"
  }
}
