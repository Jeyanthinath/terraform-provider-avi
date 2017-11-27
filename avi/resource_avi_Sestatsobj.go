/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceSeStatsObjSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "avg_bandwidth" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_connection_mem_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_connection_policy_drops" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_connections" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_connections_dropped" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_cpu_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_disk1_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_dos_bandwidth" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_dos_icmp_flood" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_dos_ip_frag_full" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_dos_ip_frag_incomplete" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_dos_ip_frag_overrun" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_dos_ip_frag_toosmall" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_dos_land" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_dos_port_scan" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_dos_rx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_dos_smurf" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_dos_tcp_non_syn_flood" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_dos_teardrop" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_dos_tx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_dos_unknown_protocol" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_dynamic_mem_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth0_bandwidth" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth0_connection_table_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth0_rx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth0_rx_bytes_dropped_non_vs" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth0_rx_pkts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth0_rx_pkts_dropped_non_vs" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth0_syn_seen_entries_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth0_tx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth0_tx_pkts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth10_bandwidth" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth10_connection_table_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth10_rx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth10_rx_bytes_dropped_non_vs" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth10_rx_pkts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth10_rx_pkts_dropped_non_vs" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth10_syn_seen_entries_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth10_tx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth10_tx_pkts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth11_bandwidth" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth11_connection_table_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth11_rx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth11_rx_bytes_dropped_non_vs" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth11_rx_pkts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth11_rx_pkts_dropped_non_vs" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth11_syn_seen_entries_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth11_tx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth11_tx_pkts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth12_bandwidth" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth12_connection_table_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth12_rx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth12_rx_bytes_dropped_non_vs" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth12_rx_pkts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth12_rx_pkts_dropped_non_vs" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth12_syn_seen_entries_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth12_tx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth12_tx_pkts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth13_bandwidth" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth13_connection_table_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth13_rx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth13_rx_bytes_dropped_non_vs" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth13_rx_pkts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth13_rx_pkts_dropped_non_vs" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth13_syn_seen_entries_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth13_tx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth13_tx_pkts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth14_bandwidth" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth14_connection_table_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth14_rx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth14_rx_bytes_dropped_non_vs" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth14_rx_pkts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth14_rx_pkts_dropped_non_vs" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth14_syn_seen_entries_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth14_tx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth14_tx_pkts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth15_bandwidth" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth15_connection_table_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth15_rx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth15_rx_bytes_dropped_non_vs" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth15_rx_pkts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth15_rx_pkts_dropped_non_vs" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth15_syn_seen_entries_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth15_tx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth15_tx_pkts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth16_bandwidth" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth16_connection_table_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth16_rx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth16_rx_bytes_dropped_non_vs" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth16_rx_pkts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth16_rx_pkts_dropped_non_vs" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth16_syn_seen_entries_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth16_tx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth16_tx_pkts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth17_bandwidth" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth17_connection_table_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth17_rx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth17_rx_bytes_dropped_non_vs" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth17_rx_pkts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth17_rx_pkts_dropped_non_vs" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth17_syn_seen_entries_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth17_tx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth17_tx_pkts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth18_bandwidth" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth18_connection_table_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth18_rx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth18_rx_bytes_dropped_non_vs" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth18_rx_pkts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth18_rx_pkts_dropped_non_vs" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth18_syn_seen_entries_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth18_tx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth18_tx_pkts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth19_bandwidth" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth19_connection_table_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth19_rx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth19_rx_bytes_dropped_non_vs" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth19_rx_pkts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth19_rx_pkts_dropped_non_vs" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth19_syn_seen_entries_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth19_tx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth19_tx_pkts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth1_bandwidth" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth1_connection_table_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth1_rx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth1_rx_bytes_dropped_non_vs" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth1_rx_pkts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth1_rx_pkts_dropped_non_vs" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth1_syn_seen_entries_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth1_tx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth1_tx_pkts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth20_bandwidth" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth20_connection_table_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth20_rx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth20_rx_bytes_dropped_non_vs" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth20_rx_pkts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth20_rx_pkts_dropped_non_vs" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth20_syn_seen_entries_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth20_tx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth20_tx_pkts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth21_bandwidth" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth21_connection_table_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth21_rx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth21_rx_bytes_dropped_non_vs" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth21_rx_pkts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth21_rx_pkts_dropped_non_vs" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth21_syn_seen_entries_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth21_tx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth21_tx_pkts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth22_bandwidth" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth22_connection_table_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth22_rx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth22_rx_bytes_dropped_non_vs" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth22_rx_pkts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth22_rx_pkts_dropped_non_vs" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth22_syn_seen_entries_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth22_tx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth22_tx_pkts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth23_bandwidth" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth23_connection_table_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth23_rx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth23_rx_bytes_dropped_non_vs" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth23_rx_pkts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth23_rx_pkts_dropped_non_vs" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth23_syn_seen_entries_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth23_tx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth23_tx_pkts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth2_bandwidth" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth2_connection_table_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth2_rx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth2_rx_bytes_dropped_non_vs" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth2_rx_pkts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth2_rx_pkts_dropped_non_vs" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth2_syn_seen_entries_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth2_tx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth2_tx_pkts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth3_bandwidth" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth3_connection_table_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth3_rx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth3_rx_bytes_dropped_non_vs" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth3_rx_pkts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth3_rx_pkts_dropped_non_vs" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth3_syn_seen_entries_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth3_tx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth3_tx_pkts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth4_bandwidth" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth4_connection_table_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth4_rx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth4_rx_bytes_dropped_non_vs" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth4_rx_pkts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth4_rx_pkts_dropped_non_vs" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth4_syn_seen_entries_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth4_tx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth4_tx_pkts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth5_bandwidth" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth5_connection_table_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth5_rx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth5_rx_bytes_dropped_non_vs" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth5_rx_pkts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth5_rx_pkts_dropped_non_vs" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth5_syn_seen_entries_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth5_tx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth5_tx_pkts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth6_bandwidth" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth6_connection_table_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth6_rx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth6_rx_bytes_dropped_non_vs" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth6_rx_pkts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth6_rx_pkts_dropped_non_vs" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth6_syn_seen_entries_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth6_tx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth6_tx_pkts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth7_bandwidth" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth7_connection_table_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth7_rx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth7_rx_bytes_dropped_non_vs" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth7_rx_pkts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth7_rx_pkts_dropped_non_vs" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth7_syn_seen_entries_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth7_tx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth7_tx_pkts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth8_bandwidth" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth8_connection_table_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth8_rx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth8_rx_bytes_dropped_non_vs" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth8_rx_pkts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth8_rx_pkts_dropped_non_vs" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth8_syn_seen_entries_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth8_tx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth8_tx_pkts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth9_bandwidth" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth9_connection_table_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth9_rx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth9_rx_bytes_dropped_non_vs" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth9_rx_pkts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth9_rx_pkts_dropped_non_vs" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth9_syn_seen_entries_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth9_tx_bytes" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_eth9_tx_pkts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_mem_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_packet_buffer_header_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_packet_buffer_large_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_packet_buffer_size" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_packet_buffer_small_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_packet_buffer_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_persistent_table_size" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_persistent_table_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_rx_bandwidth" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_rx_bytes_dropped" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_rx_pkts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_rx_pkts_dropped" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_ssl_session_cache" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_ssl_session_cache_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_tx_bandwidth" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_tx_pkts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_connection_mem_total" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_connection_table_size" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth0_max_bw" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth0_rx_bytes_absolute" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth0_tx_bytes_absolute" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth10_max_bw" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth10_rx_bytes_absolute" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth10_tx_bytes_absolute" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth11_max_bw" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth11_rx_bytes_absolute" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth11_tx_bytes_absolute" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth12_max_bw" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth12_rx_bytes_absolute" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth12_tx_bytes_absolute" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth13_max_bw" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth13_rx_bytes_absolute" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth13_tx_bytes_absolute" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth14_max_bw" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth14_rx_bytes_absolute" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth14_tx_bytes_absolute" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth15_max_bw" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth15_rx_bytes_absolute" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth15_tx_bytes_absolute" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth16_max_bw" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth16_rx_bytes_absolute" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth16_tx_bytes_absolute" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth17_max_bw" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth17_rx_bytes_absolute" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth17_tx_bytes_absolute" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth18_max_bw" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth18_rx_bytes_absolute" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth18_tx_bytes_absolute" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth19_max_bw" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth19_rx_bytes_absolute" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth19_tx_bytes_absolute" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth1_max_bw" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth1_rx_bytes_absolute" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth1_tx_bytes_absolute" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth20_max_bw" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth20_rx_bytes_absolute" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth20_tx_bytes_absolute" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth21_max_bw" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth21_rx_bytes_absolute" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth21_tx_bytes_absolute" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth22_max_bw" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth22_rx_bytes_absolute" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth22_tx_bytes_absolute" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth23_max_bw" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth23_rx_bytes_absolute" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth23_tx_bytes_absolute" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth2_max_bw" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth2_rx_bytes_absolute" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth2_tx_bytes_absolute" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth3_max_bw" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth3_rx_bytes_absolute" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth3_tx_bytes_absolute" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth4_max_bw" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth4_rx_bytes_absolute" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth4_tx_bytes_absolute" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth5_max_bw" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth5_rx_bytes_absolute" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth5_tx_bytes_absolute" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth6_max_bw" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth6_rx_bytes_absolute" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth6_tx_bytes_absolute" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth7_max_bw" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth7_rx_bytes_absolute" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth7_tx_bytes_absolute" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth8_max_bw" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth8_rx_bytes_absolute" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth8_tx_bytes_absolute" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth9_max_bw" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth9_rx_bytes_absolute" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_eth9_tx_bytes_absolute" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_num_vs" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_se_bandwidth" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_synseen_entries_size" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_total_memory" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "node_obj_id" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "pct_connections_dropped" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "pct_rx_bytes_dropped" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "pct_rx_pkts_dropped" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "pct_syn_cache_usage" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_cache_object_allocation_failure" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_connection_dropped_memory_limit" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_connection_dropped_packet_buffer_stressed" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_connection_dropped_persistence_table_limit" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_eth0_connection_dropped_syn_seen_limit" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_eth0_connection_dropped_table_limit" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_eth10_connection_dropped_syn_seen_limit" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_eth10_connection_dropped_table_limit" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_eth11_connection_dropped_syn_seen_limit" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_eth11_connection_dropped_table_limit" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_eth12_connection_dropped_syn_seen_limit" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_eth12_connection_dropped_table_limit" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_eth13_connection_dropped_syn_seen_limit" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_eth13_connection_dropped_table_limit" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_eth14_connection_dropped_syn_seen_limit" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_eth14_connection_dropped_table_limit" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_eth15_connection_dropped_syn_seen_limit" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_eth15_connection_dropped_table_limit" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_eth16_connection_dropped_syn_seen_limit" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_eth16_connection_dropped_table_limit" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_eth17_connection_dropped_syn_seen_limit" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_eth17_connection_dropped_table_limit" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_eth18_connection_dropped_syn_seen_limit" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_eth18_connection_dropped_table_limit" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_eth19_connection_dropped_syn_seen_limit" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_eth19_connection_dropped_table_limit" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_eth1_connection_dropped_syn_seen_limit" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_eth1_connection_dropped_table_limit" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_eth20_connection_dropped_syn_seen_limit" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_eth20_connection_dropped_table_limit" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_eth21_connection_dropped_syn_seen_limit" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_eth21_connection_dropped_table_limit" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_eth22_connection_dropped_syn_seen_limit" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_eth22_connection_dropped_table_limit" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_eth23_connection_dropped_syn_seen_limit" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_eth23_connection_dropped_table_limit" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_eth2_connection_dropped_syn_seen_limit" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_eth2_connection_dropped_table_limit" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_eth3_connection_dropped_syn_seen_limit" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_eth3_connection_dropped_table_limit" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_eth4_connection_dropped_syn_seen_limit" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_eth4_connection_dropped_table_limit" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_eth5_connection_dropped_syn_seen_limit" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_eth5_connection_dropped_table_limit" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_eth6_connection_dropped_syn_seen_limit" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_eth6_connection_dropped_table_limit" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_eth7_connection_dropped_syn_seen_limit" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_eth7_connection_dropped_table_limit" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_eth8_connection_dropped_syn_seen_limit" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_eth8_connection_dropped_table_limit" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_eth9_connection_dropped_syn_seen_limit" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_eth9_connection_dropped_table_limit" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_packet_buffer_allocation_failure" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_packet_dropped_packet_buffer_stressed" :&schema.Schema{
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


