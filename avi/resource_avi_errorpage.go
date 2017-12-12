/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceErrorPageSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enable": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"error_page_body_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"error_redirect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"index": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"match": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceHTTPStatusMatchSchema()},
		},
	}
}
