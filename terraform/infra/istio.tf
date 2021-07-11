resource "kubernetes_namespace" "istio-system" {
  metadata {
    name = "istio-system"
  }
}

resource "helm_release" "istio-base" {
  name       = "istio-base"
  namespace = kubernetes_namespace.istio-system.metadata[0].name
  chart = "./istio-1.10.2/manifests/charts/base"
}


resource "helm_release" "istiod" {
  depends_on = [
    helm_release.istio-base
  ]
  name       = "istiod"
  namespace = kubernetes_namespace.istio-system.metadata[0].name
  chart = "./istio-1.10.2/manifests/charts/istio-control/istio-discovery"
}

resource "helm_release" "ingress-gateway" {
  depends_on = [
    helm_release.istiod
  ]
  name       = "istio-ingress"
  namespace = kubernetes_namespace.istio-system.metadata[0].name
  chart = "./istio-1.10.2/manifests/charts/gateways/istio-ingress"
}
