/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeResourceDelRespSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"se_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"status": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"svc_ids": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSeMgrSvcIdSchema()},
		},
	}
}
