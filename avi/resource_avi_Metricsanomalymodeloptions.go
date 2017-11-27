/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceMetricsAnomalyModelOptionsSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "allowed_entity_types" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                              Elem:&schema.Schema{Type: schema.TypeString},                             },
             "alpha" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "anomaly_model" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "beta" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "delta" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "gamma" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "period" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "seasonality" :&schema.Schema{
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


