/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceNetworkSecurityPolicyActionRLParamSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"burst_size": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"max_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
		},
	}
}
