/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceVSDataScriptsSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "index" :&schema.Schema{
                             Type: schema.TypeInt,            
                              Required: true,                              
                                                        },
             "vs_datascript_set_ref" :&schema.Schema{
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


