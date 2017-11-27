/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceStaticIPResyncReqSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "current_ip" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceStaticIPInfoSchema(),                             },
             "se_name" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "se_uuid" :&schema.Schema{
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

