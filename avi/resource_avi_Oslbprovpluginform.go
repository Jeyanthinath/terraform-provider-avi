/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceOsLbProvPluginFormSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "neutron_host_ip" :&schema.Schema{
                             Type: schema.TypeSet,            
                              Required: true,                              
                             Elem: ResourceIpAddrSchema(),                             },
             "neutron_host_passwd" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "neutron_host_user" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "op_type" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "prov_host_ip" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceIpAddrSchema(),                             },
             "prov_name" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "prov_svc_passwd" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "prov_svc_user" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "restart_neutron" :&schema.Schema{
                             Type: schema.TypeBool,            
                             Optional: true,                              
                                                        },
                                "url": &schema.Schema{
                                Type:     schema.TypeString,
                                Optional: true,
                                Computed: true,
                            },
        },
    }
}


