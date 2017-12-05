/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi


import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceServerSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "autoscaling_group_name" :&schema.Schema{
                             Type: schema.TypeString, 
                             Optional: true,
                                                                                                                                            },
             //"availability_zone" :&schema.Schema{
             //                Type: schema.TypeString,
             //                Optional: true,
             //                                                                                                                               },
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
             "enabled" :&schema.Schema{
                             Type: schema.TypeBool, 
                             Optional: true,
                             Default: true,                                                                                                                },
             "external_orchestration_id" :&schema.Schema{
                             Type: schema.TypeString, 
                             Optional: true,
                                                                                                                                            },
             "external_uuid" :&schema.Schema{
                             Type: schema.TypeString, 
                             Optional: true,
                                                                                                                                            },
             "hostname" :&schema.Schema{
                             Type: schema.TypeString, 
                             Optional: true,
                                                                                                                                            },
             "ip" :&schema.Schema{
                             Type: schema.TypeSet, 
                             Required: true,
                                                                                     Elem: ResourceIpAddrSchema(),                             },
             "location" :&schema.Schema{
                             Type: schema.TypeSet, 
                             Optional: true,
                                                                                                                 Elem: ResourceGeoLocationSchema(),                             },
             "mac_address" :&schema.Schema{
                             Type: schema.TypeString, 
                             Optional: true,
                                                                                                                                            },
             "nw_ref" :&schema.Schema{
                             Type: schema.TypeString, 
                             Optional: true,
                                                                                                                  Elem:&schema.Schema{Type: schema.TypeString},                             },
             "port" :&schema.Schema{
                             Type: schema.TypeInt, 
                             Optional: true,
                                                                                                                                            },
             "prst_hdr_val" :&schema.Schema{
                             Type: schema.TypeString, 
                             Optional: true,
                                                                                                                                            },
             "ratio" :&schema.Schema{
                             Type: schema.TypeInt, 
                             Optional: true,
                             Default: 1,
                                                                                                                                            },
             "resolve_server_by_dns" :&schema.Schema{
                             Type: schema.TypeBool, 
                             Optional: true,
                                                                                                                                            },
             "rewrite_host_header" :&schema.Schema{
                             Type: schema.TypeBool, 
                             Optional: true,
                                                                                                                                            },
             "server_node" :&schema.Schema{
                             Type: schema.TypeString, 
                             Optional: true,
                                                                                                                                            },
             "static" :&schema.Schema{
                             Type: schema.TypeBool, 
                             Optional: true,
                                                                                                                                            },
             "verify_network" :&schema.Schema{
                             Type: schema.TypeBool, 
                             Optional: true,
                                                                                                                                            },
             "vm_ref" :&schema.Schema{
                             Type: schema.TypeString, 
                             Optional: true,
                                                                                                                  Elem:&schema.Schema{Type: schema.TypeString},                             },
                                "url": &schema.Schema{
                                Type:     schema.TypeString,
                                Optional: true,
                                Computed: true,
                            },
        },
    }
}


