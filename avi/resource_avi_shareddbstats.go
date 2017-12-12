/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSharedDbStatsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"age_deletes": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"age_updates_to_redis": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"create_exists": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"creates": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"creates_to_redis": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"delete_from_redis_invalid_keytype": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"deletes": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"deletes_from_redis": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"deletes_from_redis_invalid_handle": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"deletes_from_redis_nokey_exist": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"deletes_nokey_exist": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"deletes_to_redis": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"destroy": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"init": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"lookup_failures": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"se_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"selective_destroy": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"syncs_to_redis": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"syncs_to_redis_failures": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"updates_from_redis": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"updates_from_redis_create_exists": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"updates_from_redis_find_exists": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"updates_from_redis_invalid_handle": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"updates_from_redis_invalid_keytype": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"updates_from_redis_invalid_value": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
		},
	}
}
