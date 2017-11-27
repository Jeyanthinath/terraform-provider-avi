/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceMetricsFiltersSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "op" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "rel_param" :&schema.Schema{
                             Type: schema.TypeSet,            
                              Required: true,                              
                             Elem: ResourceRelationalOperatorsParamsSchema(),                             },
                                "url": &schema.Schema{
                                Type:     schema.TypeString,
                                Optional: true,
                                Computed: true,
                            },
        },
    }
}


