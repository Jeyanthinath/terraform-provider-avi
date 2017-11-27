/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceOverallInfoSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "available" :&schema.Schema{
                             Type: schema.TypeInt,            
                              Required: true,                              
                                                        },
             "free_percent" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "path" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "size" :&schema.Schema{
                             Type: schema.TypeInt,            
                              Required: true,                              
                                                        },
             "used" :&schema.Schema{
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


