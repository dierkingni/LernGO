resource "google_service_account" "service_account" {
  account_id                   = "githubactions-cloudfunctions"
  display_name                 = "githubactions-cloudfunctions"
  create_ignore_already_exists = "true"
}

resource "google_service_account" "cloudfunctions_service_account" {
  account_id                   = "cloudfunctions-account"
  display_name                 = "cloudfunctions-account"
  create_ignore_already_exists = "true"
}



resource "google_project_iam_binding" "service_account" {
  project = "mms-clp-playground2402-a-x11e"
  role    = "roles/storage.objectUser"
  members = [
    "serviceAccount:cloudfunctions-account@mms-clp-playground2402-a-x11e.iam.gserviceaccount.com",
  ]
}

resource "google_project_iam_binding" "github-service_account" {
  project = "mms-clp-playground2402-a-x11e"
  for_each = toset([
   "roles/cloudfunctions.developer",
   "roles/iam.serviceAccountUser",
   "roles/iam.workloadIdentityUser"
  ])
  role = each.key
  members = [
    "serviceAccount:githubactions-cloudfunctions@mms-clp-playground2402-a-x11e.iam.gserviceaccount.com",
  ]
}

