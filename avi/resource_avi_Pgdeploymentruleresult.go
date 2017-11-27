/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourcePGDeploymentRuleResultSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "metric_value" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "result" :&schema.Schema{
                             Type: schema.TypeBool,            
                             Optional: true,                              
                                                        },
             "rule" :&schema.Schema{
                             Type: schema.TypeSet,            
                              Required: true,                              
                             Elem: ResourcePGDeploymentRuleSchema(),                             },
                                "url": &schema.Schema{
                                Type:     schema.TypeString,
                                Optional: true,
                                Computed: true,
                            },
        },
    }
}


