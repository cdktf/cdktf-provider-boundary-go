// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package databoundaryaccount

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataBoundaryAccountScopeList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataBoundaryAccountScopeList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataBoundaryAccountScopeList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataBoundaryAccountScopeList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataBoundaryAccountScopeList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataBoundaryAccountScopeList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataBoundaryAccountScopeListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

