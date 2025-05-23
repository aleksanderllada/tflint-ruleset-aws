// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsEcsAccountSettingDefaultInvalidNameRule checks the pattern is valid
type AwsEcsAccountSettingDefaultInvalidNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsEcsAccountSettingDefaultInvalidNameRule returns new rule with default attributes
func NewAwsEcsAccountSettingDefaultInvalidNameRule() *AwsEcsAccountSettingDefaultInvalidNameRule {
	return &AwsEcsAccountSettingDefaultInvalidNameRule{
		resourceType:  "aws_ecs_account_setting_default",
		attributeName: "name",
		enum: []string{
			"serviceLongArnFormat",
			"taskLongArnFormat",
			"containerInstanceLongArnFormat",
			"awsvpcTrunking",
			"containerInsights",
			"fargateFIPSMode",
			"tagResourceAuthorization",
			"fargateTaskRetirementWaitPeriod",
			"guardDutyActivate",
			"defaultLogDriverMode",
		},
	}
}

// Name returns the rule name
func (r *AwsEcsAccountSettingDefaultInvalidNameRule) Name() string {
	return "aws_ecs_account_setting_default_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsEcsAccountSettingDefaultInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsEcsAccountSettingDefaultInvalidNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsEcsAccountSettingDefaultInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsEcsAccountSettingDefaultInvalidNameRule) Check(runner tflint.Runner) error {
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

		err := runner.EvaluateExpr(attribute.Expr, func (val string) error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as name`, truncateLongMessage(val)),
					attribute.Expr.Range(),
				)
			}
			return nil
		}, nil)
		if err != nil {
			return err
		}
	}

	return nil
}
