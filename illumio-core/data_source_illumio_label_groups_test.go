package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var providerDSLabelGsL *schema.Provider

func TestAccIllumioLabelGsL_Read(t *testing.T) {
	listAttr := map[string]interface{}{}
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerDSLabelGsL),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioLabelGsLDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioDataSourceLabelGsLExists("data.illumio-core_label_groups.test", listAttr),
					testAccCheckIllumioListDataSourceSize(listAttr, "5"),
				),
			},
		},
	})
}

func testAccCheckIllumioLabelGsLDataSourceConfig_basic() string {
	return `
	data "illumio-core_label_groups" "test" {
		max_results = "5"
	}
	`
}

func testAccCheckIllumioDataSourceLabelGsLExists(name string, listAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("List of Label Groups %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		listAttr["length"] = rs.Primary.Attributes["items.#"]

		return nil
	}
}