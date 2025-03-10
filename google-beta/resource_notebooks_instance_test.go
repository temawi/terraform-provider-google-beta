package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNotebooksInstance_create_vm_image(t *testing.T) {
	t.Parallel()

	prefix := fmt.Sprintf("%d", randInt(t))
	name := fmt.Sprintf("tf-%s", prefix)

	vcrTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNotebooksInstance_create_vm_image(name),
			},
			{
				ResourceName:            "google_notebooks_instance.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"vm_image", "metadata"},
			},
		},
	})
}

func TestAccNotebooksInstance_update(t *testing.T) {
	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNotebooksInstance_basic(context),
			},
			{
				ResourceName:            "google_notebooks_instance.instance",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"vm_image"},
			},
			{
				Config: testAccNotebooksInstance_update(context),
			},
			{
				ResourceName:            "google_notebooks_instance.instance",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"vm_image"},
			},
			{
				Config: testAccNotebooksInstance_basic(context),
			},
			{
				ResourceName:            "google_notebooks_instance.instance",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"vm_image"},
			},
		},
	})
}

func testAccNotebooksInstance_create_vm_image(name string) string {
	return fmt.Sprintf(`

resource "google_notebooks_instance" "test" {
  name = "%s"
  location = "us-west1-a"
  machine_type = "e2-medium"
  metadata = {
    proxy-mode = "service_account"
    terraform  = "true"
  }

  nic_type = "VIRTIO_NET"

  reservation_affinity {
    consume_reservation_type = "NO_RESERVATION"
  }

  vm_image {
    project      = "deeplearning-platform-release"
    image_family = "tf-latest-cpu"
  }
}
`, name)
}

func testAccNotebooksInstance_basic(context map[string]interface{}) string {
	return Nprintf(`
resource "google_notebooks_instance" "instance" {
  name = "tf-test-notebooks-instance%{random_suffix}"
  location = "us-central1-a"
  machine_type = "e2-medium"

  vm_image {
    project      = "deeplearning-platform-release"
    image_family = "tf-latest-cpu"
  }
}
`, context)
}

func testAccNotebooksInstance_update(context map[string]interface{}) string {
	return Nprintf(`
resource "google_notebooks_instance" "instance" {
  name = "tf-test-notebooks-instance%{random_suffix}"
  location = "us-central1-a"
  machine_type = "e2-medium"

  nic_type = "VIRTIO_NET"

  reservation_affinity {
    consume_reservation_type = "NO_RESERVATION"
  }

  vm_image {
    project      = "deeplearning-platform-release"
    image_family = "tf-latest-cpu"
  }

  labels = {
    key = "value"
  }
}
`, context)
}
