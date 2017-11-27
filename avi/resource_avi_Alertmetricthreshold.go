/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceAlertMetricThresholdSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "comparator" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "threshold" :&schema.Schema{
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


