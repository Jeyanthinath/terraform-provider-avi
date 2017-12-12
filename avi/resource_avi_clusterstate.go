/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceClusterStateSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cluster_state": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Set: func(v interface{}) int { return 0 }, Elem: ResourceClusterOperationalStatusSchema()},
			"node_states": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceClusterNodeStateSchema()},
			"service_states": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceClusterServiceStateSchema()},
		},
	}
}
