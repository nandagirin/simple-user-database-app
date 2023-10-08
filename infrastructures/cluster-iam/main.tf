terraform {}

provider "google" {
  project = var.project_id
}

resource "google_iam_workload_identity_pool" "wip" {
  workload_identity_pool_id = "ci-wip"
  display_name              = "CI WIP"
}

resource "google_iam_workload_identity_pool_provider" "wip_provider" {
  workload_identity_pool_id          = google_iam_workload_identity_pool.wip.workload_identity_pool_id
  workload_identity_pool_provider_id = "ci-wipp"
  display_name                       = "CI WIPP"
  attribute_mapping = {
    "google.subject"       = "assertion.sub"
    "attribute.actor"      = "assertion.actor"
    "attribute.repository" = "assertion.repository"
  }
  oidc {
    issuer_uri = var.issuer_uri
  }
}

resource "google_service_account" "ci_sa" {
  account_id   = "sa-ci-pipeline"
  display_name = "CI Service Account"
}

module "service_account-iam-bindings" {
  source = "terraform-google-modules/iam/google//modules/service_accounts_iam"

  service_accounts = [google_service_account.ci_sa.email]
  project          = var.project_id
  mode             = "additive"
  bindings = {
    "roles/iam.workloadIdentityUser" = [
      "principalSet://iam.googleapis.com/${google_iam_workload_identity_pool.wip.name}/attribute.repository/${var.allowed_repo}",
    ]
  }
}

module "project_iam_bindings" {
  source   = "terraform-google-modules/iam/google//modules/projects_iam"
  projects = [var.project_id]
  mode     = "additive"

  bindings = {
    "roles/container.developer" = [
      "serviceAccount:${google_service_account.ci_sa.email}",
    ]
  }
}
