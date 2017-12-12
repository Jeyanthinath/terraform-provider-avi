/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceContainerTaskIdFilterKeySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"microservice_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"task_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
