/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourcePortGroupSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "ip" :&schema.Schema{
                             Type: schema.TypeSet,            
                              Required: true,                              
                             Elem: ResourceIpAddrSchema(),                             },
             "mac_addr" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "port_group" :&schema.Schema{
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


