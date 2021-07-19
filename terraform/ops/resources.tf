# resource "kubernetes_manifest" "gateway_demo_gateway" {
#   depends_on = [ data.helm_template.ingress-gateway ]
#   provider = kubernetes-alpha
#   manifest = {
#     "apiVersion" = "networking.istio.io/v1alpha3"
#     "kind" = "Gateway"
#     "metadata" = {
#       "name" = "demo-gateway"
#       "namespace" = "istio-system"
#     }
#     "spec" = {
#       "selector" = {
#         "istio" = "ingressgateway"
#       }
#       "servers" = [
#         {
#           "hosts" = [
#             "*",
#           ]
#           "port" = {
#             "name" = "http"
#             "number" = 80
#             "protocol" = "HTTP"
#           }
#         },
#       ]
#     }
#   }
# }
#
# resource "kubernetes_manifest" "authorizationpolicy_require_jwt" {
#   depends_on = [ kubernetes_manifest.gateway_demo_gateway ]
#   provider = kubernetes-alpha
#   manifest = {
#     "apiVersion" = "security.istio.io/v1beta1"
#     "kind" = "AuthorizationPolicy"
#     "metadata" = {
#       "name" = "require-jwt"
#       "namespace" = "istio-system"
#     }
#     "spec" = {
#       "action" = "ALLOW"
#       "rules" = [
#         {
#           "from" = [
#             {
#               "source" = {
#                 "requestPrincipals" = [
#                   "test/*",
#                 ]
#               }
#             },
#           ]
#         },
#       ]
#       "selector" = {
#         "matchLabels" = {
#           "istio" = "ingressgateway"
#         }
#       }
#     }
#   }
# }
#
# resource "kubernetes_manifest" "virtualService_jwt" {
#   depends_on = [ kubernetes_manifest.gateway_demo_gateway ]
#   provider = kubernetes-alpha
#   manifest = {
#     "apiVersion" = "networking.istio.io/v1alpha3"
#     "kind" = "VirtualService"
#     "metadata" = {
#       "name" = "jwk"
#       "namespace" = "istio-system"
#     }
#     "spec" = {
#       "hosts" = ["*"]
#       "gateways" = ["istio-system/demo-gateway"]
#       "http" = [
#         {
#           "name" = "jwk"
#           "match" = [
#             {
#               "uri" = {
#                 "prefix" = "/jwk/"
#               }
#             }
#           ]
#           "rewrite" = {
#             "uri" = "/"
#           }
#           "route" = [
#             {
#               "destination" = {
#                 "port" = {
#                   number = "80"
#                 }
#                 "host" = "jwk.default.svc.cluster.local"
#               }
#             }
#           ]
#         }
#       ]  
#     }
#   }
# }
#
# resource "kubernetes_manifest" "authorizationpolicy_allow_without_jwt" {
#   depends_on = [ kubernetes_manifest.gateway_demo_gateway ]
#   provider = kubernetes-alpha
#   manifest = {
#     "apiVersion" = "security.istio.io/v1beta1"
#     "kind" = "AuthorizationPolicy"
#     "metadata" = {
#       "name" = "require-jwt"
#       "namespace" = "istio-system"
#     }
#     "spec" = {
#       "rules" = [
#         {
#           "to" = [
#             {
#               "operation" = {
#                 "methods" = ["POST"]
#                 "paths" = ["jwk/jwt/generate"]
#               }
#             },
#           ]
#         },
#       ]
#       "selector" = {
#         "matchLabels" = {
#           "istio" = "ingressgateway"
#         }
#       }
#       "action" = "ALLOW"
#     }
#   }
# }