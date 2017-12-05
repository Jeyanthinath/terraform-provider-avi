/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi


import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceURIParamSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "tokens" :&schema.Schema{
                             Type: schema.TypeList, 
                             Optional: true,
                                                                                                                 Elem: ResourceURIParamTokenSchema(),                             },
             "type" :&schema.Schema{
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


