/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceAutoScaleStateSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"autoscale_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true},
			"autoscale_in_progress": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"autoscale_policy_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"intelligent_autoscale": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"last_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SCALEOUT"},
			"last_autoscale_ts": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_num_scalein": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"last_num_scaleout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"last_scalein_reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SYSERR_SUCCESS"},
			"last_scalein_ts": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_scaleout_reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SYSERR_SUCCESS"},
			"last_scaleout_ts": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"launch_error": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SYSERR_SUCCESS"},
			"launch_retries": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_oper_capacity": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_oper_servers": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"min_oper_capacity": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"min_oper_servers": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"obj_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
