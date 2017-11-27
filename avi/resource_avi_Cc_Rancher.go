/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceCC_RancherSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "access_err" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "cfg" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceRancherConfigurationSchema(),                             },
             "dns_info" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceCC_DnsInfoSchema(),                             },
             "hosts" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceCC_HostSchema(),                             },
             "ipam_info" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceCC_IpamInfoSchema(),                             },
             "sshinfo" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceCC_SshInfoSchema(),                             },
                                "url": &schema.Schema{
                                Type:     schema.TypeString,
                                Optional: true,
                                Computed: true,
                            },
        },
    }
}


