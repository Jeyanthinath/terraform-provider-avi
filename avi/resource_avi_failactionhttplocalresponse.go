/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceFailActionHTTPLocalResponseSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"file": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceHTTPLocalFileSchema()},
			"status_code": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "FAIL_HTTP_STATUS_CODE_503"},
		},
	}
}
