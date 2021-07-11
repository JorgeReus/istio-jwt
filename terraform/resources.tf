resource "kubernetes_manifest" "gateway_demo_gateway" {
  provider = kubernetes-alpha
  manifest = {
    "apiVersion" = "networking.istio.io/v1alpha3"
    "kind" = "Gateway"
    "metadata" = {
      "name" = "demo-gateway"
      "namespace" = "default"
    }
    "spec" = {
      "selector" = {
        "istio" = "ingressgateway"
      }
      "servers" = [
        {
          "hosts" = [
            "*",
          ]
          "port" = {
            "name" = "http"
            "number" = 80
            "protocol" = "HTTP"
          }
        },
      ]
    }
  }
}

resource "kubernetes_manifest" "authorizationpolicy_require_jwt" {
  provider = kubernetes-alpha
  manifest = {
    "apiVersion" = "security.istio.io/v1beta1"
    "kind" = "AuthorizationPolicy"
    "metadata" = {
      "name" = "require-jwt"
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
                  "test/*",
                ]
              }
            },
          ]
        },
      ]
      "selector" = {
        "matchLabels" = {
          "istio" = "ingressgateway"
        }
      }
    }
  }
}
