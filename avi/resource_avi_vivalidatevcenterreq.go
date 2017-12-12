/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVIValidateVcenterReqSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"admin": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Set: func(v interface{}) int { return 0 }, Elem: ResourceVIAdminCredentialsSchema()},
			"cloud_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"controller_cluster_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"controller_ips": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"controller_vm_mor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"datacenter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"management_network": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcenter_template_se_location": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
