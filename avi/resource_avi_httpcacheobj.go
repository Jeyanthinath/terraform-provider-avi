/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceHttpCacheObjSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ae_type_bm": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"body_size": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"ce_top": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ce_type_bm": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ctype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"data_size": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"date_time": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"etag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"exp_age": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"exp_age_hrt": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true},
			"handle": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"has_vary": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true},
			"hdr_size": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"in_time": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"init_age": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"is_chunked": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true},
			"is_expired": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true},
			"is_purged": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true},
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"key_extn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_mod_time": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"last_used": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"mbuf_head": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"mcache_out": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true},
			"meta_size": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"must_reval": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true},
			"no_exp_info": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true},
			"no_txm": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true},
			"proc_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"proxy_reval": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true},
			"raw_extn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"raw_key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"refcnt": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"reuse_cnt": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"reval": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
		},
	}
}
