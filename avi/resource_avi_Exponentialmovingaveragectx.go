/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceExponentialMovingAverageCtxSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "average" :&schema.Schema{
                             Type: schema.,            
                              Required: true,                              
                                                        },
             "deviation" :&schema.Schema{
                             Type: schema.,            
                              Required: true,                              
                                                        },
             "max_std_dev" :&schema.Schema{
                             Type: schema.,            
                              Required: true,                              
                                                        },
             "normal_high" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "normal_low" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "prediction" :&schema.Schema{
                             Type: schema.,            
                              Required: true,                              
                                                        },
             "prediction_interval_high" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "prediction_interval_low" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "std_dev" :&schema.Schema{
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


