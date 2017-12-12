/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceWafRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enable": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"exclude_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceWafExcludeListEntrySchema()},
			"index": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rule_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
