/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourcePropertySchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "name" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "value" :&schema.Schema{
                             Type: schema.TypeString,            
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


