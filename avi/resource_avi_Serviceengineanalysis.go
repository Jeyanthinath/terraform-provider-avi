/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceServiceEngineAnalysisSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "result" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "task" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceSeAnalysisTaskSchema(),                             },
                                "url": &schema.Schema{
                                Type:     schema.TypeString,
                                Optional: true,
                                Computed: true,
                            },
        },
    }
}


