/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVCALoginSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "iam.vchs.vmware.com"},
			"instance": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"orgnization": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
