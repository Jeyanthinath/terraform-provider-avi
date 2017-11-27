/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceServiceEngineMibEntrySchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "id" :&schema.Schema{
                             Type: schema.TypeInt,            
                              Required: true,                              
                                                        },
             "mgmt_ip" :&schema.Schema{
                             Type: schema.TypeSet,            
                              Required: true,                              
                             Elem: ResourceIpAddrSchema(),                             },
             "name" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "status" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "uuid" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                             Computed:  true,  
                                                        },
                                "url": &schema.Schema{
                                Type:     schema.TypeString,
                                Optional: true,
                                Computed: true,
                            },
        },
    }
}


