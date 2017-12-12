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
			"apic_disable_enable_vs_params": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceAPICDisableEnableVsParamsSchema()},
			"apic_txn_params": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceAPICTransactionFlapParamsSchema()},
			"arp_table_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceArpTableFilterSchema()},
			"connection_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceConnectionClearFilterSchema()},
			"connpool_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceConnpoolFilterSchema()},
			//"cps_doser_filter": &schema.Schema{
			//	Type:     schema.TypeSet,
			//	Optional: true,
			//	Set:      func(v interface{}) int { return 0 }, Elem: ResourceCpsDoserFilterSchema()},
			//"dispatcher_stat_clear": &schema.Schema{
			//	Type:     schema.TypeSet,
			//	Optional: true,
			//	Set:      func(v interface{}) int { return 0 }, Elem: ResourceDispatcherStatClearSchema()},
			//"dispatcher_table_dump_clear": &schema.Schema{
			//	Type:     schema.TypeSet,
			//	Optional: true,
			//	Set:      func(v interface{}) int { return 0 }, Elem: ResourceDispatcherTableDumpClearSchema()},
			"event_gen_params": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceEventGenParamsSchema()},
			//"ft_entry_delete": &schema.Schema{
			//	Type:     schema.TypeSet,
			//	Optional: true,
			//	Set:      func(v interface{}) int { return 0 }, Elem: ResourceFlowtableEntryFilterSchema()},
			//"gs_ops": &schema.Schema{
			//	Type:     schema.TypeSet,
			//	Optional: true,
			//	Set:      func(v interface{}) int { return 0 }, Elem: ResourceGslbSiteOpsSchema()},
			"httpcache_obj_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceHttpCacheObjFilterSchema()},
			"nsx_fault_params": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceNsxFaultParamsSchema()},
			//"persistence_filter": &schema.Schema{
			//	Type:     schema.TypeSet,
			//	Optional: true,
			//	Set:      func(v interface{}) int { return 0 }, Elem: ResourcePersistenceFilterSchema()},
			"pool_server_state": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourcePoolServerStateSchema()},
			"rediscover_vcenter_param": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceRediscoverVcenterParamSchema()},
			"resync": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceGslbSiteOpsResyncSchema()},
			"retry_placement_params": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceRetryPlacementParamsSchema()},
			"se_fault_inject_exhaust_param": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSEFaultInjectExhaustParamSchema()},
			"se_switchover_params": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSeSwitchoverParamsSchema()},
			"vcenter_login": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceVcenterLoginSchema()},
			"vi_delete_network_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceVIDeleteNetworkFilterSchema()},
			"vi_retrieve_pg_names": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceVIRetrievePGNamesSchema()},
		},
	}
}
