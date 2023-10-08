terraform {}

provider "google" {
  project = var.project_id
}

data "google_client_config" "default" {}

provider "kubernetes" {
  host                   = "https://${module.gke.endpoint}"
  token                  = data.google_client_config.default.access_token
  cluster_ca_certificate = base64decode(module.gke.ca_certificate)
}

module "gke" {
  source                     = "terraform-google-modules/kubernetes-engine/google//modules/beta-autopilot-private-cluster"
  project_id                 = var.project_id
  name                       = "sandbox"
  region                     = "us-central1"
  zones                      = ["us-central1-a"]
  network                    = "sandbox"
  subnetwork                 = "sandbox-us-central1"
  ip_range_pods              = "pod-range"
  ip_range_services          = "service-range"
  horizontal_pod_autoscaling = true
  enable_private_endpoint    = false
  enable_private_nodes       = true
  master_ipv4_cidr_block     = "10.0.0.0/28"
}

#
# GLOBAL K8S RESOURCES
#
locals {
  envs = ["dev", "stg", "prd"]
}

resource "kubernetes_namespace" "ns" {
  for_each = { for env in local.envs : env => env }
  metadata {
    name = each.value
  }
}

resource "kubernetes_secret" "registry_secret" {
  for_each = { for env in local.envs : env => env }

  metadata {
    name      = "dockerconfigjson"
    namespace = kubernetes_namespace.ns[each.key].metadata.0.name
  }

  type = "kubernetes.io/dockerconfigjson"

  data = {
    ".dockerconfigjson" = jsonencode({
      auths = {
        "${var.registry_server}" = {
          "username" = var.registry_username
          "password" = var.registry_password
          "auth"     = base64encode("${var.registry_username}:${var.registry_password}")
        }
      }
    })
  }
}

#
# SERVICES SECRET
#
resource "random_password" "admin_pass" {
  for_each = { for env in local.envs : env => env }
  length   = 16
}
resource "random_password" "jwt_secret" {
  for_each = { for env in local.envs : env => env }
  length   = 64
}

resource "kubernetes_secret" "auth_secret" {
  for_each = { for env in local.envs : env => env }

  metadata {
    name      = "auth"
    namespace = kubernetes_namespace.ns[each.key].metadata.0.name
  }

  type = "Opaque"

  data = {
    ADMIN_PASS = random_password.admin_pass[each.key].result
    JWT_SECRET = random_password.jwt_secret[each.key].result
  }
}

resource "kubernetes_secret" "user_secret" {
  for_each = { for env in local.envs : env => env }

  metadata {
    name      = "user"
    namespace = kubernetes_namespace.ns[each.key].metadata.0.name
  }

  type = "Opaque"

  data = {
    JWT_SECRET = random_password.jwt_secret[each.key].result
  }
}

#
# IP Reservation for GKE ingress
#
resource "google_compute_global_address" "ingress" {
  for_each = { for env in local.envs : env => env }
  name     = "sandbox-ingress-${each.key}"
}
