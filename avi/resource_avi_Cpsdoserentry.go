/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceCpsDoserEntrySchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "hits" :&schema.Schema{
                             Type: schema.TypeInt,            
                             Optional: true,                              
                                                        },
             "ip" :&schema.Schema{
                             Type: schema.TypeSet,            
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

