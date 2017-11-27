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
func ResourceHTTPPolicySetSchema() map[string]*schema.Schema {
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
             "http_request_policy" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceHTTPRequestPolicySchema(),                             },
             "http_response_policy" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceHTTPResponsePolicySchema(),                             },
             "http_security_policy" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceHTTPSecurityPolicySchema(),                             },
             "is_internal_policy" :&schema.Schema{
                             Type: schema.TypeBool,            
                             Optional: true,                              
                                                        },
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


func resourceAviHTTPPolicySet() *schema.Resource {
    return &schema.Resource{
        Create: resourceAviHTTPPolicySetCreate,
        Read:   ResourceAviHTTPPolicySetRead,
        Update: resourceAviHTTPPolicySetUpdate,
        Delete: resourceAviHTTPPolicySetDelete,
        Schema: ResourceHTTPPolicySetSchema(),
    }
}


func ResourceAviHTTPPolicySetRead(d *schema.ResourceData, meta interface{}) error {
    s := ResourceHTTPPolicySetSchema()
    log.Printf("[INFO] ResourceAviHTTPPolicySetRead Avi Client %v\n", d)
    client := meta.(*clients.AviClient)
    var obj interface{}
    if uuid, ok := d.GetOk("uuid"); ok {
        path := "api/httppolicyset/" + uuid.(string)
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
    log.Printf("ResourceAviHTTPPolicySetRead CURRENT obj %v\n", d)

    log.Printf("ResourceAviHTTPPolicySetRead Read API obj %v\n", obj)
    if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
        log.Printf("[INFO] ResourceAviHTTPPolicySetRead Converted obj %v\n", tObj)
        //err = d.Set("obj", tObj)
        if err != nil {
            log.Printf("[ERROR] in setting read object %v\n", err)
        }
    }
    log.Printf("[INFO] ResourceAviHTTPPolicySetRead Updated %v\n", d)
    return nil
}

func resourceAviHTTPPolicySetCreate(d *schema.ResourceData, meta interface{}) error {
    s := ResourceVrfContextSchema()
    err := ApiCreateOrUpdate(d, meta, "httppolicyset", s)
    log.Printf("[DEBUG] created object %v: %v", "httppolicyset", d)
    if err == nil {
        err = ResourceAviVrfContextRead(d, meta)
    }
    log.Printf("[DEBUG] created object %v: %v", "httppolicyset", d)
    return err
}

func resourceAviHTTPPolicySetUpdate(d *schema.ResourceData, meta interface{}) error {
    s := ResourcePoolSchema()
    err := ApiCreateOrUpdate(d, meta, "httppolicyset", s)
    if err == nil {
        err = ResourceAviVrfContextRead(d, meta)
    }
    log.Printf("[DEBUG] updated object %v: %v", "httppolicyset", d)
    return err
}

func resourceAviHTTPPolicySetDelete(d *schema.ResourceData, meta interface{}) error {
    s := ResourcePoolSchema()
    err := ApiCreateOrUpdate(d, meta, "httppolicyset", s)
    log.Printf("[DEBUG] created object %v: %v", "httppolicyset", d)
    if err == nil {
        err = ResourceAviVrfContextRead(d, meta)
    }
    log.Printf("[DEBUG] created object %v: %v", "httppolicyset", d)
    return err
}


