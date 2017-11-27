/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceServerL7MetricsObjSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "apdexr" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_application_response_time" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_complete_responses" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_error_responses" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_errors_excluded" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_frustrated_responses" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_pool_complete_responses" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_pool_error_responses" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_reqs_per_session" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_resp_1xx" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_resp_2xx" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_resp_3xx" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_resp_4xx" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_resp_4xx_errors" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_resp_4xx_errors_excluded" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_resp_5xx" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_resp_5xx_errors" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_resp_5xx_errors_excluded" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_resp_latency" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_satisfactory_responses" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_server_count" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_timeouts" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_tolerated_responses" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "avg_total_requests" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "max_concurrent_sessions" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "node_obj_id" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "pct_response_errors" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_application_response_time" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_finished_sessions" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_get_reqs" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_get_resp_latency" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_get_resp_latency_bucket1" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_get_resp_latency_bucket2" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_lb_fail_count" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_other_reqs" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_other_resp_latency" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_other_resp_latency_bucket1" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_other_resp_latency_bucket2" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_post_reqs" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_post_resp_latency" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_post_resp_latency_bucket1" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_post_resp_latency_bucket2" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_reqs_finished_sessions" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_resp_1xx" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_resp_2xx" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_resp_3xx" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_resp_4xx" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_resp_5xx" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "sum_total_responses" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
                                "url": &schema.Schema{
                                Type:     schema.TypeString,
                                Optional: true,
                                Computed: true,
                            },
        },
    }
}


