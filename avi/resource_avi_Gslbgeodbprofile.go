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
func ResourceGslbGeoDbProfileSchema() map[string]*schema.Schema {
    return map[string]*schema.Schema{
             "description" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "entries" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceGslbGeoDbEntrySchema(),                             },
             "is_federated" :&schema.Schema{
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


func resourceAviGslbGeoDbProfile() *schema.Resource {
    return &schema.Resource{
        Create: resourceAviGslbGeoDbProfileCreate,
        Read:   ResourceAviGslbGeoDbProfileRead,
        Update: resourceAviGslbGeoDbProfileUpdate,
        Delete: resourceAviGslbGeoDbProfileDelete,
        Schema: ResourceGslbGeoDbProfileSchema(),
    }
}


func ResourceAviGslbGeoDbProfileRead(d *schema.ResourceData, meta interface{}) error {
    s := ResourceGslbGeoDbProfileSchema()
    log.Printf("[INFO] ResourceAviGslbGeoDbProfileRead Avi Client %v\n", d)
    client := meta.(*clients.AviClient)
    var obj interface{}
    if uuid, ok := d.GetOk("uuid"); ok {
        path := "api/gslbgeodbprofile/" + uuid.(string)
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
    log.Printf("ResourceAviGslbGeoDbProfileRead CURRENT obj %v\n", d)

    log.Printf("ResourceAviGslbGeoDbProfileRead Read API obj %v\n", obj)
    if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
        log.Printf("[INFO] ResourceAviGslbGeoDbProfileRead Converted obj %v\n", tObj)
        //err = d.Set("obj", tObj)
        if err != nil {
            log.Printf("[ERROR] in setting read object %v\n", err)
        }
    }
    log.Printf("[INFO] ResourceAviGslbGeoDbProfileRead Updated %v\n", d)
    return nil
}

func resourceAviGslbGeoDbProfileCreate(d *schema.ResourceData, meta interface{}) error {
    s := ResourceVrfContextSchema()
    err := ApiCreateOrUpdate(d, meta, "gslbgeodbprofile", s)
    log.Printf("[DEBUG] created object %v: %v", "gslbgeodbprofile", d)
    if err == nil {
        err = ResourceAviVrfContextRead(d, meta)
    }
    log.Printf("[DEBUG] created object %v: %v", "gslbgeodbprofile", d)
    return err
}

func resourceAviGslbGeoDbProfileUpdate(d *schema.ResourceData, meta interface{}) error {
    s := ResourcePoolSchema()
    err := ApiCreateOrUpdate(d, meta, "gslbgeodbprofile", s)
    if err == nil {
        err = ResourceAviVrfContextRead(d, meta)
    }
    log.Printf("[DEBUG] updated object %v: %v", "gslbgeodbprofile", d)
    return err
}

func resourceAviGslbGeoDbProfileDelete(d *schema.ResourceData, meta interface{}) error {
    s := ResourcePoolSchema()
    err := ApiCreateOrUpdate(d, meta, "gslbgeodbprofile", s)
    log.Printf("[DEBUG] created object %v: %v", "gslbgeodbprofile", d)
    if err == nil {
        err = ResourceAviVrfContextRead(d, meta)
    }
    log.Printf("[DEBUG] created object %v: %v", "gslbgeodbprofile", d)
    return err
}


