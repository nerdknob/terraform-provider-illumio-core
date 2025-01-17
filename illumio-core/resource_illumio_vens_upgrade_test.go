// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var providerVENsUpgrade *schema.Provider

func TestAccIllumioVENsUpgrade_CreateUpdate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerVENsUpgrade),
		// CheckDestroy is ignored as illumio-core_vens_upgrade does not support delete operation
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioVENsUpgradeConfig_basic("63bf19d1-1efa-49ec-b712-c51d5c0aa552"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioVENsUpgradeExists("illumio-core_vens_upgrade.test"),
				),
			},
		},
	})
}

func testAccCheckIllumioVENsUpgradeConfig_basic(id string) string {
	return fmt.Sprintf(`
	resource "illumio-core_vens_upgrade" "test" {
		release = "21.2.0-7828"
		vens {
		  href = "/orgs/1/vens/%s"
		}
	  }
	`, id)
}

func testAccCheckIllumioVENsUpgradeExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("VENs Upgrade %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		return nil
	}
}
