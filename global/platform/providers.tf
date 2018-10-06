provider "google" {
  credentials = "${file("google-credentials.json")}"
  project     = "${var.gcloud-project}"
  region      = "${var.gcloud-region}"
}