/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSecureChannelTokenSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"expiry_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"metadata": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSecureChannelMetadataSchema()},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"node_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true},
		},
	}
}
