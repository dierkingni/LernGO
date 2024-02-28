resource "google_service_account" "service_account" {
  account_id   = "githubactions-cloudfunctions-556655823"
  display_name = "githubactions-cloudfunctions"
}

 
resource "google_project_iam_binding" "service-account" {
  project = "mms-clp-playground2402-a-i2ar"
  role    = "roles/storage.objectUser"
  members = [
    "serviceAccount:mms-clp-playground2402-a-i2ar@appspot.gserviceaccount.com",
  ]
}

resource "google_project_iam_binding" "github-service-account" {
  project = "mms-clp-playground2402-a-i2ar"
  for_each = toset([
    "roles/storage.admin",
    "roles/resourcemanager.projectIamAdmin",
    "roles/cloudfunctions.admin",
  ])
  role = each.key
  members = [
    "serviceAccount:githubactions@mms-clp-playground2402-a-i2ar.iam.gserviceaccount.com",
  ]
}