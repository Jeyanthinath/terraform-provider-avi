/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceRetrieveVcenterSubfoldersRspSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "cloud_uuid" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "folder_info" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceVISubfoldersSchema(),                             },
             "status" :&schema.Schema{
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


