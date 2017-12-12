/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSslKeySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"session_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
