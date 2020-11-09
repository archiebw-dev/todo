# todo-api

A simple backend API written in go using the echo framework. It can run using an in memory DB for testing purposes or using google cloud firestore db.

## Deploy

- Update `./bootstrap/environment.sh` with google project id.
- Run `./bootstrap/bootstrap.sh` to create the service principal and assign necessary roles.
- Run `cd ./deploy/terraform && terraform init && terraform apply`
