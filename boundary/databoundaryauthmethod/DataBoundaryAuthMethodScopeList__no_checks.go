// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package databoundaryauthmethod

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataBoundaryAuthMethodScopeList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataBoundaryAuthMethodScopeList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataBoundaryAuthMethodScopeList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataBoundaryAuthMethodScopeList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataBoundaryAuthMethodScopeList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataBoundaryAuthMethodScopeListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

