/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceAnomalyzerEWMASchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "alpha" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "anomalous_z_score" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "num_samples" :&schema.Schema{
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


