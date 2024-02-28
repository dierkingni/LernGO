provider "google" {
  project = "mms-clp-playground2402-a-i2ar"
  region  = "europe-west4"
}

resource "google_storage_bucket" "my-bucket" {
  name                        = "calculator-bucket"
  location                    = "europe-west4"
  force_destroy               = true
  uniform_bucket_level_access = true


}

resource "google_iam_workload_identity_pool" "example" {
  workload_identity_pool_id = "example-pool"
  display_name              = "Name of pool"
  description               = "Identity pool for automated test"
  disabled                  = true
  provider = 
}