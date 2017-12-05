/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi


import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceReplaceStringVarSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "type" :&schema.Schema{
                             Type: schema.TypeString, 
                             Optional: true,
                                                                                                                                            },
             "val" :&schema.Schema{
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

