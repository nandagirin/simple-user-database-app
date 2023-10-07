terraform {}

provider "google" {
  project = var.project_id
}

module "vpc" {
  source  = "terraform-google-modules/network/google"
  version = "~> 7.3"

  project_id   = var.project_id
  network_name = "sandbox"
  routing_mode = "REGIONAL"

  subnets = [
    {
      subnet_name           = "sandbox-us-central1"
      subnet_ip             = "10.10.10.0/24"
      subnet_region         = "us-central1"
      subnet_private_access = "true"
    },
  ]

  secondary_ranges = {
    "sandbox-us-central1" = [
      {
        range_name    = "pod-range"
        ip_cidr_range = "192.168.0.0/16"
      },
      {
        range_name    = "service-range"
        ip_cidr_range = "172.16.0.0/22"
      },
    ]
  }
}

module "cloud_nat" {
  source        = "terraform-google-modules/cloud-nat/google"
  version       = "~> 1.2"
  project_id    = var.project_id
  region        = "us-central1"
  network       = "sandbox"
  create_router = true
  router        = "sandbox-router"
}
