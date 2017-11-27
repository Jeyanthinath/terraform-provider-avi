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
func ResourceIpamDnsProviderProfileSchema() map[string]*schema.Schema {
    return map[string]*schema.Schema{
             "allocate_ip_in_vrf" :&schema.Schema{
                             Type: schema.TypeBool,            
                             Optional: true,                              
                                                        },
             "aws_profile" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceIpamDnsAwsProfileSchema(),                             },
             "azure_profile" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceIpamDnsAzureProfileSchema(),                             },
             "custom_profile" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceIpamDnsCustomProfileSchema(),                             },
             "gcp_profile" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceIpamDnsGCPProfileSchema(),                             },
             "infoblox_profile" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceIpamDnsInfobloxProfileSchema(),                             },
             "internal_profile" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceIpamDnsInternalProfileSchema(),                             },
             "name" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "openstack_profile" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceIpamDnsOpenstackProfileSchema(),                             },
             "proxy_configuration" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceProxyConfigurationSchema(),                             },
             "tenant_ref" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                              Elem:&schema.Schema{Type: schema.TypeString},                             },
             "type" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
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


func resourceAviIpamDnsProviderProfile() *schema.Resource {
    return &schema.Resource{
        Create: resourceAviIpamDnsProviderProfileCreate,
        Read:   ResourceAviIpamDnsProviderProfileRead,
        Update: resourceAviIpamDnsProviderProfileUpdate,
        Delete: resourceAviIpamDnsProviderProfileDelete,
        Schema: ResourceIpamDnsProviderProfileSchema(),
    }
}


func ResourceAviIpamDnsProviderProfileRead(d *schema.ResourceData, meta interface{}) error {
    s := ResourceIpamDnsProviderProfileSchema()
    log.Printf("[INFO] ResourceAviIpamDnsProviderProfileRead Avi Client %v\n", d)
    client := meta.(*clients.AviClient)
    var obj interface{}
    if uuid, ok := d.GetOk("uuid"); ok {
        path := "api/ipamdnsproviderprofile/" + uuid.(string)
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
    log.Printf("ResourceAviIpamDnsProviderProfileRead CURRENT obj %v\n", d)

    log.Printf("ResourceAviIpamDnsProviderProfileRead Read API obj %v\n", obj)
    if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
        log.Printf("[INFO] ResourceAviIpamDnsProviderProfileRead Converted obj %v\n", tObj)
        //err = d.Set("obj", tObj)
        if err != nil {
            log.Printf("[ERROR] in setting read object %v\n", err)
        }
    }
    log.Printf("[INFO] ResourceAviIpamDnsProviderProfileRead Updated %v\n", d)
    return nil
}

func resourceAviIpamDnsProviderProfileCreate(d *schema.ResourceData, meta interface{}) error {
    s := ResourceVrfContextSchema()
    err := ApiCreateOrUpdate(d, meta, "ipamdnsproviderprofile", s)
    log.Printf("[DEBUG] created object %v: %v", "ipamdnsproviderprofile", d)
    if err == nil {
        err = ResourceAviVrfContextRead(d, meta)
    }
    log.Printf("[DEBUG] created object %v: %v", "ipamdnsproviderprofile", d)
    return err
}

func resourceAviIpamDnsProviderProfileUpdate(d *schema.ResourceData, meta interface{}) error {
    s := ResourcePoolSchema()
    err := ApiCreateOrUpdate(d, meta, "ipamdnsproviderprofile", s)
    if err == nil {
        err = ResourceAviVrfContextRead(d, meta)
    }
    log.Printf("[DEBUG] updated object %v: %v", "ipamdnsproviderprofile", d)
    return err
}

func resourceAviIpamDnsProviderProfileDelete(d *schema.ResourceData, meta interface{}) error {
    s := ResourcePoolSchema()
    err := ApiCreateOrUpdate(d, meta, "ipamdnsproviderprofile", s)
    log.Printf("[DEBUG] created object %v: %v", "ipamdnsproviderprofile", d)
    if err == nil {
        err = ResourceAviVrfContextRead(d, meta)
    }
    log.Printf("[DEBUG] created object %v: %v", "ipamdnsproviderprofile", d)
    return err
}


