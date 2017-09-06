package logentries

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	lexp "github.com/terraform-providers/terraform-provider-logentries/logentries/expect"
)

func TestAccDataSourceLogentriesLogSet(t *testing.T) {
	var logSetResource LogSetResource

	logSetName := fmt.Sprintf("terraform-test-%s", acctest.RandStringFromCharSet(8, acctest.CharSetAlphaNum))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceLogentriesLogSet_basic(logSetName),
				Check: lexp.TestCheckResourceExpectation(
					"logentries_logset.logset",
					&logSetResource,
					testAccCheckLogentriesLogSetExists,
					map[string]lexp.TestExpectValue{
						"name": lexp.Equals(logSetName),
					},
				),
			},
			{
				Config: testAccDataSourceLogentriesLogSet_basic(logSetName),
				Check:  resource.TestCheckResourceAttr("data.logentries_logset.logset", "name", logSetName),
			},
		},
	})
}

func testAccResourceLogentriesLogSet_basic(logSetName string) string {
	return fmt.Sprintf(`
resource "logentries_logset" "logset" {
  name = "%s"
}
`, logSetName)
}

func testAccDataSourceLogentriesLogSet_basic(logSetName string) string {
	return fmt.Sprintf(`
resource "logentries_logset" "logset" {
  name = "%s"
}

data "logentries_logset" "logset" {
  name = "%s"
}
`, logSetName, logSetName)
}
