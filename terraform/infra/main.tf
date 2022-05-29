resource "kubernetes_namespace" "istio-system" {
  metadata {
    name = "istio-system"
  }
}

resource "helm_release" "istio_base" {
  repository = "https://istio-release.storage.googleapis.com/charts"
  version    = var.istio_version
  name       = "istio-base"
  chart      = "base"
  namespace  = kubernetes_namespace.istio-system.metadata[0].name
}

resource "helm_release" "istiod" {
  depends_on = [
    helm_release.istio_base
  ]
  repository = "https://istio-release.storage.googleapis.com/charts"
  version    = var.istio_version
  name       = "istiod"
  chart      = "istiod"
  namespace  = kubernetes_namespace.istio-system.metadata[0].name
  values     = [yamlencode(var.mesh_config)]
#   values     = [<<EOF
# meshConfig:
#   extensionProviders:
#   - name: "custom-ext-authz-http"
#     envoyExtAuthzHttp:
#       service: "authz-service.default.svc.cluster.local"
#       port: "8080"
#       includeRequestHeadersInCheck: ["Authorization"]
#   - name: "custom-ext-authz-grpc"
#     envoyExtAuthzGrpc:
#       service: "authz-service.default.svc.cluster.local"
#       port: "9090"
# EOF
#   ]
}

resource "helm_release" "gateway" {
  depends_on = [
    helm_release.istiod
  ]
  repository = "https://istio-release.storage.googleapis.com/charts"
  version    = var.istio_version
  name       = "istio-gateway"
  chart      = "gateway"
  namespace  = kubernetes_namespace.istio-system.metadata[0].name
}
