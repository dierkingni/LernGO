resource "google_project_iam_binding" "service-account" {
  project = "mms-clp-playground2402-a-i2ar"
  //role    = "roles/storage.objectCreator"
  role = "roles/storage.objectUser"

  members = [
    "serviceAccount:mms-clp-playground2402-a-i2ar@appspot.gserviceaccount.com",
  ]
}