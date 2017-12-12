/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceHTTPPoliciesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"http_policy_set_ref": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"index": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
		},
	}
}
