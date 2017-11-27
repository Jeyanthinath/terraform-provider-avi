/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceEventDetailsFilterSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "comparator" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "event_details_key" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "event_details_value" :&schema.Schema{
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


