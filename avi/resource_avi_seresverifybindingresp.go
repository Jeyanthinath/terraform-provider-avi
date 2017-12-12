/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeResVerifyBindingRespSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"binding_status": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"rpc_status": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"se_binding_info": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSeAllocInfoSchema()},
		},
	}
}
