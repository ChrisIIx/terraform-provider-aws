// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsBudgetsBudgetInvalidBudgetTypeRule checks the pattern is valid
type AwsBudgetsBudgetInvalidBudgetTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsBudgetsBudgetInvalidBudgetTypeRule returns new rule with default attributes
func NewAwsBudgetsBudgetInvalidBudgetTypeRule() *AwsBudgetsBudgetInvalidBudgetTypeRule {
	return &AwsBudgetsBudgetInvalidBudgetTypeRule{
		resourceType:  "aws_budgets_budget",
		attributeName: "budget_type",
		enum: []string{
			"USAGE",
			"COST",
			"RI_UTILIZATION",
			"RI_COVERAGE",
			"SAVINGS_PLANS_UTILIZATION",
			"SAVINGS_PLANS_COVERAGE",
		},
	}
}

// Name returns the rule name
func (r *AwsBudgetsBudgetInvalidBudgetTypeRule) Name() string {
	return "aws_budgets_budget_invalid_budget_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsBudgetsBudgetInvalidBudgetTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsBudgetsBudgetInvalidBudgetTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsBudgetsBudgetInvalidBudgetTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsBudgetsBudgetInvalidBudgetTypeRule) Check(runner *tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as budget_type`, truncateLongMessage(val)),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
