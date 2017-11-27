/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceVirtualServiceAnalysisSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "result" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "task" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceVsAnalysisTaskSchema(),                             },
                                "url": &schema.Schema{
                                Type:     schema.TypeString,
                                Optional: true,
                                Computed: true,
                            },
        },
    }
}


