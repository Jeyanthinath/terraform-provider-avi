/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSocketInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"so_error": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"so_fibnum": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"so_gencnt": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"so_incqlen": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"so_linger": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"so_oobmark": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"so_options": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"so_pcb": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"so_proto": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"so_qlen": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"so_qlimit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"so_qstate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"so_rcv": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSocketBufferInfoSchema()},
			"so_ref_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"so_snd": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSocketBufferInfoSchema()},
			"so_starttime": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceTimeStampSchema()},
			"so_state": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"so_type": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"so_usr_dbg_flags": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"so_usr_flags": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"so_usr_l7_handle": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"so_usr_proxy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"so_usr_service_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"so_usr_vserver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"so_vnet": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}
