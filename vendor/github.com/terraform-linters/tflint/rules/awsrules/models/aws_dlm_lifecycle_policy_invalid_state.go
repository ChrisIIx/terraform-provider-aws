// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsDlmLifecyclePolicyInvalidStateRule checks the pattern is valid
type AwsDlmLifecyclePolicyInvalidStateRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsDlmLifecyclePolicyInvalidStateRule returns new rule with default attributes
func NewAwsDlmLifecyclePolicyInvalidStateRule() *AwsDlmLifecyclePolicyInvalidStateRule {
	return &AwsDlmLifecyclePolicyInvalidStateRule{
		resourceType:  "aws_dlm_lifecycle_policy",
		attributeName: "state",
		enum: []string{
			"ENABLED",
			"DISABLED",
		},
	}
}

// Name returns the rule name
func (r *AwsDlmLifecyclePolicyInvalidStateRule) Name() string {
	return "aws_dlm_lifecycle_policy_invalid_state"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDlmLifecyclePolicyInvalidStateRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDlmLifecyclePolicyInvalidStateRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDlmLifecyclePolicyInvalidStateRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDlmLifecyclePolicyInvalidStateRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as state`, truncateLongMessage(val)),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
