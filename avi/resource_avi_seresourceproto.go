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

func ResourceSeResourceProtoSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"active_tags": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString}},
		"at_curr_ver": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"az": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"azure_info": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Set:      func(v interface{}) int { return 0 }, Elem: ResourceAzureInfoSchema()},
		"bootup_consumers": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString}},
		"bootup_failed": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"came_online_at": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		"container_mode": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"container_type": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "CONTAINER_TYPE_HOST"},
		"controller_created": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"create_cookie": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"created_at": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		"del_pending": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"del_start_ticks": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		"disconnected": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"disk_gb": &schema.Schema{
			Type:     schema.TypeInt,
			Required: true},
		"enable_state": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "SE_STATE_ENABLED"},
		"flavor": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"gateway_up": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"host_uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"hypervisor": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"in_use": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"ip_mac_addr": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString}},
		"ip_masks": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeInt}},
		"ips": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceIpAddrSchema()},
		"last_reboot_ticks": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		"masked_ips": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceIpAddrSchema()},
		"max_ips_per_vnic": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		"max_vnics": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		"memory": &schema.Schema{
			Type:     schema.TypeInt,
			Required: true},
		"memory_free": &schema.Schema{
			Type:     schema.TypeInt,
			Required: true},
		"memory_inuse": &schema.Schema{
			Type:     schema.TypeInt,
			Required: true},
		"mgmt_net_uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"next_vnic_gc_ticks": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		"next_vnic_op_ticks": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		"online": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"pending_consumers": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourcePendingConsumerSchema()},
		"reason_code": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "SE_DEREG_POWERED_OFF"},
		"reboot_attempts": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		"ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"resources_consumed": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceSeResourceConsumedProtoSchema()},
		"se_group_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"static_ip": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Set:      func(v interface{}) int { return 0 }, Elem: ResourceIpAddrSchema()},
		"unused_since": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		"upgrading": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"vcpus": &schema.Schema{
			Type:     schema.TypeInt,
			Required: true},
		"vcpus_free": &schema.Schema{
			Type:     schema.TypeInt,
			Required: true},
		"vcpus_inuse": &schema.Schema{
			Type:     schema.TypeInt,
			Required: true},
		"version": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "0.0.0"},
		"vinfra_discovered": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"vnic_op": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Set:      func(v interface{}) int { return 0 }, Elem: ResourceSeVnicOpProtoSchema()},
		"vnic_op_consumers": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString}},
		"vnics": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceSeResourceVnicSchema()},
		"warm_starting": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"went_headless_at": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		"went_offline_at": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
	}
}

func resourceAviSeResourceProto() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviSeResourceProtoCreate,
		Read:   ResourceAviSeResourceProtoRead,
		Update: resourceAviSeResourceProtoUpdate,
		Delete: resourceAviSeResourceProtoDelete,
		Schema: ResourceSeResourceProtoSchema(),
	}
}

func ResourceAviSeResourceProtoRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSeResourceProtoSchema()
	log.Printf("[INFO] ResourceAviSeResourceProtoRead Avi Client %v\n", d)
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/seresourceproto/" + uuid.(string)
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
	log.Printf("ResourceAviSeResourceProtoRead CURRENT obj %v\n", d)

	log.Printf("ResourceAviSeResourceProtoRead Read API obj %v\n", obj)
	if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
		log.Printf("[INFO] ResourceAviSeResourceProtoRead Converted obj %v\n", tObj)
		//err = d.Set("obj", tObj)
		if err != nil {
			log.Printf("[ERROR] in setting read object %v\n", err)
		}
	}
	log.Printf("[INFO] ResourceAviSeResourceProtoRead Updated %v\n", d)
	return nil
}

func resourceAviSeResourceProtoCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSeResourceProtoSchema()
	err := ApiCreateOrUpdate(d, meta, "seresourceproto", s)
	log.Printf("[DEBUG] created object %v: %v", "seresourceproto", d)
	if err == nil {
		err = ResourceAviSeResourceProtoRead(d, meta)
	}
	log.Printf("[DEBUG] created object %v: %v", "seresourceproto", d)
	return err
}

func resourceAviSeResourceProtoUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSeResourceProtoSchema()
	err := ApiCreateOrUpdate(d, meta, "seresourceproto", s)
	if err == nil {
		err = ResourceAviSeResourceProtoRead(d, meta)
	}
	log.Printf("[DEBUG] updated object %v: %v", "seresourceproto", d)
	return err
}

func resourceAviSeResourceProtoDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "seresourceproto"
	log.Println("[INFO] ResourceAviSeResourceProtoRead Avi Client")
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil {
			log.Println("[INFO] resourceAviSeResourceProtoDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
