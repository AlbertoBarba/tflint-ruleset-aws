// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsRoute53DelegationSetInvalidReferenceNameRule checks the pattern is valid
type AwsRoute53DelegationSetInvalidReferenceNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsRoute53DelegationSetInvalidReferenceNameRule returns new rule with default attributes
func NewAwsRoute53DelegationSetInvalidReferenceNameRule() *AwsRoute53DelegationSetInvalidReferenceNameRule {
	return &AwsRoute53DelegationSetInvalidReferenceNameRule{
		resourceType:  "aws_route53_delegation_set",
		attributeName: "reference_name",
		max:           128,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsRoute53DelegationSetInvalidReferenceNameRule) Name() string {
	return "aws_route53_delegation_set_invalid_reference_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsRoute53DelegationSetInvalidReferenceNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsRoute53DelegationSetInvalidReferenceNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsRoute53DelegationSetInvalidReferenceNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsRoute53DelegationSetInvalidReferenceNameRule) Check(runner tflint.Runner) error {
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
					"reference_name must be 128 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"reference_name must be 1 characters or higher",
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
