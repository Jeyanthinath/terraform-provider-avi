/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourcensxExcludeMemberSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"member": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourcensxMemberSchema()},
		},
	}
}
