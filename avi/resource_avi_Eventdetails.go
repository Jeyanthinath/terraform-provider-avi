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
             "add_networks_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceRmAddNetworksEventDetailsSchema(),                             },
             "all_seupgrade_event_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceAllSeUpgradeEventDetailsSchema(),                             },
             "anomaly_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceAnomalyEventDetailsSchema(),                             },
             "apic_agent_bd_vrf_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceApicAgentBridgeDomainVrfChangeSchema(),                             },
             "apic_agent_generic_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceApicAgentGenericEventDetailsSchema(),                             },
             "apic_agent_vs_network_error" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceApicAgentVsNetworkErrorSchema(),                             },
             "avg_uptime_change_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceAvgUptimeChangeDetailsSchema(),                             },
             "aws_asg_notif_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceAWSASGNotifDetailsSchema(),                             },
             "aws_infra_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceAWSSetupSchema(),                             },
             "azure_info" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceAzureSetupSchema(),                             },
             "bind_vs_se_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceRmBindVsSeEventDetailsSchema(),                             },
             "bm_infra_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceBMSetupSchema(),                             },
             "bootup_fail_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceRmSeBootupFailEventDetailsSchema(),                             },
             "cc_cluster_vip_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceCloudClusterVipSchema(),                             },
             "cc_dns_update_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceCloudDnsUpdateSchema(),                             },
             "cc_health_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceCloudHealthSchema(),                             },
             "cc_infra_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceCloudGenericSchema(),                             },
             "cc_ip_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceCloudIpChangeSchema(),                             },
             "cc_parkintf_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceCloudVipParkingIntfSchema(),                             },
             "cc_se_vm_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceCloudSeVmChangeSchema(),                             },
             "cc_sync_services_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceCloudSyncServicesSchema(),                             },
             "cc_tenant_del_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceCloudTenantsDeletedSchema(),                             },
             "cc_vip_update_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceCloudVipUpdateSchema(),                             },
             "cc_vnic_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceCloudVnicChangeSchema(),                             },
             "cluster_config_failed_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceClusterConfigFailedEventSchema(),                             },
             "cluster_leader_failover_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceClusterLeaderFailoverEventSchema(),                             },
             "cluster_node_add_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceClusterNodeAddEventSchema(),                             },
             "cluster_node_db_failed_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceClusterNodeDbFailedEventSchema(),                             },
             "cluster_node_remove_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceClusterNodeRemoveEventSchema(),                             },
             "cluster_node_shutdown_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceClusterNodeShutdownEventSchema(),                             },
             "cluster_node_started_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceClusterNodeStartedEventSchema(),                             },
             "cluster_service_critical_failure_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceClusterServiceCriticalFailureEventSchema(),                             },
             "cluster_service_failed_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceClusterServiceFailedEventSchema(),                             },
             "cluster_service_restored_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceClusterServiceRestoredEventSchema(),                             },
             "cluster_warm_reboot_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceClusterWarmRebootEventSchema(),                             },
             "cntlr_host_list_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceVinfraCntlrHostUnreachableListSchema(),                             },
             "config_action_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceConfigActionDetailsSchema(),                             },
             "config_create_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceConfigCreateDetailsSchema(),                             },
             "config_delete_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceConfigDeleteDetailsSchema(),                             },
             "config_password_change_request_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceConfigUserPasswordChangeRequestSchema(),                             },
             "config_update_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceConfigUpdateDetailsSchema(),                             },
             "config_user_authrz_rule_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceConfigUserAuthrzByRuleSchema(),                             },
             "config_user_login_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceConfigUserLoginSchema(),                             },
             "config_user_logout_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceConfigUserLogoutSchema(),                             },
             "config_user_not_authrz_rule_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceConfigUserNotAuthrzByRuleSchema(),                             },
             "container_cloud_setup" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceContainerCloudSetupSchema(),                             },
             "container_cloud_sevice" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceContainerCloudServiceSchema(),                             },
             "cs_infra_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceCloudStackSetupSchema(),                             },
             "delete_se_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceRmDeleteSeEventDetailsSchema(),                             },
             "disable_se_migrate_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceDisableSeMigrateEventDetailsSchema(),                             },
             "disc_summary" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceVinfraDiscSummaryDetailsSchema(),                             },
             "dns_sync_info" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceDNSVsSyncInfoSchema(),                             },
             "docker_ucp_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceDockerUCPSetupSchema(),                             },
             "dos_attack_event_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceDosAttackEventDetailsSchema(),                             },
             "gcp_info" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceGCPSetupSchema(),                             },
             "glb_info" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceGslbStatusSchema(),                             },
             "gs_info" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceGslbServiceStatusSchema(),                             },
             "host_unavail_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceHostUnavailEventDetailsSchema(),                             },
             "hs_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceHealthScoreDetailsSchema(),                             },
             "ip_fail_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceRmSeIpFailEventDetailsSchema(),                             },
             "license_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceLicenseDetailsSchema(),                             },
             "license_expiry_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceLicenseExpiryDetailsSchema(),                             },
             "marathon_service_port_conflict_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceMarathonServicePortConflictSchema(),                             },
             "mesos_infra_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceMesosSetupSchema(),                             },
             "metric_threshold_up_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceMetricThresoldUpDetailsSchema(),                             },
             "metrics_db_disk_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceMetricsDbDiskEventDetailsSchema(),                             },
             "mgmt_nw_change_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceVinfraMgmtNwChangeDetailsSchema(),                             },
             "modify_networks_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceRmModifyNetworksEventDetailsSchema(),                             },
             "network_subnet_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceNetworkSubnetInfoSchema(),                             },
             "nw_subnet_clash_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceNetworkSubnetClashSchema(),                             },
             "nw_summarized_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSummarizedInfoSchema(),                             },
             "os_infra_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceOpenStackClusterSetupSchema(),                             },
             "os_ip_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceOpenStackIpChangeSchema(),                             },
             "os_lbaudit_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceOpenStackLbProvAuditCheckSchema(),                             },
             "os_lbplugin_op_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceOpenStackLbPluginOpSchema(),                             },
             "os_se_vm_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceOpenStackSeVmChangeSchema(),                             },
             "os_sync_services_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceOpenStackSyncServicesSchema(),                             },
             "os_vnic_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceOpenStackVnicChangeSchema(),                             },
             "pool_deployment_failure_info" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourcePoolDeploymentFailureInfoSchema(),                             },
             "pool_deployment_success_info" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourcePoolDeploymentSuccessInfoSchema(),                             },
             "pool_server_delete_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceVinfraPoolServerDeleteDetailsSchema(),                             },
             "rebalance_migrate_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceRebalanceMigrateEventDetailsSchema(),                             },
             "rebalance_scalein_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceRebalanceScaleinEventDetailsSchema(),                             },
             "rebalance_scaleout_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceRebalanceScaleoutEventDetailsSchema(),                             },
             "reboot_se_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceRmRebootSeEventDetailsSchema(),                             },
             "scheduler_action_info" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSchedulerActionDetailsSchema(),                             },
             "se_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeMgrEventDetailsSchema(),                             },
             "se_dupip_event_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeDupipEventDetailsSchema(),                             },
             "se_gateway_heartbeat_failed_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeGatewayHeartbeatFailedDetailsSchema(),                             },
             "se_gateway_heartbeat_success_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeGatewayHeartbeatSuccessDetailsSchema(),                             },
             "se_geo_db_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeGeoDbDetailsSchema(),                             },
             "se_hb_event_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeHBEventDetailsSchema(),                             },
             "se_hm_gs_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeHmEventGSDetailsSchema(),                             },
             "se_hm_gsgroup_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeHmEventGslbPoolDetailsSchema(),                             },
             "se_hm_pool_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeHmEventPoolDetailsSchema(),                             },
             "se_hm_vs_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeHmEventVsDetailsSchema(),                             },
             "se_ip_added_event_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeIpAddedEventDetailsSchema(),                             },
             "se_ip_removed_event_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeIpRemovedEventDetailsSchema(),                             },
             "se_ipfailure_event_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeIpfailureEventDetailsSchema(),                             },
             "se_persistence_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSePersistenceEventDetailsSchema(),                             },
             "se_pool_lb_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSePoolLbEventDetailsSchema(),                             },
             "se_thresh_event_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeThreshEventDetailsSchema(),                             },
             "se_version_check_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeVersionCheckFailedEventSchema(),                             },
             "se_vnic_down_event_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeVnicDownEventDetailsSchema(),                             },
             "se_vnic_tx_queue_stall_event_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeVnicTxQueueStallEventDetailsSchema(),                             },
             "semigrate_event_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeMigrateEventDetailsSchema(),                             },
             "server_autoscale_failed_info" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceServerAutoScaleFailedInfoSchema(),                             },
             "server_autoscalein_complete_info" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceServerAutoScaleInCompleteInfoSchema(),                             },
             "server_autoscalein_info" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceServerAutoScaleInInfoSchema(),                             },
             "server_autoscaleout_complete_info" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceServerAutoScaleOutCompleteInfoSchema(),                             },
             "server_autoscaleout_info" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceServerAutoScaleOutInfoSchema(),                             },
             "seupgrade_disrupted_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeUpgradeVsDisruptedEventDetailsSchema(),                             },
             "seupgrade_event_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeUpgradeEventDetailsSchema(),                             },
             "seupgrade_migrate_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeUpgradeMigrateEventDetailsSchema(),                             },
             "seupgrade_scalein_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeUpgradeScaleinEventDetailsSchema(),                             },
             "seupgrade_scaleout_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeUpgradeScaleoutEventDetailsSchema(),                             },
             "spawn_se_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceRmSpawnSeEventDetailsSchema(),                             },
             "ssl_expire_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSSLExpireDetailsSchema(),                             },
             "ssl_export_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSSLExportDetailsSchema(),                             },
             "ssl_renew_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSSLRenewDetailsSchema(),                             },
             "ssl_renew_failed_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSSLRenewFailedDetailsSchema(),                             },
             "switchover_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSwitchoverEventDetailsSchema(),                             },
             "switchover_fail_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSwitchoverFailEventDetailsSchema(),                             },
             "system_upgrade_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSystemUpgradeDetailsSchema(),                             },
             "unbind_vs_se_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceRmUnbindVsSeEventDetailsSchema(),                             },
             "vca_infra_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceVCASetupSchema(),                             },
             "vcenter_connectivity_status" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceVinfraVcenterConnectivityStatusSchema(),                             },
             "vcenter_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceVinfraVcenterBadCredentialsSchema(),                             },
             "vcenter_disc_failure" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceVinfraVcenterDiscoveryFailureSchema(),                             },
             "vcenter_obj_delete_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceVinfraVcenterObjDeleteDetailsSchema(),                             },
             "vip_dns_info" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceDNSRegisterInfoSchema(),                             },
             "vm_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceVinfraVmDetailsSchema(),                             },
             "vs_awaitingse_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceVsAwaitingSeEventDetailsSchema(),                             },
             "vs_error_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceVsErrorEventDetailsSchema(),                             },
             "vs_fsm_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceVsFsmEventDetailsSchema(),                             },
             "vs_initialplacement_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceVsInitialPlacementEventDetailsSchema(),                             },
             "vs_migrate_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceVsMigrateEventDetailsSchema(),                             },
             "vs_pool_nw_fltr_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceVsPoolNwFilterEventDetailsSchema(),                             },
             "vs_scalein_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceVsScaleInEventDetailsSchema(),                             },
             "vs_scaleout_details" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceVsScaleOutEventDetailsSchema(),                             },
                                "url": &schema.Schema{
                                Type:     schema.TypeString,
                                Optional: true,
                                Computed: true,
                            },
        },
    }
}


