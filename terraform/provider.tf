provider "google" {
  region      = "us-central1"
  zone        = "us-central1-c"
}

provider "helm" {
  kubernetes {
    config_path = "~/.kube/config"
  }
}

provider "kubernetes-alpha" {
  config_path = "~/.kube/config"
}
