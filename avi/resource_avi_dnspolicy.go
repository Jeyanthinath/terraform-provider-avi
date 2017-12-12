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

func ResourceDnsPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"created_by": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"description": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"rule": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceDnsRuleSchema()},
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

func resourceAviDnsPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviDnsPolicyCreate,
		Read:   ResourceAviDnsPolicyRead,
		Update: resourceAviDnsPolicyUpdate,
		Delete: resourceAviDnsPolicyDelete,
		Schema: ResourceDnsPolicySchema(),
	}
}

func ResourceAviDnsPolicyRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceDnsPolicySchema()
	log.Printf("[INFO] ResourceAviDnsPolicyRead Avi Client %v\n", d)
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/dnspolicy/" + uuid.(string)
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
	log.Printf("ResourceAviDnsPolicyRead CURRENT obj %v\n", d)

	log.Printf("ResourceAviDnsPolicyRead Read API obj %v\n", obj)
	if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
		log.Printf("[INFO] ResourceAviDnsPolicyRead Converted obj %v\n", tObj)
		//err = d.Set("obj", tObj)
		if err != nil {
			log.Printf("[ERROR] in setting read object %v\n", err)
		}
	}
	log.Printf("[INFO] ResourceAviDnsPolicyRead Updated %v\n", d)
	return nil
}

func resourceAviDnsPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceDnsPolicySchema()
	err := ApiCreateOrUpdate(d, meta, "dnspolicy", s)
	log.Printf("[DEBUG] created object %v: %v", "dnspolicy", d)
	if err == nil {
		err = ResourceAviDnsPolicyRead(d, meta)
	}
	log.Printf("[DEBUG] created object %v: %v", "dnspolicy", d)
	return err
}

func resourceAviDnsPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceDnsPolicySchema()
	err := ApiCreateOrUpdate(d, meta, "dnspolicy", s)
	if err == nil {
		err = ResourceAviDnsPolicyRead(d, meta)
	}
	log.Printf("[DEBUG] updated object %v: %v", "dnspolicy", d)
	return err
}

func resourceAviDnsPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "dnspolicy"
	log.Println("[INFO] ResourceAviDnsPolicyRead Avi Client")
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil {
			log.Println("[INFO] resourceAviDnsPolicyDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}