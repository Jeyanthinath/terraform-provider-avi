/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"log"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceCloudSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"apic_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Set:      func(v interface{}) int { return 0 }, Elem: ResourceAPICConfigurationSchema()},
		"apic_mode": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"aws_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Set:      func(v interface{}) int { return 0 }, Elem: ResourceAwsConfigurationSchema()},
		"azure_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Set:      func(v interface{}) int { return 0 }, Elem: ResourceAzureConfigurationSchema()},
		"cloudstack_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Set:      func(v interface{}) int { return 0 }, Elem: ResourceCloudStackConfigurationSchema()},
		"custom_tags": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceCustomTagSchema()},
		"dhcp_enabled": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"dns_provider_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"docker_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Set:      func(v interface{}) int { return 0 }, Elem: ResourceDockerConfigurationSchema()},
		"east_west_dns_provider_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"east_west_ipam_provider_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"enable_vip_static_routes": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"ipam_provider_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"license_type": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"linuxserver_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Set:      func(v interface{}) int { return 0 }, Elem: ResourceLinuxServerConfigurationSchema()},
		"mesos_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Set:      func(v interface{}) int { return 0 }, Elem: ResourceMesosConfigurationSchema()},
		"mtu": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1500,
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true},
		"nsx_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Set:      func(v interface{}) int { return 0 }, Elem: ResourceNsxConfigurationSchema()},
		"obj_name_prefix": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"openstack_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Set:      func(v interface{}) int { return 0 }, Elem: ResourceOpenStackConfigurationSchema()},
		"oshiftk8s_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Set:      func(v interface{}) int { return 0 }, Elem: ResourceOShiftK8SConfigurationSchema()},
		"prefer_static_routes": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"proxy_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Set:      func(v interface{}) int { return 0 }, Elem: ResourceProxyConfigurationSchema()},
		"rancher_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Set:      func(v interface{}) int { return 0 }, Elem: ResourceRancherConfigurationSchema()},
		"state_based_dns_registration": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"tenant_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true},
		"vca_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Set:      func(v interface{}) int { return 0 }, Elem: ResourcevCloudAirConfigurationSchema()},
		"vcenter_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Set:      func(v interface{}) int { return 0 }, Elem: ResourcevCenterConfigurationSchema()},
		"vtype": &schema.Schema{
			Type:     schema.TypeString,
			Required: true},
	}
}

func resourceAviCloud() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviCloudCreate,
		Read:   ResourceAviCloudRead,
		Update: resourceAviCloudUpdate,
		Delete: resourceAviCloudDelete,
		Schema: ResourceCloudSchema(),
	}
}

func ResourceAviCloudRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCloudSchema()
	log.Printf("[INFO] ResourceAviCloudRead Avi Client %v\n", d)
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/cloud/" + uuid.(string)
		err := client.AviSession.Get(path, &obj)
		if err != nil {
			d.SetId("")
			return nil
		}
	} else {
		d.SetId("")
		return nil
	}
	// no need to set the ID
	log.Printf("ResourceAviCloudRead CURRENT obj %v\n", d)

	log.Printf("ResourceAviCloudRead Read API obj %v\n", obj)
	if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
		log.Printf("[INFO] ResourceAviCloudRead Converted obj %v\n", tObj)
		//err = d.Set("obj", tObj)
		if err != nil {
			log.Printf("[ERROR] in setting read object %v\n", err)
		}
	}
	log.Printf("[INFO] ResourceAviCloudRead Updated %v\n", d)
	return nil
}

func resourceAviCloudCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCloudSchema()
	err := ApiCreateOrUpdate(d, meta, "cloud", s)
	log.Printf("[DEBUG] created object %v: %v", "cloud", d)
	if err == nil {
		err = ResourceAviCloudRead(d, meta)
	}
	log.Printf("[DEBUG] created object %v: %v", "cloud", d)
	return err
}

func resourceAviCloudUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCloudSchema()
	err := ApiCreateOrUpdate(d, meta, "cloud", s)
	if err == nil {
		err = ResourceAviCloudRead(d, meta)
	}
	log.Printf("[DEBUG] updated object %v: %v", "cloud", d)
	return err
}

func resourceAviCloudDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "cloud"
	log.Println("[INFO] ResourceAviCloudRead Avi Client")
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil {
			log.Println("[INFO] resourceAviCloudDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
