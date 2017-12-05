/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi


import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceIpAddrRangeSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "begin" :&schema.Schema{
                             Type: schema.TypeSet, 
                             Required: true,
                                                                                     Elem: ResourceIpAddrSchema(),                             },
             "end" :&schema.Schema{
                             Type: schema.TypeSet, 
                             Required: true,
                                                                                     Elem: ResourceIpAddrSchema(),                             },
                                "url": &schema.Schema{
                                Type:     schema.TypeString,
                                Optional: true,
                                Computed: true,
                            },
        },
    }
}


