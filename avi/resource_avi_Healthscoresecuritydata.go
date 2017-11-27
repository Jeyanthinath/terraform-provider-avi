/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceHealthScoreSecurityDataSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "dos_attack_level_data" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceDosAttackLevelDataSchema(),                             },
             "reason" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "reason_attr" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "se_dos_attack_level_data" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceServiceEngineDosAttackLevelDataSchema(),                             },
             "serviceengine_dos_attack_level_data" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceServiceEngineDosAttackLevelDataSchema(),                             },
             "ssl_score_data" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSslScoreDataSchema(),                             },
             "value" :&schema.Schema{
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


