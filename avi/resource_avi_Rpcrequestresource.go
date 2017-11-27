/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceRPCRequestResourceSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "analytics_profile" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceAnalyticsProfileSchema(),                             },
             "apic_configuration" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceAPICConfigurationSchema(),                             },
             "apic_transaction" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceAPICTransactionSchema(),                             },
             "application_persistence_profile" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceApplicationPersistenceProfileSchema(),                             },
             "application_profile" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceApplicationProfileSchema(),                             },
             "auth_profile" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceAuthProfileSchema(),                             },
             "auto_scale_launch_config" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceAutoScaleLaunchConfigSchema(),                             },
             "cloud" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceCloudSchema(),                             },
             "cloud_props" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceCloudPropertiesSchema(),                             },
             "controller_props" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceControllerPropertiesSchema(),                             },
             "debug_controller" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceDebugControllerSchema(),                             },
             "debug_service_engine" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceDebugServiceEngineSchema(),                             },
             "dns_policy" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceDnsPolicySchema(),                             },
             "errorpage_body" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceErrorPageBodySchema(),                             },
             "errorpage_profile" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceErrorPageProfileSchema(),                             },
             "gs_ops" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceGslbSiteOpsSchema(),                             },
             "gslb_dns_geo_update" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceGslbDnsGeoUpdateSchema(),                             },
             "gslb_dns_update" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceGslbDnsUpdateSchema(),                             },
             "gslb_geodb" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceGslbGeoDbProfileSchema(),                             },
             "gslb_service" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceGslbServiceSchema(),                             },
             "gslb_service_runtime" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceGslbServiceRuntimeSchema(),                             },
             "hardwaresecuritymodule_group" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceHardwareSecurityModuleGroupSchema(),                             },
             "health_monitor" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceHealthMonitorSchema(),                             },
             "http_policy_set" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceHTTPPolicySetSchema(),                             },
             "ip_addr_group" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceIpAddrGroupSchema(),                             },
             "ipam_dns_records" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceDnsRecordSchema(),                             },
             "microservice" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceMicroServiceSchema(),                             },
             "microservice_group" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceMicroServiceGroupSchema(),                             },
             "network" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceNetworkSchema(),                             },
             "network_profile" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceNetworkProfileSchema(),                             },
             "network_security_policy" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceNetworkSecurityPolicySchema(),                             },
             "pki_profile" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourcePKIProfileSchema(),                             },
             "pool" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourcePoolSchema(),                             },
             "pool_group" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourcePoolGroupSchema(),                             },
             "pool_group_deployment_policy" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourcePoolGroupDeploymentPolicySchema(),                             },
             "priority_labels" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourcePriorityLabelsSchema(),                             },
             "se_props" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSePropertiesSchema(),                             },
             "server_auto_scale_policy" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceServerAutoScalePolicySchema(),                             },
             "service_engine" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceServiceEngineSchema(),                             },
             "service_engine_group" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceServiceEngineGroupSchema(),                             },
             "ssl_key_and_certificate" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSSLKeyAndCertificateSchema(),                             },
             "ssl_profile" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSSLProfileSchema(),                             },
             "string_group" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceStringGroupSchema(),                             },
             "system_configuration" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSystemConfigurationSchema(),                             },
             "traffic_clone_profile" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceTrafficCloneProfileSchema(),                             },
             "vrf_context" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceVrfContextSchema(),                             },
             "vsdatascriptset" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceVSDataScriptSetSchema(),                             },
             "waf_policy" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceWafPolicySchema(),                             },
             "waf_profile" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceWafProfileSchema(),                             },
                                "url": &schema.Schema{
                                Type:     schema.TypeString,
                                Optional: true,
                                Computed: true,
                            },
        },
    }
}


