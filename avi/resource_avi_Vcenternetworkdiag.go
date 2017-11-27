/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceVcenterNetworkDiagSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "mgmt_nw_diag" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceVcenterMgmtNwDiagSchema(),                             },
             "overlapping_subnets" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceOverlappingSubnetInfoSchema(),                             },
                                "url": &schema.Schema{
                                Type:     schema.TypeString,
                                Optional: true,
                                Computed: true,
                            },
        },
    }
}

