/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceServiceEngineRuntimeSummarySchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "active_tags" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                              Elem:&schema.Schema{Type: schema.TypeString},                             },
             "at_curr_ver" :&schema.Schema{
                             Type: schema.TypeBool,            
                             Optional: true,                              
                                                        },
             "gateway_up" :&schema.Schema{
                             Type: schema.TypeBool,            
                             Optional: true,                              
                                                        },
             "hb_status" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeHbStatusSchema(),                             },
             "inband_mgmt" :&schema.Schema{
                             Type: schema.TypeBool,            
                             Optional: true,                              
                                                        },
             "migrate_state" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "online_since" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "oper_status" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceOperationalStatusSchema(),                             },
             "power_state" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "se_connected" :&schema.Schema{
                             Type: schema.TypeBool,            
                             Optional: true,                              
                                                        },
             "version" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "vinfra_discovered" :&schema.Schema{
                             Type: schema.TypeBool,            
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

