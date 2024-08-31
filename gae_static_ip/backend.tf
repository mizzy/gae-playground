terraform {
  backend "gcs" {
    bucket = "tf-state-mizzy"
    prefix = "gae-playground/gae-static-ip"
  }
}
