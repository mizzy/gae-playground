resource "google_compute_network" "gae_static_ip" {
  name                    = "gae-static-ip-network"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "gae_static_ip" {
  name          = "gae-static-ip-subnet"
  network       = google_compute_network.gae_static_ip.id
  ip_cidr_range = "10.0.0.0/28" // Error: Error waiting to create Connector: Error waiting for Creating Connector: Error code 3, message: Operation failed: Subnets used for VPC connectors must have a netmask of 28.
}

resource "google_compute_router" "gae_static_ip" {
  name    = "gae-static-ip-router"
  network = google_compute_network.gae_static_ip.id
}

resource "google_compute_address" "gae_static_ip" {
  name = "gae-static-ip"
}

resource "google_compute_router_nat" "gae_static_ip" {
  name                               = "gae-static-ip-nat"
  router                             = google_compute_router.gae_static_ip.name
  nat_ip_allocate_option             = "MANUAL_ONLY"
  nat_ips                            = [google_compute_address.gae_static_ip.self_link]
  source_subnetwork_ip_ranges_to_nat = "ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES"
}

resource "google_vpc_access_connector" "gae_static_ip" {
  name          = "gae-static-ip"
  machine_type  = "e2-micro"
  min_instances = 2
  max_instances = 3

  subnet {
    name = google_compute_subnetwork.gae_static_ip.name
  }
}
