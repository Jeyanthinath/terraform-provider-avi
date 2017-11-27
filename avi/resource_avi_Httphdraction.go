/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceHTTPHdrActionSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "action" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "cookie" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceHTTPCookieDataSchema(),                             },
             "hdr" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceHTTPHdrDataSchema(),                             },
                                "url": &schema.Schema{
                                Type:     schema.TypeString,
                                Optional: true,
                                Computed: true,
                            },
        },
    }
}


