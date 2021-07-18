resource "kubernetes_deployment" "app" {
  metadata {
    name = var.app-name
    labels = {
      app = var.app-name
    }
  }

  spec {
    replicas = 1

    selector {
      match_labels = {
        app = var.app-name
      }
    }

    template {
      metadata {
        labels = {
          app = var.app-name
        }
      }

      spec {
        container {
          image = docker_registry_image.jwt-service.name
          name  = var.app-name

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

          liveness_probe {
            http_get {
              path = "/healthz"
              port = 80
            }
            
            # 300 seconds of unhealthiness to destroy
            failure_threshold = 30
            period_seconds = 10
          }

          readiness_probe {
            http_get {
              path = "/healthz"
              port = 8080
            }
            # 10 seconds to start
            failure_threshold = 5
            period_seconds = 2
          }
        }
      }
    }
  }
}

resource "kubernetes_service" "svc" {
  metadata {
    name = kubernetes_deployment.app.metadata.0.name
  }
  spec {
    selector = {
      app = kubernetes_deployment.app.metadata.0.name
    }
    port {
      port        = 80
      target_port = 8080
    }

    type = "ClusterIP"
  }
}
