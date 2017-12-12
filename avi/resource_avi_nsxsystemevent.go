/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourcensxSystemEventSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"eventcode": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"eventid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"eventmetadata": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourcensxEventMetaDataSchema()},
			"eventsource": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"isresourceuniversal": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"message": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"module": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"reportername": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"reportertype": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"severity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sourcetype": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"timestamp": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}
