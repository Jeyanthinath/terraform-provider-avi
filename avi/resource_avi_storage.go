/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceStorageSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"capacity": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"pci_bus": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"pci_channel": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"speed": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
