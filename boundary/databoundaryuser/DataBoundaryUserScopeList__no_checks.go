// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package databoundaryuser

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataBoundaryUserScopeList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataBoundaryUserScopeList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataBoundaryUserScopeList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataBoundaryUserScopeList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataBoundaryUserScopeList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataBoundaryUserScopeList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataBoundaryUserScopeListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

