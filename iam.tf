resource "google_project_iam_binding" "service-account" {
  project = "mms-clp-playground2402-a-i2ar"
  //role    = "roles/storage.objectCreator"
  role = "roles/storage.objectUser"

  members = [
    "serviceAccount:mms-clp-playground2402-a-i2ar@appspot.gserviceaccount.com",
  ]
}

resource "google_project_iam_binding" "github-service-account" {
  project = "mms-clp-playground2402-a-i2ar"
  //role    = "roles/storage.objectCreator"
  role = [
    "roles/storage.admin",
    "roles/resourcemanager.projectIamAdmin",
    "roles/cloudfunctions.admin",
    "roles/cloudfunctions.serviceAgent"
  ]

  members = [
    "serviceAccount:githubactions@mms-clp-playground2402-a-i2ar.iam.gserviceaccount.com",
  ]
}