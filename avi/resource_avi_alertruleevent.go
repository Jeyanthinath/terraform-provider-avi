/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceAlertRuleEventSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"event_details": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceEventDetailsFilterSchema()},
			"event_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"not_cond": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}
