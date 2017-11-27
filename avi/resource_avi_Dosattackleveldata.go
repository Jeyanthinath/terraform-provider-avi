/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceDosAttackLevelDataSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "avg_complete_conns" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_dos_total_req" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_total_connections" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "pct_application_dos_attacks" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "pct_connections_dos_attacks" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "pct_dos_bandwidth" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "pct_dos_rx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "pct_network_dos_attacks" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "pct_pkts_dos_attacks" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "pct_policy_drops" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "reason" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "reason_attr" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "ref" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                              Elem:&schema.Schema{Type: schema.TypeString},                             },
             "serviceengine_dos_attack_level_scores" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceServiceEngineDosAttackLevelSchema(),                             },
             "serviceengine_dosattklevel_scores" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceServiceEngineDosAttackLevelSchema(),                             },
             "value" :&schema.Schema{
                             Type: schema.,            
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


