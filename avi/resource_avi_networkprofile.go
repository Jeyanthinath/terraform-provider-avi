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
func ResourceNetworkProfileSchema() map[string]*schema.Schema {
    return map[string]*schema.Schema{
             "description" :&schema.Schema{
                             Type: schema.TypeString, 
                             Optional: true,
                                                                                                                                            },
             "name" :&schema.Schema{
                             Type: schema.TypeString, 
                             Required: true,
                                                                                                                },
             "profile" :&schema.Schema{
                             Type: schema.TypeSet, 
                             Required: true,
                                                                                     Elem: ResourceNetworkProfileUnionSchema(),                             },
             "tenant_ref" :&schema.Schema{
                             Type: schema.TypeString, 
                             Optional: true,
                                                                                                                  Elem:&schema.Schema{Type: schema.TypeString},                             },
             "uuid" :&schema.Schema{
                             Type: schema.TypeString, 
                             Optional: true,
                                                                                    Computed:  true,                                                         },
                                "url": &schema.Schema{
                                Type:     schema.TypeString,
                                Optional: true,
                                Computed: true,
                            },
    }
}


func resourceAviNetworkProfile() *schema.Resource {
    return &schema.Resource{
        Create: resourceAviNetworkProfileCreate,
        Read:   ResourceAviNetworkProfileRead,
        Update: resourceAviNetworkProfileUpdate,
        Delete: resourceAviNetworkProfileDelete,
        Schema: ResourceNetworkProfileSchema(),
    }
}


func ResourceAviNetworkProfileRead(d *schema.ResourceData, meta interface{}) error {
    s := ResourceNetworkProfileSchema()
    log.Printf("[INFO] ResourceAviNetworkProfileRead Avi Client %v\n", d)
    client := meta.(*clients.AviClient)
    var obj interface{}
    if uuid, ok := d.GetOk("uuid"); ok {
        path := "api/networkprofile/" + uuid.(string)
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
    log.Printf("ResourceAviNetworkProfileRead CURRENT obj %v\n", d)

    log.Printf("ResourceAviNetworkProfileRead Read API obj %v\n", obj)
    if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
        log.Printf("[INFO] ResourceAviNetworkProfileRead Converted obj %v\n", tObj)
        //err = d.Set("obj", tObj)
        if err != nil {
            log.Printf("[ERROR] in setting read object %v\n", err)
        }
    }
    log.Printf("[INFO] ResourceAviNetworkProfileRead Updated %v\n", d)
    return nil
}

func resourceAviNetworkProfileCreate(d *schema.ResourceData, meta interface{}) error {
    s := ResourceNetworkProfileSchema()
    err := ApiCreateOrUpdate(d, meta, "networkprofile", s)
    log.Printf("[DEBUG] created object %v: %v", "networkprofile", d)
    if err == nil {
        err = ResourceAviNetworkProfileRead(d, meta)
    }
    log.Printf("[DEBUG] created object %v: %v", "networkprofile", d)
    return err
}

func resourceAviNetworkProfileUpdate(d *schema.ResourceData, meta interface{}) error {
    s := ResourceNetworkProfileSchema()
    err := ApiCreateOrUpdate(d, meta, "networkprofile", s)
    if err == nil {
        err = ResourceAviNetworkProfileRead(d, meta)
    }
    log.Printf("[DEBUG] updated object %v: %v", "networkprofile", d)
    return err
}

func resourceAviNetworkProfileDelete(d *schema.ResourceData, meta interface{}) error {
    s := ResourceNetworkProfileSchema()
    err := ApiCreateOrUpdate(d, meta, "networkprofile", s)
    log.Printf("[DEBUG] created object %v: %v", "networkprofile", d)
    if err == nil {
        err = ResourceAviNetworkProfileRead(d, meta)
    }
    log.Printf("[DEBUG] created object %v: %v", "networkprofile", d)
    return err
}


