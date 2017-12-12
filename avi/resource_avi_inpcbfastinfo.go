/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceInpcbFastInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"debug_flags": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"inp_ack_delta": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"inp_fast_flags": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"inp_fast_state": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"inp_idle_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"inp_last_pkt_tick": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"inp_next_seq": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"inp_seq_delta": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"inp_snat_faddr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"inp_snat_fport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"inp_snat_laddr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"inp_snat_lport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"inp_starttime": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceTimeStampSchema()},
			"inp_syns_sent": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"inp_timer_shift": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"inp_ts_offset": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"inp_tsecr": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tx_vnic_hdl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}
