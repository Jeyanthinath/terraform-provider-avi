/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceSeAgentGraphDBRuntimeSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "analyticsprofile" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeAgentGraphDBNodeInfoSchema(),                             },
             "applicationpersistenceprofile" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeAgentGraphDBNodeInfoSchema(),                             },
             "applicationprofile" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeAgentGraphDBNodeInfoSchema(),                             },
             "cloud" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeAgentGraphDBNodeInfoSchema(),                             },
             "graph_version" :&schema.Schema{
                             Type: schema.TypeInt,            
                             Optional: true,                              
                                                        },
             "gslb" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeAgentGraphDBNodeInfoSchema(),                             },
             "gslbgeodbprofile" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeAgentGraphDBNodeInfoSchema(),                             },
             "gslbhealthmonitor" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeAgentGraphDBNodeInfoSchema(),                             },
             "gslbservice" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeAgentGraphDBNodeInfoSchema(),                             },
             "healthmonitor" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeAgentGraphDBNodeInfoSchema(),                             },
             "httprequestpolicy" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeAgentGraphDBNodeInfoSchema(),                             },
             "httpresponsepolicy" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeAgentGraphDBNodeInfoSchema(),                             },
             "httpsecuritypolicy" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeAgentGraphDBNodeInfoSchema(),                             },
             "ipaddrgroup" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeAgentGraphDBNodeInfoSchema(),                             },
             "microservice" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeAgentGraphDBNodeInfoSchema(),                             },
             "networkprofile" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeAgentGraphDBNodeInfoSchema(),                             },
             "networksecuritypolicy" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeAgentGraphDBNodeInfoSchema(),                             },
             "pool" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeAgentGraphDBNodeInfoSchema(),                             },
             "se_ref" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                              Elem:&schema.Schema{Type: schema.TypeString},                             },
             "serviceenginegroup" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeAgentGraphDBNodeInfoSchema(),                             },
             "sslkeyandcertificate" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeAgentGraphDBNodeInfoSchema(),                             },
             "sslprofile" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeAgentGraphDBNodeInfoSchema(),                             },
             "stringgroup" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeAgentGraphDBNodeInfoSchema(),                             },
             "tenant" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeAgentGraphDBNodeInfoSchema(),                             },
             "total_obj" :&schema.Schema{
                             Type: schema.TypeInt,            
                              Required: true,                              
                                                        },
             "total_obj_active" :&schema.Schema{
                             Type: schema.TypeInt,            
                             Optional: true,                              
                                                        },
             "total_obj_awaiting_dp" :&schema.Schema{
                             Type: schema.TypeInt,            
                             Optional: true,                              
                                                        },
             "total_obj_error" :&schema.Schema{
                             Type: schema.TypeInt,            
                             Optional: true,                              
                                                        },
             "total_obj_ew_subnet_error" :&schema.Schema{
                             Type: schema.TypeInt,            
                             Optional: true,                              
                                                        },
             "virtualservice" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeAgentGraphDBNodeInfoSchema(),                             },
             "vsdatascriptset" :&schema.Schema{
                             Type: schema.TypeSet,            
                             Optional: true,                              
                             Elem: ResourceSeAgentGraphDBNodeInfoSchema(),                             },
                                "url": &schema.Schema{
                                Type:     schema.TypeString,
                                Optional: true,
                                Computed: true,
                            },
        },
    }
}


