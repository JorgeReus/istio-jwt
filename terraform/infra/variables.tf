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
            service                      = "authz-service.default.svc.cluster.local"
            port                         = "8080"
            includeRequestHeadersInCheck = ["Authorization"]
          }
        },
        {
          name = "custom-ext-authz-grpc",
          envoyExtAuthzGrpc = {
            service = "authz-service.default.svc.cluster.local"
            port    = "9090"
          }
        }
      ]
    }
  }
}

variable "registry_name" {
  default = "istio-test-registry"
}

variable "auth_service_name" {
  default = "auth-service"
}

variable "auth_service_image_version" {
  default = "latest"
}

variable "authz_service_name" {
  default = "authz-service"
}

variable "authz_service_image_version" {
  default = "latest"
}

variable "private_jwk" {
  default = "eyJ1c2UiOiJzaWciLCJrdHkiOiJFQyIsImtpZCI6IkFyU3lpVmhGbzFEVlFZUjBQNEVWR1Z1UnZTeFNZYWtsVGFQcDlNT1ZfUE09IiwiY3J2IjoiUC0yNTYiLCJhbGciOiJFUzI1NiIsIngiOiJ6TzhLLVpTQkdmSnNIazJiWC1ySy04Q2tkR21zazBWTDh2UURJeTcweVBFIiwieSI6InZLZGhncVc4OHUzdHZhTVVTbFhoeUUxZWVSc2tnd21QRGp0SEpzTW5vT1kiLCJkIjoiU3hCTXhaWFByOHNRNGRRRGpSQ2sybDk1UkVHR1ZkTV9HTUVyMWJoU2VNdyJ9"
}

variable "public_jwk" {
  default = "eyJ1c2UiOiJzaWciLCJrdHkiOiJFQyIsImtpZCI6IkFyU3lpVmhGbzFEVlFZUjBQNEVWR1Z1UnZTeFNZYWtsVGFQcDlNT1ZfUE09IiwiY3J2IjoiUC0yNTYiLCJhbGciOiJFUzI1NiIsIngiOiJ6TzhLLVpTQkdmSnNIazJiWC1ySy04Q2tkR21zazBWTDh2UURJeTcweVBFIiwieSI6InZLZGhncVc4OHUzdHZhTVVTbFhoeUUxZWVSc2tnd21QRGp0SEpzTW5vT1kifQ=="
}
