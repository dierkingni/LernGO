provider "google" {
  project = "mms-clp-playground2402-a-i2ar"
  region  = "europe-west4"
}
terraform {
  required_version = ">= 1.0"
  required_providers {
    google = {
      source  = "hashicorp/google"
    }
  }
}

resource "google_storage_bucket" "my-bucket"{
    name          = "calculator-bucket"
    location      = "europe-west4"
    force_destroy = true
    uniform_bucket_level_access = true

    
}