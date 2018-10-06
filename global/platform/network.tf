# Create the network
resource "google_compute_network" "neovencia-network" {
  name = "${var.platform-name}"
}

# Create firewall
#       SSH - 22
#       HTTP - 80
#       HTTPS - 443
resource "google_compute_firewall" "ssh" {
  name = "${var.platform-name}-ssh"
  network = "${google_compute_network.neovencia-network.name}"

  allow {
      protocol = "icmp"
  }

  allow {
      protocol = "tcp"
      ports = ["22", "80", "443"]
  }

  source_ranges = ["0.0.0.0/0"] #CAUTION!
}

# Creates root DNS zone
resource "google_dns_managed_zone" "neovencia-root-zone" {
  name = "neovencia"
  dns_name = "neovencia.com."
  description = "neovencia root DNS Zone"
}

resource "google_dns_managed_zone" "neovencia-prod-zone" {
  name = "neovencia-prod"
  dns_name = "prod.neovencia.com."
  description = "neovencia prod DNS Zone"
}

resource "google_dns_managed_zone" "neovencia-test-zone" {
  name = "neovencia-test"
  dns_name = "test.neovencia.com."
  description = "neovencia test DNS Zone"
}

# Create a new subnet for platform and region
resource "google_compute_subnetwork" "neovencia-subnet" {
    name = "dev-${var.platform-name}-${var.gcloud-region}"
    ip_cidr_range = "10.1.2.0/24"
    network = "${google_compute_network.neovencia-network.self_link}"
    region = "${var.gcloud-region}"
}
