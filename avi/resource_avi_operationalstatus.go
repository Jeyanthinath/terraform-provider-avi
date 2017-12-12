/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceOperationalStatusSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"last_changed_time": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceTimeStampSchema()},
			"reason": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"reason_code": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"reason_code_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
