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

resource "kubernetes_manifest" "gateway_api_gateway" {
  manifest = {
    "apiVersion" = "networking.istio.io/v1beta1"
    "kind"       = "Gateway"
    "metadata" = {
      "namespace" = "default"
      "name"      = "istio-gateway-default"
    }
    "spec" = {
      "selector" = {
        "app"   = "istio-gateway"
        "istio" = "gateway"
      }
      "servers" = [
        {
          "hosts" = [
            "*",
          ]
          "port" = {
            "name"     = "http"
            "number"   = 80
            "protocol" = "HTTP"
          }
        },
      ]
    }
  }
}

resource "kubernetes_manifest" "request_authentication" {
  manifest = {
    "apiVersion" = "security.istio.io/v1beta1"
    "kind" = "RequestAuthentication"
    "metadata" = {
      "name" = "default-jwt-api"
      "namespace" = "istio-system"
    }
    "spec" = {
      "jwtRules" = [
          {
            "issuer" = "my.jwt.issuer",
            "jwksUri" = "http://auth-service.default.svc.cluster.local/jwk/public"
          }
      ]
      "selector" = {
        "matchLabels" = {
          "app"   = "istio-gateway"
          "istio" = "gateway"
        }
      }
    }
  }
}

resource "kubernetes_manifest" "ext_authz" {
  manifest = {
    "apiVersion" = "security.istio.io/v1beta1"
    "kind" = "AuthorizationPolicy"
    "metadata" = {
      "name" = "ext-authz"
      "namespace" = "istio-system"
    }
    "spec" = {
      "selector" = {
        "matchLabels" = {
          "app"   = "istio-gateway"
          "istio" = "gateway"
        }
      }
      action = "CUSTOM"
      provider = {
        name = "custom-ext-authz-grpc"
      }
      "rules" = [
        {
          "to" = [
            {
              "operation" = {
                "paths" = [
                  "/nginx/",
                ]
              }
            },
          ]
        },
      ]
    }
  }
}

resource "kubernetes_manifest" "authorization_policy_default" {
  manifest = {
    "apiVersion" = "security.istio.io/v1beta1"
    "kind" = "AuthorizationPolicy"
    "metadata" = {
      "name" = "require-default-jwt-api"
      "namespace" = "istio-system"
    }
    "spec" = {
      "action" = "ALLOW"
      "rules" = [
        {
          "from" = [
            {
              "source" = {
                "requestPrincipals" = [
                  "my.jwt.issuer/*",
                ]
              }
            },
          ]
        },
        {
          "to" = [
            {
              "operation" = {
                "methods" = [
                  "GET",
                  "POST",
                ]
                "paths" = [
                  "/auth/*",
                ]
              }
            },
          ]
        },
      ]
      "selector" = {
        "matchLabels" = {
          "app"   = "istio-gateway"
          "istio" = "gateway"
        }
      }
    }
  }
}

resource "kubernetes_manifest" "nginx" {
  manifest = {
    "apiVersion" = "networking.istio.io/v1alpha3"
    "kind" = "VirtualService"
    "metadata" = {
      "name" = "nginx"
      "namespace" = "default"
    }
    "spec" = {
      "gateways" = [
        "default/istio-gateway-default",
      ]
      "hosts" = [
        "*",
      ]
      "http" = [
        {
          "match" = [
            {
              "uri" = {
                "prefix" = "/nginx/"
              }
            },
          ]
          "name" = "nginx"
          "rewrite" = {
            "uri" = "/"
          }
          "route" = [
            {
              "destination" = {
                "host" = "nginx.default.svc.cluster.local"
                "port" = {
                  "number" = 80
                }
              }
            },
          ]
        },
      ]
    }
  }
}
