/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceHttpCacheStatsDetailSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"available_size": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"etype_adds": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHttpCacheETypeStatsSchema()},
			"etype_objects": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHttpCacheETypeStatsSchema()},
			"fetch": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHttpCacheStatSchema()},
			"global_evicts": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"incoming": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Set: func(v interface{}) int { return 0 }, Elem: ResourceHttpCacheStatsObjSchema()},
			"local_evicts": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"lookups": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"mcopy_fail": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"meta_size": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"msplit_fail": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"outgoing": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Set: func(v interface{}) int { return 0 }, Elem: ResourceHttpCacheStatsObjSchema()},
			"proc_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"reval": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHttpCacheStatSchema()},
			"se_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"shm_fail": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"store": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHttpCacheStatSchema()},
			"store_out": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHttpCacheStatSchema()},
		},
	}
}
