/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
        "github.com/hashicorp/terraform/helper/schema"
)
 func ResourceSslScoreDataSchema() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
             "cert_chain_verified" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "cert_chain_verified_score" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "disable_client_renegotiation" :&schema.Schema{
                             Type: schema.TypeBool,            
                             Optional: true,                              
                                                        },
             "disable_client_renegotiation_score" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "earliest_cert_expiry" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "earliest_cert_expiry_score" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "hsts_enabled" :&schema.Schema{
                             Type: schema.TypeBool,            
                             Optional: true,                              
                                                        },
             "hsts_enabled_score" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "min_cipher_strength" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "min_cipher_strength_score" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "min_ssl_protocol_strength" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "min_ssl_protocol_strength_score" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "pfs_support" :&schema.Schema{
                             Type: schema.TypeBool,            
                             Optional: true,                              
                                                        },
             "pfs_support_score" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "reason" :&schema.Schema{
                             Type: schema.TypeString,            
                              Required: true,                              
                                                        },
             "reason_attr" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "ref" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                              Elem:&schema.Schema{Type: schema.TypeString},                             },
             "self_signed_cert" :&schema.Schema{
                             Type: schema.TypeBool,            
                             Optional: true,                              
                                                        },
             "self_signed_cert_score" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "ssl_enabled" :&schema.Schema{
                             Type: schema.TypeBool,            
                             Optional: true,                              
                                                        },
             "value" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "weak_signature_algorithm" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "weak_signature_algorithm_score" :&schema.Schema{
                             Type: schema.,            
                             Optional: true,                              
                                                        },
             "weakest_enc_algo" :&schema.Schema{
                             Type: schema.TypeString,            
                             Optional: true,                              
                                                        },
             "weakest_enc_algo_score" :&schema.Schema{
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


