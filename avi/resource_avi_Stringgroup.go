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
func ResourceStringGroupSchema() map[string]*schema.Schema {
    return map[string]*schema.Schema{
             "description" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "kv" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceKeyValueSchema(),                             },
             "name" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "tenant_ref" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                              Elem:&schema.Schema{Type: schema.TypeString},                             },
             "type" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
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


func resourceAviStringGroup() *schema.Resource {
    return &schema.Resource{
        Create: resourceAviStringGroupCreate,
        Read:   ResourceAviStringGroupRead,
        Update: resourceAviStringGroupUpdate,
        Delete: resourceAviStringGroupDelete,
        Schema: ResourceStringGroupSchema(),
    }
}


func ResourceAviStringGroupRead(d *schema.ResourceData, meta interface{}) error {
    s := ResourceStringGroupSchema()
    log.Printf("[INFO] ResourceAviStringGroupRead Avi Client %v\n", d)
    client := meta.(*clients.AviClient)
    var obj interface{}
    if uuid, ok := d.GetOk("uuid"); ok {
        path := "api/stringgroup/" + uuid.(string)
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
    log.Printf("ResourceAviStringGroupRead CURRENT obj %v\n", d)

    log.Printf("ResourceAviStringGroupRead Read API obj %v\n", obj)
    if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
        log.Printf("[INFO] ResourceAviStringGroupRead Converted obj %v\n", tObj)
        //err = d.Set("obj", tObj)
        if err != nil {
            log.Printf("[ERROR] in setting read object %v\n", err)
        }
    }
    log.Printf("[INFO] ResourceAviStringGroupRead Updated %v\n", d)
    return nil
}

func resourceAviStringGroupCreate(d *schema.ResourceData, meta interface{}) error {
    s := ResourceVrfContextSchema()
    err := ApiCreateOrUpdate(d, meta, "stringgroup", s)
    log.Printf("[DEBUG] created object %v: %v", "stringgroup", d)
    if err == nil {
        err = ResourceAviVrfContextRead(d, meta)
    }
    log.Printf("[DEBUG] created object %v: %v", "stringgroup", d)
    return err
}

func resourceAviStringGroupUpdate(d *schema.ResourceData, meta interface{}) error {
    s := ResourcePoolSchema()
    err := ApiCreateOrUpdate(d, meta, "stringgroup", s)
    if err == nil {
        err = ResourceAviVrfContextRead(d, meta)
    }
    log.Printf("[DEBUG] updated object %v: %v", "stringgroup", d)
    return err
}

func resourceAviStringGroupDelete(d *schema.ResourceData, meta interface{}) error {
    s := ResourcePoolSchema()
    err := ApiCreateOrUpdate(d, meta, "stringgroup", s)
    log.Printf("[DEBUG] created object %v: %v", "stringgroup", d)
    if err == nil {
        err = ResourceAviVrfContextRead(d, meta)
    }
    log.Printf("[DEBUG] created object %v: %v", "stringgroup", d)
    return err
}


