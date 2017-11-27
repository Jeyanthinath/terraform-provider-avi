/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceRPCRequestFilterSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "apic_all_tenant_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceAPICEpgAllTenantFilterSchema(),                             },
             "apic_cli_login" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceApicCliLoginSchema(),                             },
             "apic_epg_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceAPICEpgFilterSchema(),                             },
             "arp_table_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceArpTableFilterSchema(),                             },
             "candidate_se_host_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceCandidateSeHostFilterSchema(),                             },
             "cc_params" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceGslbSiteOpsConsistencyCheckerSchema(),                             },
             "con_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeConsumerFilterSchema(),                             },
             "connection_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceConnectionFilterSchema(),                             },
             "connpool_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceConnpoolFilterSchema(),                             },
             "core_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceCoreFilterSchema(),                             },
             "corenum_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceCorenumFilterSchema(),                             },
             "cps_doser_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceCpsDoserFilterSchema(),                             },
             "delete_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceDeleteFilterSchema(),                             },
             "flowtable_entry_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceFlowtableEntryFilterSchema(),                             },
             "flowtable_intf_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceFlowtableIntfFilterSchema(),                             },
             "geo_location_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceGeoLocationFilterSchema(),                             },
             "glb_params_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceGlbParamsFilterSchema(),                             },
             "gs_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceGsFilterSchema(),                             },
             "gs_params_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceGSParamsFilterSchema(),                             },
             "httpcache_obj_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceHttpCacheObjFilterSchema(),                             },
             "listeningsock_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceListeningsockFilterSchema(),                             },
             "metrics_agent_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceMetricsAgentFilterSchema(),                             },
             "metrics_mgr_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceMetricsMgrFilterSchema(),                             },
             "ms_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeMicroServiceFilterSchema(),                             },
             "nsx_internal_params" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceNsxInternalParamsSchema(),                             },
             "nsx_sg_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceNsxSgFilterSchema(),                             },
             "persistence_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourcePersistenceFilterSchema(),                             },
             "placement_status_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourcePlacementStatusFilterSchema(),                             },
             "res_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeResourceFilterSchema(),                             },
             "se_metrics_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeMetricsFilterSchema(),                             },
             "se_params_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeParamsFilterSchema(),                             },
             "server_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceServerFilterSchema(),                             },
             "tcp_stat_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceTcpStatFilterSchema(),                             },
             "udp_stat_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceUdpStatFilterSchema(),                             },
             "vcenter_login" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceVcenterLoginSchema(),                             },
             "vi_datastore_filtler" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceVIDatastoreFilterSchema(),                             },
             "vi_host_resource_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceVIHostResourceFilterSchema(),                             },
             "vi_network_subnet_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceVINetworkSubnetFilterSchema(),                             },
             "vi_redis_datastore_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceVIRedisDatastoreFilterSchema(),                             },
             "vi_retrieve_pg_names" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceVIRetrievePGNamesSchema(),                             },
             "vi_subfolder_filtler" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceVISubfolderFilterSchema(),                             },
             "vip_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeVipFilterSchema(),                             },
             "vrf_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourcePlacementVrfFilterSchema(),                             },
             "vs_metrics_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceVsMetricsFilterSchema(),                             },
             "vstype_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceVstypeFilterSchema(),                             },
                                "url": &schema.Schema{
                                Type:     schema.TypeString,
                                Optional: true,
                                Computed: true,
                            },
        },
    }
}


