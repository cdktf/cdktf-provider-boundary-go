// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package aliastarget

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AliasTarget) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (a *jsiiProxy_AliasTarget) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (a *jsiiProxy_AliasTarget) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_AliasTarget) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_AliasTarget) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_AliasTarget) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_AliasTarget) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_AliasTarget) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_AliasTarget) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_AliasTarget) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_AliasTarget) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_AliasTarget) validateImportFromParameters(id *string) error {
	return nil
}

func (a *jsiiProxy_AliasTarget) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_AliasTarget) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (a *jsiiProxy_AliasTarget) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (a *jsiiProxy_AliasTarget) validateMoveToIdParameters(id *string) error {
	return nil
}

func (a *jsiiProxy_AliasTarget) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateAliasTarget_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateAliasTarget_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAliasTarget_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateAliasTarget_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_AliasTarget) validateSetAuthorizeSessionHostIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AliasTarget) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AliasTarget) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AliasTarget) validateSetDescriptionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AliasTarget) validateSetDestinationIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AliasTarget) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_AliasTarget) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AliasTarget) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_AliasTarget) validateSetScopeIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AliasTarget) validateSetTypeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AliasTarget) validateSetValueParameters(val *string) error {
	return nil
}

func validateNewAliasTargetParameters(scope constructs.Construct, id *string, config *AliasTargetConfig) error {
	return nil
}

