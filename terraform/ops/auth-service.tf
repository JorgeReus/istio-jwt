resource "kubernetes_manifest" "auth_service" {
  manifest = {
    "apiVersion" = "networking.istio.io/v1alpha3"
    "kind"       = "VirtualService"
    "metadata" = {
      "name"      = "auth-service"
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
                "prefix" = "/auth/"
              }
            },
          ]
          "name" = "auth-service"
          "rewrite" = {
            "uri" = "/"
          }
          "route" = [
            {
              "destination" = {
                "host" = "auth-service.default.svc.cluster.local"
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
