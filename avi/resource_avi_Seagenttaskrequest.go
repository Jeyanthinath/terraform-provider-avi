/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceSeAgentTaskRequestSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "analyticsprofile_object" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceAnalyticsProfileSchema(),                             },
             "applicationpersistenceprofile_object" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceApplicationPersistenceProfileSchema(),                             },
             "applicationprofile_object" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceApplicationProfileSchema(),                             },
             "authprofile_object" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceAuthProfileSchema(),                             },
             "autoscalelaunchconfig_object" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceAutoScaleLaunchConfigSchema(),                             },
             "cloud_object" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceCloudSchema(),                             },
             "debugserviceengine_object" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceDebugServiceEngineSchema(),                             },
             "dnspolicy_object" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceDnsPolicySchema(),                             },
             "errorpagebody_object" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceErrorPageBodySchema(),                             },
             "errorpageprofile_object" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceErrorPageProfileSchema(),                             },
             "gslb_object" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceGslbSchema(),                             },
             "gslbgeodbprofile_object" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceGslbGeoDbProfileSchema(),                             },
             "gslbservice_object" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceGslbServiceSchema(),                             },
             "gslbserviceruntime_object" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceGslbServiceRuntimeSchema(),                             },
             "hardwaresecuritymodulegroup_object" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceHardwareSecurityModuleGroupSchema(),                             },
             "healthmonitor_object" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceHealthMonitorSchema(),                             },
             "httppolicyset_object" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceHTTPPolicySetSchema(),                             },
             "ipaddrgroup_object" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceIpAddrGroupSchema(),                             },
             "method" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "microservice_object" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceMicroServiceSchema(),                             },
             "microservicegroup_object" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceMicroServiceGroupSchema(),                             },
             "networkprofile_object" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceNetworkProfileSchema(),                             },
             "networksecuritypolicy_object" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceNetworkSecurityPolicySchema(),                             },
             "obj_type" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "pkiprofile_object" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourcePKIProfileSchema(),                             },
             "pool_object" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourcePoolSchema(),                             },
             "poolgroup_object" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourcePoolGroupSchema(),                             },
             "prioritylabels_object" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourcePriorityLabelsSchema(),                             },
             "se_agent_object_uuid" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "seagentvnicrequest_object" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceSeAgentVnicRequestSchema(),                             },
             "seheadlessonlinereq_object" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceSeHeadlessOnlineReqSchema(),                             },
             "seproperties_object" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceSePropertiesSchema(),                             },
             "serverautoscalepolicy_object" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceServerAutoScalePolicySchema(),                             },
             "serverstateupdateinfo_object" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceServerStateUpdateInfoSchema(),                             },
             "serverupdatereq_object" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceServerUpdateReqSchema(),                             },
             "serviceenginegroup_object" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceServiceEngineGroupSchema(),                             },
             "setrolesrequest_object" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceSetRolesRequestSchema(),                             },
             "sslkeyandcertificate_object" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceSSLKeyAndCertificateSchema(),                             },
             "sslprofile_object" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceSSLProfileSchema(),                             },
             "stringgroup_object" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceStringGroupSchema(),                             },
             "tenant_object" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceTenantSchema(),                             },
             "trafficcloneprofile_object" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceTrafficCloneProfileSchema(),                             },
             "virtualservicese_object" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceVirtualServiceSeSchema(),                             },
             "vrfcontext_object" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceVrfContextSchema(),                             },
             "vsdatascriptset_object" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceVSDataScriptSetSchema(),                             },
             "wafpolicy_object" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceWafPolicySchema(),                             },
             "wafprofile_object" :&schema.Schema{
                             Type: schema.TypeList,            
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


