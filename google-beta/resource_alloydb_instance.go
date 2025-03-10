// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAlloydbInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlloydbInstanceCreate,
		Read:   resourceAlloydbInstanceRead,
		Update: resourceAlloydbInstanceUpdate,
		Delete: resourceAlloydbInstanceDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAlloydbInstanceImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"cluster": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description: `Identifies the alloydb cluster. Must be in the format
'projects/{project}/locations/{location}/clusters/{cluster_id}'`,
			},
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The ID of the alloydb instance.`,
			},
			"instance_type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateEnum([]string{"PRIMARY", "READ_POOL"}),
				Description:  `The type of the instance. Possible values: ["PRIMARY", "READ_POOL"]`,
			},
			"annotations": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `Annotations to allow client tools to store small amount of arbitrary data. This is distinct from labels.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"availability_type": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ValidateFunc: validateEnum([]string{"AVAILABILITY_TYPE_UNSPECIFIED", "ZONAL", "REGIONAL", ""}),
				Description:  `Availability type of an Instance. Defaults to REGIONAL for both primary and read instances. Note that primary and read instances can have different availability types. Possible values: ["AVAILABILITY_TYPE_UNSPECIFIED", "ZONAL", "REGIONAL"]`,
			},
			"database_flags": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `Database flags. Set at instance level. * They are copied from primary instance on read instance creation. * Read instances can set new or override existing flags that are relevant for reads, e.g. for enabling columnar cache on a read instance. Flags set on read instance may or may not be present on primary.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"display_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `User-settable and human-readable display name for the Instance.`,
			},
			"gce_zone": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The Compute Engine zone that the instance should serve from, per https://cloud.google.com/compute/docs/regions-zones This can ONLY be specified for ZONAL instances. If present for a REGIONAL instance, an error will be thrown. If this is absent for a ZONAL instance, instance is created in a random zone with available capacity.`,
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `User-defined labels for the alloydb instance.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"machine_config": {
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				Description: `Configurations for the machines that host the underlying database engine.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cpu_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Optional:    true,
							Description: `The number of CPU's in the VM instance.`,
						},
					},
				},
			},
			"read_pool_config": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Read pool specific config.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"node_count": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: `Read capacity, i.e. number of nodes in a read pool instance.`,
						},
					},
				},
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time the Instance was created in UTC.`,
			},
			"ip_address": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The IP address for the Instance. This is the connection endpoint for an end-user application.`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The name of the instance resource.`,
			},
			"reconciling": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: `Set to true if the current state of Instance does not match the user's intended state, and the service is actively updating the resource to reconcile them. This can happen due to user-triggered updates or system actions like failover or maintenance.`,
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The current state of the alloydb instance.`,
			},
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The system-generated UID of the resource.`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time the Instance was updated in UTC.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceAlloydbInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	var project string
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	labelsProp, err := expandAlloydbInstanceLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	annotationsProp, err := expandAlloydbInstanceAnnotations(d.Get("annotations"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("annotations"); !isEmptyValue(reflect.ValueOf(annotationsProp)) && (ok || !reflect.DeepEqual(v, annotationsProp)) {
		obj["annotations"] = annotationsProp
	}
	displayNameProp, err := expandAlloydbInstanceDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	gceZoneProp, err := expandAlloydbInstanceGceZone(d.Get("gce_zone"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("gce_zone"); !isEmptyValue(reflect.ValueOf(gceZoneProp)) && (ok || !reflect.DeepEqual(v, gceZoneProp)) {
		obj["gceZone"] = gceZoneProp
	}
	databaseFlagsProp, err := expandAlloydbInstanceDatabaseFlags(d.Get("database_flags"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("database_flags"); !isEmptyValue(reflect.ValueOf(databaseFlagsProp)) && (ok || !reflect.DeepEqual(v, databaseFlagsProp)) {
		obj["databaseFlags"] = databaseFlagsProp
	}
	availabilityTypeProp, err := expandAlloydbInstanceAvailabilityType(d.Get("availability_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("availability_type"); !isEmptyValue(reflect.ValueOf(availabilityTypeProp)) && (ok || !reflect.DeepEqual(v, availabilityTypeProp)) {
		obj["availabilityType"] = availabilityTypeProp
	}
	instanceTypeProp, err := expandAlloydbInstanceInstanceType(d.Get("instance_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("instance_type"); !isEmptyValue(reflect.ValueOf(instanceTypeProp)) && (ok || !reflect.DeepEqual(v, instanceTypeProp)) {
		obj["instanceType"] = instanceTypeProp
	}
	readPoolConfigProp, err := expandAlloydbInstanceReadPoolConfig(d.Get("read_pool_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("read_pool_config"); !isEmptyValue(reflect.ValueOf(readPoolConfigProp)) && (ok || !reflect.DeepEqual(v, readPoolConfigProp)) {
		obj["readPoolConfig"] = readPoolConfigProp
	}
	machineConfigProp, err := expandAlloydbInstanceMachineConfig(d.Get("machine_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("machine_config"); !isEmptyValue(reflect.ValueOf(machineConfigProp)) && (ok || !reflect.DeepEqual(v, machineConfigProp)) {
		obj["machineConfig"] = machineConfigProp
	}

	url, err := replaceVars(d, config, "{{AlloydbBasePath}}{{cluster}}/instances?instanceId={{instance_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Instance: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Instance: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{cluster}}/instances/{{instance_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = alloydbOperationWaitTime(
		config, res, project, "Creating Instance", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Instance: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Instance %q: %#v", d.Id(), res)

	return resourceAlloydbInstanceRead(d, meta)
}

func resourceAlloydbInstanceRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{AlloydbBasePath}}{{cluster}}/instances/{{instance_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("AlloydbInstance %q", d.Id()))
	}

	if err := d.Set("name", flattenAlloydbInstanceName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("create_time", flattenAlloydbInstanceCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("update_time", flattenAlloydbInstanceUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("uid", flattenAlloydbInstanceUid(res["uid"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("labels", flattenAlloydbInstanceLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("annotations", flattenAlloydbInstanceAnnotations(res["annotations"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("state", flattenAlloydbInstanceState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("gce_zone", flattenAlloydbInstanceGceZone(res["gceZone"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("reconciling", flattenAlloydbInstanceReconciling(res["reconciling"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("database_flags", flattenAlloydbInstanceDatabaseFlags(res["databaseFlags"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("availability_type", flattenAlloydbInstanceAvailabilityType(res["availabilityType"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("instance_type", flattenAlloydbInstanceInstanceType(res["instanceType"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("ip_address", flattenAlloydbInstanceIpAddress(res["ipAddress"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("read_pool_config", flattenAlloydbInstanceReadPoolConfig(res["readPoolConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("machine_config", flattenAlloydbInstanceMachineConfig(res["machineConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}

	return nil
}

func resourceAlloydbInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	var project string
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	labelsProp, err := expandAlloydbInstanceLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	annotationsProp, err := expandAlloydbInstanceAnnotations(d.Get("annotations"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("annotations"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, annotationsProp)) {
		obj["annotations"] = annotationsProp
	}
	displayNameProp, err := expandAlloydbInstanceDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	gceZoneProp, err := expandAlloydbInstanceGceZone(d.Get("gce_zone"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("gce_zone"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, gceZoneProp)) {
		obj["gceZone"] = gceZoneProp
	}
	databaseFlagsProp, err := expandAlloydbInstanceDatabaseFlags(d.Get("database_flags"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("database_flags"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, databaseFlagsProp)) {
		obj["databaseFlags"] = databaseFlagsProp
	}
	availabilityTypeProp, err := expandAlloydbInstanceAvailabilityType(d.Get("availability_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("availability_type"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, availabilityTypeProp)) {
		obj["availabilityType"] = availabilityTypeProp
	}
	readPoolConfigProp, err := expandAlloydbInstanceReadPoolConfig(d.Get("read_pool_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("read_pool_config"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, readPoolConfigProp)) {
		obj["readPoolConfig"] = readPoolConfigProp
	}
	machineConfigProp, err := expandAlloydbInstanceMachineConfig(d.Get("machine_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("machine_config"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, machineConfigProp)) {
		obj["machineConfig"] = machineConfigProp
	}

	url, err := replaceVars(d, config, "{{AlloydbBasePath}}{{cluster}}/instances/{{instance_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Instance %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("labels") {
		updateMask = append(updateMask, "labels")
	}

	if d.HasChange("annotations") {
		updateMask = append(updateMask, "annotations")
	}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("gce_zone") {
		updateMask = append(updateMask, "gceZone")
	}

	if d.HasChange("database_flags") {
		updateMask = append(updateMask, "databaseFlags")
	}

	if d.HasChange("availability_type") {
		updateMask = append(updateMask, "availabilityType")
	}

	if d.HasChange("read_pool_config") {
		updateMask = append(updateMask, "readPoolConfig")
	}

	if d.HasChange("machine_config") {
		updateMask = append(updateMask, "machineConfig")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating Instance %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Instance %q: %#v", d.Id(), res)
	}

	err = alloydbOperationWaitTime(
		config, res, project, "Updating Instance", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceAlloydbInstanceRead(d, meta)
}

func resourceAlloydbInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	var project string
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := replaceVars(d, config, "{{AlloydbBasePath}}{{cluster}}/instances/{{instance_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Instance %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Instance")
	}

	err = alloydbOperationWaitTime(
		config, res, project, "Deleting Instance", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Instance %q: %#v", d.Id(), res)
	return nil
}

func resourceAlloydbInstanceImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := parseImportId([]string{
		"(?P<cluster>.+)/instances/(?P<instance_id>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{cluster}}/instances/{{instance_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenAlloydbInstanceName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAlloydbInstanceCreateTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAlloydbInstanceUpdateTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAlloydbInstanceUid(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAlloydbInstanceLabels(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAlloydbInstanceAnnotations(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAlloydbInstanceState(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAlloydbInstanceGceZone(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAlloydbInstanceReconciling(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAlloydbInstanceDatabaseFlags(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAlloydbInstanceAvailabilityType(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAlloydbInstanceInstanceType(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAlloydbInstanceIpAddress(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAlloydbInstanceReadPoolConfig(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["node_count"] =
		flattenAlloydbInstanceReadPoolConfigNodeCount(original["nodeCount"], d, config)
	return []interface{}{transformed}
}
func flattenAlloydbInstanceReadPoolConfigNodeCount(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := stringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenAlloydbInstanceMachineConfig(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["cpu_count"] =
		flattenAlloydbInstanceMachineConfigCpuCount(original["cpuCount"], d, config)
	return []interface{}{transformed}
}
func flattenAlloydbInstanceMachineConfigCpuCount(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := stringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func expandAlloydbInstanceLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandAlloydbInstanceAnnotations(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandAlloydbInstanceDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbInstanceGceZone(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbInstanceDatabaseFlags(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandAlloydbInstanceAvailabilityType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbInstanceInstanceType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbInstanceReadPoolConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedNodeCount, err := expandAlloydbInstanceReadPoolConfigNodeCount(original["node_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNodeCount); val.IsValid() && !isEmptyValue(val) {
		transformed["nodeCount"] = transformedNodeCount
	}

	return transformed, nil
}

func expandAlloydbInstanceReadPoolConfigNodeCount(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbInstanceMachineConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCpuCount, err := expandAlloydbInstanceMachineConfigCpuCount(original["cpu_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCpuCount); val.IsValid() && !isEmptyValue(val) {
		transformed["cpuCount"] = transformedCpuCount
	}

	return transformed, nil
}

func expandAlloydbInstanceMachineConfigCpuCount(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
