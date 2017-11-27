/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceApicDevicePkgVerSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "MajorVersion" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "MinorVersion" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "Model" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "Vendor" :&schema.Schema{
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


