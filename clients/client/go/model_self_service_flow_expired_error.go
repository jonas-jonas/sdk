/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v0.2.0-alpha.48
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// SelfServiceFlowExpiredError Is sent when a flow is expired
type SelfServiceFlowExpiredError struct {
	// The status code
	Code *int64 `json:"code,omitempty"`
	// Debug information  This field is often not exposed to protect against leaking sensitive information.
	Debug *string `json:"debug,omitempty"`
	// Further error details
	Details map[string]interface{} `json:"details,omitempty"`
	// The error ID  Useful when trying to identify various errors in application logic.
	Id *string `json:"id,omitempty"`
	// Error message  The error's message.
	Message string `json:"message"`
	// A human-readable reason for the error
	Reason *string `json:"reason,omitempty"`
	// The request ID  The request ID is often exposed internally in order to trace errors across service architectures. This is often a UUID.
	Request *string `json:"request,omitempty"`
	// A Duration represents the elapsed time between two instants as an int64 nanosecond count. The representation limits the largest representable duration to approximately 290 years.
	Since *int64 `json:"since,omitempty"`
	// The status description
	Status *string `json:"status,omitempty"`
	UseFlowId *string `json:"use_flow_id,omitempty"`
}

// NewSelfServiceFlowExpiredError instantiates a new SelfServiceFlowExpiredError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelfServiceFlowExpiredError(message string) *SelfServiceFlowExpiredError {
	this := SelfServiceFlowExpiredError{}
	this.Message = message
	return &this
}

// NewSelfServiceFlowExpiredErrorWithDefaults instantiates a new SelfServiceFlowExpiredError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelfServiceFlowExpiredErrorWithDefaults() *SelfServiceFlowExpiredError {
	this := SelfServiceFlowExpiredError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *SelfServiceFlowExpiredError) GetCode() int64 {
	if o == nil || o.Code == nil {
		var ret int64
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceFlowExpiredError) GetCodeOk() (*int64, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *SelfServiceFlowExpiredError) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given int64 and assigns it to the Code field.
func (o *SelfServiceFlowExpiredError) SetCode(v int64) {
	o.Code = &v
}

// GetDebug returns the Debug field value if set, zero value otherwise.
func (o *SelfServiceFlowExpiredError) GetDebug() string {
	if o == nil || o.Debug == nil {
		var ret string
		return ret
	}
	return *o.Debug
}

// GetDebugOk returns a tuple with the Debug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceFlowExpiredError) GetDebugOk() (*string, bool) {
	if o == nil || o.Debug == nil {
		return nil, false
	}
	return o.Debug, true
}

// HasDebug returns a boolean if a field has been set.
func (o *SelfServiceFlowExpiredError) HasDebug() bool {
	if o != nil && o.Debug != nil {
		return true
	}

	return false
}

// SetDebug gets a reference to the given string and assigns it to the Debug field.
func (o *SelfServiceFlowExpiredError) SetDebug(v string) {
	o.Debug = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *SelfServiceFlowExpiredError) GetDetails() map[string]interface{} {
	if o == nil || o.Details == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceFlowExpiredError) GetDetailsOk() (map[string]interface{}, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *SelfServiceFlowExpiredError) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given map[string]interface{} and assigns it to the Details field.
func (o *SelfServiceFlowExpiredError) SetDetails(v map[string]interface{}) {
	o.Details = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SelfServiceFlowExpiredError) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceFlowExpiredError) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SelfServiceFlowExpiredError) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SelfServiceFlowExpiredError) SetId(v string) {
	o.Id = &v
}

// GetMessage returns the Message field value
func (o *SelfServiceFlowExpiredError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SelfServiceFlowExpiredError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SelfServiceFlowExpiredError) SetMessage(v string) {
	o.Message = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *SelfServiceFlowExpiredError) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceFlowExpiredError) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *SelfServiceFlowExpiredError) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *SelfServiceFlowExpiredError) SetReason(v string) {
	o.Reason = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *SelfServiceFlowExpiredError) GetRequest() string {
	if o == nil || o.Request == nil {
		var ret string
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceFlowExpiredError) GetRequestOk() (*string, bool) {
	if o == nil || o.Request == nil {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *SelfServiceFlowExpiredError) HasRequest() bool {
	if o != nil && o.Request != nil {
		return true
	}

	return false
}

// SetRequest gets a reference to the given string and assigns it to the Request field.
func (o *SelfServiceFlowExpiredError) SetRequest(v string) {
	o.Request = &v
}

// GetSince returns the Since field value if set, zero value otherwise.
func (o *SelfServiceFlowExpiredError) GetSince() int64 {
	if o == nil || o.Since == nil {
		var ret int64
		return ret
	}
	return *o.Since
}

// GetSinceOk returns a tuple with the Since field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceFlowExpiredError) GetSinceOk() (*int64, bool) {
	if o == nil || o.Since == nil {
		return nil, false
	}
	return o.Since, true
}

// HasSince returns a boolean if a field has been set.
func (o *SelfServiceFlowExpiredError) HasSince() bool {
	if o != nil && o.Since != nil {
		return true
	}

	return false
}

// SetSince gets a reference to the given int64 and assigns it to the Since field.
func (o *SelfServiceFlowExpiredError) SetSince(v int64) {
	o.Since = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SelfServiceFlowExpiredError) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceFlowExpiredError) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SelfServiceFlowExpiredError) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SelfServiceFlowExpiredError) SetStatus(v string) {
	o.Status = &v
}

// GetUseFlowId returns the UseFlowId field value if set, zero value otherwise.
func (o *SelfServiceFlowExpiredError) GetUseFlowId() string {
	if o == nil || o.UseFlowId == nil {
		var ret string
		return ret
	}
	return *o.UseFlowId
}

// GetUseFlowIdOk returns a tuple with the UseFlowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceFlowExpiredError) GetUseFlowIdOk() (*string, bool) {
	if o == nil || o.UseFlowId == nil {
		return nil, false
	}
	return o.UseFlowId, true
}

// HasUseFlowId returns a boolean if a field has been set.
func (o *SelfServiceFlowExpiredError) HasUseFlowId() bool {
	if o != nil && o.UseFlowId != nil {
		return true
	}

	return false
}

// SetUseFlowId gets a reference to the given string and assigns it to the UseFlowId field.
func (o *SelfServiceFlowExpiredError) SetUseFlowId(v string) {
	o.UseFlowId = &v
}

func (o SelfServiceFlowExpiredError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Debug != nil {
		toSerialize["debug"] = o.Debug
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["message"] = o.Message
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.Request != nil {
		toSerialize["request"] = o.Request
	}
	if o.Since != nil {
		toSerialize["since"] = o.Since
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.UseFlowId != nil {
		toSerialize["use_flow_id"] = o.UseFlowId
	}
	return json.Marshal(toSerialize)
}

type NullableSelfServiceFlowExpiredError struct {
	value *SelfServiceFlowExpiredError
	isSet bool
}

func (v NullableSelfServiceFlowExpiredError) Get() *SelfServiceFlowExpiredError {
	return v.value
}

func (v *NullableSelfServiceFlowExpiredError) Set(val *SelfServiceFlowExpiredError) {
	v.value = val
	v.isSet = true
}

func (v NullableSelfServiceFlowExpiredError) IsSet() bool {
	return v.isSet
}

func (v *NullableSelfServiceFlowExpiredError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelfServiceFlowExpiredError(val *SelfServiceFlowExpiredError) *NullableSelfServiceFlowExpiredError {
	return &NullableSelfServiceFlowExpiredError{value: val, isSet: true}
}

func (v NullableSelfServiceFlowExpiredError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelfServiceFlowExpiredError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


