// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BoundaryProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (b *jsiiProxy_BoundaryProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateBoundaryProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateBoundaryProvider_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateBoundaryProvider_IsTerraformProviderParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_BoundaryProvider) validateSetTlsInsecureParameters(val interface{}) error {
	return nil
}

func validateNewBoundaryProviderParameters(scope constructs.Construct, id *string, config *BoundaryProviderConfig) error {
	return nil
}

