provider "kubernetes" {
  config_path = "~/.kube/config"
  experiments {
    manifest_resource = true
  }
}

# provider "docker" {
#   host = "unix:///var/run/docker.sock"
#   registry_auth {
#     address     = "istio-test-registry"
#     # config_file = pathexpand("~/.docker/config.json")
#   }
# }
