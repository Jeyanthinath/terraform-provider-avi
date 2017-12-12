/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceWafLogSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"latency_request_body_phase": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"latency_request_header_phase": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"latency_response_body_phase": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"latency_response_header_phase": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"rule_logs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceWafRuleLogSchema()},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
