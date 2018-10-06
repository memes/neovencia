resource "google_container_cluster" "neovencia-cluster" {
  name = "neovencia-cluster"
  network = "${google_compute_network.neovencia-network.name}"
  subnetwork = "${google_compute_subnetwork.neovencia-subnet.name}"
  zone = "${var.gcloud-zone}"

  initial_node_count = 1

  master_auth {
      username = "kube-admin"
      password = "8PW6A4dApy$Bn#m_"
  }

  node_config {
      machine_type = "n1-standard-1"

      oauth_scopes = [
          "https://www.googleapis.com/auth/projecthosting",
          "https://www.googleapis.com/auth/devstorage.full_control",
          "https://www.googleapis.com/auth/monitoring",
          "https://www.googleapis.com/auth/logging.write",
          "https://www.googleapis.com/auth/compute",
          "https://www.googleapis.com/auth/cloud-platform"
      ]
  }
}