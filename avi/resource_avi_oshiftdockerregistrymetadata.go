/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceOshiftDockerRegistryMetaDataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"registry_namespace": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "default"},
			"registry_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "docker-registry"},
			"registry_vip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceIpAddrSchema()},
		},
	}
}
