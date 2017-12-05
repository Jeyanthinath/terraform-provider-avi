/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi


import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceIPNetworkSubnetSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "network_ref" :&schema.Schema{
                             Type: schema.TypeString, 
                             Optional: true,
                                                                                                                  Elem:&schema.Schema{Type: schema.TypeString},                             },
             "subnet" :&schema.Schema{
                             Type: schema.TypeSet, 
                             Optional: true,
                                                                                                                 Elem: ResourceIpAddrPrefixSchema(),                             },
             "subnet6" :&schema.Schema{
                             Type: schema.TypeSet, 
                             Optional: true,
                                                                                                                 Elem: ResourceIpAddrPrefixSchema(),                             },
             //"subnet6_uuid" :&schema.Schema{
             //                Type: schema.TypeString,
             //                Optional: true,
             //                                                                                                                               },
             //"subnet_uuid" :&schema.Schema{
             //                Type: schema.TypeString,
             //                Optional: true,
             //                                                                                                                               },
                                "url": &schema.Schema{
                                Type:     schema.TypeString,
                                Optional: true,
                                Computed: true,
                            },
        },
    }
}


