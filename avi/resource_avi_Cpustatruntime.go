/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceCpuStatRuntimeSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "free_memory" :&schema.Schema{
                             Type: schema.TypeInt,            
                             Optional: true,                              
                                                        },
             "idle_cpu" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "process_cpu_utilization" :&schema.Schema{
                             Type: schema.TypeList,            
                             Optional: true,                              
                             Elem: ResourceProcCpuStatSchema(),                             },
             "se_uuid" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "total_cpu_utilization" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "total_linux_cpu_utilization" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "total_memory" :&schema.Schema{
                             Type: schema.TypeInt,            
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


