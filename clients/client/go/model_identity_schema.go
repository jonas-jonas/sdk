/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.48
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// IdentitySchema Together the name and identity uuid are a unique index constraint. This prevents a user from having schemas with the same name. This also allows schemas to have the same name across the system.
type IdentitySchema struct {
	// The Schema's Creation Date
	CreatedAt time.Time `json:"created_at"`
	Id string `json:"id"`
	// The schema name  This is set by the user and is for them to easily recognise their schema
	Name string `json:"name"`
	// The actual Identity JSON Schema
	Schema map[string]interface{} `json:"schema,omitempty"`
	// Last Time Schema was Updated
	UpdatedAt time.Time `json:"updated_at"`
}

// NewIdentitySchema instantiates a new IdentitySchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentitySchema(createdAt time.Time, id string, name string, updatedAt time.Time) *IdentitySchema {
	this := IdentitySchema{}
	this.CreatedAt = createdAt
	this.Id = id
	this.Name = name
	this.UpdatedAt = updatedAt
	return &this
}

// NewIdentitySchemaWithDefaults instantiates a new IdentitySchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentitySchemaWithDefaults() *IdentitySchema {
	this := IdentitySchema{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *IdentitySchema) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *IdentitySchema) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *IdentitySchema) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetId returns the Id field value
func (o *IdentitySchema) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IdentitySchema) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *IdentitySchema) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *IdentitySchema) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IdentitySchema) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IdentitySchema) SetName(v string) {
	o.Name = v
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *IdentitySchema) GetSchema() map[string]interface{} {
	if o == nil || o.Schema == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySchema) GetSchemaOk() (map[string]interface{}, bool) {
	if o == nil || o.Schema == nil {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *IdentitySchema) HasSchema() bool {
	if o != nil && o.Schema != nil {
		return true
	}

	return false
}

// SetSchema gets a reference to the given map[string]interface{} and assigns it to the Schema field.
func (o *IdentitySchema) SetSchema(v map[string]interface{}) {
	o.Schema = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *IdentitySchema) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *IdentitySchema) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *IdentitySchema) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o IdentitySchema) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Schema != nil {
		toSerialize["schema"] = o.Schema
	}
	if true {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableIdentitySchema struct {
	value *IdentitySchema
	isSet bool
}

func (v NullableIdentitySchema) Get() *IdentitySchema {
	return v.value
}

func (v *NullableIdentitySchema) Set(val *IdentitySchema) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentitySchema) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentitySchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentitySchema(val *IdentitySchema) *NullableIdentitySchema {
	return &NullableIdentitySchema{value: val, isSet: true}
}

func (v NullableIdentitySchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentitySchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


