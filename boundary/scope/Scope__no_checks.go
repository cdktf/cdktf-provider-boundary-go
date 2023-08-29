// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package scope

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Scope) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (s *jsiiProxy_Scope) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Scope) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Scope) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Scope) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Scope) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Scope) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Scope) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Scope) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Scope) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Scope) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Scope) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateScope_IsConstructParameters(x interface{}) error {
	return nil
}

func validateScope_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateScope_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Scope) validateSetAutoCreateAdminRoleParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Scope) validateSetAutoCreateDefaultRoleParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Scope) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Scope) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Scope) validateSetDescriptionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Scope) validateSetGlobalScopeParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Scope) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Scope) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Scope) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Scope) validateSetScopeIdParameters(val *string) error {
	return nil
}

func validateNewScopeParameters(scope constructs.Construct, id *string, config *ScopeConfig) error {
	return nil
}

