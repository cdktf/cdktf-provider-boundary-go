// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package databoundarygroup

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataBoundaryGroupScopeList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataBoundaryGroupScopeList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataBoundaryGroupScopeList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataBoundaryGroupScopeList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataBoundaryGroupScopeList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataBoundaryGroupScopeListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

