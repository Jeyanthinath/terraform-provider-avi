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
func ResourceNetworkSecurityPolicySchema() map[string]*schema.Schema {
    return map[string]*schema.Schema{
             "cloud_config_cksum" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "created_by" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "description" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "name" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "rules" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceNetworkSecurityRuleSchema(),                             },
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


func resourceAviNetworkSecurityPolicy() *schema.Resource {
    return &schema.Resource{
        Create: resourceAviNetworkSecurityPolicyCreate,
        Read:   ResourceAviNetworkSecurityPolicyRead,
        Update: resourceAviNetworkSecurityPolicyUpdate,
        Delete: resourceAviNetworkSecurityPolicyDelete,
        Schema: ResourceNetworkSecurityPolicySchema(),
    }
}


func ResourceAviNetworkSecurityPolicyRead(d *schema.ResourceData, meta interface{}) error {
    s := ResourceNetworkSecurityPolicySchema()
    log.Printf("[INFO] ResourceAviNetworkSecurityPolicyRead Avi Client %v\n", d)
    client := meta.(*clients.AviClient)
    var obj interface{}
    if uuid, ok := d.GetOk("uuid"); ok {
        path := "api/networksecuritypolicy/" + uuid.(string)
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
    log.Printf("ResourceAviNetworkSecurityPolicyRead CURRENT obj %v\n", d)

    log.Printf("ResourceAviNetworkSecurityPolicyRead Read API obj %v\n", obj)
    if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
        log.Printf("[INFO] ResourceAviNetworkSecurityPolicyRead Converted obj %v\n", tObj)
        //err = d.Set("obj", tObj)
        if err != nil {
            log.Printf("[ERROR] in setting read object %v\n", err)
        }
    }
    log.Printf("[INFO] ResourceAviNetworkSecurityPolicyRead Updated %v\n", d)
    return nil
}

func resourceAviNetworkSecurityPolicyCreate(d *schema.ResourceData, meta interface{}) error {
    s := ResourceVrfContextSchema()
    err := ApiCreateOrUpdate(d, meta, "networksecuritypolicy", s)
    log.Printf("[DEBUG] created object %v: %v", "networksecuritypolicy", d)
    if err == nil {
        err = ResourceAviVrfContextRead(d, meta)
    }
    log.Printf("[DEBUG] created object %v: %v", "networksecuritypolicy", d)
    return err
}

func resourceAviNetworkSecurityPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
    s := ResourcePoolSchema()
    err := ApiCreateOrUpdate(d, meta, "networksecuritypolicy", s)
    if err == nil {
        err = ResourceAviVrfContextRead(d, meta)
    }
    log.Printf("[DEBUG] updated object %v: %v", "networksecuritypolicy", d)
    return err
}

func resourceAviNetworkSecurityPolicyDelete(d *schema.ResourceData, meta interface{}) error {
    s := ResourcePoolSchema()
    err := ApiCreateOrUpdate(d, meta, "networksecuritypolicy", s)
    log.Printf("[DEBUG] created object %v: %v", "networksecuritypolicy", d)
    if err == nil {
        err = ResourceAviVrfContextRead(d, meta)
    }
    log.Printf("[DEBUG] created object %v: %v", "networksecuritypolicy", d)
    return err
}


