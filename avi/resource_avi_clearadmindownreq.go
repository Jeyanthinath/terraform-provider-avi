/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceClearAdminDownReqSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"consumer_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"se_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
