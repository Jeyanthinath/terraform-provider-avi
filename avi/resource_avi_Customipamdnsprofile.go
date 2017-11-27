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
func ResourceCustomIpamDnsProfileSchema() map[string]*schema.Schema {
    return map[string]*schema.Schema{
             "name" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "script_params" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceCustomParamsSchema(),                             },
             "script_uri" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
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


func resourceAviCustomIpamDnsProfile() *schema.Resource {
    return &schema.Resource{
        Create: resourceAviCustomIpamDnsProfileCreate,
        Read:   ResourceAviCustomIpamDnsProfileRead,
        Update: resourceAviCustomIpamDnsProfileUpdate,
        Delete: resourceAviCustomIpamDnsProfileDelete,
        Schema: ResourceCustomIpamDnsProfileSchema(),
    }
}


func ResourceAviCustomIpamDnsProfileRead(d *schema.ResourceData, meta interface{}) error {
    s := ResourceCustomIpamDnsProfileSchema()
    log.Printf("[INFO] ResourceAviCustomIpamDnsProfileRead Avi Client %v\n", d)
    client := meta.(*clients.AviClient)
    var obj interface{}
    if uuid, ok := d.GetOk("uuid"); ok {
        path := "api/customipamdnsprofile/" + uuid.(string)
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
    log.Printf("ResourceAviCustomIpamDnsProfileRead CURRENT obj %v\n", d)

    log.Printf("ResourceAviCustomIpamDnsProfileRead Read API obj %v\n", obj)
    if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
        log.Printf("[INFO] ResourceAviCustomIpamDnsProfileRead Converted obj %v\n", tObj)
        //err = d.Set("obj", tObj)
        if err != nil {
            log.Printf("[ERROR] in setting read object %v\n", err)
        }
    }
    log.Printf("[INFO] ResourceAviCustomIpamDnsProfileRead Updated %v\n", d)
    return nil
}

func resourceAviCustomIpamDnsProfileCreate(d *schema.ResourceData, meta interface{}) error {
    s := ResourceVrfContextSchema()
    err := ApiCreateOrUpdate(d, meta, "customipamdnsprofile", s)
    log.Printf("[DEBUG] created object %v: %v", "customipamdnsprofile", d)
    if err == nil {
        err = ResourceAviVrfContextRead(d, meta)
    }
    log.Printf("[DEBUG] created object %v: %v", "customipamdnsprofile", d)
    return err
}

func resourceAviCustomIpamDnsProfileUpdate(d *schema.ResourceData, meta interface{}) error {
    s := ResourcePoolSchema()
    err := ApiCreateOrUpdate(d, meta, "customipamdnsprofile", s)
    if err == nil {
        err = ResourceAviVrfContextRead(d, meta)
    }
    log.Printf("[DEBUG] updated object %v: %v", "customipamdnsprofile", d)
    return err
}

func resourceAviCustomIpamDnsProfileDelete(d *schema.ResourceData, meta interface{}) error {
    s := ResourcePoolSchema()
    err := ApiCreateOrUpdate(d, meta, "customipamdnsprofile", s)
    log.Printf("[DEBUG] created object %v: %v", "customipamdnsprofile", d)
    if err == nil {
        err = ResourceAviVrfContextRead(d, meta)
    }
    log.Printf("[DEBUG] created object %v: %v", "customipamdnsprofile", d)
    return err
}


