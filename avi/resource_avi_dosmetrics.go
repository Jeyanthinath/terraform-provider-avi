/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceDosMetricsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"attacker_info": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMaliciousIpInfoSchema()},
			"node_obj_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
