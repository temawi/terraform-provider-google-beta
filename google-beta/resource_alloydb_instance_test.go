package google

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccAlloydbInstance_update(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckAlloydbInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccAlloydbInstance_alloydbInstanceBasicExample(context),
			},
			{
				ResourceName:            "google_alloydb_instance.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"cluster", "instance_id", "reconciling", "update_time"},
			},
			{
				Config: testAccAlloydbInstance_update(context),
			},
			{
				ResourceName:            "google_alloydb_instance.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"cluster", "instance_id", "reconciling", "update_time"},
			},
		},
	})
}

func testAccAlloydbInstance_update(context map[string]interface{}) string {
	return Nprintf(`
resource "google_alloydb_instance" "default" {
  provider      = google-beta
  cluster       = google_alloydb_cluster.default.name
  instance_id   = "tf-test-alloydb-instance%{random_suffix}"
  instance_type = "PRIMARY"

  machine_config {
    cpu_count = 4
  }

  labels = {
	test = "tf-test-alloydb-instance%{random_suffix}"
  }

  depends_on = [google_service_networking_connection.vpc_connection]
}

resource "google_alloydb_cluster" "default" {
  provider   = google-beta
  cluster_id = "tf-test-alloydb-cluster%{random_suffix}"
  location   = "us-central1"
  network    = "projects/${data.google_project.project.number}/global/networks/${google_compute_network.default.name}"

  initial_user {
    password = "tf-test-alloydb-cluster%{random_suffix}"
  }
}

data "google_project" "project" {
  provider = google-beta
}

resource "google_compute_network" "default" {
  provider = google-beta
  name = "tf-test-alloydb-cluster%{random_suffix}"
}

resource "google_compute_global_address" "private_ip_alloc" {
  provider = google-beta
  name          =  "tf-test-alloydb-cluster%{random_suffix}"
  address_type  = "INTERNAL"
  purpose       = "VPC_PEERING"
  prefix_length = 16
  network       = google_compute_network.default.id
}

resource "google_service_networking_connection" "vpc_connection" {
  provider   = google-beta
  network                 = google_compute_network.default.id
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.private_ip_alloc.name]
}
`, context)
}
