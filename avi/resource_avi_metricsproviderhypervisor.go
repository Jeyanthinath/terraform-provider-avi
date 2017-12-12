/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceMetricsProviderHypervisorSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"kvm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vmware": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
