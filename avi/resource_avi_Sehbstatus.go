/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceSeHbStatusSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "last_hb_req_sent" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "last_hb_resp_recv" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "num_hb_misses" :&schema.Schema{
                             Type: schema.TypeInt,            
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


