variable "auth_service_name" {
  default = "auth-service"
}

variable "auth_service_image_version" {
  default = "latest"
}

resource "kubernetes_deployment" "auth" {
  metadata {
    name = var.auth_service_name
    labels = {
      app = var.auth_service_name
    }
  }

  spec {
    replicas = 1

    selector {
      match_labels = {
        app = var.auth_service_name
      }
    }

    template {
      metadata {
        labels = {
          app = var.auth_service_name
        }
      }
      spec {
        container {
          image = "${var.registry_name}:5050/${var.auth_service_name}:${var.auth_service_image_version}"
          name  = var.auth_service_name
          env {
            name = "JWT_ISSUER"
            value = "my.jwt.issuer"
          }
          env {
            name = "JWT_AUDIENCE"
            value = "my.jwt.audience"
          }
          env {
            name = "SCHEME_DOMAIN_NAME"
            value = "http://localhost:8080"
          }
          env {
            name = "PRIVATE_JWK"
            value = "eyJ1c2UiOiJzaWciLCJrdHkiOiJFQyIsImtpZCI6IkFyU3lpVmhGbzFEVlFZUjBQNEVWR1Z1UnZTeFNZYWtsVGFQcDlNT1ZfUE09IiwiY3J2IjoiUC0yNTYiLCJhbGciOiJFUzI1NiIsIngiOiJ6TzhLLVpTQkdmSnNIazJiWC1ySy04Q2tkR21zazBWTDh2UURJeTcweVBFIiwieSI6InZLZGhncVc4OHUzdHZhTVVTbFhoeUUxZWVSc2tnd21QRGp0SEpzTW5vT1kiLCJkIjoiU3hCTXhaWFByOHNRNGRRRGpSQ2sybDk1UkVHR1ZkTV9HTUVyMWJoU2VNdyJ9"
          }

          env {
            name = "PUBLIC_JWK"
            value = "eyJ1c2UiOiJzaWciLCJrdHkiOiJFQyIsImtpZCI6IkFyU3lpVmhGbzFEVlFZUjBQNEVWR1Z1UnZTeFNZYWtsVGFQcDlNT1ZfUE09IiwiY3J2IjoiUC0yNTYiLCJhbGciOiJFUzI1NiIsIngiOiJ6TzhLLVpTQkdmSnNIazJiWC1ySy04Q2tkR21zazBWTDh2UURJeTcweVBFIiwieSI6InZLZGhncVc4OHUzdHZhTVVTbFhoeUUxZWVSc2tnd21QRGp0SEpzTW5vT1kifQ=="
          }
          resources {
            limits = {
              cpu    = "1"
              memory = "1024Mi"
            }
            requests = {
              cpu    = "250m"
              memory = "50Mi"
            }
          }
        }
      }
    }
  }
}

resource "kubernetes_service" "auth" {
  metadata {
    name = kubernetes_deployment.auth.metadata.0.name
  }
  spec {
    selector = {
      app = kubernetes_deployment.auth.metadata.0.name
    }
    port {
      name        = "http"
      port        = 80
      target_port = 8080
    }

    type = "ClusterIP"
  }
}
