data "google_client_config" "provider" {}

data "helm_template" "ingress-gateway" {
  name       = "istio-ingress"
  namespace  = "istio-system"
  chart = "../infra/istio-1.10.2/manifests/charts/gateways/istio-ingress"
}