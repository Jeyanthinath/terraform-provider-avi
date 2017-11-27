/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceAnomalyEventDetailsSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "deviation" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "metric_id" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "metric_timestamp" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "model" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "models" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                              Elem:&schema.Schema{Type: schema.TypeString},                             },
             "node_id" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "obj_type" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "pool_name" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "pool_uuid" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "priority" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "server" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "value" :&schema.Schema{
                             Type: schema.,            
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


