/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.55
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// SubmitSelfServiceSettingsFlowBody - struct for SubmitSelfServiceSettingsFlowBody
type SubmitSelfServiceSettingsFlowBody struct {
	SubmitSelfServiceSettingsFlowWithOidcMethodBody *SubmitSelfServiceSettingsFlowWithOidcMethodBody
	SubmitSelfServiceSettingsFlowWithPasswordMethodBody *SubmitSelfServiceSettingsFlowWithPasswordMethodBody
	SubmitSelfServiceSettingsFlowWithProfileMethodBody *SubmitSelfServiceSettingsFlowWithProfileMethodBody
	SubmitSelfServiceSettingsFlowWithTotpMethodBody *SubmitSelfServiceSettingsFlowWithTotpMethodBody
}

// SubmitSelfServiceSettingsFlowWithOidcMethodBodyAsSubmitSelfServiceSettingsFlowBody is a convenience function that returns SubmitSelfServiceSettingsFlowWithOidcMethodBody wrapped in SubmitSelfServiceSettingsFlowBody
func SubmitSelfServiceSettingsFlowWithOidcMethodBodyAsSubmitSelfServiceSettingsFlowBody(v *SubmitSelfServiceSettingsFlowWithOidcMethodBody) SubmitSelfServiceSettingsFlowBody {
	return SubmitSelfServiceSettingsFlowBody{
		SubmitSelfServiceSettingsFlowWithOidcMethodBody: v,
	}
}

// SubmitSelfServiceSettingsFlowWithPasswordMethodBodyAsSubmitSelfServiceSettingsFlowBody is a convenience function that returns SubmitSelfServiceSettingsFlowWithPasswordMethodBody wrapped in SubmitSelfServiceSettingsFlowBody
func SubmitSelfServiceSettingsFlowWithPasswordMethodBodyAsSubmitSelfServiceSettingsFlowBody(v *SubmitSelfServiceSettingsFlowWithPasswordMethodBody) SubmitSelfServiceSettingsFlowBody {
	return SubmitSelfServiceSettingsFlowBody{
		SubmitSelfServiceSettingsFlowWithPasswordMethodBody: v,
	}
}

// SubmitSelfServiceSettingsFlowWithProfileMethodBodyAsSubmitSelfServiceSettingsFlowBody is a convenience function that returns SubmitSelfServiceSettingsFlowWithProfileMethodBody wrapped in SubmitSelfServiceSettingsFlowBody
func SubmitSelfServiceSettingsFlowWithProfileMethodBodyAsSubmitSelfServiceSettingsFlowBody(v *SubmitSelfServiceSettingsFlowWithProfileMethodBody) SubmitSelfServiceSettingsFlowBody {
	return SubmitSelfServiceSettingsFlowBody{
		SubmitSelfServiceSettingsFlowWithProfileMethodBody: v,
	}
}

// SubmitSelfServiceSettingsFlowWithTotpMethodBodyAsSubmitSelfServiceSettingsFlowBody is a convenience function that returns SubmitSelfServiceSettingsFlowWithTotpMethodBody wrapped in SubmitSelfServiceSettingsFlowBody
func SubmitSelfServiceSettingsFlowWithTotpMethodBodyAsSubmitSelfServiceSettingsFlowBody(v *SubmitSelfServiceSettingsFlowWithTotpMethodBody) SubmitSelfServiceSettingsFlowBody {
	return SubmitSelfServiceSettingsFlowBody{
		SubmitSelfServiceSettingsFlowWithTotpMethodBody: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *SubmitSelfServiceSettingsFlowBody) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SubmitSelfServiceSettingsFlowWithOidcMethodBody
	err = newStrictDecoder(data).Decode(&dst.SubmitSelfServiceSettingsFlowWithOidcMethodBody)
	if err == nil {
		jsonSubmitSelfServiceSettingsFlowWithOidcMethodBody, _ := json.Marshal(dst.SubmitSelfServiceSettingsFlowWithOidcMethodBody)
		if string(jsonSubmitSelfServiceSettingsFlowWithOidcMethodBody) == "{}" { // empty struct
			dst.SubmitSelfServiceSettingsFlowWithOidcMethodBody = nil
		} else {
			match++
		}
	} else {
		dst.SubmitSelfServiceSettingsFlowWithOidcMethodBody = nil
	}

	// try to unmarshal data into SubmitSelfServiceSettingsFlowWithPasswordMethodBody
	err = newStrictDecoder(data).Decode(&dst.SubmitSelfServiceSettingsFlowWithPasswordMethodBody)
	if err == nil {
		jsonSubmitSelfServiceSettingsFlowWithPasswordMethodBody, _ := json.Marshal(dst.SubmitSelfServiceSettingsFlowWithPasswordMethodBody)
		if string(jsonSubmitSelfServiceSettingsFlowWithPasswordMethodBody) == "{}" { // empty struct
			dst.SubmitSelfServiceSettingsFlowWithPasswordMethodBody = nil
		} else {
			match++
		}
	} else {
		dst.SubmitSelfServiceSettingsFlowWithPasswordMethodBody = nil
	}

	// try to unmarshal data into SubmitSelfServiceSettingsFlowWithProfileMethodBody
	err = newStrictDecoder(data).Decode(&dst.SubmitSelfServiceSettingsFlowWithProfileMethodBody)
	if err == nil {
		jsonSubmitSelfServiceSettingsFlowWithProfileMethodBody, _ := json.Marshal(dst.SubmitSelfServiceSettingsFlowWithProfileMethodBody)
		if string(jsonSubmitSelfServiceSettingsFlowWithProfileMethodBody) == "{}" { // empty struct
			dst.SubmitSelfServiceSettingsFlowWithProfileMethodBody = nil
		} else {
			match++
		}
	} else {
		dst.SubmitSelfServiceSettingsFlowWithProfileMethodBody = nil
	}

	// try to unmarshal data into SubmitSelfServiceSettingsFlowWithTotpMethodBody
	err = newStrictDecoder(data).Decode(&dst.SubmitSelfServiceSettingsFlowWithTotpMethodBody)
	if err == nil {
		jsonSubmitSelfServiceSettingsFlowWithTotpMethodBody, _ := json.Marshal(dst.SubmitSelfServiceSettingsFlowWithTotpMethodBody)
		if string(jsonSubmitSelfServiceSettingsFlowWithTotpMethodBody) == "{}" { // empty struct
			dst.SubmitSelfServiceSettingsFlowWithTotpMethodBody = nil
		} else {
			match++
		}
	} else {
		dst.SubmitSelfServiceSettingsFlowWithTotpMethodBody = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.SubmitSelfServiceSettingsFlowWithOidcMethodBody = nil
		dst.SubmitSelfServiceSettingsFlowWithPasswordMethodBody = nil
		dst.SubmitSelfServiceSettingsFlowWithProfileMethodBody = nil
		dst.SubmitSelfServiceSettingsFlowWithTotpMethodBody = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(SubmitSelfServiceSettingsFlowBody)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(SubmitSelfServiceSettingsFlowBody)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SubmitSelfServiceSettingsFlowBody) MarshalJSON() ([]byte, error) {
	if src.SubmitSelfServiceSettingsFlowWithOidcMethodBody != nil {
		return json.Marshal(&src.SubmitSelfServiceSettingsFlowWithOidcMethodBody)
	}

	if src.SubmitSelfServiceSettingsFlowWithPasswordMethodBody != nil {
		return json.Marshal(&src.SubmitSelfServiceSettingsFlowWithPasswordMethodBody)
	}

	if src.SubmitSelfServiceSettingsFlowWithProfileMethodBody != nil {
		return json.Marshal(&src.SubmitSelfServiceSettingsFlowWithProfileMethodBody)
	}

	if src.SubmitSelfServiceSettingsFlowWithTotpMethodBody != nil {
		return json.Marshal(&src.SubmitSelfServiceSettingsFlowWithTotpMethodBody)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SubmitSelfServiceSettingsFlowBody) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.SubmitSelfServiceSettingsFlowWithOidcMethodBody != nil {
		return obj.SubmitSelfServiceSettingsFlowWithOidcMethodBody
	}

	if obj.SubmitSelfServiceSettingsFlowWithPasswordMethodBody != nil {
		return obj.SubmitSelfServiceSettingsFlowWithPasswordMethodBody
	}

	if obj.SubmitSelfServiceSettingsFlowWithProfileMethodBody != nil {
		return obj.SubmitSelfServiceSettingsFlowWithProfileMethodBody
	}

	if obj.SubmitSelfServiceSettingsFlowWithTotpMethodBody != nil {
		return obj.SubmitSelfServiceSettingsFlowWithTotpMethodBody
	}

	// all schemas are nil
	return nil
}

type NullableSubmitSelfServiceSettingsFlowBody struct {
	value *SubmitSelfServiceSettingsFlowBody
	isSet bool
}

func (v NullableSubmitSelfServiceSettingsFlowBody) Get() *SubmitSelfServiceSettingsFlowBody {
	return v.value
}

func (v *NullableSubmitSelfServiceSettingsFlowBody) Set(val *SubmitSelfServiceSettingsFlowBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitSelfServiceSettingsFlowBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitSelfServiceSettingsFlowBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitSelfServiceSettingsFlowBody(val *SubmitSelfServiceSettingsFlowBody) *NullableSubmitSelfServiceSettingsFlowBody {
	return &NullableSubmitSelfServiceSettingsFlowBody{value: val, isSet: true}
}

func (v NullableSubmitSelfServiceSettingsFlowBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmitSelfServiceSettingsFlowBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


