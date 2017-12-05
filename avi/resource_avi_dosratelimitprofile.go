/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi


import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceDosRateLimitProfileSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "dos_profile" :&schema.Schema{
                             Type: schema.TypeSet, 
                             Optional: true,
                                                                                                                 Elem: ResourceDosThresholdProfileSchema(),                             },
             "rl_profile" :&schema.Schema{
                             Type: schema.TypeSet, 
                             Optional: true,
                                                                                                                 Elem: ResourceRateLimiterProfileSchema(),                             },
                                "url": &schema.Schema{
                                Type:     schema.TypeString,
                                Optional: true,
                                Computed: true,
                            },
        },
    }
}


