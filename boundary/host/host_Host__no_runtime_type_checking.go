//go:build no_runtime_type_checking

package host

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_Host) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (h *jsiiProxy_Host) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (h *jsiiProxy_Host) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (h *jsiiProxy_Host) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (h *jsiiProxy_Host) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (h *jsiiProxy_Host) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (h *jsiiProxy_Host) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (h *jsiiProxy_Host) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (h *jsiiProxy_Host) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (h *jsiiProxy_Host) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (h *jsiiProxy_Host) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (h *jsiiProxy_Host) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateHost_IsConstructParameters(x interface{}) error {
	return nil
}

func validateHost_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateHost_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Host) validateSetAddressParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Host) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Host) validateSetDescriptionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Host) validateSetHostCatalogIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Host) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Host) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Host) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Host) validateSetTypeParameters(val *string) error {
	return nil
}

func validateNewHostParameters(scope constructs.Construct, id *string, config *HostConfig) error {
	return nil
}

