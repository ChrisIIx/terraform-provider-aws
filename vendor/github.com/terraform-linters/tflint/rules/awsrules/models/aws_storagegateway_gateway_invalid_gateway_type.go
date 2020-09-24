// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsStoragegatewayGatewayInvalidGatewayTypeRule checks the pattern is valid
type AwsStoragegatewayGatewayInvalidGatewayTypeRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsStoragegatewayGatewayInvalidGatewayTypeRule returns new rule with default attributes
func NewAwsStoragegatewayGatewayInvalidGatewayTypeRule() *AwsStoragegatewayGatewayInvalidGatewayTypeRule {
	return &AwsStoragegatewayGatewayInvalidGatewayTypeRule{
		resourceType:  "aws_storagegateway_gateway",
		attributeName: "gateway_type",
		max:           20,
		min:           2,
	}
}

// Name returns the rule name
func (r *AwsStoragegatewayGatewayInvalidGatewayTypeRule) Name() string {
	return "aws_storagegateway_gateway_invalid_gateway_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsStoragegatewayGatewayInvalidGatewayTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsStoragegatewayGatewayInvalidGatewayTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsStoragegatewayGatewayInvalidGatewayTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsStoragegatewayGatewayInvalidGatewayTypeRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"gateway_type must be 20 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"gateway_type must be 2 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
