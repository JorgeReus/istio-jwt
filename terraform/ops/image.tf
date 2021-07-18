resource "docker_registry_image" "jwt-service" {
  name = "gcr.io/${data.google_client_config.provider.project}/jwk-istio-demo:latest"
  build {
    context = "../../"
  }
}
