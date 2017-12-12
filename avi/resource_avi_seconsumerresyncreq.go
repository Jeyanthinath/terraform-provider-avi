/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeConsumerResyncReqSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"apic_txn_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"consumer_info": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Set: func(v interface{}) int { return 0 }, Elem: ResourceSeResourceFindReqSchema()},
			"se_alloc_info": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSeAllocInfoSchema()},
		},
	}
}
