/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceVirtualServiceResourcesScoreDataSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "avg_bandwidth_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "ha_index" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_open_conns_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "multivip_index" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "pool_resources_scores" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourcePoolResourcesScoreSchema(),                             },
             "reason" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "reason_attr" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "ref" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                              Elem:&schema.Schema{Type: schema.TypeString},                             },
             "serviceengine_resources_scores" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceServiceEngineResourcesScoreSchema(),                             },
                                "url": &schema.Schema{
                                Type:     schema.TypeString,
                                Optional: true,
                                Computed: true,
                            },
        },
    }
}


