/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceBgpProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"community": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"hold_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  180,
			},
			"ibgp": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true},
			"ip_communities": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpCommunitySchema()},
			"keepalive_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60,
			},
			"local_as": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"peers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceBgpPeerSchema()},
			"send_community": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"shutdown": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}
