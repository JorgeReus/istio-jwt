variable "istio_version" {
  default = "1.13.4"
}

variable "mesh_config" {
  default = {
    meshConfig = {
      extensionProviders = [
        {
          name = "custom-ext-authz-http",
          envoyExtAuthzHttp = {
            service = "authz-service.default.svc.cluster.local"
            port = "8080"
            includeRequestHeadersInCheck = ["Authorization"]
          }
        },
        {
          name = "custom-ext-authz-grpc",
          envoyExtAuthzGrpc = {
            service = "authz-service.default.svc.cluster.local"
            port = "9090"
          }
        }
      ]
    }
  }
}
