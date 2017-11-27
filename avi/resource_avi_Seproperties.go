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
func ResourceSePropertiesSchema() map[string]*schema.Schema {
    return map[string]*schema.Schema{
             "se_agent_properties" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeAgentPropertiesSchema(),                             },
             "se_bootup_properties" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeBootupPropertiesSchema(),                             },
             "se_runtime_properties" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeRuntimePropertiesSchema(),                             },
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


func resourceAviSeProperties() *schema.Resource {
    return &schema.Resource{
        Create: resourceAviSePropertiesCreate,
        Read:   ResourceAviSePropertiesRead,
        Update: resourceAviSePropertiesUpdate,
        Delete: resourceAviSePropertiesDelete,
        Schema: ResourceSePropertiesSchema(),
    }
}


func ResourceAviSePropertiesRead(d *schema.ResourceData, meta interface{}) error {
    s := ResourceSePropertiesSchema()
    log.Printf("[INFO] ResourceAviSePropertiesRead Avi Client %v\n", d)
    client := meta.(*clients.AviClient)
    var obj interface{}
    if uuid, ok := d.GetOk("uuid"); ok {
        path := "api/seproperties/" + uuid.(string)
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
    log.Printf("ResourceAviSePropertiesRead CURRENT obj %v\n", d)

    log.Printf("ResourceAviSePropertiesRead Read API obj %v\n", obj)
    if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
        log.Printf("[INFO] ResourceAviSePropertiesRead Converted obj %v\n", tObj)
        //err = d.Set("obj", tObj)
        if err != nil {
            log.Printf("[ERROR] in setting read object %v\n", err)
        }
    }
    log.Printf("[INFO] ResourceAviSePropertiesRead Updated %v\n", d)
    return nil
}

func resourceAviSePropertiesCreate(d *schema.ResourceData, meta interface{}) error {
    s := ResourceVrfContextSchema()
    err := ApiCreateOrUpdate(d, meta, "seproperties", s)
    log.Printf("[DEBUG] created object %v: %v", "seproperties", d)
    if err == nil {
        err = ResourceAviVrfContextRead(d, meta)
    }
    log.Printf("[DEBUG] created object %v: %v", "seproperties", d)
    return err
}

func resourceAviSePropertiesUpdate(d *schema.ResourceData, meta interface{}) error {
    s := ResourcePoolSchema()
    err := ApiCreateOrUpdate(d, meta, "seproperties", s)
    if err == nil {
        err = ResourceAviVrfContextRead(d, meta)
    }
    log.Printf("[DEBUG] updated object %v: %v", "seproperties", d)
    return err
}

func resourceAviSePropertiesDelete(d *schema.ResourceData, meta interface{}) error {
    s := ResourcePoolSchema()
    err := ApiCreateOrUpdate(d, meta, "seproperties", s)
    log.Printf("[DEBUG] created object %v: %v", "seproperties", d)
    if err == nil {
        err = ResourceAviVrfContextRead(d, meta)
    }
    log.Printf("[DEBUG] created object %v: %v", "seproperties", d)
    return err
}


