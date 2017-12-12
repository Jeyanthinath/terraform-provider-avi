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
			"apic_all_tenant_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceAPICEpgAllTenantFilterSchema()},
			"apic_cli_login": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceApicCliLoginSchema()},
			"apic_epg_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceAPICEpgFilterSchema()},
			"arp_table_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceArpTableFilterSchema()},
			"candidate_se_host_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceCandidateSeHostFilterSchema()},
			"cc_params": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceGslbSiteOpsConsistencyCheckerSchema()},
			"con_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSeConsumerFilterSchema()},
			"connection_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceConnectionFilterSchema()},
			"connpool_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceConnpoolFilterSchema()},
			"core_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceCoreFilterSchema()},
			"corenum_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceCorenumFilterSchema()},
			//"cps_doser_filter": &schema.Schema{
			//	Type:     schema.TypeSet,
			//	Optional: true,
			//	Set:      func(v interface{}) int { return 0 }, Elem: ResourceCpsDoserFilterSchema()},
			"delete_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceDeleteFilterSchema()},
			//"flowtable_entry_filter": &schema.Schema{
			//	Type:     schema.TypeSet,
			//	Optional: true,
			//	Set:      func(v interface{}) int { return 0 }, Elem: ResourceFlowtableEntryFilterSchema()},
			//"flowtable_intf_filter": &schema.Schema{
			//	Type:     schema.TypeSet,
			//	Optional: true,
			//	Set:      func(v interface{}) int { return 0 }, Elem: ResourceFlowtableIntfFilterSchema()},
			//"geo_location_filter": &schema.Schema{
			//	Type:     schema.TypeSet,
			//	Optional: true,
			//	Set:      func(v interface{}) int { return 0 }, Elem: ResourceGeoLocationFilterSchema()},
			//"glb_params_filter": &schema.Schema{
			//	Type:     schema.TypeSet,
			//	Optional: true,
			//	Set:      func(v interface{}) int { return 0 }, Elem: ResourceGlbParamsFilterSchema()},
			"gs_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceGsFilterSchema()},
			//"gs_params_filter": &schema.Schema{
			//	Type:     schema.TypeSet,
			//	Optional: true,
			//	Set:      func(v interface{}) int { return 0 }, Elem: ResourceGSParamsFilterSchema()},
			"httpcache_obj_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceHttpCacheObjFilterSchema()},
			"listeningsock_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceListeningsockFilterSchema()},
			//"metrics_agent_filter": &schema.Schema{
			//	Type:     schema.TypeSet,
			//	Optional: true,
			//	Set:      func(v interface{}) int { return 0 }, Elem: ResourceMetricsAgentFilterSchema()},
			//"metrics_mgr_filter": &schema.Schema{
			//	Type:     schema.TypeSet,
			//	Optional: true,
			//	Set:      func(v interface{}) int { return 0 }, Elem: ResourceMetricsMgrFilterSchema()},
			"ms_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSeMicroServiceFilterSchema()},
			"nsx_internal_params": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceNsxInternalParamsSchema()},
			"nsx_sg_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceNsxSgFilterSchema()},
			//"persistence_filter": &schema.Schema{
			//	Type:     schema.TypeSet,
			//	Optional: true,
			//	Set:      func(v interface{}) int { return 0 }, Elem: ResourcePersistenceFilterSchema()},
			"placement_status_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourcePlacementStatusFilterSchema()},
			"res_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSeResourceFilterSchema()},
			"se_metrics_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSeMetricsFilterSchema()},
			"se_params_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSeParamsFilterSchema()},
			"server_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceServerFilterSchema()},
			"tcp_stat_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceTcpStatFilterSchema()},
			"udp_stat_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceUdpStatFilterSchema()},
			"vcenter_login": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceVcenterLoginSchema()},
			"vi_datastore_filtler": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceVIDatastoreFilterSchema()},
			"vi_host_resource_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceVIHostResourceFilterSchema()},
			"vi_network_subnet_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceVINetworkSubnetFilterSchema()},
			"vi_redis_datastore_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceVIRedisDatastoreFilterSchema()},
			"vi_retrieve_pg_names": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceVIRetrievePGNamesSchema()},
			"vi_subfolder_filtler": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceVISubfolderFilterSchema()},
			"vip_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSeVipFilterSchema()},
			"vrf_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourcePlacementVrfFilterSchema()},
			"vs_metrics_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceVsMetricsFilterSchema()},
			"vstype_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceVstypeFilterSchema()},
		},
	}
}
