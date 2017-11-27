/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceProxyConfigurationSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "host" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "port" :&schema.Schema{
                             Type: schema.TypeInt,            
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


