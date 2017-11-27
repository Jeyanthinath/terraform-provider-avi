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
func ResourceHardwareSecurityModuleGroupSchema() map[string]*schema.Schema {
    return map[string]*schema.Schema{
             "hsm" :&schema.Schema{
                             Type: schema.TypeSet,            
                              Required: true,                              
                             Elem: ResourceHardwareSecurityModuleSchema(),                             },
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


func resourceAviHardwareSecurityModuleGroup() *schema.Resource {
    return &schema.Resource{
        Create: resourceAviHardwareSecurityModuleGroupCreate,
        Read:   ResourceAviHardwareSecurityModuleGroupRead,
        Update: resourceAviHardwareSecurityModuleGroupUpdate,
        Delete: resourceAviHardwareSecurityModuleGroupDelete,
        Schema: ResourceHardwareSecurityModuleGroupSchema(),
    }
}


func ResourceAviHardwareSecurityModuleGroupRead(d *schema.ResourceData, meta interface{}) error {
    s := ResourceHardwareSecurityModuleGroupSchema()
    log.Printf("[INFO] ResourceAviHardwareSecurityModuleGroupRead Avi Client %v\n", d)
    client := meta.(*clients.AviClient)
    var obj interface{}
    if uuid, ok := d.GetOk("uuid"); ok {
        path := "api/hardwaresecuritymodulegroup/" + uuid.(string)
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
    log.Printf("ResourceAviHardwareSecurityModuleGroupRead CURRENT obj %v\n", d)

    log.Printf("ResourceAviHardwareSecurityModuleGroupRead Read API obj %v\n", obj)
    if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
        log.Printf("[INFO] ResourceAviHardwareSecurityModuleGroupRead Converted obj %v\n", tObj)
        //err = d.Set("obj", tObj)
        if err != nil {
            log.Printf("[ERROR] in setting read object %v\n", err)
        }
    }
    log.Printf("[INFO] ResourceAviHardwareSecurityModuleGroupRead Updated %v\n", d)
    return nil
}

func resourceAviHardwareSecurityModuleGroupCreate(d *schema.ResourceData, meta interface{}) error {
    s := ResourceVrfContextSchema()
    err := ApiCreateOrUpdate(d, meta, "hardwaresecuritymodulegroup", s)
    log.Printf("[DEBUG] created object %v: %v", "hardwaresecuritymodulegroup", d)
    if err == nil {
        err = ResourceAviVrfContextRead(d, meta)
    }
    log.Printf("[DEBUG] created object %v: %v", "hardwaresecuritymodulegroup", d)
    return err
}

func resourceAviHardwareSecurityModuleGroupUpdate(d *schema.ResourceData, meta interface{}) error {
    s := ResourcePoolSchema()
    err := ApiCreateOrUpdate(d, meta, "hardwaresecuritymodulegroup", s)
    if err == nil {
        err = ResourceAviVrfContextRead(d, meta)
    }
    log.Printf("[DEBUG] updated object %v: %v", "hardwaresecuritymodulegroup", d)
    return err
}

func resourceAviHardwareSecurityModuleGroupDelete(d *schema.ResourceData, meta interface{}) error {
    s := ResourcePoolSchema()
    err := ApiCreateOrUpdate(d, meta, "hardwaresecuritymodulegroup", s)
    log.Printf("[DEBUG] created object %v: %v", "hardwaresecuritymodulegroup", d)
    if err == nil {
        err = ResourceAviVrfContextRead(d, meta)
    }
    log.Printf("[DEBUG] created object %v: %v", "hardwaresecuritymodulegroup", d)
    return err
}


