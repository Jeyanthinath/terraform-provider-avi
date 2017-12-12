/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeAssignReqBatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"child_reqs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSeResourceFindReqSchema()},
			"parent_req": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSeAssignReqSchema()},
		},
	}
}
