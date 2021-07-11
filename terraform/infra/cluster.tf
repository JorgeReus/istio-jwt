resource "google_service_account" "default" {
  account_id   = "istio-k8s-demo"
  display_name = "Service Account"
}

resource "google_container_cluster" "primary" {
  name               = "istio-jwt-demo"
  location           = "us-central1-a"
  initial_node_count = 1
  node_config {
    machine_type = "n1-standard-2"
    service_account = google_service_account.default.email
    oauth_scopes = [
      "https://www.googleapis.com/auth/cloud-platform"
    ]
  }
  timeouts {
    create = "30m"
    update = "40m"
  }
}



