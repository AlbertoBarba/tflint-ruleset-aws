// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCloudfrontFunctionInvalidRuntimeRule checks the pattern is valid
type AwsCloudfrontFunctionInvalidRuntimeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsCloudfrontFunctionInvalidRuntimeRule returns new rule with default attributes
func NewAwsCloudfrontFunctionInvalidRuntimeRule() *AwsCloudfrontFunctionInvalidRuntimeRule {
	return &AwsCloudfrontFunctionInvalidRuntimeRule{
		resourceType:  "aws_cloudfront_function",
		attributeName: "runtime",
		enum: []string{
			"cloudfront-js-1.0",
		},
	}
}

// Name returns the rule name
func (r *AwsCloudfrontFunctionInvalidRuntimeRule) Name() string {
	return "aws_cloudfront_function_invalid_runtime"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudfrontFunctionInvalidRuntimeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudfrontFunctionInvalidRuntimeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudfrontFunctionInvalidRuntimeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudfrontFunctionInvalidRuntimeRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" is an invalid value as runtime`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}