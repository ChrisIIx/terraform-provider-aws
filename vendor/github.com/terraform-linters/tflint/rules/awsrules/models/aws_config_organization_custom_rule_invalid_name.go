// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsConfigOrganizationCustomRuleInvalidNameRule checks the pattern is valid
type AwsConfigOrganizationCustomRuleInvalidNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsConfigOrganizationCustomRuleInvalidNameRule returns new rule with default attributes
func NewAwsConfigOrganizationCustomRuleInvalidNameRule() *AwsConfigOrganizationCustomRuleInvalidNameRule {
	return &AwsConfigOrganizationCustomRuleInvalidNameRule{
		resourceType:  "aws_config_organization_custom_rule",
		attributeName: "name",
		max:           64,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsConfigOrganizationCustomRuleInvalidNameRule) Name() string {
	return "aws_config_organization_custom_rule_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsConfigOrganizationCustomRuleInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsConfigOrganizationCustomRuleInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsConfigOrganizationCustomRuleInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsConfigOrganizationCustomRuleInvalidNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"name must be 64 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"name must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
