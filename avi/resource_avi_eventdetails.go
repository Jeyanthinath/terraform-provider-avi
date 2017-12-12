/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"add_networks_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceRmAddNetworksEventDetailsSchema()},
			"all_seupgrade_event_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceAllSeUpgradeEventDetailsSchema()},
			"anomaly_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceAnomalyEventDetailsSchema()},
			"apic_agent_bd_vrf_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceApicAgentBridgeDomainVrfChangeSchema()},
			"apic_agent_generic_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceApicAgentGenericEventDetailsSchema()},
			"apic_agent_vs_network_error": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceApicAgentVsNetworkErrorSchema()},
			"avg_uptime_change_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceAvgUptimeChangeDetailsSchema()},
			"aws_asg_notif_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceAWSASGNotifDetailsSchema()},
			"aws_infra_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceAWSSetupSchema()},
			"azure_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceAzureSetupSchema()},
			"bind_vs_se_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceRmBindVsSeEventDetailsSchema()},
			"bm_infra_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceBMSetupSchema()},
			"bootup_fail_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceRmSeBootupFailEventDetailsSchema()},
			"cc_cluster_vip_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceCloudClusterVipSchema()},
			"cc_dns_update_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceCloudDnsUpdateSchema()},
			"cc_health_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceCloudHealthSchema()},
			"cc_infra_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceCloudGenericSchema()},
			"cc_ip_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceCloudIpChangeSchema()},
			"cc_parkintf_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceCloudVipParkingIntfSchema()},
			"cc_se_vm_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceCloudSeVmChangeSchema()},
			"cc_sync_services_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceCloudSyncServicesSchema()},
			"cc_tenant_del_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceCloudTenantsDeletedSchema()},
			"cc_vip_update_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceCloudVipUpdateSchema()},
			"cc_vnic_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceCloudVnicChangeSchema()},
			"cluster_config_failed_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceClusterConfigFailedEventSchema()},
			"cluster_leader_failover_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceClusterLeaderFailoverEventSchema()},
			"cluster_node_add_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceClusterNodeAddEventSchema()},
			"cluster_node_db_failed_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceClusterNodeDbFailedEventSchema()},
			"cluster_node_remove_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceClusterNodeRemoveEventSchema()},
			"cluster_node_shutdown_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceClusterNodeShutdownEventSchema()},
			"cluster_node_started_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceClusterNodeStartedEventSchema()},
			"cluster_service_critical_failure_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceClusterServiceCriticalFailureEventSchema()},
			"cluster_service_failed_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceClusterServiceFailedEventSchema()},
			"cluster_service_restored_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceClusterServiceRestoredEventSchema()},
			"cluster_warm_reboot_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceClusterWarmRebootEventSchema()},
			"cntlr_host_list_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceVinfraCntlrHostUnreachableListSchema()},
			"config_action_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceConfigActionDetailsSchema()},
			"config_create_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceConfigCreateDetailsSchema()},
			"config_delete_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceConfigDeleteDetailsSchema()},
			"config_password_change_request_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceConfigUserPasswordChangeRequestSchema()},
			"config_update_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceConfigUpdateDetailsSchema()},
			"config_user_authrz_rule_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceConfigUserAuthrzByRuleSchema()},
			"config_user_login_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceConfigUserLoginSchema()},
			"config_user_logout_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceConfigUserLogoutSchema()},
			"config_user_not_authrz_rule_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceConfigUserNotAuthrzByRuleSchema()},
			"container_cloud_setup": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceContainerCloudSetupSchema()},
			"container_cloud_sevice": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceContainerCloudServiceSchema()},
			"cs_infra_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceCloudStackSetupSchema()},
			"delete_se_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceRmDeleteSeEventDetailsSchema()},
			"disable_se_migrate_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceDisableSeMigrateEventDetailsSchema()},
			"disc_summary": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceVinfraDiscSummaryDetailsSchema()},
			"dns_sync_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceDNSVsSyncInfoSchema()},
			"docker_ucp_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceDockerUCPSetupSchema()},
			"dos_attack_event_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceDosAttackEventDetailsSchema()},
			"gcp_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceGCPSetupSchema()},
			"glb_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceGslbStatusSchema()},
			"gs_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceGslbServiceStatusSchema()},
			"host_unavail_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceHostUnavailEventDetailsSchema()},
			"hs_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceHealthScoreDetailsSchema()},
			"ip_fail_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceRmSeIpFailEventDetailsSchema()},
			"license_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceLicenseDetailsSchema()},
			"license_expiry_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceLicenseExpiryDetailsSchema()},
			"marathon_service_port_conflict_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceMarathonServicePortConflictSchema()},
			"mesos_infra_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceMesosSetupSchema()},
			"metric_threshold_up_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceMetricThresoldUpDetailsSchema()},
			"metrics_db_disk_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceMetricsDbDiskEventDetailsSchema()},
			"mgmt_nw_change_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceVinfraMgmtNwChangeDetailsSchema()},
			"modify_networks_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceRmModifyNetworksEventDetailsSchema()},
			"network_subnet_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceNetworkSubnetInfoSchema()},
			"nw_subnet_clash_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceNetworkSubnetClashSchema()},
			"nw_summarized_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSummarizedInfoSchema()},
			"os_infra_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceOpenStackClusterSetupSchema()},
			"os_ip_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceOpenStackIpChangeSchema()},
			"os_lbaudit_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceOpenStackLbProvAuditCheckSchema()},
			"os_lbplugin_op_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceOpenStackLbPluginOpSchema()},
			"os_se_vm_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceOpenStackSeVmChangeSchema()},
			"os_sync_services_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceOpenStackSyncServicesSchema()},
			"os_vnic_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceOpenStackVnicChangeSchema()},
			"pool_deployment_failure_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourcePoolDeploymentFailureInfoSchema()},
			"pool_deployment_success_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourcePoolDeploymentSuccessInfoSchema()},
			"pool_server_delete_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceVinfraPoolServerDeleteDetailsSchema()},
			"rebalance_migrate_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceRebalanceMigrateEventDetailsSchema()},
			"rebalance_scalein_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceRebalanceScaleinEventDetailsSchema()},
			"rebalance_scaleout_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceRebalanceScaleoutEventDetailsSchema()},
			"reboot_se_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceRmRebootSeEventDetailsSchema()},
			"scheduler_action_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSchedulerActionDetailsSchema()},
			"se_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSeMgrEventDetailsSchema()},
			"se_dupip_event_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSeDupipEventDetailsSchema()},
			"se_gateway_heartbeat_failed_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSeGatewayHeartbeatFailedDetailsSchema()},
			"se_gateway_heartbeat_success_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSeGatewayHeartbeatSuccessDetailsSchema()},
			"se_geo_db_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSeGeoDbDetailsSchema()},
			"se_hb_event_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSeHBEventDetailsSchema()},
			"se_hm_gs_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSeHmEventGSDetailsSchema()},
			"se_hm_gsgroup_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSeHmEventGslbPoolDetailsSchema()},
			"se_hm_pool_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSeHmEventPoolDetailsSchema()},
			"se_hm_vs_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSeHmEventVsDetailsSchema()},
			"se_ip_added_event_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSeIpAddedEventDetailsSchema()},
			"se_ip_removed_event_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSeIpRemovedEventDetailsSchema()},
			"se_ipfailure_event_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSeIpfailureEventDetailsSchema()},
			"se_persistence_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSePersistenceEventDetailsSchema()},
			"se_pool_lb_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSePoolLbEventDetailsSchema()},
			"se_thresh_event_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSeThreshEventDetailsSchema()},
			"se_version_check_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSeVersionCheckFailedEventSchema()},
			"se_vnic_down_event_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSeVnicDownEventDetailsSchema()},
			"se_vnic_tx_queue_stall_event_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSeVnicTxQueueStallEventDetailsSchema()},
			"semigrate_event_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSeMigrateEventDetailsSchema()},
			"server_autoscale_failed_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceServerAutoScaleFailedInfoSchema()},
			"server_autoscalein_complete_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceServerAutoScaleInCompleteInfoSchema()},
			"server_autoscalein_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceServerAutoScaleInInfoSchema()},
			"server_autoscaleout_complete_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceServerAutoScaleOutCompleteInfoSchema()},
			"server_autoscaleout_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceServerAutoScaleOutInfoSchema()},
			"seupgrade_disrupted_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSeUpgradeVsDisruptedEventDetailsSchema()},
			"seupgrade_event_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSeUpgradeEventDetailsSchema()},
			"seupgrade_migrate_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSeUpgradeMigrateEventDetailsSchema()},
			"seupgrade_scalein_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSeUpgradeScaleinEventDetailsSchema()},
			"seupgrade_scaleout_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSeUpgradeScaleoutEventDetailsSchema()},
			"spawn_se_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceRmSpawnSeEventDetailsSchema()},
			"ssl_expire_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSSLExpireDetailsSchema()},
			"ssl_export_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSSLExportDetailsSchema()},
			"ssl_renew_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSSLRenewDetailsSchema()},
			"ssl_renew_failed_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSSLRenewFailedDetailsSchema()},
			"switchover_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSwitchoverEventDetailsSchema()},
			"switchover_fail_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSwitchoverFailEventDetailsSchema()},
			"system_upgrade_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSystemUpgradeDetailsSchema()},
			"unbind_vs_se_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceRmUnbindVsSeEventDetailsSchema()},
			"vca_infra_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceVCASetupSchema()},
			"vcenter_connectivity_status": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceVinfraVcenterConnectivityStatusSchema()},
			"vcenter_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceVinfraVcenterBadCredentialsSchema()},
			"vcenter_disc_failure": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceVinfraVcenterDiscoveryFailureSchema()},
			"vcenter_obj_delete_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceVinfraVcenterObjDeleteDetailsSchema()},
			"vip_dns_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceDNSRegisterInfoSchema()},
			"vm_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceVinfraVmDetailsSchema()},
			"vs_awaitingse_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceVsAwaitingSeEventDetailsSchema()},
			"vs_error_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceVsErrorEventDetailsSchema()},
			"vs_fsm_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceVsFsmEventDetailsSchema()},
			"vs_initialplacement_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceVsInitialPlacementEventDetailsSchema()},
			"vs_migrate_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceVsMigrateEventDetailsSchema()},
			"vs_pool_nw_fltr_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceVsPoolNwFilterEventDetailsSchema()},
			"vs_scalein_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceVsScaleInEventDetailsSchema()},
			"vs_scaleout_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceVsScaleOutEventDetailsSchema()},
		},
	}
}
