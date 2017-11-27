/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceServerOperationalAnalysisSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "pool_ref" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                              Elem:&schema.Schema{Type: schema.TypeString},                             },
             "server_ip" :&schema.Schema{
                             Type: schema.TypeSet,            
                              Required: true,                              
                             Elem: ResourceIpAddrSchema(),                             },
             "server_oper_status" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceOperationalStatusSchema(),                             },
             "server_port" :&schema.Schema{
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


