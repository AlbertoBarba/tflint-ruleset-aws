// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AwsSpotFleetRequestInvalidAllocationStrategyRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_spot_fleet_request" "foo" {
	allocation_strategy = "highestPrice"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAwsSpotFleetRequestInvalidAllocationStrategyRule(),
					Message: `"highestPrice" is an invalid value as allocation_strategy`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_spot_fleet_request" "foo" {
	allocation_strategy = "lowestPrice"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewAwsSpotFleetRequestInvalidAllocationStrategyRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}