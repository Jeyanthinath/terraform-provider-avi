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
func ResourceWafProfileSchema() map[string]*schema.Schema {
    return map[string]*schema.Schema{
             "config" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceWafConfigSchema(),                             },
             "description" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "files" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceWafDataFileSchema(),                             },
             "name" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "tenant_ref" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                              Elem:&schema.Schema{Type: schema.TypeString},                             },
             "uuid" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                             Computed:  true,  
                                                        },
                                "url": &schema.Schema{
                                Type:     schema.TypeString,
                                Optional: true,
                                Computed: true,
                            },
    }
}


func resourceAviWafProfile() *schema.Resource {
    return &schema.Resource{
        Create: resourceAviWafProfileCreate,
        Read:   ResourceAviWafProfileRead,
        Update: resourceAviWafProfileUpdate,
        Delete: resourceAviWafProfileDelete,
        Schema: ResourceWafProfileSchema(),
    }
}


func ResourceAviWafProfileRead(d *schema.ResourceData, meta interface{}) error {
    s := ResourceWafProfileSchema()
    log.Printf("[INFO] ResourceAviWafProfileRead Avi Client %v\n", d)
    client := meta.(*clients.AviClient)
    var obj interface{}
    if uuid, ok := d.GetOk("uuid"); ok {
        path := "api/wafprofile/" + uuid.(string)
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
    log.Printf("ResourceAviWafProfileRead CURRENT obj %v\n", d)

    log.Printf("ResourceAviWafProfileRead Read API obj %v\n", obj)
    if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
        log.Printf("[INFO] ResourceAviWafProfileRead Converted obj %v\n", tObj)
        //err = d.Set("obj", tObj)
        if err != nil {
            log.Printf("[ERROR] in setting read object %v\n", err)
        }
    }
    log.Printf("[INFO] ResourceAviWafProfileRead Updated %v\n", d)
    return nil
}

func resourceAviWafProfileCreate(d *schema.ResourceData, meta interface{}) error {
    s := ResourceVrfContextSchema()
    err := ApiCreateOrUpdate(d, meta, "wafprofile", s)
    log.Printf("[DEBUG] created object %v: %v", "wafprofile", d)
    if err == nil {
        err = ResourceAviVrfContextRead(d, meta)
    }
    log.Printf("[DEBUG] created object %v: %v", "wafprofile", d)
    return err
}

func resourceAviWafProfileUpdate(d *schema.ResourceData, meta interface{}) error {
    s := ResourcePoolSchema()
    err := ApiCreateOrUpdate(d, meta, "wafprofile", s)
    if err == nil {
        err = ResourceAviVrfContextRead(d, meta)
    }
    log.Printf("[DEBUG] updated object %v: %v", "wafprofile", d)
    return err
}

func resourceAviWafProfileDelete(d *schema.ResourceData, meta interface{}) error {
    s := ResourcePoolSchema()
    err := ApiCreateOrUpdate(d, meta, "wafprofile", s)
    log.Printf("[DEBUG] created object %v: %v", "wafprofile", d)
    if err == nil {
        err = ResourceAviVrfContextRead(d, meta)
    }
    log.Printf("[DEBUG] created object %v: %v", "wafprofile", d)
    return err
}


