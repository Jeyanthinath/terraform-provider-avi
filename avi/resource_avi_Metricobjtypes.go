/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceMetricObjTypesSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "anomaly" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceMetricsAnomalyObjSchema(),                             },
             "app_insights" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceMetricsAppInsightsObjSchema(),                             },
             "collection" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceMetricCollectionsSchema(),                             },
             "controller_stats" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceControllerStatsObjSchema(),                             },
             "dns_client" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceMetricsVserverDNSObjSchema(),                             },
             "dns_server" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceMetricsServerDNSObjSchema(),                             },
             "dos" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceMetricsDosAnalyticsObjSchema(),                             },
             "healthscore" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceMetricsHealthScoreObjSchema(),                             },
             "l4_client" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceVserverL4MetricsObjSchema(),                             },
             "l4_server" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceServerL4MetricsObjSchema(),                             },
             "l7_client" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceVserverL7MetricsObjSchema(),                             },
             "l7_server" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceServerL7MetricsObjSchema(),                             },
             "process_stats" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceProcessStatsObjSchema(),                             },
             "rum" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceMetricsRumAnalyticsObjSchema(),                             },
             "rum_browser" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceMetricsRumPreaggBrowserObjSchema(),                             },
             "rum_country" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceMetricsRumPreaggCountryObjSchema(),                             },
             "rum_devtype" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceMetricsRumPreaggDevtypeObjSchema(),                             },
             "rum_ipgroup" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceMetricsRumPreaggIPGroupObjSchema(),                             },
             "rum_lang" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceMetricsRumPreaggLangObjSchema(),                             },
             "rum_os" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceMetricsRumPreaggOsObjSchema(),                             },
             "rum_url" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceMetricsRumPreaggUrlObjSchema(),                             },
             "samples_blob" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceMetricsRumSamplesBlobObjSchema(),                             },
             "samples_dimension" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceMetricsRumSamplesDimensionObjSchema(),                             },
             "se_if" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceMetricsSeIfStatsObjSchema(),                             },
             "se_stats" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeStatsObjSchema(),                             },
             "service_insights" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceMetricsServiceInsightsObjSchema(),                             },
             "source_insights" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceMetricsSourceInsightsObjSchema(),                             },
             "tenant_stats" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceMetricsTenantStatsObjSchema(),                             },
             "user_metrics" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceMetricsUserMetricsObjSchema(),                             },
             "vm_stats" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceVmStatsObjSchema(),                             },
             "waf_group" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceMetricsWafGroupMetricsObjSchema(),                             },
             "waf_rule" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceMetricsWafRuleMetricsObjSchema(),                             },
             "waf_tag" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceMetricsWafTagMetricsObjSchema(),                             },
                                "url": &schema.Schema{
                                Type:     schema.TypeString,
                                Optional: true,
                                Computed: true,
                            },
        },
    }
}


