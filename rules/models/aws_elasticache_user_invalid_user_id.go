// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsElastiCacheUserInvalidUserIDRule checks the pattern is valid
type AwsElastiCacheUserInvalidUserIDRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	min           int
	pattern       *regexp.Regexp
}

// NewAwsElastiCacheUserInvalidUserIDRule returns new rule with default attributes
func NewAwsElastiCacheUserInvalidUserIDRule() *AwsElastiCacheUserInvalidUserIDRule {
	return &AwsElastiCacheUserInvalidUserIDRule{
		resourceType:  "aws_elasticache_user",
		attributeName: "user_id",
		min:           1,
		pattern:       regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9\-]*$`),
	}
}

// Name returns the rule name
func (r *AwsElastiCacheUserInvalidUserIDRule) Name() string {
	return "aws_elasticache_user_invalid_user_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsElastiCacheUserInvalidUserIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsElastiCacheUserInvalidUserIDRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsElastiCacheUserInvalidUserIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsElastiCacheUserInvalidUserIDRule) Check(runner tflint.Runner) error {
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
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"user_id must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-zA-Z][a-zA-Z0-9\-]*$`),
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
