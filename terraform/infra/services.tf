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
            name  = "JWT_ISSUER"
            value = "my.jwt.issuer"
          }
          env {
            name  = "JWT_AUDIENCE"
            value = "my.jwt.audience"
          }
          env {
            name  = "SCHEME_DOMAIN_NAME"
            value = "http://localhost:8080"
          }
          env {
            name  = "PRIVATE_JWK"
            value = var.private_jwk
          }

          env {
            name  = "PUBLIC_JWK"
            value = var.public_jwk
          }

          env {
            name  = "PORT"
            value = 8080
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
