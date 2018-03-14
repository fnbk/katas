package features

import (
	"github.com/DATA-DOG/godog"
)

func SuiteContext(suite *godog.Suite) {

	api := &pimAPI{}

	// show product
	suite.BeforeScenario(api.resetResponse)
	suite.Step(`^a product with the ID "([^"]*)" and StructureID "([^"]*)" exists$`, api.aProductWithTheIDAndStructureIDExists)
	suite.Step(`^a structure with the ID "([^"]*)" exists$`, api.aStructureWithTheIDExists)
	suite.Step(`^product and structure with attributes and settings in the database:$`,
		api.productAndStructureWithAttributesAndSettingsInTheDatabase)
	suite.Step(`^I send "([^"]*)" request to "([^"]*)"$`, api.iSendRequestTo)
	suite.Step(`^the response code should be (\d+)$`, api.theResponseCodeShouldBe)
	suite.Step(`^the response should match json:$`, api.theResponseShouldMatchJson)

	// list product IDs
	suite.BeforeScenario(api.resetResponse)
	suite.Step(`^there are IDs in the product database:$`, api.thereAreIDsInTheProductDatabase)
	suite.Step(`^I send "([^"]*)" request to "([^"]*)"$`, api.iSendRequestTo)
	suite.Step(`^the response code should be (\d+)$`, api.theResponseCodeShouldBe)
	suite.Step(`^the response should match json:$`, api.theResponseShouldMatchJson)
}
