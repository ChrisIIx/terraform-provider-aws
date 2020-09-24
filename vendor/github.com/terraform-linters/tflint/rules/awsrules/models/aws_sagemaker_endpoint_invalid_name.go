// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsSagemakerEndpointInvalidNameRule checks the pattern is valid
type AwsSagemakerEndpointInvalidNameRule struct {
	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsSagemakerEndpointInvalidNameRule returns new rule with default attributes
func NewAwsSagemakerEndpointInvalidNameRule() *AwsSagemakerEndpointInvalidNameRule {
	return &AwsSagemakerEndpointInvalidNameRule{
		resourceType:  "aws_sagemaker_endpoint",
		attributeName: "name",
		max:           63,
		pattern:       regexp.MustCompile(`^[a-zA-Z0-9](-*[a-zA-Z0-9])*`),
	}
}

// Name returns the rule name
func (r *AwsSagemakerEndpointInvalidNameRule) Name() string {
	return "aws_sagemaker_endpoint_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSagemakerEndpointInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSagemakerEndpointInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSagemakerEndpointInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSagemakerEndpointInvalidNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"name must be 63 characters or less",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-zA-Z0-9](-*[a-zA-Z0-9])*`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
