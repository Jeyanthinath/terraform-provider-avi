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
func ResourceSystemConfigurationSchema() map[string]*schema.Schema {
    return map[string]*schema.Schema{
             "admin_auth_configuration" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceAdminAuthConfigurationSchema(),                             },
             "dns_configuration" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceDNSConfigurationSchema(),                             },
             "dns_virtualservice_refs" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                              Elem:&schema.Schema{Type: schema.TypeString},                             },
             "docker_mode" :&schema.Schema{
                             Type: schema.TypeBool,            
                             Optional: true,                              
                                                        },
             "email_configuration" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceEmailConfigurationSchema(),                             },
             "global_tenant_config" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceTenantConfigurationSchema(),                             },
             "linux_configuration" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceLinuxConfigurationSchema(),                             },
             "mgmt_ip_access_control" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceMgmtIpAccessControlSchema(),                             },
             "ntp_configuration" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceNTPConfigurationSchema(),                             },
             "portal_configuration" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourcePortalConfigurationSchema(),                             },
             "proxy_configuration" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceProxyConfigurationSchema(),                             },
             "snmp_configuration" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSnmpConfigurationSchema(),                             },
             "ssh_ciphers" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                              Elem:&schema.Schema{Type: schema.TypeString},                             },
             "ssh_hmacs" :&schema.Schema{
                             Type: schema.TypeList,            
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


func resourceAviSystemConfiguration() *schema.Resource {
    return &schema.Resource{
        Create: resourceAviSystemConfigurationCreate,
        Read:   ResourceAviSystemConfigurationRead,
        Update: resourceAviSystemConfigurationUpdate,
        Delete: resourceAviSystemConfigurationDelete,
        Schema: ResourceSystemConfigurationSchema(),
    }
}


func ResourceAviSystemConfigurationRead(d *schema.ResourceData, meta interface{}) error {
    s := ResourceSystemConfigurationSchema()
    log.Printf("[INFO] ResourceAviSystemConfigurationRead Avi Client %v\n", d)
    client := meta.(*clients.AviClient)
    var obj interface{}
    if uuid, ok := d.GetOk("uuid"); ok {
        path := "api/systemconfiguration/" + uuid.(string)
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
    log.Printf("ResourceAviSystemConfigurationRead CURRENT obj %v\n", d)

    log.Printf("ResourceAviSystemConfigurationRead Read API obj %v\n", obj)
    if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
        log.Printf("[INFO] ResourceAviSystemConfigurationRead Converted obj %v\n", tObj)
        //err = d.Set("obj", tObj)
        if err != nil {
            log.Printf("[ERROR] in setting read object %v\n", err)
        }
    }
    log.Printf("[INFO] ResourceAviSystemConfigurationRead Updated %v\n", d)
    return nil
}

func resourceAviSystemConfigurationCreate(d *schema.ResourceData, meta interface{}) error {
    s := ResourceVrfContextSchema()
    err := ApiCreateOrUpdate(d, meta, "systemconfiguration", s)
    log.Printf("[DEBUG] created object %v: %v", "systemconfiguration", d)
    if err == nil {
        err = ResourceAviVrfContextRead(d, meta)
    }
    log.Printf("[DEBUG] created object %v: %v", "systemconfiguration", d)
    return err
}

func resourceAviSystemConfigurationUpdate(d *schema.ResourceData, meta interface{}) error {
    s := ResourcePoolSchema()
    err := ApiCreateOrUpdate(d, meta, "systemconfiguration", s)
    if err == nil {
        err = ResourceAviVrfContextRead(d, meta)
    }
    log.Printf("[DEBUG] updated object %v: %v", "systemconfiguration", d)
    return err
}

func resourceAviSystemConfigurationDelete(d *schema.ResourceData, meta interface{}) error {
    s := ResourcePoolSchema()
    err := ApiCreateOrUpdate(d, meta, "systemconfiguration", s)
    log.Printf("[DEBUG] created object %v: %v", "systemconfiguration", d)
    if err == nil {
        err = ResourceAviVrfContextRead(d, meta)
    }
    log.Printf("[DEBUG] created object %v: %v", "systemconfiguration", d)
    return err
}


