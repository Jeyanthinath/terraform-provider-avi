/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceMesosSeResourcesSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "attribute_key" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "attribute_value" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "cpu" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "memory" :&schema.Schema{
                             Type: schema.TypeInt,            
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


