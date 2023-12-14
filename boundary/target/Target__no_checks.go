// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package target

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_Target) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (t *jsiiProxy_Target) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (t *jsiiProxy_Target) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_Target) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_Target) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_Target) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_Target) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_Target) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_Target) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_Target) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_Target) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_Target) validateImportFromParameters(id *string) error {
	return nil
}

func (t *jsiiProxy_Target) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_Target) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (t *jsiiProxy_Target) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (t *jsiiProxy_Target) validateMoveToIdParameters(id *string) error {
	return nil
}

func (t *jsiiProxy_Target) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateTarget_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateTarget_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTarget_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateTarget_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Target) validateSetAddressParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Target) validateSetBrokeredCredentialSourceIdsParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_Target) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Target) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Target) validateSetDefaultClientPortParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Target) validateSetDefaultPortParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Target) validateSetDescriptionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Target) validateSetEgressWorkerFilterParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Target) validateSetEnableSessionRecordingParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Target) validateSetHostSourceIdsParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_Target) validateSetIngressWorkerFilterParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Target) validateSetInjectedApplicationCredentialSourceIdsParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_Target) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Target) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Target) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Target) validateSetScopeIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Target) validateSetSessionConnectionLimitParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Target) validateSetSessionMaxSecondsParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Target) validateSetStorageBucketIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Target) validateSetTypeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Target) validateSetWorkerFilterParameters(val *string) error {
	return nil
}

func validateNewTargetParameters(scope constructs.Construct, id *string, config *TargetConfig) error {
	return nil
}

