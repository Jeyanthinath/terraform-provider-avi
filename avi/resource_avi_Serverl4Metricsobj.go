/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceServerL4MetricsObjSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "apdexc" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_available_capacity" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_bandwidth" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_capacity" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_complete_conns" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_connections_dropped" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_errored_connections" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_est_capacity" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_goodput" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_health_status" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_lossy_connections" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_new_established_conns" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_open_conns" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_pool_bandwidth" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_pool_complete_conns" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_pool_errored_connections" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_pool_new_established_conns" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_pool_open_conns" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_rx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_rx_goodput" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_rx_pkts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_server_count" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_server_uptime" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_syns_sent" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_total_rtt" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_tx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_tx_goodput" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_tx_pkts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_uptime" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_available_capacity" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_capacity" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_open_conns" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_rx_bytes_absolute" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_rx_pkts_absolute" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_tx_bytes_absolute" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_tx_pkts_absolute" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "node_obj_id" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "pct_connection_errors" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "pct_connection_saturation" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_conn_duration" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_connection_errors" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_connection_setup_time" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_connections_dropped" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_dup_ack_retransmits" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_finished_conns" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_health_check_failures" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_lb_fail_count" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_lossy_connections" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_lossy_req" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_num_state_changes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_out_of_orders" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_rtt" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_rtt_valid_connections" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_rx_tcp_resets" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_rx_zero_window_size_events" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_sack_retransmits" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_timeout_retransmits" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_tx_zero_window_size_events" :&schema.Schema{
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


