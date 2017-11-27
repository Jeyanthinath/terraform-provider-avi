/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceVsScaleoutParamsSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "admin_up" :&schema.Schema{
                             Type: schema.TypeBool,            
                             Optional: true,                              
                                                        },
             "new_vcpus" :&schema.Schema{
                             Type: schema.TypeInt,            
                             Optional: true,                              
                                                        },
             "to_host_ref" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                              Elem:&schema.Schema{Type: schema.TypeString},                             },
             "to_new_se" :&schema.Schema{
                             Type: schema.TypeBool,            
                             Optional: true,                              
                                                        },
             "to_se_ref" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                              Elem:&schema.Schema{Type: schema.TypeString},                             },
             "uuid" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                             Computed:  true,  
                                                        },
             "vip_id" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
                                "url": &schema.Schema{
                                Type:     schema.TypeString,
                                Optional: true,
                                Computed: true,
                            },
        },
    }
}

