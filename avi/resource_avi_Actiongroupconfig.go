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
func ResourceActionGroupConfigSchema() map[string]*schema.Schema {
    return map[string]*schema.Schema{
             "action_script_config_ref" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                              Elem:&schema.Schema{Type: schema.TypeString},                             },
             "autoscale_trigger_notification" :&schema.Schema{
                             Type: schema.TypeBool,            
                             Optional: true,                              
                                                        },
             "description" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "email_config_ref" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                              Elem:&schema.Schema{Type: schema.TypeString},                             },
             "external_only" :&schema.Schema{
                             Type: schema.TypeBool,            
                              Required: true,                              
                                                        },
             "level" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "name" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "snmp_trap_profile_ref" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                              Elem:&schema.Schema{Type: schema.TypeString},                             },
             "syslog_config_ref" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                              Elem:&schema.Schema{Type: schema.TypeString},                             },
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


func resourceAviActionGroupConfig() *schema.Resource {
    return &schema.Resource{
        Create: resourceAviActionGroupConfigCreate,
        Read:   ResourceAviActionGroupConfigRead,
        Update: resourceAviActionGroupConfigUpdate,
        Delete: resourceAviActionGroupConfigDelete,
        Schema: ResourceActionGroupConfigSchema(),
    }
}


func ResourceAviActionGroupConfigRead(d *schema.ResourceData, meta interface{}) error {
    s := ResourceActionGroupConfigSchema()
    log.Printf("[INFO] ResourceAviActionGroupConfigRead Avi Client %v\n", d)
    client := meta.(*clients.AviClient)
    var obj interface{}
    if uuid, ok := d.GetOk("uuid"); ok {
        path := "api/actiongroupconfig/" + uuid.(string)
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
    log.Printf("ResourceAviActionGroupConfigRead CURRENT obj %v\n", d)

    log.Printf("ResourceAviActionGroupConfigRead Read API obj %v\n", obj)
    if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
        log.Printf("[INFO] ResourceAviActionGroupConfigRead Converted obj %v\n", tObj)
        //err = d.Set("obj", tObj)
        if err != nil {
            log.Printf("[ERROR] in setting read object %v\n", err)
        }
    }
    log.Printf("[INFO] ResourceAviActionGroupConfigRead Updated %v\n", d)
    return nil
}

func resourceAviActionGroupConfigCreate(d *schema.ResourceData, meta interface{}) error {
    s := ResourceVrfContextSchema()
    err := ApiCreateOrUpdate(d, meta, "actiongroupconfig", s)
    log.Printf("[DEBUG] created object %v: %v", "actiongroupconfig", d)
    if err == nil {
        err = ResourceAviVrfContextRead(d, meta)
    }
    log.Printf("[DEBUG] created object %v: %v", "actiongroupconfig", d)
    return err
}

func resourceAviActionGroupConfigUpdate(d *schema.ResourceData, meta interface{}) error {
    s := ResourcePoolSchema()
    err := ApiCreateOrUpdate(d, meta, "actiongroupconfig", s)
    if err == nil {
        err = ResourceAviVrfContextRead(d, meta)
    }
    log.Printf("[DEBUG] updated object %v: %v", "actiongroupconfig", d)
    return err
}

func resourceAviActionGroupConfigDelete(d *schema.ResourceData, meta interface{}) error {
    s := ResourcePoolSchema()
    err := ApiCreateOrUpdate(d, meta, "actiongroupconfig", s)
    log.Printf("[DEBUG] created object %v: %v", "actiongroupconfig", d)
    if err == nil {
        err = ResourceAviVrfContextRead(d, meta)
    }
    log.Printf("[DEBUG] created object %v: %v", "actiongroupconfig", d)
    return err
}


