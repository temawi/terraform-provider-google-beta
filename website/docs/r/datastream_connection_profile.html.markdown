---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This file is automatically generated by Magic Modules and manual
#     changes will be clobbered when the file is regenerated.
#
#     Please read more about how to change this file in
#     .github/CONTRIBUTING.md.
#
# ----------------------------------------------------------------------------
subcategory: "Datastream"
page_title: "Google: google_datastream_connection_profile"
description: |-
  A set of reusable connection configurations to be used as a source or destination for a stream.
---

# google\_datastream\_connection\_profile

A set of reusable connection configurations to be used as a source or destination for a stream.


To get more information about ConnectionProfile, see:

* [API documentation](https://cloud.google.com/datastream/docs/reference/rest/v1/projects.locations.connectionProfiles)
* How-to Guides
    * [Official Documentation](https://cloud.google.com/datastream/docs/create-connection-profiles)

~> **Warning:** All arguments including `oracle_profile.password`, `mysql_profile.password`, `mysql_profile.ssl_config.client_key`, `mysql_profile.ssl_config.client_certificate`, `mysql_profile.ssl_config.ca_certificate`, `postgresql_profile.password`, `forward_ssh_connectivity.password`, and `forward_ssh_connectivity.private_key` will be stored in the raw
state as plain-text. [Read more about sensitive data in state](https://www.terraform.io/language/state/sensitive-data).

<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_working_dir=datastream_connection_profile_basic&cloudshell_image=gcr.io%2Fgraphite-cloud-shell-images%2Fterraform%3Alatest&open_in_editor=main.tf&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Datastream Connection Profile Basic


```hcl
resource "google_datastream_connection_profile" "default" {
	display_name          = "Connection profile"
	location              = "us-central1"
	connection_profile_id = "my-profile"

	gcs_profile {
		bucket    = "my-bucket"
		root_path = "/path"
	}
}
```
<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_working_dir=datastream_connection_profile_bigquery_private_connection&cloudshell_image=gcr.io%2Fgraphite-cloud-shell-images%2Fterraform%3Alatest&open_in_editor=main.tf&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Datastream Connection Profile Bigquery Private Connection


```hcl
resource "google_datastream_private_connection" "private_connection" {
	display_name          = "Connection profile"
	location              = "us-central1"
	private_connection_id = "my-connection"

	labels = {
		key = "value"
	}

	vpc_peering_config {
		vpc = google_compute_network.default.id
		subnet = "10.0.0.0/29"
	}
}

resource "google_compute_network" "default" {
	name = "my-network"
}

resource "google_datastream_connection_profile" "default" {
	display_name          = "Connection profile"
	location              = "us-central1"
	connection_profile_id = "my-profile"

	bigquery_profile {}

	private_connectivity {
		private_connection = google_datastream_private_connection.private_connection.id
	}
}
```
<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_working_dir=datastream_connection_profile_full&cloudshell_image=gcr.io%2Fgraphite-cloud-shell-images%2Fterraform%3Alatest&open_in_editor=main.tf&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Datastream Connection Profile Full


```hcl
resource "google_datastream_connection_profile" "default" {
	display_name          = "Connection profile"
	location              = "us-central1"
	connection_profile_id = "my-profile"

	gcs_profile {
		bucket    = "my-bucket"
		root_path = "/path"
	}

	forward_ssh_connectivity {
		hostname = "google.com"
		username = "my-user"
		port     = 8022
		password = "swordfish"
	}
	labels = {
		key = "value"
	}
}
```
## Example Usage - Datastream Connection Profile Postgres


```hcl
resource "google_sql_database_instance" "instance" {
    name             = "my-instance"
    database_version = "POSTGRES_14"
    region           = "us-central1"
    settings {
        tier = "db-f1-micro"

        ip_configuration {

            // Datastream IPs will vary by region.
            authorized_networks {
                value = "34.71.242.81"
            }

            authorized_networks {
                value = "34.72.28.29"
            }

            authorized_networks {
                value = "34.67.6.157"
            }

            authorized_networks {
                value = "34.67.234.134"
            }

            authorized_networks {
                value = "34.72.239.218"
            }
        }
    }

    deletion_protection  = "true"
}

resource "google_sql_database" "db" {
    instance = google_sql_database_instance.instance.name
    name     = "db"
}

resource "random_password" "pwd" {
    length = 16
    special = false
}

resource "google_sql_user" "user" {
    name = "user"
    instance = google_sql_database_instance.instance.name
    password = random_password.pwd.result
}

resource "google_datastream_connection_profile" "default" {
    display_name          = "Connection profile"
    location              = "us-central1"
    connection_profile_id = "my-profile"

    postgresql_profile {
        hostname = google_sql_database_instance.instance.public_ip_address
        username = google_sql_user.user.name
        password = google_sql_user.user.password
        database = google_sql_database.db.name
    }
}
```

## Argument Reference

The following arguments are supported:


* `display_name` -
  (Required)
  Display name.

* `connection_profile_id` -
  (Required)
  The connection profile identifier.

* `location` -
  (Required)
  The name of the location this connection profile is located in.


- - -


* `labels` -
  (Optional)
  Labels.

* `oracle_profile` -
  (Optional)
  Oracle database profile.
  Structure is [documented below](#nested_oracle_profile).

* `gcs_profile` -
  (Optional)
  Cloud Storage bucket profile.
  Structure is [documented below](#nested_gcs_profile).

* `mysql_profile` -
  (Optional)
  MySQL database profile.
  Structure is [documented below](#nested_mysql_profile).

* `bigquery_profile` -
  (Optional)
  BigQuery warehouse profile.

* `postgresql_profile` -
  (Optional)
  PostgreSQL database profile.
  Structure is [documented below](#nested_postgresql_profile).

* `forward_ssh_connectivity` -
  (Optional)
  Forward SSH tunnel connectivity.
  Structure is [documented below](#nested_forward_ssh_connectivity).

* `private_connectivity` -
  (Optional)
  Private connectivity.
  Structure is [documented below](#nested_private_connectivity).

* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.


<a name="nested_oracle_profile"></a>The `oracle_profile` block supports:

* `hostname` -
  (Required)
  Hostname for the Oracle connection.

* `port` -
  (Optional)
  Port for the Oracle connection.

* `username` -
  (Required)
  Username for the Oracle connection.

* `password` -
  (Required)
  Password for the Oracle connection.
  **Note**: This property is sensitive and will not be displayed in the plan.

* `database_service` -
  (Required)
  Database for the Oracle connection.

* `connection_attributes` -
  (Optional)
  Connection string attributes

<a name="nested_gcs_profile"></a>The `gcs_profile` block supports:

* `bucket` -
  (Required)
  The Cloud Storage bucket name.

* `root_path` -
  (Optional)
  The root path inside the Cloud Storage bucket.

<a name="nested_mysql_profile"></a>The `mysql_profile` block supports:

* `hostname` -
  (Required)
  Hostname for the MySQL connection.

* `port` -
  (Optional)
  Port for the MySQL connection.

* `username` -
  (Required)
  Username for the MySQL connection.

* `password` -
  (Required)
  Password for the MySQL connection.
  **Note**: This property is sensitive and will not be displayed in the plan.

* `ssl_config` -
  (Optional)
  SSL configuration for the MySQL connection.
  Structure is [documented below](#nested_ssl_config).


<a name="nested_ssl_config"></a>The `ssl_config` block supports:

* `client_key` -
  (Optional)
  PEM-encoded private key associated with the Client Certificate.
  If this field is used then the 'client_certificate' and the
  'ca_certificate' fields are mandatory.
  **Note**: This property is sensitive and will not be displayed in the plan.

* `client_key_set` -
  Indicates whether the clientKey field is set.

* `client_certificate` -
  (Optional)
  PEM-encoded certificate that will be used by the replica to
  authenticate against the source database server. If this field
  is used then the 'clientKey' and the 'caCertificate' fields are
  mandatory.
  **Note**: This property is sensitive and will not be displayed in the plan.

* `client_certificate_set` -
  Indicates whether the clientCertificate field is set.

* `ca_certificate` -
  (Optional)
  PEM-encoded certificate of the CA that signed the source database
  server's certificate.
  **Note**: This property is sensitive and will not be displayed in the plan.

* `ca_certificate_set` -
  Indicates whether the clientKey field is set.

<a name="nested_postgresql_profile"></a>The `postgresql_profile` block supports:

* `hostname` -
  (Required)
  Hostname for the PostgreSQL connection.

* `port` -
  (Optional)
  Port for the PostgreSQL connection.

* `username` -
  (Required)
  Username for the PostgreSQL connection.

* `password` -
  (Required)
  Password for the PostgreSQL connection.
  **Note**: This property is sensitive and will not be displayed in the plan.

* `database` -
  (Required)
  Database for the PostgreSQL connection.

<a name="nested_forward_ssh_connectivity"></a>The `forward_ssh_connectivity` block supports:

* `hostname` -
  (Required)
  Hostname for the SSH tunnel.

* `username` -
  (Required)
  Username for the SSH tunnel.

* `port` -
  (Optional)
  Port for the SSH tunnel.

* `password` -
  (Optional)
  SSH password.
  **Note**: This property is sensitive and will not be displayed in the plan.

* `private_key` -
  (Optional)
  SSH private key.
  **Note**: This property is sensitive and will not be displayed in the plan.

<a name="nested_private_connectivity"></a>The `private_connectivity` block supports:

* `private_connection` -
  (Required)
  A reference to a private connection resource. Format: `projects/{project}/locations/{location}/privateConnections/{name}`

## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `projects/{{project}}/locations/{{location}}/connectionProfiles/{{connection_profile_id}}`

* `name` -
  The resource's name.


## Timeouts

This resource provides the following
[Timeouts](/docs/configuration/resources.html#timeouts) configuration options:

- `create` - Default is 20 minutes.
- `update` - Default is 20 minutes.
- `delete` - Default is 20 minutes.

## Import


ConnectionProfile can be imported using any of these accepted formats:

```
$ terraform import google_datastream_connection_profile.default projects/{{project}}/locations/{{location}}/connectionProfiles/{{connection_profile_id}}
$ terraform import google_datastream_connection_profile.default {{project}}/{{location}}/{{connection_profile_id}}
$ terraform import google_datastream_connection_profile.default {{location}}/{{connection_profile_id}}
```

## User Project Overrides

This resource supports [User Project Overrides](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#user_project_override).
