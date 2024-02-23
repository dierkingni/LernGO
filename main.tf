locals {
    project_id = "mms-clp-playground2402-a-i2ar"
}

terraform {
  required_version = ">= 1.0"
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 4.69.1"
    }
  }
}

resource "google_storage_bucket" "my-bucket"{
    name          = "calculator-bucket"
    location      = "europe-west4"
    force_destroy = true
    uniform_bucket_level_access = true

    
}