provider "google" {
  project = "mms-clp-playground2402-a-x11e"
  region  = "europe-west4"
}

resource "google_storage_bucket" "my-bucket" {
  name                        = "calculator-bucket"
  location                    = "europe-west4"
  force_destroy               = true
  uniform_bucket_level_access = true


}