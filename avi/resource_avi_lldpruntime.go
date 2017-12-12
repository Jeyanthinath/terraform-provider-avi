/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceLldpRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"interface_lldp_entry": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceInterfaceLldpEntrySchema()},
			"namespace_lldp_summary": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceNamespaceLldpEntrySchema()},
		},
	}
}
