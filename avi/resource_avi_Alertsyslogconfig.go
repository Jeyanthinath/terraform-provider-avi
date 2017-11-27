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
func ResourceAlertSyslogConfigSchema() map[string]*schema.Schema {
    return map[string]*schema.Schema{
             "description" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "name" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "syslog_servers" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceAlertSyslogServerSchema(),                             },
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


func resourceAviAlertSyslogConfig() *schema.Resource {
    return &schema.Resource{
        Create: resourceAviAlertSyslogConfigCreate,
        Read:   ResourceAviAlertSyslogConfigRead,
        Update: resourceAviAlertSyslogConfigUpdate,
        Delete: resourceAviAlertSyslogConfigDelete,
        Schema: ResourceAlertSyslogConfigSchema(),
    }
}


func ResourceAviAlertSyslogConfigRead(d *schema.ResourceData, meta interface{}) error {
    s := ResourceAlertSyslogConfigSchema()
    log.Printf("[INFO] ResourceAviAlertSyslogConfigRead Avi Client %v\n", d)
    client := meta.(*clients.AviClient)
    var obj interface{}
    if uuid, ok := d.GetOk("uuid"); ok {
        path := "api/alertsyslogconfig/" + uuid.(string)
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
    log.Printf("ResourceAviAlertSyslogConfigRead CURRENT obj %v\n", d)

    log.Printf("ResourceAviAlertSyslogConfigRead Read API obj %v\n", obj)
    if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
        log.Printf("[INFO] ResourceAviAlertSyslogConfigRead Converted obj %v\n", tObj)
        //err = d.Set("obj", tObj)
        if err != nil {
            log.Printf("[ERROR] in setting read object %v\n", err)
        }
    }
    log.Printf("[INFO] ResourceAviAlertSyslogConfigRead Updated %v\n", d)
    return nil
}

func resourceAviAlertSyslogConfigCreate(d *schema.ResourceData, meta interface{}) error {
    s := ResourceVrfContextSchema()
    err := ApiCreateOrUpdate(d, meta, "alertsyslogconfig", s)
    log.Printf("[DEBUG] created object %v: %v", "alertsyslogconfig", d)
    if err == nil {
        err = ResourceAviVrfContextRead(d, meta)
    }
    log.Printf("[DEBUG] created object %v: %v", "alertsyslogconfig", d)
    return err
}

func resourceAviAlertSyslogConfigUpdate(d *schema.ResourceData, meta interface{}) error {
    s := ResourcePoolSchema()
    err := ApiCreateOrUpdate(d, meta, "alertsyslogconfig", s)
    if err == nil {
        err = ResourceAviVrfContextRead(d, meta)
    }
    log.Printf("[DEBUG] updated object %v: %v", "alertsyslogconfig", d)
    return err
}

func resourceAviAlertSyslogConfigDelete(d *schema.ResourceData, meta interface{}) error {
    s := ResourcePoolSchema()
    err := ApiCreateOrUpdate(d, meta, "alertsyslogconfig", s)
    log.Printf("[DEBUG] created object %v: %v", "alertsyslogconfig", d)
    if err == nil {
        err = ResourceAviVrfContextRead(d, meta)
    }
    log.Printf("[DEBUG] created object %v: %v", "alertsyslogconfig", d)
    return err
}


