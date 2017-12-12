/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeReservedVsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"detail": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceVirtualServiceRuntimeDetailSchema()},
			//"internal": &schema.Schema{
			//	Type:     schema.TypeSet,
			//	Optional: true,
			//	Set:      func(v interface{}) int { return 0 }, Elem: ResourceVirtualServiceInternalSchema()},
			"proc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tcpstat": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceTcpStatRuntimeSchema()},
			"traffic_clone_stats": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceTrafficCloneRuntimeSchema()},
			"udpstat": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceUdpStatRuntimeSchema()},
		},
	}
}
