# ory/client-php

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed
with a valid Personal Access Token. Public APIs are mostly used in browsers.



## Installation & Usage

### Requirements

PHP 7.3 and later.
Should also work with PHP 8.0 but has not been tested.

### Composer

To install the bindings via [Composer](https://getcomposer.org/), add the following to `composer.json`:

```json
{
  "repositories": [
    {
      "type": "vcs",
      "url": "https://github.com/ory/client-php.git"
    }
  ],
  "require": {
    "ory/client-php": "*@dev"
  }
}
```

Then run `composer install`

### Manual Installation

Download the files and include `autoload.php`:

```php
<?php
require_once('/path/to/ory/client-php/vendor/autoload.php');
```

## Getting Started

Please follow the [installation procedure](#installation--usage) and then run the following:

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');




$apiInstance = new Ory\Client\Api\AdminApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client()
);
$consentChallenge = 'consentChallenge_example'; // string
$acceptConsentRequest = new \Ory\Client\Model\AcceptConsentRequest(); // \Ory\Client\Model\AcceptConsentRequest

try {
    $result = $apiInstance->acceptConsentRequest($consentChallenge, $acceptConsentRequest);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AdminApi->acceptConsentRequest: ', $e->getMessage(), PHP_EOL;
}

```

## API Endpoints

All URIs are relative to *https://playground.projects.oryapis.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AdminApi* | [**acceptConsentRequest**](docs/Api/AdminApi.md#acceptconsentrequest) | **PUT** /oauth2/auth/requests/consent/accept | Accept a Consent Request
*AdminApi* | [**acceptLoginRequest**](docs/Api/AdminApi.md#acceptloginrequest) | **PUT** /oauth2/auth/requests/login/accept | Accept a Login Request
*AdminApi* | [**acceptLogoutRequest**](docs/Api/AdminApi.md#acceptlogoutrequest) | **PUT** /oauth2/auth/requests/logout/accept | Accept a Logout Request
*AdminApi* | [**createJsonWebKeySet**](docs/Api/AdminApi.md#createjsonwebkeyset) | **POST** /keys/{set} | Generate a New JSON Web Key
*AdminApi* | [**createOAuth2Client**](docs/Api/AdminApi.md#createoauth2client) | **POST** /clients | Create an OAuth 2.0 Client
*AdminApi* | [**deleteJsonWebKey**](docs/Api/AdminApi.md#deletejsonwebkey) | **DELETE** /keys/{set}/{kid} | Delete a JSON Web Key
*AdminApi* | [**deleteJsonWebKeySet**](docs/Api/AdminApi.md#deletejsonwebkeyset) | **DELETE** /keys/{set} | Delete a JSON Web Key Set
*AdminApi* | [**deleteOAuth2Client**](docs/Api/AdminApi.md#deleteoauth2client) | **DELETE** /clients/{id} | Deletes an OAuth 2.0 Client
*AdminApi* | [**deleteOAuth2Token**](docs/Api/AdminApi.md#deleteoauth2token) | **DELETE** /oauth2/tokens | Delete OAuth2 Access Tokens from a Client
*AdminApi* | [**deleteTrustedJwtGrantIssuer**](docs/Api/AdminApi.md#deletetrustedjwtgrantissuer) | **DELETE** /trust/grants/jwt-bearer/issuers/{id} | Delete a Trusted OAuth2 JWT Bearer Grant Type Issuer
*AdminApi* | [**flushInactiveOAuth2Tokens**](docs/Api/AdminApi.md#flushinactiveoauth2tokens) | **POST** /oauth2/flush | Flush Expired OAuth2 Access Tokens
*AdminApi* | [**getConsentRequest**](docs/Api/AdminApi.md#getconsentrequest) | **GET** /oauth2/auth/requests/consent | Get Consent Request Information
*AdminApi* | [**getJsonWebKey**](docs/Api/AdminApi.md#getjsonwebkey) | **GET** /keys/{set}/{kid} | Fetch a JSON Web Key
*AdminApi* | [**getJsonWebKeySet**](docs/Api/AdminApi.md#getjsonwebkeyset) | **GET** /keys/{set} | Retrieve a JSON Web Key Set
*AdminApi* | [**getLoginRequest**](docs/Api/AdminApi.md#getloginrequest) | **GET** /oauth2/auth/requests/login | Get a Login Request
*AdminApi* | [**getLogoutRequest**](docs/Api/AdminApi.md#getlogoutrequest) | **GET** /oauth2/auth/requests/logout | Get a Logout Request
*AdminApi* | [**getOAuth2Client**](docs/Api/AdminApi.md#getoauth2client) | **GET** /clients/{id} | Get an OAuth 2.0 Client
*AdminApi* | [**getTrustedJwtGrantIssuer**](docs/Api/AdminApi.md#gettrustedjwtgrantissuer) | **GET** /trust/grants/jwt-bearer/issuers/{id} | Get a Trusted OAuth2 JWT Bearer Grant Type Issuer
*AdminApi* | [**introspectOAuth2Token**](docs/Api/AdminApi.md#introspectoauth2token) | **POST** /oauth2/introspect | Introspect OAuth2 Tokens
*AdminApi* | [**listOAuth2Clients**](docs/Api/AdminApi.md#listoauth2clients) | **GET** /clients | List OAuth 2.0 Clients
*AdminApi* | [**listSubjectConsentSessions**](docs/Api/AdminApi.md#listsubjectconsentsessions) | **GET** /oauth2/auth/sessions/consent | Lists All Consent Sessions of a Subject
*AdminApi* | [**listTrustedJwtGrantIssuers**](docs/Api/AdminApi.md#listtrustedjwtgrantissuers) | **GET** /trust/grants/jwt-bearer/issuers | List Trusted OAuth2 JWT Bearer Grant Type Issuers
*AdminApi* | [**patchOAuth2Client**](docs/Api/AdminApi.md#patchoauth2client) | **PATCH** /clients/{id} | Patch an OAuth 2.0 Client
*AdminApi* | [**rejectConsentRequest**](docs/Api/AdminApi.md#rejectconsentrequest) | **PUT** /oauth2/auth/requests/consent/reject | Reject a Consent Request
*AdminApi* | [**rejectLoginRequest**](docs/Api/AdminApi.md#rejectloginrequest) | **PUT** /oauth2/auth/requests/login/reject | Reject a Login Request
*AdminApi* | [**rejectLogoutRequest**](docs/Api/AdminApi.md#rejectlogoutrequest) | **PUT** /oauth2/auth/requests/logout/reject | Reject a Logout Request
*AdminApi* | [**revokeAuthenticationSession**](docs/Api/AdminApi.md#revokeauthenticationsession) | **DELETE** /oauth2/auth/sessions/login | Invalidates All Login Sessions of a Certain User Invalidates a Subject&#39;s Authentication Session
*AdminApi* | [**revokeConsentSessions**](docs/Api/AdminApi.md#revokeconsentsessions) | **DELETE** /oauth2/auth/sessions/consent | Revokes Consent Sessions of a Subject for a Specific OAuth 2.0 Client
*AdminApi* | [**trustJwtGrantIssuer**](docs/Api/AdminApi.md#trustjwtgrantissuer) | **POST** /trust/grants/jwt-bearer/issuers | Trust an OAuth2 JWT Bearer Grant Type Issuer
*AdminApi* | [**updateJsonWebKey**](docs/Api/AdminApi.md#updatejsonwebkey) | **PUT** /keys/{set}/{kid} | Update a JSON Web Key
*AdminApi* | [**updateJsonWebKeySet**](docs/Api/AdminApi.md#updatejsonwebkeyset) | **PUT** /keys/{set} | Update a JSON Web Key Set
*AdminApi* | [**updateOAuth2Client**](docs/Api/AdminApi.md#updateoauth2client) | **PUT** /clients/{id} | Update an OAuth 2.0 Client
*MetadataApi* | [**getVersion**](docs/Api/MetadataApi.md#getversion) | **GET** /version | Return Running Software Version.
*MetadataApi* | [**isAlive**](docs/Api/MetadataApi.md#isalive) | **GET** /health/alive | Check HTTP Server Status
*MetadataApi* | [**isReady**](docs/Api/MetadataApi.md#isready) | **GET** /health/ready | Check HTTP Server and Database Status
*PublicApi* | [**disconnectUser**](docs/Api/PublicApi.md#disconnectuser) | **GET** /oauth2/sessions/logout | OpenID Connect Front-Backchannel Enabled Logout
*PublicApi* | [**discoverOpenIDConfiguration**](docs/Api/PublicApi.md#discoveropenidconfiguration) | **GET** /.well-known/openid-configuration | OpenID Connect Discovery
*PublicApi* | [**dynamicClientRegistrationCreateOAuth2Client**](docs/Api/PublicApi.md#dynamicclientregistrationcreateoauth2client) | **POST** /connect/register | Register an OAuth 2.0 Client using the OpenID / OAuth2 Dynamic Client Registration Management Protocol
*PublicApi* | [**dynamicClientRegistrationDeleteOAuth2Client**](docs/Api/PublicApi.md#dynamicclientregistrationdeleteoauth2client) | **DELETE** /connect/register/{id} | Deletes an OAuth 2.0 Client using the OpenID / OAuth2 Dynamic Client Registration Management Protocol
*PublicApi* | [**dynamicClientRegistrationGetOAuth2Client**](docs/Api/PublicApi.md#dynamicclientregistrationgetoauth2client) | **GET** /connect/register/{id} | Get an OAuth 2.0 Client using the OpenID / OAuth2 Dynamic Client Registration Management Protocol
*PublicApi* | [**dynamicClientRegistrationUpdateOAuth2Client**](docs/Api/PublicApi.md#dynamicclientregistrationupdateoauth2client) | **PUT** /connect/register/{id} | Update an OAuth 2.0 Client using the OpenID / OAuth2 Dynamic Client Registration Management Protocol
*PublicApi* | [**oauth2Token**](docs/Api/PublicApi.md#oauth2token) | **POST** /oauth2/token | The OAuth 2.0 Token Endpoint
*PublicApi* | [**oauthAuth**](docs/Api/PublicApi.md#oauthauth) | **GET** /oauth2/auth | The OAuth 2.0 Authorize Endpoint
*PublicApi* | [**revokeOAuth2Token**](docs/Api/PublicApi.md#revokeoauth2token) | **POST** /oauth2/revoke | Revoke OAuth2 Tokens
*PublicApi* | [**userinfo**](docs/Api/PublicApi.md#userinfo) | **GET** /userinfo | OpenID Connect Userinfo
*PublicApi* | [**wellKnown**](docs/Api/PublicApi.md#wellknown) | **GET** /.well-known/jwks.json | JSON Web Keys Discovery
*ReadApi* | [**getCheck**](docs/Api/ReadApi.md#getcheck) | **GET** /relation-tuples/check/openapi | # Check a relation tuple
*ReadApi* | [**getCheckMirrorStatus**](docs/Api/ReadApi.md#getcheckmirrorstatus) | **GET** /relation-tuples/check | # Check a relation tuple
*ReadApi* | [**getExpand**](docs/Api/ReadApi.md#getexpand) | **GET** /relation-tuples/expand | # Expand a Relation Tuple
*ReadApi* | [**getRelationTuples**](docs/Api/ReadApi.md#getrelationtuples) | **GET** /relation-tuples | # Query relation tuples
*ReadApi* | [**postCheck**](docs/Api/ReadApi.md#postcheck) | **POST** /relation-tuples/check/openapi | # Check a relation tuple
*ReadApi* | [**postCheckMirrorStatus**](docs/Api/ReadApi.md#postcheckmirrorstatus) | **POST** /relation-tuples/check | # Check a relation tuple
*V0alpha2Api* | [**adminCreateIdentity**](docs/Api/V0alpha2Api.md#admincreateidentity) | **POST** /admin/identities | Create an Identity
*V0alpha2Api* | [**adminCreateSelfServiceRecoveryLink**](docs/Api/V0alpha2Api.md#admincreateselfservicerecoverylink) | **POST** /admin/recovery/link | Create a Recovery Link
*V0alpha2Api* | [**adminDeleteIdentity**](docs/Api/V0alpha2Api.md#admindeleteidentity) | **DELETE** /admin/identities/{id} | Delete an Identity
*V0alpha2Api* | [**adminDeleteIdentitySessions**](docs/Api/V0alpha2Api.md#admindeleteidentitysessions) | **DELETE** /admin/identities/{id}/sessions | Calling this endpoint irrecoverably and permanently deletes and invalidates all sessions that belong to the given Identity.
*V0alpha2Api* | [**adminExtendSession**](docs/Api/V0alpha2Api.md#adminextendsession) | **PATCH** /admin/sessions/{id}/extend | Calling this endpoint extends the given session ID. If &#x60;session.earliest_possible_extend&#x60; is set it will only extend the session after the specified time has passed.
*V0alpha2Api* | [**adminGetIdentity**](docs/Api/V0alpha2Api.md#admingetidentity) | **GET** /admin/identities/{id} | Get an Identity
*V0alpha2Api* | [**adminListIdentities**](docs/Api/V0alpha2Api.md#adminlistidentities) | **GET** /admin/identities | List Identities
*V0alpha2Api* | [**adminListIdentitySessions**](docs/Api/V0alpha2Api.md#adminlistidentitysessions) | **GET** /admin/identities/{id}/sessions | This endpoint returns all sessions that belong to the given Identity.
*V0alpha2Api* | [**adminPatchIdentity**](docs/Api/V0alpha2Api.md#adminpatchidentity) | **PATCH** /admin/identities/{id} | Partially updates an Identity&#39;s field using [JSON Patch](https://jsonpatch.com/)
*V0alpha2Api* | [**adminUpdateIdentity**](docs/Api/V0alpha2Api.md#adminupdateidentity) | **PUT** /admin/identities/{id} | Update an Identity
*V0alpha2Api* | [**createProject**](docs/Api/V0alpha2Api.md#createproject) | **POST** /projects | # Create a Project
*V0alpha2Api* | [**createProjectApiKey**](docs/Api/V0alpha2Api.md#createprojectapikey) | **POST** /projects/{project}/tokens | # Create API Token
*V0alpha2Api* | [**createSelfServiceLogoutFlowUrlForBrowsers**](docs/Api/V0alpha2Api.md#createselfservicelogoutflowurlforbrowsers) | **GET** /self-service/logout/browser | Create a Logout URL for Browsers
*V0alpha2Api* | [**deleteProjectApiKey**](docs/Api/V0alpha2Api.md#deleteprojectapikey) | **DELETE** /projects/{project}/tokens/{token_id} | # Delete API Token
*V0alpha2Api* | [**getIdentitySchema**](docs/Api/V0alpha2Api.md#getidentityschema) | **GET** /schemas/{id} | 
*V0alpha2Api* | [**getProject**](docs/Api/V0alpha2Api.md#getproject) | **GET** /projects/{project_id} | # Get a Project
*V0alpha2Api* | [**getProjectMembers**](docs/Api/V0alpha2Api.md#getprojectmembers) | **GET** /projects/{project_id}/members | Get all members associated with this project.
*V0alpha2Api* | [**getSelfServiceError**](docs/Api/V0alpha2Api.md#getselfserviceerror) | **GET** /self-service/errors | Get Self-Service Errors
*V0alpha2Api* | [**getSelfServiceLoginFlow**](docs/Api/V0alpha2Api.md#getselfserviceloginflow) | **GET** /self-service/login/flows | Get Login Flow
*V0alpha2Api* | [**getSelfServiceRecoveryFlow**](docs/Api/V0alpha2Api.md#getselfservicerecoveryflow) | **GET** /self-service/recovery/flows | Get Recovery Flow
*V0alpha2Api* | [**getSelfServiceRegistrationFlow**](docs/Api/V0alpha2Api.md#getselfserviceregistrationflow) | **GET** /self-service/registration/flows | Get Registration Flow
*V0alpha2Api* | [**getSelfServiceSettingsFlow**](docs/Api/V0alpha2Api.md#getselfservicesettingsflow) | **GET** /self-service/settings/flows | Get Settings Flow
*V0alpha2Api* | [**getSelfServiceVerificationFlow**](docs/Api/V0alpha2Api.md#getselfserviceverificationflow) | **GET** /self-service/verification/flows | # Get Verification Flow
*V0alpha2Api* | [**getWebAuthnJavaScript**](docs/Api/V0alpha2Api.md#getwebauthnjavascript) | **GET** /.well-known/ory/webauthn.js | Get WebAuthn JavaScript
*V0alpha2Api* | [**initializeSelfServiceLoginFlowForBrowsers**](docs/Api/V0alpha2Api.md#initializeselfserviceloginflowforbrowsers) | **GET** /self-service/login/browser | Initialize Login Flow for Browsers
*V0alpha2Api* | [**initializeSelfServiceLoginFlowWithoutBrowser**](docs/Api/V0alpha2Api.md#initializeselfserviceloginflowwithoutbrowser) | **GET** /self-service/login/api | Initialize Login Flow for APIs, Services, Apps, ...
*V0alpha2Api* | [**initializeSelfServiceRecoveryFlowForBrowsers**](docs/Api/V0alpha2Api.md#initializeselfservicerecoveryflowforbrowsers) | **GET** /self-service/recovery/browser | Initialize Recovery Flow for Browsers
*V0alpha2Api* | [**initializeSelfServiceRecoveryFlowWithoutBrowser**](docs/Api/V0alpha2Api.md#initializeselfservicerecoveryflowwithoutbrowser) | **GET** /self-service/recovery/api | Initialize Recovery Flow for APIs, Services, Apps, ...
*V0alpha2Api* | [**initializeSelfServiceRegistrationFlowForBrowsers**](docs/Api/V0alpha2Api.md#initializeselfserviceregistrationflowforbrowsers) | **GET** /self-service/registration/browser | Initialize Registration Flow for Browsers
*V0alpha2Api* | [**initializeSelfServiceRegistrationFlowWithoutBrowser**](docs/Api/V0alpha2Api.md#initializeselfserviceregistrationflowwithoutbrowser) | **GET** /self-service/registration/api | Initialize Registration Flow for APIs, Services, Apps, ...
*V0alpha2Api* | [**initializeSelfServiceSettingsFlowForBrowsers**](docs/Api/V0alpha2Api.md#initializeselfservicesettingsflowforbrowsers) | **GET** /self-service/settings/browser | Initialize Settings Flow for Browsers
*V0alpha2Api* | [**initializeSelfServiceSettingsFlowWithoutBrowser**](docs/Api/V0alpha2Api.md#initializeselfservicesettingsflowwithoutbrowser) | **GET** /self-service/settings/api | Initialize Settings Flow for APIs, Services, Apps, ...
*V0alpha2Api* | [**initializeSelfServiceVerificationFlowForBrowsers**](docs/Api/V0alpha2Api.md#initializeselfserviceverificationflowforbrowsers) | **GET** /self-service/verification/browser | Initialize Verification Flow for Browser Clients
*V0alpha2Api* | [**initializeSelfServiceVerificationFlowWithoutBrowser**](docs/Api/V0alpha2Api.md#initializeselfserviceverificationflowwithoutbrowser) | **GET** /self-service/verification/api | Initialize Verification Flow for APIs, Services, Apps, ...
*V0alpha2Api* | [**listIdentitySchemas**](docs/Api/V0alpha2Api.md#listidentityschemas) | **GET** /schemas | 
*V0alpha2Api* | [**listProjectApiKeys**](docs/Api/V0alpha2Api.md#listprojectapikeys) | **GET** /projects/{project}/tokens | # List a Project&#39;s API Tokens
*V0alpha2Api* | [**listProjects**](docs/Api/V0alpha2Api.md#listprojects) | **GET** /projects | # List All Projects
*V0alpha2Api* | [**listSessions**](docs/Api/V0alpha2Api.md#listsessions) | **GET** /sessions | This endpoints returns all other active sessions that belong to the logged-in user. The current session can be retrieved by calling the &#x60;/sessions/whoami&#x60; endpoint.
*V0alpha2Api* | [**patchProject**](docs/Api/V0alpha2Api.md#patchproject) | **PATCH** /projects/{project_id} | Patch an Ory Cloud Project Configuration&#x60;
*V0alpha2Api* | [**purgeProject**](docs/Api/V0alpha2Api.md#purgeproject) | **DELETE** /projects/{project_id} | # Irrecoverably Purge a Project
*V0alpha2Api* | [**removeProjectMember**](docs/Api/V0alpha2Api.md#removeprojectmember) | **DELETE** /projects/{project_id}/members/{member_id} | Remove a member associated with this project. This also sets their invite status to &#x60;REMOVED&#x60;.
*V0alpha2Api* | [**revokeSession**](docs/Api/V0alpha2Api.md#revokesession) | **DELETE** /sessions/{id} | Calling this endpoint invalidates the specified session. The current session cannot be revoked. Session data are not deleted.
*V0alpha2Api* | [**revokeSessions**](docs/Api/V0alpha2Api.md#revokesessions) | **DELETE** /sessions | Calling this endpoint invalidates all except the current session that belong to the logged-in user. Session data are not deleted.
*V0alpha2Api* | [**submitSelfServiceLoginFlow**](docs/Api/V0alpha2Api.md#submitselfserviceloginflow) | **POST** /self-service/login | Submit a Login Flow
*V0alpha2Api* | [**submitSelfServiceLogoutFlow**](docs/Api/V0alpha2Api.md#submitselfservicelogoutflow) | **GET** /self-service/logout | Complete Self-Service Logout
*V0alpha2Api* | [**submitSelfServiceLogoutFlowWithoutBrowser**](docs/Api/V0alpha2Api.md#submitselfservicelogoutflowwithoutbrowser) | **DELETE** /self-service/logout/api | Perform Logout for APIs, Services, Apps, ...
*V0alpha2Api* | [**submitSelfServiceRecoveryFlow**](docs/Api/V0alpha2Api.md#submitselfservicerecoveryflow) | **POST** /self-service/recovery | Complete Recovery Flow
*V0alpha2Api* | [**submitSelfServiceRegistrationFlow**](docs/Api/V0alpha2Api.md#submitselfserviceregistrationflow) | **POST** /self-service/registration | Submit a Registration Flow
*V0alpha2Api* | [**submitSelfServiceSettingsFlow**](docs/Api/V0alpha2Api.md#submitselfservicesettingsflow) | **POST** /self-service/settings | Complete Settings Flow
*V0alpha2Api* | [**submitSelfServiceVerificationFlow**](docs/Api/V0alpha2Api.md#submitselfserviceverificationflow) | **POST** /self-service/verification | # Complete Verification Flow
*V0alpha2Api* | [**toSession**](docs/Api/V0alpha2Api.md#tosession) | **GET** /sessions/whoami | Check Who the Current HTTP Session Belongs To
*V0alpha2Api* | [**updateProject**](docs/Api/V0alpha2Api.md#updateproject) | **PUT** /projects/{project_id} | # Update an Ory Cloud Project Configuration
*WriteApi* | [**createRelationTuple**](docs/Api/WriteApi.md#createrelationtuple) | **PUT** /admin/relation-tuples | # Create a Relation Tuple
*WriteApi* | [**deleteRelationTuples**](docs/Api/WriteApi.md#deleterelationtuples) | **DELETE** /admin/relation-tuples | # Delete Relation Tuples
*WriteApi* | [**patchRelationTuples**](docs/Api/WriteApi.md#patchrelationtuples) | **PATCH** /admin/relation-tuples | # Patch Multiple Relation Tuples

## Models

- [AcceptConsentRequest](docs/Model/AcceptConsentRequest.md)
- [AcceptLoginRequest](docs/Model/AcceptLoginRequest.md)
- [ActiveProject](docs/Model/ActiveProject.md)
- [AdminCreateIdentityBody](docs/Model/AdminCreateIdentityBody.md)
- [AdminCreateIdentityImportCredentialsOidc](docs/Model/AdminCreateIdentityImportCredentialsOidc.md)
- [AdminCreateIdentityImportCredentialsOidcConfig](docs/Model/AdminCreateIdentityImportCredentialsOidcConfig.md)
- [AdminCreateIdentityImportCredentialsOidcProvider](docs/Model/AdminCreateIdentityImportCredentialsOidcProvider.md)
- [AdminCreateIdentityImportCredentialsPassword](docs/Model/AdminCreateIdentityImportCredentialsPassword.md)
- [AdminCreateIdentityImportCredentialsPasswordConfig](docs/Model/AdminCreateIdentityImportCredentialsPasswordConfig.md)
- [AdminCreateSelfServiceRecoveryLinkBody](docs/Model/AdminCreateSelfServiceRecoveryLinkBody.md)
- [AdminIdentityImportCredentials](docs/Model/AdminIdentityImportCredentials.md)
- [AdminUpdateIdentityBody](docs/Model/AdminUpdateIdentityBody.md)
- [AuthenticatorAssuranceLevel](docs/Model/AuthenticatorAssuranceLevel.md)
- [CloudAccount](docs/Model/CloudAccount.md)
- [CnameSettings](docs/Model/CnameSettings.md)
- [CompletedRequest](docs/Model/CompletedRequest.md)
- [ConsentRequest](docs/Model/ConsentRequest.md)
- [ConsentRequestSession](docs/Model/ConsentRequestSession.md)
- [CreateCustomHostnameBody](docs/Model/CreateCustomHostnameBody.md)
- [CreateProjectBody](docs/Model/CreateProjectBody.md)
- [CreateSubscriptionPayload](docs/Model/CreateSubscriptionPayload.md)
- [ErrorAuthenticatorAssuranceLevelNotSatisfied](docs/Model/ErrorAuthenticatorAssuranceLevelNotSatisfied.md)
- [ExpandTree](docs/Model/ExpandTree.md)
- [FlushInactiveOAuth2TokensRequest](docs/Model/FlushInactiveOAuth2TokensRequest.md)
- [FlushLoginConsentRequest](docs/Model/FlushLoginConsentRequest.md)
- [GenericError](docs/Model/GenericError.md)
- [GenericErrorContent](docs/Model/GenericErrorContent.md)
- [GetCheckResponse](docs/Model/GetCheckResponse.md)
- [GetManagedIdentitySchemaLocation](docs/Model/GetManagedIdentitySchemaLocation.md)
- [GetRelationTuplesResponse](docs/Model/GetRelationTuplesResponse.md)
- [HealthNotReadyStatus](docs/Model/HealthNotReadyStatus.md)
- [HealthStatus](docs/Model/HealthStatus.md)
- [Identity](docs/Model/Identity.md)
- [IdentityCredentials](docs/Model/IdentityCredentials.md)
- [IdentityCredentialsOidc](docs/Model/IdentityCredentialsOidc.md)
- [IdentityCredentialsOidcProvider](docs/Model/IdentityCredentialsOidcProvider.md)
- [IdentityCredentialsPassword](docs/Model/IdentityCredentialsPassword.md)
- [IdentityCredentialsType](docs/Model/IdentityCredentialsType.md)
- [IdentitySchemaContainer](docs/Model/IdentitySchemaContainer.md)
- [IdentitySchemaPreset](docs/Model/IdentitySchemaPreset.md)
- [IdentityState](docs/Model/IdentityState.md)
- [InlineObject3](docs/Model/InlineObject3.md)
- [InlineResponse200](docs/Model/InlineResponse200.md)
- [InlineResponse2001](docs/Model/InlineResponse2001.md)
- [InlineResponse503](docs/Model/InlineResponse503.md)
- [InvitePayload](docs/Model/InvitePayload.md)
- [IsOwnerForProjectBySlug](docs/Model/IsOwnerForProjectBySlug.md)
- [IsOwnerForProjectBySlugPayload](docs/Model/IsOwnerForProjectBySlugPayload.md)
- [JSONWebKey](docs/Model/JSONWebKey.md)
- [JSONWebKeySet](docs/Model/JSONWebKeySet.md)
- [JsonError](docs/Model/JsonError.md)
- [JsonPatch](docs/Model/JsonPatch.md)
- [JsonWebKeySetGeneratorRequest](docs/Model/JsonWebKeySetGeneratorRequest.md)
- [KetoNamespace](docs/Model/KetoNamespace.md)
- [LoginRequest](docs/Model/LoginRequest.md)
- [LogoutRequest](docs/Model/LogoutRequest.md)
- [ManagedIdentitySchema](docs/Model/ManagedIdentitySchema.md)
- [ManagedIdentitySchemaValidationResult](docs/Model/ManagedIdentitySchemaValidationResult.md)
- [NeedsPrivilegedSessionError](docs/Model/NeedsPrivilegedSessionError.md)
- [NormalizedProject](docs/Model/NormalizedProject.md)
- [NormalizedProjectRevision](docs/Model/NormalizedProjectRevision.md)
- [NormalizedProjectRevisionHook](docs/Model/NormalizedProjectRevisionHook.md)
- [NormalizedProjectRevisionIdentitySchema](docs/Model/NormalizedProjectRevisionIdentitySchema.md)
- [NormalizedProjectRevisionThirdPartyProvider](docs/Model/NormalizedProjectRevisionThirdPartyProvider.md)
- [NullPlan](docs/Model/NullPlan.md)
- [OAuth2Client](docs/Model/OAuth2Client.md)
- [OAuth2TokenIntrospection](docs/Model/OAuth2TokenIntrospection.md)
- [Oauth2TokenResponse](docs/Model/Oauth2TokenResponse.md)
- [OauthTokenResponse](docs/Model/OauthTokenResponse.md)
- [OpenIDConnectContext](docs/Model/OpenIDConnectContext.md)
- [Pagination](docs/Model/Pagination.md)
- [PatchDelta](docs/Model/PatchDelta.md)
- [PatchDocument](docs/Model/PatchDocument.md)
- [PreviousConsentSession](docs/Model/PreviousConsentSession.md)
- [Project](docs/Model/Project.md)
- [ProjectApiKey](docs/Model/ProjectApiKey.md)
- [ProjectHost](docs/Model/ProjectHost.md)
- [ProjectInvite](docs/Model/ProjectInvite.md)
- [ProjectMetadata](docs/Model/ProjectMetadata.md)
- [ProjectServiceIdentity](docs/Model/ProjectServiceIdentity.md)
- [ProjectServiceOAuth2](docs/Model/ProjectServiceOAuth2.md)
- [ProjectServicePermission](docs/Model/ProjectServicePermission.md)
- [ProjectServices](docs/Model/ProjectServices.md)
- [ProvisionMockSubscriptionPayload](docs/Model/ProvisionMockSubscriptionPayload.md)
- [QuotaCustomDomains](docs/Model/QuotaCustomDomains.md)
- [QuotaProjectMemberSeats](docs/Model/QuotaProjectMemberSeats.md)
- [RecoveryAddress](docs/Model/RecoveryAddress.md)
- [RefreshTokenHookRequest](docs/Model/RefreshTokenHookRequest.md)
- [RefreshTokenHookResponse](docs/Model/RefreshTokenHookResponse.md)
- [RejectRequest](docs/Model/RejectRequest.md)
- [RelationQuery](docs/Model/RelationQuery.md)
- [RelationTuple](docs/Model/RelationTuple.md)
- [RequestWasHandledResponse](docs/Model/RequestWasHandledResponse.md)
- [RevokedSessions](docs/Model/RevokedSessions.md)
- [SchemaPatch](docs/Model/SchemaPatch.md)
- [SelfServiceBrowserLocationChangeRequiredError](docs/Model/SelfServiceBrowserLocationChangeRequiredError.md)
- [SelfServiceError](docs/Model/SelfServiceError.md)
- [SelfServiceFlowExpiredError](docs/Model/SelfServiceFlowExpiredError.md)
- [SelfServiceLoginFlow](docs/Model/SelfServiceLoginFlow.md)
- [SelfServiceLogoutUrl](docs/Model/SelfServiceLogoutUrl.md)
- [SelfServiceRecoveryFlow](docs/Model/SelfServiceRecoveryFlow.md)
- [SelfServiceRecoveryFlowState](docs/Model/SelfServiceRecoveryFlowState.md)
- [SelfServiceRecoveryLink](docs/Model/SelfServiceRecoveryLink.md)
- [SelfServiceRegistrationFlow](docs/Model/SelfServiceRegistrationFlow.md)
- [SelfServiceSettingsFlow](docs/Model/SelfServiceSettingsFlow.md)
- [SelfServiceSettingsFlowState](docs/Model/SelfServiceSettingsFlowState.md)
- [SelfServiceVerificationFlow](docs/Model/SelfServiceVerificationFlow.md)
- [SelfServiceVerificationFlowState](docs/Model/SelfServiceVerificationFlowState.md)
- [Session](docs/Model/Session.md)
- [SessionAuthenticationMethod](docs/Model/SessionAuthenticationMethod.md)
- [SessionDevice](docs/Model/SessionDevice.md)
- [SettingsProfileFormConfig](docs/Model/SettingsProfileFormConfig.md)
- [StripeCustomerResponse](docs/Model/StripeCustomerResponse.md)
- [SubjectSet](docs/Model/SubjectSet.md)
- [SubmitSelfServiceFlowWithWebAuthnRegistrationMethod](docs/Model/SubmitSelfServiceFlowWithWebAuthnRegistrationMethod.md)
- [SubmitSelfServiceLoginFlowBody](docs/Model/SubmitSelfServiceLoginFlowBody.md)
- [SubmitSelfServiceLoginFlowWithLookupSecretMethodBody](docs/Model/SubmitSelfServiceLoginFlowWithLookupSecretMethodBody.md)
- [SubmitSelfServiceLoginFlowWithOidcMethodBody](docs/Model/SubmitSelfServiceLoginFlowWithOidcMethodBody.md)
- [SubmitSelfServiceLoginFlowWithPasswordMethodBody](docs/Model/SubmitSelfServiceLoginFlowWithPasswordMethodBody.md)
- [SubmitSelfServiceLoginFlowWithTotpMethodBody](docs/Model/SubmitSelfServiceLoginFlowWithTotpMethodBody.md)
- [SubmitSelfServiceLoginFlowWithWebAuthnMethodBody](docs/Model/SubmitSelfServiceLoginFlowWithWebAuthnMethodBody.md)
- [SubmitSelfServiceLogoutFlowWithoutBrowserBody](docs/Model/SubmitSelfServiceLogoutFlowWithoutBrowserBody.md)
- [SubmitSelfServiceRecoveryFlowBody](docs/Model/SubmitSelfServiceRecoveryFlowBody.md)
- [SubmitSelfServiceRecoveryFlowWithLinkMethodBody](docs/Model/SubmitSelfServiceRecoveryFlowWithLinkMethodBody.md)
- [SubmitSelfServiceRegistrationFlowBody](docs/Model/SubmitSelfServiceRegistrationFlowBody.md)
- [SubmitSelfServiceRegistrationFlowWithOidcMethodBody](docs/Model/SubmitSelfServiceRegistrationFlowWithOidcMethodBody.md)
- [SubmitSelfServiceRegistrationFlowWithPasswordMethodBody](docs/Model/SubmitSelfServiceRegistrationFlowWithPasswordMethodBody.md)
- [SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody](docs/Model/SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody.md)
- [SubmitSelfServiceSettingsFlowBody](docs/Model/SubmitSelfServiceSettingsFlowBody.md)
- [SubmitSelfServiceSettingsFlowWithLookupMethodBody](docs/Model/SubmitSelfServiceSettingsFlowWithLookupMethodBody.md)
- [SubmitSelfServiceSettingsFlowWithOidcMethodBody](docs/Model/SubmitSelfServiceSettingsFlowWithOidcMethodBody.md)
- [SubmitSelfServiceSettingsFlowWithPasswordMethodBody](docs/Model/SubmitSelfServiceSettingsFlowWithPasswordMethodBody.md)
- [SubmitSelfServiceSettingsFlowWithProfileMethodBody](docs/Model/SubmitSelfServiceSettingsFlowWithProfileMethodBody.md)
- [SubmitSelfServiceSettingsFlowWithTotpMethodBody](docs/Model/SubmitSelfServiceSettingsFlowWithTotpMethodBody.md)
- [SubmitSelfServiceSettingsFlowWithWebAuthnMethodBody](docs/Model/SubmitSelfServiceSettingsFlowWithWebAuthnMethodBody.md)
- [SubmitSelfServiceVerificationFlowBody](docs/Model/SubmitSelfServiceVerificationFlowBody.md)
- [SubmitSelfServiceVerificationFlowWithLinkMethodBody](docs/Model/SubmitSelfServiceVerificationFlowWithLinkMethodBody.md)
- [Subscription](docs/Model/Subscription.md)
- [SuccessfulProjectUpdate](docs/Model/SuccessfulProjectUpdate.md)
- [SuccessfulSelfServiceLoginWithoutBrowser](docs/Model/SuccessfulSelfServiceLoginWithoutBrowser.md)
- [SuccessfulSelfServiceRegistrationWithoutBrowser](docs/Model/SuccessfulSelfServiceRegistrationWithoutBrowser.md)
- [TokenPagination](docs/Model/TokenPagination.md)
- [TokenPaginationHeaders](docs/Model/TokenPaginationHeaders.md)
- [TrustJwtGrantIssuerBody](docs/Model/TrustJwtGrantIssuerBody.md)
- [TrustedJsonWebKey](docs/Model/TrustedJsonWebKey.md)
- [TrustedJwtGrantIssuer](docs/Model/TrustedJwtGrantIssuer.md)
- [UiContainer](docs/Model/UiContainer.md)
- [UiNode](docs/Model/UiNode.md)
- [UiNodeAnchorAttributes](docs/Model/UiNodeAnchorAttributes.md)
- [UiNodeAttributes](docs/Model/UiNodeAttributes.md)
- [UiNodeImageAttributes](docs/Model/UiNodeImageAttributes.md)
- [UiNodeInputAttributes](docs/Model/UiNodeInputAttributes.md)
- [UiNodeMeta](docs/Model/UiNodeMeta.md)
- [UiNodeScriptAttributes](docs/Model/UiNodeScriptAttributes.md)
- [UiNodeTextAttributes](docs/Model/UiNodeTextAttributes.md)
- [UiText](docs/Model/UiText.md)
- [UpdateCustomHostnameBody](docs/Model/UpdateCustomHostnameBody.md)
- [UpdateProject](docs/Model/UpdateProject.md)
- [UpdateSubscriptionPayload](docs/Model/UpdateSubscriptionPayload.md)
- [UserinfoResponse](docs/Model/UserinfoResponse.md)
- [VerifiableIdentityAddress](docs/Model/VerifiableIdentityAddress.md)
- [Version](docs/Model/Version.md)
- [Warning](docs/Model/Warning.md)
- [WellKnown](docs/Model/WellKnown.md)

## Authorization

### basic

- **Type**: HTTP basic authentication


### oauth2

- **Type**: `OAuth`
- **Flow**: `accessCode`
- **Authorization URL**: `https://hydra.demo.ory.sh/oauth2/auth`
- **Scopes**: 
    - **offline**: A scope required when requesting refresh tokens (alias for `offline_access`)
    - **offline_access**: A scope required when requesting refresh tokens
    - **openid**: Request an OpenID Connect ID Token


### oryAccessToken

- **Type**: Bearer authentication

## Tests

To run the tests, use:

```bash
composer install
vendor/bin/phpunit
```

## Author

support@ory.sh

## About this package

This PHP package is automatically generated by the [OpenAPI Generator](https://openapi-generator.tech) project:

- API version: `v0.2.0-alpha.48`
- Build package: `org.openapitools.codegen.languages.PhpClientCodegen`
