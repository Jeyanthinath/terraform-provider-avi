/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceDebugFilterUnionSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "alert_debug_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceAlertMgrDebugFilterSchema(),                             },
             "autoscale_mgr_debug_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceAutoScaleMgrDebugFilterSchema(),                             },
             "cloud_connector_debug_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceCloudConnectorDebugFilterSchema(),                             },
             "hs_debug_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceHSMgrDebugFilterSchema(),                             },
             "mesos_metrics_debug_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceMesosMetricsDebugFilterSchema(),                             },
             "metrics_debug_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceMetricsMgrDebugFilterSchema(),                             },
             "se_mgr_debug_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeMgrDebugFilterSchema(),                             },
             "state_cache_mgr_debug_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceStateCacheMgrDebugFilterSchema(),                             },
             "type" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "vs_debug_filter" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceVsDebugFilterSchema(),                             },
                                "url": &schema.Schema{
                                Type:     schema.TypeString,
                                Optional: true,
                                Computed: true,
                            },
        },
    }
}


