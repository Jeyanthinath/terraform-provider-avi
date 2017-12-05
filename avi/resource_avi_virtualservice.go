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
func ResourceVirtualServiceSchema() map[string]*schema.Schema {
    return map[string]*schema.Schema{
             "active_standby_se_tag" :&schema.Schema{
                             Type: schema.TypeString, 
                             Optional: true,
                             Default: "ACTIVE_STANDBY_SE_1",                                                                                                                },
             "analytics_policy" :&schema.Schema{
                             Type: schema.TypeSet, 
                             Optional: true,
                                                                                                                 Elem: ResourceAnalyticsPolicySchema(),                             },
             "analytics_profile_ref" :&schema.Schema{
                             Type: schema.TypeString, 
                             Optional: true,
                                                                                                                  Elem:&schema.Schema{Type: schema.TypeString},                             },
             "application_profile_ref" :&schema.Schema{
                             Type: schema.TypeString, 
                             Optional: true,
                                                                                                                  Elem:&schema.Schema{Type: schema.TypeString},                             },
             //"auto_allocate_floating_ip" :&schema.Schema{
             //                Type: schema.TypeBool,
             //                Optional: true,
             //                                                                                                                               },
             //"auto_allocate_ip" :&schema.Schema{
             //                Type: schema.TypeBool,
             //                Optional: true,
             //                                                                                                                               },
             //"availability_zone" :&schema.Schema{
             //                Type: schema.TypeString,
             //                Optional: true,
             //                                                                                                                               },
             //"avi_allocated_fip" :&schema.Schema{
             //                Type: schema.TypeBool,
             //                Optional: true,
             //                                                                                                                               },
             //"avi_allocated_vip" :&schema.Schema{
             //                Type: schema.TypeBool,
             //                Optional: true,
             //                                                                                                                               },
             "client_auth" :&schema.Schema{
                             Type: schema.TypeSet, 
                             Optional: true,
                                                                                                                 Elem: ResourceHTTPClientAuthenticationParamsSchema(),                             },
             //"close_client_conn_on_config_update" :&schema.Schema{
             //                Type: schema.TypeBool,
             //                Optional: true,
             //                                                                                                                               },
             "cloud_config_cksum" :&schema.Schema{
                             Type: schema.TypeString, 
                             Optional: true,
                                                                                                                                            },
             //"cloud_ref" :&schema.Schema{
             //                Type: schema.TypeString,
             //                Optional: true,
             //                                                                                                     Elem:&schema.Schema{Type: schema.TypeString},                             },
             "cloud_type" :&schema.Schema{
                             Type: schema.TypeString, 
                             Optional: true,
                             Default: "CLOUD_NONE",                                                                                                                },
             "connections_rate_limit" :&schema.Schema{
                             Type: schema.TypeSet, 
                             Optional: true,
                                                                                                                 Elem: ResourceRateProfileSchema(),                             },
             "content_rewrite" :&schema.Schema{
                             Type: schema.TypeSet, 
                             Optional: true,
                                                                                                                 Elem: ResourceContentRewriteProfileSchema(),                             },
             "created_by" :&schema.Schema{
                             Type: schema.TypeString, 
                             Optional: true,
                                                                                                                                            },
             "delay_fairness" :&schema.Schema{
                             Type: schema.TypeBool, 
                             Optional: true,
                                                                                                                                            },
             "description" :&schema.Schema{
                             Type: schema.TypeString, 
                             Optional: true,
                                                                                                                                            },
             "discovered_network_ref" :&schema.Schema{
                             Type: schema.TypeList, 
                             Optional: true,
                                                                                                                  Elem:&schema.Schema{Type: schema.TypeString},                             },
             "discovered_networks" :&schema.Schema{
                             Type: schema.TypeList, 
                             Optional: true,
                                                                                                                 Elem: ResourceDiscoveredNetworkSchema(),                             },
             "discovered_subnet" :&schema.Schema{
                             Type: schema.TypeList, 
                             Optional: true,
                                                                                                                 Elem: ResourceIpAddrPrefixSchema(),                             },
             "dns_info" :&schema.Schema{
                             Type: schema.TypeList, 
                             Optional: true,
                                                                                                                 Elem: ResourceDnsInfoSchema(),                             },
             "dns_policies" :&schema.Schema{
                             Type: schema.TypeList, 
                             Optional: true,
                                                                                                                 Elem: ResourceDnsPoliciesSchema(),                             },
             "east_west_placement" :&schema.Schema{
                             Type: schema.TypeBool, 
                             Optional: true,
                                                                                                                                            },
             "enable_autogw" :&schema.Schema{
                             Type: schema.TypeBool, 
                             Optional: true,
                             Default: true,                                                                                                                },
             "enable_rhi" :&schema.Schema{
                             Type: schema.TypeBool, 
                             Optional: true,
                                                                                                                                            },
             "enable_rhi_snat" :&schema.Schema{
                             Type: schema.TypeBool, 
                             Optional: true,
                                                                                                                                            },
             "enabled" :&schema.Schema{
                             Type: schema.TypeBool, 
                             Optional: true,
                             Default: true,                                                                                                                },
             //"error_page_profile_ref" :&schema.Schema{
             //                Type: schema.TypeString,
             //                Optional: true,
             //                                                                                                     Elem:&schema.Schema{Type: schema.TypeString},                             },
             "floating_ip" :&schema.Schema{
                             Type: schema.TypeSet, 
                             Optional: true,
                                                                                                                 Elem: ResourceIpAddrSchema(),                             },
             //"floating_subnet_uuid" :&schema.Schema{
             //                Type: schema.TypeString,
             //                Optional: true,
             //                                                                                                                               },
             "flow_dist" :&schema.Schema{
                             Type: schema.TypeString, 
                             Optional: true,
                             Default: "LOAD_AWARE",                                                                                                                },
             "flow_label_type" :&schema.Schema{
                             Type: schema.TypeString, 
                             Optional: true,
                             Default: "NO_LABEL",                                                                                                                },
             "fqdn" :&schema.Schema{
                             Type: schema.TypeString, 
                             Optional: true,
                                                                                                                                            },
             "host_name_xlate" :&schema.Schema{
                             Type: schema.TypeString, 
                             Optional: true,
                                                                                                                                            },
             "http_policies" :&schema.Schema{
                             Type: schema.TypeList, 
                             Optional: true,
                                                                                                                 Elem: ResourceHTTPPoliciesSchema(),                             },
             "ign_pool_net_reach" :&schema.Schema{
                             Type: schema.TypeBool, 
                             Optional: true,
                                                                                                                                            },
             "ip_address" :&schema.Schema{
                             Type: schema.TypeSet, 
                             Optional: true,
                                                                                                                 Elem: ResourceIpAddrSchema(),                             },
             "ipam_network_subnet" :&schema.Schema{
                             Type: schema.TypeSet, 
                             Optional: true,
                                                                                                                 Elem: ResourceIPNetworkSubnetSchema(),                             },
             "limit_doser" :&schema.Schema{
                             Type: schema.TypeBool, 
                             Optional: true,
                                                                                                                                            },
             "max_cps_per_client" :&schema.Schema{
                             Type: schema.TypeInt, 
                             Optional: true,
                                                                                                                                            },
             "microservice_ref" :&schema.Schema{
                             Type: schema.TypeString, 
                             Optional: true,
                                                                                                                  Elem:&schema.Schema{Type: schema.TypeString},                             },
             "name" :&schema.Schema{
                             Type: schema.TypeString, 
                             Required: true,
                                                                                                                },
             "network_profile_ref" :&schema.Schema{
                             Type: schema.TypeString, 
                             Optional: true,
                                                                                                                  Elem:&schema.Schema{Type: schema.TypeString},                             },
             "network_ref" :&schema.Schema{
                             Type: schema.TypeString, 
                             Optional: true,
                                                                                                                  Elem:&schema.Schema{Type: schema.TypeString},                             },
             "network_security_policy_ref" :&schema.Schema{
                             Type: schema.TypeString, 
                             Optional: true,
                                                                                                                  Elem:&schema.Schema{Type: schema.TypeString},                             },
             "nsx_securitygroup" :&schema.Schema{
                             Type: schema.TypeList, 
                             Optional: true,
                                                                                                                  Elem:&schema.Schema{Type: schema.TypeString},                             },
             "performance_limits" :&schema.Schema{
                             Type: schema.TypeSet, 
                             Optional: true,
                                                                                                                 Elem: ResourcePerformanceLimitsSchema(),                             },
             "pool_group_ref" :&schema.Schema{
                             Type: schema.TypeString, 
                             Optional: true,
                                                                                                                  Elem:&schema.Schema{Type: schema.TypeString},                             },
             "pool_ref" :&schema.Schema{
                             Type: schema.TypeString, 
                             Optional: true,
                                                                                                                  Elem:&schema.Schema{Type: schema.TypeString},                             },
             //"port_uuid" :&schema.Schema{
             //                Type: schema.TypeString,
             //                Optional: true,
             //                                                                                                                               },
             "remove_listening_port_on_vs_down" :&schema.Schema{
                             Type: schema.TypeBool, 
                             Optional: true,
                                                                                                                                            },
             "requests_rate_limit" :&schema.Schema{
                             Type: schema.TypeSet, 
                             Optional: true,
                                                                                                                 Elem: ResourceRateProfileSchema(),                             },
             "scaleout_ecmp" :&schema.Schema{
                             Type: schema.TypeBool, 
                             Optional: true,
                                                                                                                                            },
             "se_group_ref" :&schema.Schema{
                             Type: schema.TypeString, 
                             Optional: true,
                                                                                                                  Elem:&schema.Schema{Type: schema.TypeString},                             },
             "server_network_profile_ref" :&schema.Schema{
                             Type: schema.TypeString, 
                             Optional: true,
                                                                                                                  Elem:&schema.Schema{Type: schema.TypeString},                             },
             "service_metadata" :&schema.Schema{
                             Type: schema.TypeString, 
                             Optional: true,
                                                                                                                                            },
             "service_pool_select" :&schema.Schema{
                             Type: schema.TypeList, 
                             Optional: true,
                                                                                                                 Elem: ResourceServicePoolSelectorSchema(),                             },
             "services" :&schema.Schema{
                             Type: schema.TypeList, 
                             Optional: true,
                                                                                                                 Elem: ResourceServiceSchema(),                             },
             "sideband_profile" :&schema.Schema{
                             Type: schema.TypeSet, 
                             Optional: true,
                                                                                                                 Elem: ResourceSidebandProfileSchema(),                             },
             "snat_ip" :&schema.Schema{
                             Type: schema.TypeList, 
                             Optional: true,
                                                                                                                 Elem: ResourceIpAddrSchema(),                             },
             "sp_pool_refs" :&schema.Schema{
                             Type: schema.TypeList, 
                             Optional: true,
                                                                                                                  Elem:&schema.Schema{Type: schema.TypeString},                             },
             "ssl_key_and_certificate_refs" :&schema.Schema{
                             Type: schema.TypeList, 
                             Optional: true,
                                                                                                                  Elem:&schema.Schema{Type: schema.TypeString},                             },
             "ssl_profile_ref" :&schema.Schema{
                             Type: schema.TypeString, 
                             Optional: true,
                                                                                                                  Elem:&schema.Schema{Type: schema.TypeString},                             },
             "ssl_sess_cache_avg_size" :&schema.Schema{
                             Type: schema.TypeInt, 
                             Optional: true,
                             Default: 1024,
                                                                                                                                            },
             "static_dns_records" :&schema.Schema{
                             Type: schema.TypeList, 
                             Optional: true,
                                                                                                                 Elem: ResourceDnsRecordSchema(),                             },
             "subnet" :&schema.Schema{
                             Type: schema.TypeSet, 
                             Optional: true,
                                                                                                                 Elem: ResourceIpAddrPrefixSchema(),                             },
             //"subnet_uuid" :&schema.Schema{
             //                Type: schema.TypeString,
             //                Optional: true,
             //                                                                                                                               },
             "tenant_ref" :&schema.Schema{
                             Type: schema.TypeString, 
                             Optional: true,
                                                                                                                  Elem:&schema.Schema{Type: schema.TypeString},                             },
             "traffic_clone_profile_ref" :&schema.Schema{
                             Type: schema.TypeString, 
                             Optional: true,
                                                                                                                  Elem:&schema.Schema{Type: schema.TypeString},                             },
             "type" :&schema.Schema{
                             Type: schema.TypeString, 
                             Optional: true,
                             Default: "VS_TYPE_NORMAL",                                                                                                                },
             "use_bridge_ip_as_vip" :&schema.Schema{
                             Type: schema.TypeBool, 
                             Optional: true,
                                                                                                                                            },
             //"use_vip_as_snat" :&schema.Schema{
             //                Type: schema.TypeBool,
             //                Optional: true,
             //                                                                                                                               },
             "uuid" :&schema.Schema{
                             Type: schema.TypeString, 
                             Optional: true,
                                                                                    Computed:  true,                                                         },
             "vh_domain_name" :&schema.Schema{
                             Type: schema.TypeList, 
                             Optional: true,
                                                                                                                  Elem:&schema.Schema{Type: schema.TypeString},                             },
             "vh_parent_vs_uuid" :&schema.Schema{
                             Type: schema.TypeString, 
                             Optional: true,
                                                                                                                                            },
             "vip" :&schema.Schema{
                             Type: schema.TypeList, 
                             Optional: true,
                                                                                                                 Elem: ResourceVipSchema(),                             },
             "vrf_context_ref" :&schema.Schema{
                             Type: schema.TypeString, 
                             Optional: true,
                                                                                                                  Elem:&schema.Schema{Type: schema.TypeString},                             },
             "vs_datascripts" :&schema.Schema{
                             Type: schema.TypeList, 
                             Optional: true,
                                                                                                                 Elem: ResourceVSDataScriptsSchema(),                             },
             "vsvip_ref" :&schema.Schema{
                             Type: schema.TypeString, 
                             Optional: true,
                                                                                                                  Elem:&schema.Schema{Type: schema.TypeString},                             },
             //"waf_policy_ref" :&schema.Schema{
             //                Type: schema.TypeString,
             //                Optional: true,
             //                                                                                                     Elem:&schema.Schema{Type: schema.TypeString},                             },
             "weight" :&schema.Schema{
                             Type: schema.TypeInt, 
                             Optional: true,
                             Default: 1,
                                                                                                                                            },
                                "url": &schema.Schema{
                                Type:     schema.TypeString,
                                Optional: true,
                                Computed: true,
                            },
    }
}


func resourceAviVirtualService() *schema.Resource {
    return &schema.Resource{
        Create: resourceAviVirtualServiceCreate,
        Read:   ResourceAviVirtualServiceRead,
        Update: resourceAviVirtualServiceUpdate,
        Delete: resourceAviVirtualServiceDelete,
        Schema: ResourceVirtualServiceSchema(),
    }
}


func ResourceAviVirtualServiceRead(d *schema.ResourceData, meta interface{}) error {
    s := ResourceVirtualServiceSchema()
    log.Printf("[INFO] ResourceAviVirtualServiceRead Avi Client %v\n", d)
    client := meta.(*clients.AviClient)
    var obj interface{}
    if uuid, ok := d.GetOk("uuid"); ok {
        path := "api/virtualservice/" + uuid.(string)
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
    log.Printf("ResourceAviVirtualServiceRead CURRENT obj %v\n", d)

    log.Printf("ResourceAviVirtualServiceRead Read API obj %v\n", obj)
    if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
        log.Printf("[INFO] ResourceAviVirtualServiceRead Converted obj %v\n", tObj)
        //err = d.Set("obj", tObj)
        if err != nil {
            log.Printf("[ERROR] in setting read object %v\n", err)
        }
    }
    log.Printf("[INFO] ResourceAviVirtualServiceRead Updated %v\n", d)
    return nil
}

func resourceAviVirtualServiceCreate(d *schema.ResourceData, meta interface{}) error {
    s := ResourceVirtualServiceSchema()
    err := ApiCreateOrUpdate(d, meta, "virtualservice", s)
    log.Printf("[DEBUG] created object %v: %v", "virtualservice", d)
    if err == nil {
        err = ResourceAviVirtualServiceRead(d, meta)
    }
    log.Printf("[DEBUG] created object %v: %v", "virtualservice", d)
    return err
}

func resourceAviVirtualServiceUpdate(d *schema.ResourceData, meta interface{}) error {
    s := ResourceVirtualServiceSchema()
    err := ApiCreateOrUpdate(d, meta, "virtualservice", s)
    if err == nil {
        err = ResourceAviVirtualServiceRead(d, meta)
    }
    log.Printf("[DEBUG] updated object %v: %v", "virtualservice", d)
    return err
}

func resourceAviVirtualServiceDelete(d *schema.ResourceData, meta interface{}) error {
    s := ResourceVirtualServiceSchema()
    err := ApiCreateOrUpdate(d, meta, "virtualservice", s)
    log.Printf("[DEBUG] created object %v: %v", "virtualservice", d)
    if err == nil {
        err = ResourceAviVirtualServiceRead(d, meta)
    }
    log.Printf("[DEBUG] created object %v: %v", "virtualservice", d)
    return err
}


