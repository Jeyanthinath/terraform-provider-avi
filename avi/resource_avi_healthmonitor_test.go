package avi

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAVIhealthmonitorCreate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAcchealthmonitorConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckhealthmonitorExists("avi_healthmonitor.testhealthmonitor"),
					resource.TestCheckResourceAttr(
						"avi_healthmonitor.testhealthmonitor", "name", "p1234")),
			},
		},
	})

}

func testAccCheckhealthmonitorExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No AVI healthmonitor ID is set")
		}
		return nil
	}

}

const testAcchealthmonitorConfig = `
`
