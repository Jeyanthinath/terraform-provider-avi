/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceMetricLogSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "end_timestamp" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "metric_id" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "report_timestamp" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "step" :&schema.Schema{
                             Type: schema.TypeInt,            
                             Optional: true,                              
                                                        },
             "time_series" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceMetricsQueryResponseSchema(),                             },
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


