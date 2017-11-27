/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceRetrieveIpsSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "count" :&schema.Schema{
                             Type: schema.TypeInt,            
                             Optional: true,                              
                                                        },
             "subnet" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceIpAddrPrefixSchema(),                             },
                                "url": &schema.Schema{
                                Type:     schema.TypeString,
                                Optional: true,
                                Computed: true,
                            },
        },
    }
}


