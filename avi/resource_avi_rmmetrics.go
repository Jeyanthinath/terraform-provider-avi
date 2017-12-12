/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceRmMetricsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"se_grp_metrics": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceRmSeGrpMetricsSchema()},
		},
	}
}
