/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceIpNexthopSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "ip" :&schema.Schema{
                             Type: schema.TypeSet,            
                              Required: true,                              
                             Elem: ResourceIpAddrSchema(),                             },
             "nexthop" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceIpAddrSchema(),                             },
                                "url": &schema.Schema{
                                Type:     schema.TypeString,
                                Optional: true,
                                Computed: true,
                            },
        },
    }
}

