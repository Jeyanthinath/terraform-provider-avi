/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi


import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceIpAddrSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "addr" :&schema.Schema{
                             Type: schema.TypeString, 
                             Required: true,
                                                                                                                },
             "type" :&schema.Schema{
                             Type: schema.TypeString, 
                             Required: true,
                                                                                                                },
                                "url": &schema.Schema{
                                Type:     schema.TypeString,
                                Optional: true,
                                Computed: true,
                            },
        },
    }
}


