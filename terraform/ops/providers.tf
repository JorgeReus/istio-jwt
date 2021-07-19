# provider "kubernetes-alpha" {
#   host  = "https://${google_container_cluster.primary.endpoint}"
#   token = data.google_client_config.provider.access_token
#   cluster_ca_certificate = base64decode(
#     google_container_cluster.primary.master_auth[0].cluster_ca_certificate,
#   )
# }


provider "kubernetes" {
  config_context = "minikube"
  config_path = "~/.kube/config"
}

provider "docker" {
  host = "unix:///var/run/docker.sock"
  registry_auth {
    address     = "gcr.io"
    config_file = pathexpand("~/.docker/config.json")
  }
}

provider "kubernetes-alpha" {
  config_path = "~/.kube/config"
  config_context_cluster   = "minikube"
}