package e2e

import (
	"fmt"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/openshift/odo/tests/e2e/helper"
)

var _ = Describe("odoServiceE2e", func() {
	//new clean project and context for each test
	var project string
	var context string

	//  current directory and project (before eny test is run) so it can restored  after all testing is done
	var originalDir string
	var originalProject string

	// Setup up state for each test spec
	// create new project (not set as active) and new context directory for each test spec
	// This is before every spec (It)
	var _ = BeforeEach(func() {
		project = helper.OcCreateRandProject()
		context = helper.CreateNewContext()
		originalDir = helper.Getwd()
		originalProject = helper.OcGetCurrentProject()
		helper.Chdir(context)
		helper.OcSwitchProject(project)
	})

	// Clean up after the test
	// This is run after every Spec (It)
	var _ = AfterEach(func() {
		helper.Chdir(originalDir)
		helper.OcSwitchProject(originalProject)
		helper.OcDeleteProject(project)
		helper.DeleteDir(context)
	})

	Context("odo service create with a spring boot application", func() {
		It("should be able to create postgresql and link it with springboot", func() {
			importOpenJDKImage()
			helper.CopyExample("source/openjdk-sb-postgresql", context)

			// Local config needs to be present in order to create service https://github.com/openshift/odo/issues/1602
			runCmdShouldPass("odo create java sb-app")

			runCmdShouldPass("odo service create dh-postgresql-apb --plan dev -p postgresql_user=luke -p postgresql_password=secret -p postgresql_database=my_data -p postgresql_version=9.6")
			waitForCmdOut("oc get serviceinstance -o name", 1, true, func(output string) bool {
				return strings.Contains(output, "dh-postgresql-apb")
			})

			// Create a URL
			runCmdShouldPass("odo url create --port 8080")

			// push removes link, this is why link needs to be run alaways after the push https://github.com/openshift/odo/issues/1596
			runCmdShouldPass("odo push")

			runCmdShouldPass("odo link dh-postgresql-apb -w --wait-for-target")

			waitForCmdOut("odo service list | sed 1d", 1, true, func(output string) bool {
				return strings.Contains(output, "dh-postgresql-apb") &&
					strings.Contains(output, "ProvisionedAndLinked")
			})

			routeURL := determineRouteURL()

			// Ping said URL
			responseStringMatchStatus := matchResponseSubString(routeURL, "Spring Boot", 30, 1)
			Expect(responseStringMatchStatus).Should(BeTrue())

			// Delete the component
			runCmdShouldPass("odo delete sb-app -f")

			// Delete the service
			runCmdShouldPass("odo service delete dh-postgresql-apb -f")
		})
	})

	// TODO: auth issue, we need to find a proper way how to test it without requiring cluster admin privileges

	// Context("odo hides a hidden service in service catalog", func() {
	// 	It("not show a hidden service in the catalog", func() {
	// 		runCmdShouldPass("oc apply -f https://github.com/openshift/library/raw/master/official/sso/templates/sso72-https.json -n openshift")
	// 		outputErr := runCmdShouldFail("odo catalog search service sso72-https")
	// 		Expect(outputErr).To(ContainSubstring("No service matched the query: sso72-https"))
	// 	})
	// })

})

func serviceInstanceStatusCmd(serviceInstanceName string) string {
	return fmt.Sprintf("oc get serviceinstance %s -o go-template='{{ (index .status.conditions 0).reason}}'", serviceInstanceName)
}
