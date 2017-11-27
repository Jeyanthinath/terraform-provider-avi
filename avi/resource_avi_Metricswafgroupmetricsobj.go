/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceMetricsWafGroupMetricsObjSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "avg_evaluated" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_flagged" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_matched" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_rejected" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_evaluated" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_flagged" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_matched" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_rejected" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "node_obj_id" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "sum_evaluated" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_flagged" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_matched" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_rejected" :&schema.Schema{
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


