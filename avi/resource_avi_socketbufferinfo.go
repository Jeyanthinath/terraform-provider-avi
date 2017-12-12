/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSocketBufferInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"sb_cc": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sb_credit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sb_flags": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sb_hiwat": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sb_lowat": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sb_mcnt": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sb_sndptroff": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sb_state": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}
