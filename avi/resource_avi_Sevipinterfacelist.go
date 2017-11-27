/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceSeVipInterfaceListSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "is_portchannel" :&schema.Schema{
                             Type: schema.TypeBool,            
                             Optional: true,                              
                                                        },
             "vip_intf_ip" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceIpAddrSchema(),                             },
             "vip_intf_mac" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "vlan_id" :&schema.Schema{
                             Type: schema.TypeInt,            
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


