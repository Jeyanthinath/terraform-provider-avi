/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceCIFTableSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "adapter" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "cif" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "if_name" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "mac_address" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "se_ref" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                              Elem:&schema.Schema{Type: schema.TypeString},                             },
                                "url": &schema.Schema{
                                Type:     schema.TypeString,
                                Optional: true,
                                Computed: true,
                            },
        },
    }
}


