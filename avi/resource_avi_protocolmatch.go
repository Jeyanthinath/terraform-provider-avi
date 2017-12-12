/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceProtocolMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match_criteria": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"protocols": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
