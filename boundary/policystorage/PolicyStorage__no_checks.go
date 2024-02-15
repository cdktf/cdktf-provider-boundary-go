// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package policystorage

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PolicyStorage) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (p *jsiiProxy_PolicyStorage) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (p *jsiiProxy_PolicyStorage) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_PolicyStorage) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_PolicyStorage) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_PolicyStorage) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_PolicyStorage) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_PolicyStorage) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_PolicyStorage) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_PolicyStorage) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_PolicyStorage) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_PolicyStorage) validateImportFromParameters(id *string) error {
	return nil
}

func (p *jsiiProxy_PolicyStorage) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_PolicyStorage) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (p *jsiiProxy_PolicyStorage) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (p *jsiiProxy_PolicyStorage) validateMoveToIdParameters(id *string) error {
	return nil
}

func (p *jsiiProxy_PolicyStorage) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validatePolicyStorage_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validatePolicyStorage_IsConstructParameters(x interface{}) error {
	return nil
}

func validatePolicyStorage_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validatePolicyStorage_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_PolicyStorage) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PolicyStorage) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PolicyStorage) validateSetDeleteAfterDaysParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_PolicyStorage) validateSetDeleteAfterOverridableParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PolicyStorage) validateSetDescriptionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PolicyStorage) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_PolicyStorage) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PolicyStorage) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_PolicyStorage) validateSetRetainForDaysParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_PolicyStorage) validateSetRetainForOverridableParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PolicyStorage) validateSetScopeIdParameters(val *string) error {
	return nil
}

func validateNewPolicyStorageParameters(scope constructs.Construct, id *string, config *PolicyStorageConfig) error {
	return nil
}

