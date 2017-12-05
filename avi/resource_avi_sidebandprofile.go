/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi


import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceSidebandProfileSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "ip" :&schema.Schema{
                             Type: schema.TypeList, 
                             Optional: true,
                                                                                                                 Elem: ResourceIpAddrSchema(),                             },
             "sideband_max_request_body_size" :&schema.Schema{
                             Type: schema.TypeInt, 
                             Optional: true,
                             Default: 1024,
                                                                                                                                            },
                                "url": &schema.Schema{
                                Type:     schema.TypeString,
                                Optional: true,
                                Computed: true,
                            },
        },
    }
}


