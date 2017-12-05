/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi


import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceDnsInfoSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "algorithm" :&schema.Schema{
                             Type: schema.TypeString, 
                             Optional: true,
                             Default: "DNS_RECORD_RESPONSE_CONSISTENT_HASH",                                                                                                                },
             "cname" :&schema.Schema{
                             Type: schema.TypeSet, 
                             Optional: true,
                                                                                                                 Elem: ResourceDnsCnameRdataSchema(),                             },
             "fqdn" :&schema.Schema{
                             Type: schema.TypeString, 
                             Optional: true,
                                                                                                                                            },
             "metadata" :&schema.Schema{
                             Type: schema.TypeString, 
                             Optional: true,
                                                                                                                                            },
             "num_records_in_response" :&schema.Schema{
                             Type: schema.TypeInt, 
                             Optional: true,
                             Default: 1,
                                                                                                                                            },
             "ttl" :&schema.Schema{
                             Type: schema.TypeInt, 
                             Optional: true,
                                                                                                                                            },
             "type" :&schema.Schema{
                             Type: schema.TypeString, 
                             Optional: true,
                             Default: "DNS_RECORD_A",                                                                                                                },
                                "url": &schema.Schema{
                                Type:     schema.TypeString,
                                Optional: true,
                                Computed: true,
                            },
        },
    }
}


