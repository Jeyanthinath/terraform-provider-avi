/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceHttpCacheStatsSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "additions" :&schema.Schema{
                             Type: schema.TypeInt,            
                              Required: true,                              
                                                        },
             "available_size" :&schema.Schema{
                             Type: schema.TypeInt,            
                              Required: true,                              
                                                        },
             "current_size" :&schema.Schema{
                             Type: schema.TypeInt,            
                              Required: true,                              
                                                        },
             "deletions" :&schema.Schema{
                             Type: schema.TypeInt,            
                              Required: true,                              
                                                        },
             "hits" :&schema.Schema{
                             Type: schema.TypeInt,            
                              Required: true,                              
                                                        },
             "lookups" :&schema.Schema{
                             Type: schema.TypeInt,            
                              Required: true,                              
                                                        },
             "objects" :&schema.Schema{
                             Type: schema.TypeInt,            
                              Required: true,                              
                                                        },
             "out_of_mem_evicts" :&schema.Schema{
                             Type: schema.TypeInt,            
                              Required: true,                              
                                                        },
             "proc_id" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "se_uuid" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "served_bytes" :&schema.Schema{
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


