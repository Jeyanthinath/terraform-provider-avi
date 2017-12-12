/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceNTPAuthenticationKeySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "NTP_AUTH_ALGORITHM_MD5"},
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"key_number": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}
