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
func ResourceApplicationPersistenceProfileSchema() map[string]*schema.Schema {
    return map[string]*schema.Schema{
             "app_cookie_persistence_profile" :&schema.Schema{
                             Type: schema.TypeSet, 
                             Optional: true,
                                                                                                                 Elem: ResourceAppCookiePersistenceProfileSchema(),                             },
             "description" :&schema.Schema{
                             Type: schema.TypeString, 
                             Optional: true,
                                                                                                                                            },
             "hdr_persistence_profile" :&schema.Schema{
                             Type: schema.TypeSet, 
                             Optional: true,
                                                                                                                 Elem: ResourceHdrPersistenceProfileSchema(),                             },
             "http_cookie_persistence_profile" :&schema.Schema{
                             Type: schema.TypeSet, 
                             Optional: true,
                                                                                                                 Elem: ResourceHttpCookiePersistenceProfileSchema(),                             },
             "ip_persistence_profile" :&schema.Schema{
                             Type: schema.TypeSet, 
                             Optional: true,
                                                                                                                 Elem: ResourceIPPersistenceProfileSchema(),                             },
             //"is_federated" :&schema.Schema{
             //                Type: schema.TypeBool,
             //                Optional: true,
             //                                                                                                                               },
             "name" :&schema.Schema{
                             Type: schema.TypeString, 
                             Required: true,
                                                                                                                },
             "persistence_type" :&schema.Schema{
                             Type: schema.TypeString, 
                             Required: true,
                                                                                                                },
             "server_hm_down_recovery" :&schema.Schema{
                             Type: schema.TypeString, 
                             Optional: true,
                             Default: "HM_DOWN_PICK_NEW_SERVER",                                                                                                                },
             "tenant_ref" :&schema.Schema{
                             Type: schema.TypeString, 
                             Optional: true,
                                                                                                                  Elem:&schema.Schema{Type: schema.TypeString},                             },
             "uuid" :&schema.Schema{
                             Type: schema.TypeString, 
                             Optional: true,
                                                                                    Computed:  true,                                                         },
                                "url": &schema.Schema{
                                Type:     schema.TypeString,
                                Optional: true,
                                Computed: true,
                            },
    }
}


func resourceAviApplicationPersistenceProfile() *schema.Resource {
    return &schema.Resource{
        Create: resourceAviApplicationPersistenceProfileCreate,
        Read:   ResourceAviApplicationPersistenceProfileRead,
        Update: resourceAviApplicationPersistenceProfileUpdate,
        Delete: resourceAviApplicationPersistenceProfileDelete,
        Schema: ResourceApplicationPersistenceProfileSchema(),
    }
}


func ResourceAviApplicationPersistenceProfileRead(d *schema.ResourceData, meta interface{}) error {
    s := ResourceApplicationPersistenceProfileSchema()
    log.Printf("[INFO] ResourceAviApplicationPersistenceProfileRead Avi Client %v\n", d)
    client := meta.(*clients.AviClient)
    var obj interface{}
    if uuid, ok := d.GetOk("uuid"); ok {
        path := "api/applicationpersistenceprofile/" + uuid.(string)
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
    log.Printf("ResourceAviApplicationPersistenceProfileRead CURRENT obj %v\n", d)

    log.Printf("ResourceAviApplicationPersistenceProfileRead Read API obj %v\n", obj)
    if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
        log.Printf("[INFO] ResourceAviApplicationPersistenceProfileRead Converted obj %v\n", tObj)
        //err = d.Set("obj", tObj)
        if err != nil {
            log.Printf("[ERROR] in setting read object %v\n", err)
        }
    }
    log.Printf("[INFO] ResourceAviApplicationPersistenceProfileRead Updated %v\n", d)
    return nil
}

func resourceAviApplicationPersistenceProfileCreate(d *schema.ResourceData, meta interface{}) error {
    s := ResourceApplicationPersistenceProfileSchema()
    err := ApiCreateOrUpdate(d, meta, "applicationpersistenceprofile", s)
    log.Printf("[DEBUG] created object %v: %v", "applicationpersistenceprofile", d)
    if err == nil {
        err = ResourceAviApplicationPersistenceProfileRead(d, meta)
    }
    log.Printf("[DEBUG] created object %v: %v", "applicationpersistenceprofile", d)
    return err
}

func resourceAviApplicationPersistenceProfileUpdate(d *schema.ResourceData, meta interface{}) error {
    s := ResourceApplicationPersistenceProfileSchema()
    err := ApiCreateOrUpdate(d, meta, "applicationpersistenceprofile", s)
    if err == nil {
        err = ResourceAviApplicationPersistenceProfileRead(d, meta)
    }
    log.Printf("[DEBUG] updated object %v: %v", "applicationpersistenceprofile", d)
    return err
}

func resourceAviApplicationPersistenceProfileDelete(d *schema.ResourceData, meta interface{}) error {
    s := ResourceApplicationPersistenceProfileSchema()
    err := ApiCreateOrUpdate(d, meta, "applicationpersistenceprofile", s)
    log.Printf("[DEBUG] created object %v: %v", "applicationpersistenceprofile", d)
    if err == nil {
        err = ResourceAviApplicationPersistenceProfileRead(d, meta)
    }
    log.Printf("[DEBUG] created object %v: %v", "applicationpersistenceprofile", d)
    return err
}


