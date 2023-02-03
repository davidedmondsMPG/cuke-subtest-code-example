package cuke_subtest_code_example

import (
	"testing"

	"github.com/cucumber/godog"
)

func TestFeaturesNoSubtests(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format: "pretty",
			Paths:  []string{"features"},
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func TestFeaturesSubtests(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func iEat(arg1 int) error {
	return godog.ErrPending
}

func thereAreGodogs(arg1 int) error {
	return godog.ErrPending
}

func thereShouldBeRemaining(arg1 int) error {
	return godog.ErrPending
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I eat (\d+)$`, iEat)
	ctx.Step(`^there are (\d+) godogs$`, thereAreGodogs)
	ctx.Step(`^there should be (\d+) remaining$`, thereShouldBeRemaining)
}
