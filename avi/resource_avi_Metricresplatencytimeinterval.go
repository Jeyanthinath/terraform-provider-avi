/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceMetricRespLatencyTimeIntervalSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "end" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "start" :&schema.Schema{
                             Type: schema.TypeString,            
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


