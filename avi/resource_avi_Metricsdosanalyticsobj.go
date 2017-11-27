/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceMetricsDosAnalyticsObjSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "avg_attack_count" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "node_obj_id" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "sum_attack_duration" :&schema.Schema{
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


