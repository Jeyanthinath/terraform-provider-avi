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
func ResourcePoolGroupDeploymentPolicySchema() map[string]*schema.Schema {
    return map[string]*schema.Schema{
             "auto_disable_old_prod_pools" :&schema.Schema{
                             Type: schema.TypeBool,            
                             Optional: true,                              
                                                        },
             "cloud_ref" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                              Elem:&schema.Schema{Type: schema.TypeString},                             },
             "description" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "evaluation_duration" :&schema.Schema{
                             Type: schema.TypeInt,            
                             Optional: true,                              
                                                        },
             "name" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "rules" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourcePGDeploymentRuleSchema(),                             },
             "scheme" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "target_test_traffic_ratio" :&schema.Schema{
                             Type: schema.TypeInt,            
                             Optional: true,                              
                                                        },
             "tenant_ref" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                              Elem:&schema.Schema{Type: schema.TypeString},                             },
             "test_traffic_ratio_rampup" :&schema.Schema{
                             Type: schema.TypeInt,            
                             Optional: true,                              
                                                        },
             "uuid" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                             Computed:  true,  
                                                        },
             "webhook_ref" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                              Elem:&schema.Schema{Type: schema.TypeString},                             },
                                "url": &schema.Schema{
                                Type:     schema.TypeString,
                                Optional: true,
                                Computed: true,
                            },
    }
}


func resourceAviPoolGroupDeploymentPolicy() *schema.Resource {
    return &schema.Resource{
        Create: resourceAviPoolGroupDeploymentPolicyCreate,
        Read:   ResourceAviPoolGroupDeploymentPolicyRead,
        Update: resourceAviPoolGroupDeploymentPolicyUpdate,
        Delete: resourceAviPoolGroupDeploymentPolicyDelete,
        Schema: ResourcePoolGroupDeploymentPolicySchema(),
    }
}


func ResourceAviPoolGroupDeploymentPolicyRead(d *schema.ResourceData, meta interface{}) error {
    s := ResourcePoolGroupDeploymentPolicySchema()
    log.Printf("[INFO] ResourceAviPoolGroupDeploymentPolicyRead Avi Client %v\n", d)
    client := meta.(*clients.AviClient)
    var obj interface{}
    if uuid, ok := d.GetOk("uuid"); ok {
        path := "api/poolgroupdeploymentpolicy/" + uuid.(string)
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
    log.Printf("ResourceAviPoolGroupDeploymentPolicyRead CURRENT obj %v\n", d)

    log.Printf("ResourceAviPoolGroupDeploymentPolicyRead Read API obj %v\n", obj)
    if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
        log.Printf("[INFO] ResourceAviPoolGroupDeploymentPolicyRead Converted obj %v\n", tObj)
        //err = d.Set("obj", tObj)
        if err != nil {
            log.Printf("[ERROR] in setting read object %v\n", err)
        }
    }
    log.Printf("[INFO] ResourceAviPoolGroupDeploymentPolicyRead Updated %v\n", d)
    return nil
}

func resourceAviPoolGroupDeploymentPolicyCreate(d *schema.ResourceData, meta interface{}) error {
    s := ResourceVrfContextSchema()
    err := ApiCreateOrUpdate(d, meta, "poolgroupdeploymentpolicy", s)
    log.Printf("[DEBUG] created object %v: %v", "poolgroupdeploymentpolicy", d)
    if err == nil {
        err = ResourceAviVrfContextRead(d, meta)
    }
    log.Printf("[DEBUG] created object %v: %v", "poolgroupdeploymentpolicy", d)
    return err
}

func resourceAviPoolGroupDeploymentPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
    s := ResourcePoolSchema()
    err := ApiCreateOrUpdate(d, meta, "poolgroupdeploymentpolicy", s)
    if err == nil {
        err = ResourceAviVrfContextRead(d, meta)
    }
    log.Printf("[DEBUG] updated object %v: %v", "poolgroupdeploymentpolicy", d)
    return err
}

func resourceAviPoolGroupDeploymentPolicyDelete(d *schema.ResourceData, meta interface{}) error {
    s := ResourcePoolSchema()
    err := ApiCreateOrUpdate(d, meta, "poolgroupdeploymentpolicy", s)
    log.Printf("[DEBUG] created object %v: %v", "poolgroupdeploymentpolicy", d)
    if err == nil {
        err = ResourceAviVrfContextRead(d, meta)
    }
    log.Printf("[DEBUG] created object %v: %v", "poolgroupdeploymentpolicy", d)
    return err
}


