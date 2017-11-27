/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceCloudStackLoginSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "access_key_id" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "api_url" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "secret_access_key" :&schema.Schema{
                             Type: schema.TypeString,            
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


