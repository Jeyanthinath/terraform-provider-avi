/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceHoltWintersCtxSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "deviation" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "level" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "prediction" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "prediction_interval_high" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "prediction_interval_low" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "trend" :&schema.Schema{
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


