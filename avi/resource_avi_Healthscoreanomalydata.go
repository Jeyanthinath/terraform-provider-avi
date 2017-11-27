/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceHealthScoreAnomalyDataSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "application_anomaly" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceApplicationAnomalyScoreDataSchema(),                             },
             "controller_anomaly" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceControllerAnomalyScoreDataSchema(),                             },
             "pool_anomaly" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourcePoolAnomalyScoreDataSchema(),                             },
             "reason" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "reason_attr" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "server_anomaly" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceServerAnomalyScoreDataSchema(),                             },
             "serviceengine_anomaly" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceServiceEngineAnomalyScoreDataSchema(),                             },
             "value" :&schema.Schema{
                             Type: schema.,            
                              Required: true,                              
                                                        },
             "virtualservice_anomaly" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceVirtualServiceAnomalyScoreDataSchema(),                             },
                                "url": &schema.Schema{
                                Type:     schema.TypeString,
                                Optional: true,
                                Computed: true,
                            },
        },
    }
}


