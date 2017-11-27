/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceRPCRequestActionParamsSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "apic_disable_enable_vs_params" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceAPICDisableEnableVsParamsSchema(),                             },
             "apic_txn_params" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceAPICTransactionFlapParamsSchema(),                             },
             "arp_table_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceArpTableFilterSchema(),                             },
             "connection_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceConnectionClearFilterSchema(),                             },
             "connpool_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceConnpoolFilterSchema(),                             },
             "cps_doser_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceCpsDoserFilterSchema(),                             },
             "dispatcher_stat_clear" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceDispatcherStatClearSchema(),                             },
             "dispatcher_table_dump_clear" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceDispatcherTableDumpClearSchema(),                             },
             "event_gen_params" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceEventGenParamsSchema(),                             },
             "ft_entry_delete" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceFlowtableEntryFilterSchema(),                             },
             "gs_ops" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceGslbSiteOpsSchema(),                             },
             "httpcache_obj_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceHttpCacheObjFilterSchema(),                             },
             "nsx_fault_params" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceNsxFaultParamsSchema(),                             },
             "persistence_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourcePersistenceFilterSchema(),                             },
             "pool_server_state" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourcePoolServerStateSchema(),                             },
             "rediscover_vcenter_param" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceRediscoverVcenterParamSchema(),                             },
             "retry_placement_params" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceRetryPlacementParamsSchema(),                             },
             "se_fault_inject_exhaust_param" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSEFaultInjectExhaustParamSchema(),                             },
             "se_switchover_params" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeSwitchoverParamsSchema(),                             },
             "vcenter_login" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceVcenterLoginSchema(),                             },
             "vi_delete_network_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceVIDeleteNetworkFilterSchema(),                             },
             "vi_retrieve_pg_names" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceVIRetrievePGNamesSchema(),                             },
                                "url": &schema.Schema{
                                Type:     schema.TypeString,
                                Optional: true,
                                Computed: true,
                            },
        },
    }
}


