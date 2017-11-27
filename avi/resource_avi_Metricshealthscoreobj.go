/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceMetricsHealthScoreObjSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "anomaly_penalty" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "health_score_value" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "performance_score_value" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "resources_penalty" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "security_penalty" :&schema.Schema{
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


