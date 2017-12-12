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

func ResourceVrfContextSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"bgp_profile": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Set:      func(v interface{}) int { return 0 }, Elem: ResourceBgpProfileSchema()},
		"cloud_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "/api/cloud?name=Default-Cloud"},
		"debugvrfcontext": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Set:      func(v interface{}) int { return 0 }, Elem: ResourceDebugVrfContextSchema()},
		"description": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"gateway_mon": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceGatewayMonitorSchema()},
		"internal_gateway_monitor": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Set:      func(v interface{}) int { return 0 }, Elem: ResourceInternalGatewayMonitorSchema()},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true},
		"static_routes": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceStaticRouteSchema()},
		"system_default": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
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

func resourceAviVrfContext() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviVrfContextCreate,
		Read:   ResourceAviVrfContextRead,
		Update: resourceAviVrfContextUpdate,
		Delete: resourceAviVrfContextDelete,
		Schema: ResourceVrfContextSchema(),
	}
}

func ResourceAviVrfContextRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceVrfContextSchema()
	log.Printf("[INFO] ResourceAviVrfContextRead Avi Client %v\n", d)
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/vrfcontext/" + uuid.(string)
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
	log.Printf("ResourceAviVrfContextRead CURRENT obj %v\n", d)

	log.Printf("ResourceAviVrfContextRead Read API obj %v\n", obj)
	if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
		log.Printf("[INFO] ResourceAviVrfContextRead Converted obj %v\n", tObj)
		//err = d.Set("obj", tObj)
		if err != nil {
			log.Printf("[ERROR] in setting read object %v\n", err)
		}
	}
	log.Printf("[INFO] ResourceAviVrfContextRead Updated %v\n", d)
	return nil
}

func resourceAviVrfContextCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceVrfContextSchema()
	err := ApiCreateOrUpdate(d, meta, "vrfcontext", s)
	log.Printf("[DEBUG] created object %v: %v", "vrfcontext", d)
	if err == nil {
		err = ResourceAviVrfContextRead(d, meta)
	}
	log.Printf("[DEBUG] created object %v: %v", "vrfcontext", d)
	return err
}

func resourceAviVrfContextUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceVrfContextSchema()
	err := ApiCreateOrUpdate(d, meta, "vrfcontext", s)
	if err == nil {
		err = ResourceAviVrfContextRead(d, meta)
	}
	log.Printf("[DEBUG] updated object %v: %v", "vrfcontext", d)
	return err
}

func resourceAviVrfContextDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "vrfcontext"
	log.Println("[INFO] ResourceAviVrfContextRead Avi Client")
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil {
			log.Println("[INFO] resourceAviVrfContextDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}