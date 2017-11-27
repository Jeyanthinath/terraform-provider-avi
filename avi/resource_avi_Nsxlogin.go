/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceNsxLoginSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "nsx_name" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "nsx_password" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "nsx_username" :&schema.Schema{
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


