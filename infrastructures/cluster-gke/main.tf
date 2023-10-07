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
