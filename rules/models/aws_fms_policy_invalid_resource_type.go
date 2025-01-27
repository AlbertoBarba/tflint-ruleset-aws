// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsFmsPolicyInvalidResourceTypeRule checks the pattern is valid
type AwsFmsPolicyInvalidResourceTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsFmsPolicyInvalidResourceTypeRule returns new rule with default attributes
func NewAwsFmsPolicyInvalidResourceTypeRule() *AwsFmsPolicyInvalidResourceTypeRule {
	return &AwsFmsPolicyInvalidResourceTypeRule{
		resourceType:  "aws_fms_policy",
		attributeName: "resource_type",
		max:           128,
		min:           1,
		pattern:       regexp.MustCompile(`^([\p{L}\p{Z}\p{N}_.:/=+\-@]*)$`),
	}
}

// Name returns the rule name
func (r *AwsFmsPolicyInvalidResourceTypeRule) Name() string {
	return "aws_fms_policy_invalid_resource_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsFmsPolicyInvalidResourceTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsFmsPolicyInvalidResourceTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsFmsPolicyInvalidResourceTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsFmsPolicyInvalidResourceTypeRule) Check(runner tflint.Runner) error {
	logger.Trace("Check `%s` rule", r.Name())

	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: r.attributeName},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}

		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		err = runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"resource_type must be 128 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"resource_type must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^([\p{L}\p{Z}\p{N}_.:/=+\-@]*)$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}
