/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceCC_AzureSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"azure_cfg": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceAzureConfigurationSchema()},
		},
	}
}
