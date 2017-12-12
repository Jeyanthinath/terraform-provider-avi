/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceTimeoutStatsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"keepalive_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"persist_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"retransmit_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
		},
	}
}
