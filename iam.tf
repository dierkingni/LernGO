resource "google_project_iam_binding" "service-account" {
  project = "mms-clp-playground2402-a-i2ar"
  for_each = toset([
    "roles/storage.objectUser",
    "roles/iam.serviceAccountUser",
    "roles/artifactregistry.reader"
  ])
  role = each.key
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
    "roles/cloudfunctions.serviceAgent"

  ])
  role = each.key
  members = [
    "serviceAccount:githubactions@mms-clp-playground2402-a-i2ar.iam.gserviceaccount.com",
  ]
}