/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceAnomalyContextSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "ema_context" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceExponentialMovingAverageCtxSchema(),                             },
             "ewma_context" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceExponentialMovingAverageCtxSchema(),                             },
             "hw_at_as_context" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceHoltWintersCtxSchema(),                             },
             "hw_at_ms_context" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceHoltWintersCtxSchema(),                             },
             "prediction_interval_high" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "prediction_interval_low" :&schema.Schema{
                             Type: schema.,            
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


