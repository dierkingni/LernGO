locals {
    project_id = "mms-clp-playground2402-a-i2ar"
}

provider "google" {
    project = local.project_id
    region  = "europe-west4"
}

resource "google_storage_bucket" "my-bucket"{
    name          = "calculator-bucket"
    location      = "europe-west4"
    force_destroy = true
}