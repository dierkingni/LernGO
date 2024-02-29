resource "google_service_account" "service_account" {
  account_id                   = "githubactions-cloudfunctions"
  display_name                 = "githubactions-cloudfunctions"
  create_ignore_already_exists = true
}


resource "google_project_iam_binding" "service-account" {
  project = "mms-clp-playground2402-a-x11e"
  role    = "roles/storage.objectUser"
  members = [
    "serviceAccount:mms-clp-playground2402-a-x11e@appspot.gserviceaccount.com",
  ]
}

resource "google_project_iam_binding" "github-service-account" {
  project = "mms-clp-playground2402-a-x11e"
  for_each = toset([
    "roles/storage.admin",
    "roles/resourcemanager.projectIamAdmin",
    "roles/cloudfunctions.admin",
    "roles/iam.serviceAccountAdmin",
    "roles/iam.workloadIdentityUser"
  ])
  role = each.key
  members = [
    "serviceAccount:githubactions-cloudfunctions@mms-clp-playground2402-a-x11e.iam.gserviceaccount.com",
  ]
}