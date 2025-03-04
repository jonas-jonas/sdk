/*
 * ORY Hydra
 *
 * Welcome to the ORY Hydra HTTP API documentation. You will find documentation for all HTTP APIs here.
 *
 * API version: v1.11.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// Oauth2TokenResponse The Access Token Response
type Oauth2TokenResponse struct {
	AccessToken *string `json:"access_token,omitempty"`
	ExpiresIn *int64 `json:"expires_in,omitempty"`
	IdToken *string `json:"id_token,omitempty"`
	RefreshToken *string `json:"refresh_token,omitempty"`
	Scope *string `json:"scope,omitempty"`
	TokenType *string `json:"token_type,omitempty"`
}

// NewOauth2TokenResponse instantiates a new Oauth2TokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOauth2TokenResponse() *Oauth2TokenResponse {
	this := Oauth2TokenResponse{}
	return &this
}

// NewOauth2TokenResponseWithDefaults instantiates a new Oauth2TokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOauth2TokenResponseWithDefaults() *Oauth2TokenResponse {
	this := Oauth2TokenResponse{}
	return &this
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *Oauth2TokenResponse) GetAccessToken() string {
	if o == nil || o.AccessToken == nil {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Oauth2TokenResponse) GetAccessTokenOk() (*string, bool) {
	if o == nil || o.AccessToken == nil {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *Oauth2TokenResponse) HasAccessToken() bool {
	if o != nil && o.AccessToken != nil {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *Oauth2TokenResponse) SetAccessToken(v string) {
	o.AccessToken = &v
}

// GetExpiresIn returns the ExpiresIn field value if set, zero value otherwise.
func (o *Oauth2TokenResponse) GetExpiresIn() int64 {
	if o == nil || o.ExpiresIn == nil {
		var ret int64
		return ret
	}
	return *o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Oauth2TokenResponse) GetExpiresInOk() (*int64, bool) {
	if o == nil || o.ExpiresIn == nil {
		return nil, false
	}
	return o.ExpiresIn, true
}

// HasExpiresIn returns a boolean if a field has been set.
func (o *Oauth2TokenResponse) HasExpiresIn() bool {
	if o != nil && o.ExpiresIn != nil {
		return true
	}

	return false
}

// SetExpiresIn gets a reference to the given int64 and assigns it to the ExpiresIn field.
func (o *Oauth2TokenResponse) SetExpiresIn(v int64) {
	o.ExpiresIn = &v
}

// GetIdToken returns the IdToken field value if set, zero value otherwise.
func (o *Oauth2TokenResponse) GetIdToken() string {
	if o == nil || o.IdToken == nil {
		var ret string
		return ret
	}
	return *o.IdToken
}

// GetIdTokenOk returns a tuple with the IdToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Oauth2TokenResponse) GetIdTokenOk() (*string, bool) {
	if o == nil || o.IdToken == nil {
		return nil, false
	}
	return o.IdToken, true
}

// HasIdToken returns a boolean if a field has been set.
func (o *Oauth2TokenResponse) HasIdToken() bool {
	if o != nil && o.IdToken != nil {
		return true
	}

	return false
}

// SetIdToken gets a reference to the given string and assigns it to the IdToken field.
func (o *Oauth2TokenResponse) SetIdToken(v string) {
	o.IdToken = &v
}

// GetRefreshToken returns the RefreshToken field value if set, zero value otherwise.
func (o *Oauth2TokenResponse) GetRefreshToken() string {
	if o == nil || o.RefreshToken == nil {
		var ret string
		return ret
	}
	return *o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Oauth2TokenResponse) GetRefreshTokenOk() (*string, bool) {
	if o == nil || o.RefreshToken == nil {
		return nil, false
	}
	return o.RefreshToken, true
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *Oauth2TokenResponse) HasRefreshToken() bool {
	if o != nil && o.RefreshToken != nil {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given string and assigns it to the RefreshToken field.
func (o *Oauth2TokenResponse) SetRefreshToken(v string) {
	o.RefreshToken = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *Oauth2TokenResponse) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Oauth2TokenResponse) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *Oauth2TokenResponse) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *Oauth2TokenResponse) SetScope(v string) {
	o.Scope = &v
}

// GetTokenType returns the TokenType field value if set, zero value otherwise.
func (o *Oauth2TokenResponse) GetTokenType() string {
	if o == nil || o.TokenType == nil {
		var ret string
		return ret
	}
	return *o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Oauth2TokenResponse) GetTokenTypeOk() (*string, bool) {
	if o == nil || o.TokenType == nil {
		return nil, false
	}
	return o.TokenType, true
}

// HasTokenType returns a boolean if a field has been set.
func (o *Oauth2TokenResponse) HasTokenType() bool {
	if o != nil && o.TokenType != nil {
		return true
	}

	return false
}

// SetTokenType gets a reference to the given string and assigns it to the TokenType field.
func (o *Oauth2TokenResponse) SetTokenType(v string) {
	o.TokenType = &v
}

func (o Oauth2TokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessToken != nil {
		toSerialize["access_token"] = o.AccessToken
	}
	if o.ExpiresIn != nil {
		toSerialize["expires_in"] = o.ExpiresIn
	}
	if o.IdToken != nil {
		toSerialize["id_token"] = o.IdToken
	}
	if o.RefreshToken != nil {
		toSerialize["refresh_token"] = o.RefreshToken
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.TokenType != nil {
		toSerialize["token_type"] = o.TokenType
	}
	return json.Marshal(toSerialize)
}

type NullableOauth2TokenResponse struct {
	value *Oauth2TokenResponse
	isSet bool
}

func (v NullableOauth2TokenResponse) Get() *Oauth2TokenResponse {
	return v.value
}

func (v *NullableOauth2TokenResponse) Set(val *Oauth2TokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOauth2TokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOauth2TokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOauth2TokenResponse(val *Oauth2TokenResponse) *NullableOauth2TokenResponse {
	return &NullableOauth2TokenResponse{value: val, isSet: true}
}

func (v NullableOauth2TokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOauth2TokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


