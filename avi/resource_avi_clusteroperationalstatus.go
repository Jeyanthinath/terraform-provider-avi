/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceClusterOperationalStatusSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"last_changed_time": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceTimeStampSchema()},
			"progress": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"reason": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"reason_code": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"up_since": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
