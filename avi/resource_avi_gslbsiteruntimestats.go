/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceGslbSiteRuntimeStatsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"num_file_cr_txed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_file_del_txed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gap_cr_rxed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gap_cr_txed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gap_del_rxed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gap_del_txed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gap_upd_rxed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gap_upd_txed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_geo_cr_rxed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_geo_cr_txed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_geo_del_rxed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_geo_del_txed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_geo_upd_rxed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_geo_upd_txed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_ghm_cr_rxed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_ghm_cr_txed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_ghm_del_rxed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_ghm_del_txed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_ghm_upd_rxed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_ghm_upd_txed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_glb_cr_rxed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_glb_cr_txed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_glb_del_rxed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_glb_del_txed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_glb_upd_rxed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_glb_upd_txed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gpki_cr_rxed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gpki_cr_txed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gpki_del_rxed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gpki_del_txed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gpki_upd_rxed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gpki_upd_txed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gs_cr_rxed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gs_cr_txed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gs_del_rxed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gs_del_txed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gs_upd_rxed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gs_upd_txed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_health_msgs_rxed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_health_msgs_txed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_of_bad_responses": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_of_events_generated": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_of_skip_outstanding_requests": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_of_timeouts": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}
