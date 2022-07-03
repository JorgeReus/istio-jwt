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
          image = "${var.registry_name}:5050/nginx:latest"
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
      name        = "http"
      port        = 80
      target_port = 80
    }

    type = "ClusterIP"
  }
}

resource "kubernetes_manifest" "nginx" {
  manifest = {
    apiVersion = "networking.istio.io/v1alpha3"
    kind       = "VirtualService"
    metadata = {
      name      = "nginx"
      namespace = "default"
    }
    spec = {
      gateways = [
        "default/istio-gateway-default",
      ]
      hosts = [
        "*",
      ]
      http = [
        {
          match = [
            {
              uri = {
                prefix = "/nginx/"
              }
            },
          ]
          name = "nginx"
          rewrite = {
            uri = "/"
          }
          route = [
            {
              destination = {
                host = "nginx.default.svc.cluster.local"
                port = {
                  number = 80
                }
              }
            }
          ]
        }
      ]
    }
  }
}
