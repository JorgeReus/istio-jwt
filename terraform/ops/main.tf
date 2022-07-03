resource "kubernetes_manifest" "gateway_api_gateway" {
  manifest = {
    apiVersion = "networking.istio.io/v1beta1"
    kind       = "Gateway"
    metadata = {
      namespace = "default"
      name      = "istio-gateway-default"
    }
    spec = {
      selector = {
        app   = "istio-gateway"
        istio = "gateway"
      }
      servers = [
        {
          hosts = [
            "*",
          ]
          port = {
            name     = "http"
            number   = 80
            protocol = "HTTP"
          }
        },
      ]
    }
  }
}

resource "kubernetes_manifest" "request_authentication" {
  manifest = {
    apiVersion = "security.istio.io/v1beta1"
    kind       = "RequestAuthentication"
    metadata = {
      name      = "default-jwt-api"
      namespace = "istio-system"
    }
    spec = {
      jwtRules = [
        {
          issuer  = "my.jwt.issuer",
          jwksUri = "http://auth-service.default.svc.cluster.local/jwk/public"
        }
      ]
      selector = {
        matchLabels = {
          app   = "istio-gateway"
          istio = "gateway"
        }
      }
    }
  }
}

resource "kubernetes_manifest" "require_jwt" {
  manifest = {
    apiVersion = "security.istio.io/v1beta1"
    kind       = "AuthorizationPolicy"
    metadata = {
      name      = "require-jwt"
      namespace = "istio-system"
    }
    spec = {
      selector = {
        matchLabels = {
          app   = "istio-gateway"
          istio = "gateway"
        }
      }
      action = "ALLOW"
      rules = [
        {
          from = [
            {
              source = {
                requestPrincipals = [
                  "my.jwt.issuer/*",
                ]
              }
            },
          ]
        },
      ]
    }
  }
}

resource "kubernetes_manifest" "ext_authz" {
  manifest = {
    apiVersion = "security.istio.io/v1beta1"
    kind       = "AuthorizationPolicy"
    metadata = {
      name      = "ext-authz"
      namespace = "istio-system"
    }
    spec = {
      selector = {
        matchLabels = {
          app   = "istio-gateway"
          istio = "gateway"
        }
      }
      action = "CUSTOM"
      provider = {
        name = "custom-ext-authz-grpc"
      }
      rules = [
        {
          to = [
            {
              operation = {
                paths = [
                  "/nginx/",
                ]
              }
            }
          ]
        }
      ]
    }
  }
}

resource "kubernetes_manifest" "authorization_policy_default" {
  manifest = {
    apiVersion = "security.istio.io/v1beta1"
    kind       = "AuthorizationPolicy"
    metadata = {
      name      = "require-default-jwt-api"
      namespace = "istio-system"
    }
    spec = {
      action = "ALLOW"
      rules = [
        {
          to = [
            {
              operation = {
                methods = [
                  "GET",
                  "POST",
                ]
                paths = [
                  "/auth/*",
                ]
              }
            }
          ]
        }
      ]
      selector = {
        matchLabels = {
          app   = "istio-gateway"
          istio = "gateway"
        }
      }
    }
  }
}
