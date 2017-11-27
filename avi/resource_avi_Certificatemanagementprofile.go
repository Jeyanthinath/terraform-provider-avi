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
func ResourceCertificateManagementProfileSchema() map[string]*schema.Schema {
    return map[string]*schema.Schema{
             "name" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "script_params" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceCustomParamsSchema(),                             },
             "script_path" :&schema.Schema{
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


func resourceAviCertificateManagementProfile() *schema.Resource {
    return &schema.Resource{
        Create: resourceAviCertificateManagementProfileCreate,
        Read:   ResourceAviCertificateManagementProfileRead,
        Update: resourceAviCertificateManagementProfileUpdate,
        Delete: resourceAviCertificateManagementProfileDelete,
        Schema: ResourceCertificateManagementProfileSchema(),
    }
}


func ResourceAviCertificateManagementProfileRead(d *schema.ResourceData, meta interface{}) error {
    s := ResourceCertificateManagementProfileSchema()
    log.Printf("[INFO] ResourceAviCertificateManagementProfileRead Avi Client %v\n", d)
    client := meta.(*clients.AviClient)
    var obj interface{}
    if uuid, ok := d.GetOk("uuid"); ok {
        path := "api/certificatemanagementprofile/" + uuid.(string)
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
    log.Printf("ResourceAviCertificateManagementProfileRead CURRENT obj %v\n", d)

    log.Printf("ResourceAviCertificateManagementProfileRead Read API obj %v\n", obj)
    if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
        log.Printf("[INFO] ResourceAviCertificateManagementProfileRead Converted obj %v\n", tObj)
        //err = d.Set("obj", tObj)
        if err != nil {
            log.Printf("[ERROR] in setting read object %v\n", err)
        }
    }
    log.Printf("[INFO] ResourceAviCertificateManagementProfileRead Updated %v\n", d)
    return nil
}

func resourceAviCertificateManagementProfileCreate(d *schema.ResourceData, meta interface{}) error {
    s := ResourceVrfContextSchema()
    err := ApiCreateOrUpdate(d, meta, "certificatemanagementprofile", s)
    log.Printf("[DEBUG] created object %v: %v", "certificatemanagementprofile", d)
    if err == nil {
        err = ResourceAviVrfContextRead(d, meta)
    }
    log.Printf("[DEBUG] created object %v: %v", "certificatemanagementprofile", d)
    return err
}

func resourceAviCertificateManagementProfileUpdate(d *schema.ResourceData, meta interface{}) error {
    s := ResourcePoolSchema()
    err := ApiCreateOrUpdate(d, meta, "certificatemanagementprofile", s)
    if err == nil {
        err = ResourceAviVrfContextRead(d, meta)
    }
    log.Printf("[DEBUG] updated object %v: %v", "certificatemanagementprofile", d)
    return err
}

func resourceAviCertificateManagementProfileDelete(d *schema.ResourceData, meta interface{}) error {
    s := ResourcePoolSchema()
    err := ApiCreateOrUpdate(d, meta, "certificatemanagementprofile", s)
    log.Printf("[DEBUG] created object %v: %v", "certificatemanagementprofile", d)
    if err == nil {
        err = ResourceAviVrfContextRead(d, meta)
    }
    log.Printf("[DEBUG] created object %v: %v", "certificatemanagementprofile", d)
    return err
}


