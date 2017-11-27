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
func ResourcePoolSchema() map[string]*schema.Schema {
    return map[string]*schema.Schema{
             "a_pool" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "ab_pool" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceAbPoolSchema(),                             },
             "ab_priority" :&schema.Schema{
                             Type: schema.TypeInt,            
                             Optional: true,                              
                                                        },
             "apic_epg_name" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "application_persistence_profile_ref" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                              Elem:&schema.Schema{Type: schema.TypeString},                             },
             "autoscale_launch_config_ref" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                              Elem:&schema.Schema{Type: schema.TypeString},                             },
             "autoscale_networks" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                              Elem:&schema.Schema{Type: schema.TypeString},                             },
             "autoscale_policy_ref" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                              Elem:&schema.Schema{Type: schema.TypeString},                             },
             "capacity_estimation" :&schema.Schema{
                             Type: schema.TypeBool,            
                             Optional: true,                              
                                                        },
             "capacity_estimation_ttfb_thresh" :&schema.Schema{
                             Type: schema.TypeInt,            
                             Optional: true,                              
                                                        },
             "cloud_config_cksum" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "cloud_ref" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                              Elem:&schema.Schema{Type: schema.TypeString},                             },
             "connection_ramp_duration" :&schema.Schema{
                             Type: schema.TypeInt,            
                             Optional: true,                              
                                                        },
             "created_by" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "default_server_port" :&schema.Schema{
                             Type: schema.TypeInt,            
                             Optional: true,                              
                                                        },
             "description" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "domain_name" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                              Elem:&schema.Schema{Type: schema.TypeString},                             },
             "east_west" :&schema.Schema{
                             Type: schema.TypeBool,            
                             Optional: true,                              
                                                        },
             "enabled" :&schema.Schema{
                             Type: schema.TypeBool,            
                             Optional: true,                              
                                                        },
             "external_autoscale_groups" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                              Elem:&schema.Schema{Type: schema.TypeString},                             },
             "fail_action" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceFailActionSchema(),                             },
             "fewest_tasks_feedback_delay" :&schema.Schema{
                             Type: schema.TypeInt,            
                             Optional: true,                              
                                                        },
             "graceful_disable_timeout" :&schema.Schema{
                             Type: schema.TypeInt,            
                             Optional: true,                              
                                                        },
             "gslb_sp_enabled" :&schema.Schema{
                             Type: schema.TypeBool,            
                             Optional: true,                              
                                                        },
             "health_monitor_refs" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                              Elem:&schema.Schema{Type: schema.TypeString},                             },
             "host_check_enabled" :&schema.Schema{
                             Type: schema.TypeBool,            
                             Optional: true,                              
                                                        },
             "inline_health_monitor" :&schema.Schema{
                             Type: schema.TypeBool,            
                             Optional: true,                              
                                                        },
             "ipaddrgroup_ref" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                              Elem:&schema.Schema{Type: schema.TypeString},                             },
             "lb_algorithm" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "lb_algorithm_consistent_hash_hdr" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "lb_algorithm_core_nonaffinity" :&schema.Schema{
                             Type: schema.TypeInt,            
                             Optional: true,                              
                                                        },
             "lb_algorithm_hash" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "lookup_server_by_name" :&schema.Schema{
                             Type: schema.TypeBool,            
                             Optional: true,                              
                                                        },
             "max_concurrent_connections_per_server" :&schema.Schema{
                             Type: schema.TypeInt,            
                             Optional: true,                              
                                                        },
             "max_conn_rate_per_server" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceRateProfileSchema(),                             },
             "name" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "networks" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceNetworkFilterSchema(),                             },
             "nsx_securitygroup" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                              Elem:&schema.Schema{Type: schema.TypeString},                             },
             "pki_profile_ref" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                              Elem:&schema.Schema{Type: schema.TypeString},                             },
             "placement_networks" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourcePlacementNetworkSchema(),                             },
             "prst_hdr_name" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "request_queue_depth" :&schema.Schema{
                             Type: schema.TypeInt,            
                             Optional: true,                              
                                                        },
             "request_queue_enabled" :&schema.Schema{
                             Type: schema.TypeBool,            
                             Optional: true,                              
                                                        },
             "rewrite_host_header_to_server_name" :&schema.Schema{
                             Type: schema.TypeBool,            
                             Optional: true,                              
                                                        },
             "rewrite_host_header_to_sni" :&schema.Schema{
                             Type: schema.TypeBool,            
                             Optional: true,                              
                                                        },
             "server_auto_scale" :&schema.Schema{
                             Type: schema.TypeBool,            
                             Optional: true,                              
                                                        },
             "server_count" :&schema.Schema{
                             Type: schema.TypeInt,            
                             Optional: true,                              
                                                        },
             "server_name" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "server_reselect" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceHTTPServerReselectSchema(),                             },
             "servers" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceServerSchema(),                             },
             "sni_enabled" :&schema.Schema{
                             Type: schema.TypeBool,            
                             Optional: true,                              
                                                        },
             "ssl_key_and_certificate_ref" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                              Elem:&schema.Schema{Type: schema.TypeString},                             },
             "ssl_profile_ref" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                              Elem:&schema.Schema{Type: schema.TypeString},                             },
             "tenant_ref" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                              Elem:&schema.Schema{Type: schema.TypeString},                             },
             "use_service_port" :&schema.Schema{
                             Type: schema.TypeBool,            
                             Optional: true,                              
                                                        },
             "uuid" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                             Computed:  true,  
                                                        },
             "vrf_ref" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                              Elem:&schema.Schema{Type: schema.TypeString},                             },
                                "url": &schema.Schema{
                                Type:     schema.TypeString,
                                Optional: true,
                                Computed: true,
                            },
    }
}


func resourceAviPool() *schema.Resource {
    return &schema.Resource{
        Create: resourceAviPoolCreate,
        Read:   ResourceAviPoolRead,
        Update: resourceAviPoolUpdate,
        Delete: resourceAviPoolDelete,
        Schema: ResourcePoolSchema(),
    }
}


func ResourceAviPoolRead(d *schema.ResourceData, meta interface{}) error {
    s := ResourcePoolSchema()
    log.Printf("[INFO] ResourceAviPoolRead Avi Client %v\n", d)
    client := meta.(*clients.AviClient)
    var obj interface{}
    if uuid, ok := d.GetOk("uuid"); ok {
        path := "api/pool/" + uuid.(string)
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
    log.Printf("ResourceAviPoolRead CURRENT obj %v\n", d)

    log.Printf("ResourceAviPoolRead Read API obj %v\n", obj)
    if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
        log.Printf("[INFO] ResourceAviPoolRead Converted obj %v\n", tObj)
        //err = d.Set("obj", tObj)
        if err != nil {
            log.Printf("[ERROR] in setting read object %v\n", err)
        }
    }
    log.Printf("[INFO] ResourceAviPoolRead Updated %v\n", d)
    return nil
}

func resourceAviPoolCreate(d *schema.ResourceData, meta interface{}) error {
    s := ResourceVrfContextSchema()
    err := ApiCreateOrUpdate(d, meta, "pool", s)
    log.Printf("[DEBUG] created object %v: %v", "pool", d)
    if err == nil {
        err = ResourceAviVrfContextRead(d, meta)
    }
    log.Printf("[DEBUG] created object %v: %v", "pool", d)
    return err
}

func resourceAviPoolUpdate(d *schema.ResourceData, meta interface{}) error {
    s := ResourcePoolSchema()
    err := ApiCreateOrUpdate(d, meta, "pool", s)
    if err == nil {
        err = ResourceAviVrfContextRead(d, meta)
    }
    log.Printf("[DEBUG] updated object %v: %v", "pool", d)
    return err
}

func resourceAviPoolDelete(d *schema.ResourceData, meta interface{}) error {
    s := ResourcePoolSchema()
    err := ApiCreateOrUpdate(d, meta, "pool", s)
    log.Printf("[DEBUG] created object %v: %v", "pool", d)
    if err == nil {
        err = ResourceAviVrfContextRead(d, meta)
    }
    log.Printf("[DEBUG] created object %v: %v", "pool", d)
    return err
}

