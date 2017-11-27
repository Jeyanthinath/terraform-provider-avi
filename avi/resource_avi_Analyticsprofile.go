/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "log"

    "github.com/avinetworks/sdk/go/clients"
        "github.com/hashicorp/terraform/helper/schema"
)
func ResourceAnalyticsProfileSchema() map[string]*schema.Schema {
    return map[string]*schema.Schema{
             "apdex_response_threshold" :&schema.Schema{
                             Type: schema.TypeInt,            
                             Optional: true,                              
                                                        },
             "apdex_response_tolerated_factor" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "apdex_rtt_threshold" :&schema.Schema{
                             Type: schema.TypeInt,            
                             Optional: true,                              
                                                        },
             "apdex_rtt_tolerated_factor" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "apdex_rum_threshold" :&schema.Schema{
                             Type: schema.TypeInt,            
                             Optional: true,                              
                                                        },
             "apdex_rum_tolerated_factor" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "apdex_server_response_threshold" :&schema.Schema{
                             Type: schema.TypeInt,            
                             Optional: true,                              
                                                        },
             "apdex_server_response_tolerated_factor" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "apdex_server_rtt_threshold" :&schema.Schema{
                             Type: schema.TypeInt,            
                             Optional: true,                              
                                                        },
             "apdex_server_rtt_tolerated_factor" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "client_log_config" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceClientLogConfigurationSchema(),                             },
             "client_log_streaming_config" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceClientLogStreamingConfigSchema(),                             },
             "conn_lossy_ooo_threshold" :&schema.Schema{
                             Type: schema.TypeInt,            
                             Optional: true,                              
                                                        },
             "conn_lossy_timeo_rexmt_threshold" :&schema.Schema{
                             Type: schema.TypeInt,            
                             Optional: true,                              
                                                        },
             "conn_lossy_total_rexmt_threshold" :&schema.Schema{
                             Type: schema.TypeInt,            
                             Optional: true,                              
                                                        },
             "conn_lossy_zero_win_size_event_threshold" :&schema.Schema{
                             Type: schema.TypeInt,            
                             Optional: true,                              
                                                        },
             "conn_server_lossy_ooo_threshold" :&schema.Schema{
                             Type: schema.TypeInt,            
                             Optional: true,                              
                                                        },
             "conn_server_lossy_timeo_rexmt_threshold" :&schema.Schema{
                             Type: schema.TypeInt,            
                             Optional: true,                              
                                                        },
             "conn_server_lossy_total_rexmt_threshold" :&schema.Schema{
                             Type: schema.TypeInt,            
                             Optional: true,                              
                                                        },
             "conn_server_lossy_zero_win_size_event_threshold" :&schema.Schema{
                             Type: schema.TypeInt,            
                             Optional: true,                              
                                                        },
             "description" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "disable_se_analytics" :&schema.Schema{
                             Type: schema.TypeBool,            
                             Optional: true,                              
                                                        },
             "disable_server_analytics" :&schema.Schema{
                             Type: schema.TypeBool,            
                             Optional: true,                              
                                                        },
             "exclude_client_close_before_request_as_error" :&schema.Schema{
                             Type: schema.TypeBool,            
                             Optional: true,                              
                                                        },
             "exclude_dns_policy_drop_as_significant" :&schema.Schema{
                             Type: schema.TypeBool,            
                             Optional: true,                              
                                                        },
             "exclude_gs_down_as_error" :&schema.Schema{
                             Type: schema.TypeBool,            
                             Optional: true,                              
                                                        },
             "exclude_http_error_codes" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                              Elem:&schema.Schema{Type: schema.TypeString},                             },
             "exclude_invalid_dns_domain_as_error" :&schema.Schema{
                             Type: schema.TypeBool,            
                             Optional: true,                              
                                                        },
             "exclude_invalid_dns_query_as_error" :&schema.Schema{
                             Type: schema.TypeBool,            
                             Optional: true,                              
                                                        },
             "exclude_no_dns_record_as_error" :&schema.Schema{
                             Type: schema.TypeBool,            
                             Optional: true,                              
                                                        },
             "exclude_no_valid_gs_member_as_error" :&schema.Schema{
                             Type: schema.TypeBool,            
                             Optional: true,                              
                                                        },
             "exclude_persistence_change_as_error" :&schema.Schema{
                             Type: schema.TypeBool,            
                             Optional: true,                              
                                                        },
             "exclude_server_dns_error_as_error" :&schema.Schema{
                             Type: schema.TypeBool,            
                             Optional: true,                              
                                                        },
             "exclude_server_tcp_reset_as_error" :&schema.Schema{
                             Type: schema.TypeBool,            
                             Optional: true,                              
                                                        },
             "exclude_syn_retransmit_as_error" :&schema.Schema{
                             Type: schema.TypeBool,            
                             Optional: true,                              
                                                        },
             "exclude_tcp_reset_as_error" :&schema.Schema{
                             Type: schema.TypeBool,            
                             Optional: true,                              
                                                        },
             "exclude_unsupported_dns_query_as_error" :&schema.Schema{
                             Type: schema.TypeBool,            
                             Optional: true,                              
                                                        },
             "hs_event_throttle_window" :&schema.Schema{
                             Type: schema.TypeInt,            
                             Optional: true,                              
                                                        },
             "hs_max_anomaly_penalty" :&schema.Schema{
                             Type: schema.TypeInt,            
                             Optional: true,                              
                                                        },
             "hs_max_resources_penalty" :&schema.Schema{
                             Type: schema.TypeInt,            
                             Optional: true,                              
                                                        },
             "hs_max_security_penalty" :&schema.Schema{
                             Type: schema.TypeInt,            
                             Optional: true,                              
                                                        },
             "hs_min_dos_rate" :&schema.Schema{
                             Type: schema.TypeInt,            
                             Optional: true,                              
                                                        },
             "hs_performance_boost" :&schema.Schema{
                             Type: schema.TypeInt,            
                             Optional: true,                              
                                                        },
             "hs_pscore_traffic_threshold_l4_client" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "hs_pscore_traffic_threshold_l4_server" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "hs_security_certscore_expired" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "hs_security_certscore_gt30d" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "hs_security_certscore_le07d" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "hs_security_certscore_le30d" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "hs_security_chain_invalidity_penalty" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "hs_security_cipherscore_eq000b" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "hs_security_cipherscore_ge128b" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "hs_security_cipherscore_lt128b" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "hs_security_encalgo_score_none" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "hs_security_encalgo_score_rc4" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "hs_security_hsts_penalty" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "hs_security_nonpfs_penalty" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "hs_security_selfsignedcert_penalty" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "hs_security_ssl30_score" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "hs_security_tls10_score" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "hs_security_tls11_score" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "hs_security_tls12_score" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "hs_security_weak_signature_algo_penalty" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "name" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "ranges" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceHTTPStatusRangeSchema(),                             },
             "resp_code_block" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                              Elem:&schema.Schema{Type: schema.TypeString},                             },
             "tenant_ref" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                              Elem:&schema.Schema{Type: schema.TypeString},                             },
             "uuid" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                             Computed:  true,  
                                                        },
                                "url": &schema.Schema{
                                Type:     schema.TypeString,
                                Optional: true,
                                Computed: true,
                            },
    }
}


func resourceAviAnalyticsProfile() *schema.Resource {
    return &schema.Resource{
        Create: resourceAviAnalyticsProfileCreate,
        Read:   ResourceAviAnalyticsProfileRead,
        Update: resourceAviAnalyticsProfileUpdate,
        Delete: resourceAviAnalyticsProfileDelete,
        Schema: ResourceAnalyticsProfileSchema(),
    }
}


func ResourceAviAnalyticsProfileRead(d *schema.ResourceData, meta interface{}) error {
    s := ResourceAnalyticsProfileSchema()
    log.Printf("[INFO] ResourceAviAnalyticsProfileRead Avi Client %v\n", d)
    client := meta.(*clients.AviClient)
    var obj interface{}
    if uuid, ok := d.GetOk("uuid"); ok {
        path := "api/analyticsprofile/" + uuid.(string)
        err := client.AviSession.Get(path, &obj)
        if err != nil {
            d.SetId("")
            return nil
        }
    } else {
        d.SetId("")
        return nil
    }
    // no need to set the ID
    log.Printf("ResourceAviAnalyticsProfileRead CURRENT obj %v\n", d)

    log.Printf("ResourceAviAnalyticsProfileRead Read API obj %v\n", obj)
    if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
        log.Printf("[INFO] ResourceAviAnalyticsProfileRead Converted obj %v\n", tObj)
        //err = d.Set("obj", tObj)
        if err != nil {
            log.Printf("[ERROR] in setting read object %v\n", err)
        }
    }
    log.Printf("[INFO] ResourceAviAnalyticsProfileRead Updated %v\n", d)
    return nil
}

func resourceAviAnalyticsProfileCreate(d *schema.ResourceData, meta interface{}) error {
    s := ResourceVrfContextSchema()
    err := ApiCreateOrUpdate(d, meta, "analyticsprofile", s)
    log.Printf("[DEBUG] created object %v: %v", "analyticsprofile", d)
    if err == nil {
        err = ResourceAviVrfContextRead(d, meta)
    }
    log.Printf("[DEBUG] created object %v: %v", "analyticsprofile", d)
    return err
}

func resourceAviAnalyticsProfileUpdate(d *schema.ResourceData, meta interface{}) error {
    s := ResourcePoolSchema()
    err := ApiCreateOrUpdate(d, meta, "analyticsprofile", s)
    if err == nil {
        err = ResourceAviVrfContextRead(d, meta)
    }
    log.Printf("[DEBUG] updated object %v: %v", "analyticsprofile", d)
    return err
}

func resourceAviAnalyticsProfileDelete(d *schema.ResourceData, meta interface{}) error {
    s := ResourcePoolSchema()
    err := ApiCreateOrUpdate(d, meta, "analyticsprofile", s)
    log.Printf("[DEBUG] created object %v: %v", "analyticsprofile", d)
    if err == nil {
        err = ResourceAviVrfContextRead(d, meta)
    }
    log.Printf("[DEBUG] created object %v: %v", "analyticsprofile", d)
    return err
}


