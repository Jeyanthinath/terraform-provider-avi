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

func ResourceServiceEngineSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"availability_zone": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"cloud_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "/api/cloud?name=Default-Cloud"},
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
		"controller_ip": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"data_vnics": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourcevNICSchema()},
		"enable_state": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "SE_STATE_ENABLED"},
		"flavor": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"host_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"hypervisor": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"mgmt_vnic": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Set:      func(v interface{}) int { return 0 }, Elem: ResourcevNICSchema()},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "VM name unknown"},
		"resources": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Set:      func(v interface{}) int { return 0 }, Elem: ResourceSeResourcesSchema()},
		"se_group_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"tenant_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true},
	}
}

func resourceAviServiceEngine() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviServiceEngineCreate,
		Read:   ResourceAviServiceEngineRead,
		Update: resourceAviServiceEngineUpdate,
		Delete: resourceAviServiceEngineDelete,
		Schema: ResourceServiceEngineSchema(),
	}
}

func ResourceAviServiceEngineRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceServiceEngineSchema()
	log.Printf("[INFO] ResourceAviServiceEngineRead Avi Client %v\n", d)
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/serviceengine/" + uuid.(string)
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
	log.Printf("ResourceAviServiceEngineRead CURRENT obj %v\n", d)

	log.Printf("ResourceAviServiceEngineRead Read API obj %v\n", obj)
	if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
		log.Printf("[INFO] ResourceAviServiceEngineRead Converted obj %v\n", tObj)
		//err = d.Set("obj", tObj)
		if err != nil {
			log.Printf("[ERROR] in setting read object %v\n", err)
		}
	}
	log.Printf("[INFO] ResourceAviServiceEngineRead Updated %v\n", d)
	return nil
}

func resourceAviServiceEngineCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceServiceEngineSchema()
	err := ApiCreateOrUpdate(d, meta, "serviceengine", s)
	log.Printf("[DEBUG] created object %v: %v", "serviceengine", d)
	if err == nil {
		err = ResourceAviServiceEngineRead(d, meta)
	}
	log.Printf("[DEBUG] created object %v: %v", "serviceengine", d)
	return err
}

func resourceAviServiceEngineUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceServiceEngineSchema()
	err := ApiCreateOrUpdate(d, meta, "serviceengine", s)
	if err == nil {
		err = ResourceAviServiceEngineRead(d, meta)
	}
	log.Printf("[DEBUG] updated object %v: %v", "serviceengine", d)
	return err
}

func resourceAviServiceEngineDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "serviceengine"
	log.Println("[INFO] ResourceAviServiceEngineRead Avi Client")
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil {
			log.Println("[INFO] resourceAviServiceEngineDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}