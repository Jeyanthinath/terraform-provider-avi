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

func ResourceTrafficCloneProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"clone_servers": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceCloneServerSchema()},
		"cloud_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "/api/cloud?name=Default-Cloud"},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true},
		"preserve_client_ip": &schema.Schema{
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

func resourceAviTrafficCloneProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviTrafficCloneProfileCreate,
		Read:   ResourceAviTrafficCloneProfileRead,
		Update: resourceAviTrafficCloneProfileUpdate,
		Delete: resourceAviTrafficCloneProfileDelete,
		Schema: ResourceTrafficCloneProfileSchema(),
	}
}

func ResourceAviTrafficCloneProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTrafficCloneProfileSchema()
	log.Printf("[INFO] ResourceAviTrafficCloneProfileRead Avi Client %v\n", d)
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/trafficcloneprofile/" + uuid.(string)
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
	log.Printf("ResourceAviTrafficCloneProfileRead CURRENT obj %v\n", d)

	log.Printf("ResourceAviTrafficCloneProfileRead Read API obj %v\n", obj)
	if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
		log.Printf("[INFO] ResourceAviTrafficCloneProfileRead Converted obj %v\n", tObj)
		//err = d.Set("obj", tObj)
		if err != nil {
			log.Printf("[ERROR] in setting read object %v\n", err)
		}
	}
	log.Printf("[INFO] ResourceAviTrafficCloneProfileRead Updated %v\n", d)
	return nil
}

func resourceAviTrafficCloneProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTrafficCloneProfileSchema()
	err := ApiCreateOrUpdate(d, meta, "trafficcloneprofile", s)
	log.Printf("[DEBUG] created object %v: %v", "trafficcloneprofile", d)
	if err == nil {
		err = ResourceAviTrafficCloneProfileRead(d, meta)
	}
	log.Printf("[DEBUG] created object %v: %v", "trafficcloneprofile", d)
	return err
}

func resourceAviTrafficCloneProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTrafficCloneProfileSchema()
	err := ApiCreateOrUpdate(d, meta, "trafficcloneprofile", s)
	if err == nil {
		err = ResourceAviTrafficCloneProfileRead(d, meta)
	}
	log.Printf("[DEBUG] updated object %v: %v", "trafficcloneprofile", d)
	return err
}

func resourceAviTrafficCloneProfileDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "trafficcloneprofile"
	log.Println("[INFO] ResourceAviTrafficCloneProfileRead Avi Client")
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil {
			log.Println("[INFO] resourceAviTrafficCloneProfileDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}