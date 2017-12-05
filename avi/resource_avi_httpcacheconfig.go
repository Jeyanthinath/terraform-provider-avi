/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi


import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceHttpCacheConfigSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "age_header" :&schema.Schema{
                             Type: schema.TypeBool, 
                             Optional: true,
                             Default: true,                                                                                                                },
             "aggressive" :&schema.Schema{
                             Type: schema.TypeBool, 
                             Optional: true,
                                                                                                                                            },
             "date_header" :&schema.Schema{
                             Type: schema.TypeBool, 
                             Optional: true,
                             Default: true,                                                                                                                },
             "default_expire" :&schema.Schema{
                             Type: schema.TypeInt, 
                             Optional: true,
                             Default: 600,
                                                                                                                                            },
             "enabled" :&schema.Schema{
                             Type: schema.TypeBool, 
                             Optional: true,
                                                                                                                                            },
             "heuristic_expire" :&schema.Schema{
                             Type: schema.TypeBool, 
                             Optional: true,
                                                                                                                                            },
             "max_cache_size" :&schema.Schema{
                             Type: schema.TypeInt, 
                             Optional: true,
                                                                                                                                            },
             "max_object_size" :&schema.Schema{
                             Type: schema.TypeInt, 
                             Optional: true,
                             Default: 4194304,
                                                                                                                                            },
             "mime_types_black_group_refs" :&schema.Schema{
                             Type: schema.TypeList, 
                             Optional: true,
                                                                                                                  Elem:&schema.Schema{Type: schema.TypeString},                             },
             "mime_types_black_list" :&schema.Schema{
                             Type: schema.TypeList, 
                             Optional: true,
                                                                                                                  Elem:&schema.Schema{Type: schema.TypeString},                             },
             "mime_types_group_refs" :&schema.Schema{
                             Type: schema.TypeList, 
                             Optional: true,
                                                                                                                  Elem:&schema.Schema{Type: schema.TypeString},                             },
             "mime_types_list" :&schema.Schema{
                             Type: schema.TypeList, 
                             Optional: true,
                                                                                                                  Elem:&schema.Schema{Type: schema.TypeString},                             },
             "min_object_size" :&schema.Schema{
                             Type: schema.TypeInt, 
                             Optional: true,
                             Default: 100,
                                                                                                                                            },
             "query_cacheable" :&schema.Schema{
                             Type: schema.TypeBool, 
                             Optional: true,
                                                                                                                                            },
             "xcache_header" :&schema.Schema{
                             Type: schema.TypeBool, 
                             Optional: true,
                             Default: true,                                                                                                                },
                                "url": &schema.Schema{
                                Type:     schema.TypeString,
                                Optional: true,
                                Computed: true,
                            },
        },
    }
}


