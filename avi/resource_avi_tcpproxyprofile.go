/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi


import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceTCPProxyProfileSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "aggressive_congestion_avoidance" :&schema.Schema{
                             Type: schema.TypeBool, 
                             Optional: true,
                                                                                                                                            },
             "automatic" :&schema.Schema{
                             Type: schema.TypeBool, 
                             Optional: true,
                             Default: true,                                                                                                                },
             "cc_algo" :&schema.Schema{
                             Type: schema.TypeString, 
                             Optional: true,
                             Default: "CC_ALGO_NEW_RENO",                                                                                                                },
             "idle_connection_timeout" :&schema.Schema{
                             Type: schema.TypeInt, 
                             Optional: true,
                             Default: 600,
                                                                                                                                            },
             "idle_connection_type" :&schema.Schema{
                             Type: schema.TypeString, 
                             Optional: true,
                             Default: "KEEP_ALIVE",                                                                                                                },
             "ignore_time_wait" :&schema.Schema{
                             Type: schema.TypeBool, 
                             Optional: true,
                                                                                                                                            },
             "ip_dscp" :&schema.Schema{
                             Type: schema.TypeInt, 
                             Optional: true,
                                                                                                                                            },
             "max_retransmissions" :&schema.Schema{
                             Type: schema.TypeInt, 
                             Optional: true,
                             Default: 8,
                                                                                                                                            },
             "max_segment_size" :&schema.Schema{
                             Type: schema.TypeInt, 
                             Optional: true,
                                                                                                                                            },
             "max_syn_retransmissions" :&schema.Schema{
                             Type: schema.TypeInt, 
                             Optional: true,
                             Default: 8,
                                                                                                                                            },
             "nagles_algorithm" :&schema.Schema{
                             Type: schema.TypeBool, 
                             Optional: true,
                                                                                                                                            },
             "receive_window" :&schema.Schema{
                             Type: schema.TypeInt, 
                             Optional: true,
                             Default: 64,
                                                                                                                                            },
             "time_wait_delay" :&schema.Schema{
                             Type: schema.TypeInt, 
                             Optional: true,
                             Default: 2000,
                                                                                                                                            },
             "use_interface_mtu" :&schema.Schema{
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


