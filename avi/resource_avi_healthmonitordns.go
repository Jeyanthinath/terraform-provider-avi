/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceHealthMonitorDNSSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"qtype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "DNS_QUERY_TYPE"},
			"query_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"rcode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "RCODE_NO_ERROR"},
			"response_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
