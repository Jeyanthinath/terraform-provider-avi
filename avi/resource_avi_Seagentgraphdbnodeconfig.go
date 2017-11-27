/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceSeAgentGraphDBNodeConfigSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "analytics_profile" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceAnalyticsProfileSchema(),                             },
             "application_persistence_profile" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceApplicationPersistenceProfileSchema(),                             },
             "application_profile" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceApplicationProfileSchema(),                             },
             "cloud" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceCloudSchema(),                             },
             "gslb" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceGslbSchema(),                             },
             "gslbgeodbprofile" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceGslbGeoDbProfileSchema(),                             },
             "gslbservice" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceGslbServiceSchema(),                             },
             "health_monitor" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceHealthMonitorSchema(),                             },
             "http_request_policy" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceHTTPRequestPolicySchema(),                             },
             "http_response_policy" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceHTTPResponsePolicySchema(),                             },
             "http_security_policy" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceHTTPSecurityPolicySchema(),                             },
             "ip_addr_group" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceIpAddrGroupSchema(),                             },
             "microservice" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceMicroServiceSchema(),                             },
             "network_profile" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceNetworkProfileSchema(),                             },
             "network_security_policy" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceNetworkSecurityPolicySchema(),                             },
             "pool" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourcePoolSchema(),                             },
             "pool_group" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourcePoolGroupSchema(),                             },
             "priority_labels" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourcePriorityLabelsSchema(),                             },
             "serviceenginegroup" :&schema.Schema{
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
             "tenant" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceTenantSchema(),                             },
             "virtual_service_se" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceVirtualServiceSeSchema(),                             },
             "vs_data_script" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceVSDataScriptSetSchema(),                             },
                                "url": &schema.Schema{
                                Type:     schema.TypeString,
                                Optional: true,
                                Computed: true,
                            },
        },
    }
}


