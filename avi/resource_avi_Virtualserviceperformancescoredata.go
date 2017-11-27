/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceVirtualServicePerformanceScoreDataSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "apdexc" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "apdexr" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_pool_performance_score" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "oper_state" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "pool_performance_scores" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourcePoolPerformanceScoreSchema(),                             },
             "reason" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "reason_attr" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "ref" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                              Elem:&schema.Schema{Type: schema.TypeString},                             },
             "rum_apdexr" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_finished_conns" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "vs_uptime" :&schema.Schema{
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


