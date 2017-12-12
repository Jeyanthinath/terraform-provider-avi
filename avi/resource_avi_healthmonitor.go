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

func ResourceHealthMonitorSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"dns_monitor": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Set:      func(v interface{}) int { return 0 }, Elem: ResourceHealthMonitorDNSSchema()},
		"external_monitor": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Set:      func(v interface{}) int { return 0 }, Elem: ResourceHealthMonitorExternalSchema()},
		"failed_checks": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  2,
		},
		"http_monitor": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Set:      func(v interface{}) int { return 0 }, Elem: ResourceHealthMonitorHttpSchema()},
		"https_monitor": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Set:      func(v interface{}) int { return 0 }, Elem: ResourceHealthMonitorHttpSchema()},
		"is_federated": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"monitor_port": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true},
		"receive_timeout": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  4,
		},
		"send_interval": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  10,
		},
		"successful_checks": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  2,
		},
		"tcp_monitor": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Set:      func(v interface{}) int { return 0 }, Elem: ResourceHealthMonitorTcpSchema()},
		"tenant_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"type": &schema.Schema{
			Type:     schema.TypeString,
			Required: true},
		"udp_monitor": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Set:      func(v interface{}) int { return 0 }, Elem: ResourceHealthMonitorUdpSchema()},
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true},
	}
}

func resourceAviHealthMonitor() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviHealthMonitorCreate,
		Read:   ResourceAviHealthMonitorRead,
		Update: resourceAviHealthMonitorUpdate,
		Delete: resourceAviHealthMonitorDelete,
		Schema: ResourceHealthMonitorSchema(),
	}
}

func ResourceAviHealthMonitorRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceHealthMonitorSchema()
	log.Printf("[INFO] ResourceAviHealthMonitorRead Avi Client %v\n", d)
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/healthmonitor/" + uuid.(string)
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
	log.Printf("ResourceAviHealthMonitorRead CURRENT obj %v\n", d)

	log.Printf("ResourceAviHealthMonitorRead Read API obj %v\n", obj)
	if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
		log.Printf("[INFO] ResourceAviHealthMonitorRead Converted obj %v\n", tObj)
		//err = d.Set("obj", tObj)
		if err != nil {
			log.Printf("[ERROR] in setting read object %v\n", err)
		}
	}
	log.Printf("[INFO] ResourceAviHealthMonitorRead Updated %v\n", d)
	return nil
}

func resourceAviHealthMonitorCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceHealthMonitorSchema()
	err := ApiCreateOrUpdate(d, meta, "healthmonitor", s)
	log.Printf("[DEBUG] created object %v: %v", "healthmonitor", d)
	if err == nil {
		err = ResourceAviHealthMonitorRead(d, meta)
	}
	log.Printf("[DEBUG] created object %v: %v", "healthmonitor", d)
	return err
}

func resourceAviHealthMonitorUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceHealthMonitorSchema()
	err := ApiCreateOrUpdate(d, meta, "healthmonitor", s)
	if err == nil {
		err = ResourceAviHealthMonitorRead(d, meta)
	}
	log.Printf("[DEBUG] updated object %v: %v", "healthmonitor", d)
	return err
}

func resourceAviHealthMonitorDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "healthmonitor"
	log.Println("[INFO] ResourceAviHealthMonitorRead Avi Client")
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil {
			log.Println("[INFO] resourceAviHealthMonitorDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
