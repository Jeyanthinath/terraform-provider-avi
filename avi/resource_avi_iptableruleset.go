/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceIptableRuleSetSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"chain": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"rules": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIptableRuleSchema()},
			"table": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
