/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceIcmpHistSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"icmp_althostaddr": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"icmp_echoreply": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"icmp_echoreq": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"icmp_ireq": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"icmp_ireqreply": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"icmp_maskreply": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"icmp_maskreq": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"icmp_paramprob": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"icmp_redirect": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"icmp_routeradvert": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"icmp_routersolicit": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"icmp_sourcequench": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"icmp_timxceed": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"icmp_tstamp": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"icmp_tstampreply": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"icmp_unreach": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
		},
	}
}
