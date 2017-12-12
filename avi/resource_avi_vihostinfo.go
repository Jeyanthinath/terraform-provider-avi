/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVIHostInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cluster_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"cpu_hz": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"dc_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"host_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"managed_object_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"mem": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"mgmt_portgroups": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"num_cpu_cores": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"num_cpu_packages": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"num_cpu_threads": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"pnics": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCdpLldpInfoSchema()},
			"portgroupmors": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"portgroups": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"powerstate": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
