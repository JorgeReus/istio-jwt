variable "authz_service_name" {
  default = "authz-service"
}

variable "authz_service_image_version" {
  default = "latest"
}

resource "kubernetes_deployment" "authz" {
  metadata {
    name = var.authz_service_name
    labels = {
      app = var.authz_service_name
    }
  }

  spec {
    replicas = 1

    selector {
      match_labels = {
        app = var.authz_service_name
      }
    }

    template {
      metadata {
        labels = {
          app = var.authz_service_name
        }
      }

      spec {
        container {
          image = "${var.registry_name}:5050/${var.authz_service_name}:${var.authz_service_image_version}"
          name  = var.authz_service_name
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

resource "kubernetes_service" "authz" {
  metadata {
    name = kubernetes_deployment.authz.metadata.0.name
  }
  spec {
    selector = {
      app = kubernetes_deployment.authz.metadata.0.name
    }
    port {
      name        = "http"
      port        = 8080
      target_port = 8080
    }
    port {
      name        = "grpc"
      port        = 9090
      target_port = 9090
    }

    type = "ClusterIP"
  }
}
