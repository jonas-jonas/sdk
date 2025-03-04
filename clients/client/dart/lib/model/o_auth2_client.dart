//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//
// @dart=2.12

// ignore_for_file: unused_element, unused_import
// ignore_for_file: always_put_required_named_parameters_first
// ignore_for_file: constant_identifier_names
// ignore_for_file: lines_longer_than_80_chars

part of openapi.api;

class OAuth2Client {
  /// Returns a new [OAuth2Client] instance.
  OAuth2Client({
    this.allowedCorsOrigins = const [],
    this.audience = const [],
    this.backchannelLogoutSessionRequired,
    this.backchannelLogoutUri,
    this.clientId,
    this.clientName,
    this.clientSecret,
    this.clientSecretExpiresAt,
    this.clientUri,
    this.contacts = const [],
    this.createdAt,
    this.frontchannelLogoutSessionRequired,
    this.frontchannelLogoutUri,
    this.grantTypes = const [],
    this.jwks,
    this.jwksUri,
    this.logoUri,
    this.metadata,
    this.owner,
    this.policyUri,
    this.postLogoutRedirectUris = const [],
    this.redirectUris = const [],
    this.registrationAccessToken,
    this.registrationClientUri,
    this.requestObjectSigningAlg,
    this.requestUris = const [],
    this.responseTypes = const [],
    this.scope,
    this.sectorIdentifierUri,
    this.subjectType,
    this.tokenEndpointAuthMethod,
    this.tokenEndpointAuthSigningAlg,
    this.tosUri,
    this.updatedAt,
    this.userinfoSignedResponseAlg,
  });

  List<String> allowedCorsOrigins;

  List<String> audience;

  /// Boolean value specifying whether the RP requires that a sid (session ID) Claim be included in the Logout Token to identify the RP session with the OP when the backchannel_logout_uri is used. If omitted, the default value is false.
  ///
  /// Please note: This property should have been non-nullable! Since the specification file
  /// does not include a default value (using the "default:" property), however, the generated
  /// source code must fall back to having a nullable type.
  /// Consider adding a "default:" property in the specification file to hide this note.
  ///
  bool? backchannelLogoutSessionRequired;

  /// RP URL that will cause the RP to log itself out when sent a Logout Token by the OP.
  ///
  /// Please note: This property should have been non-nullable! Since the specification file
  /// does not include a default value (using the "default:" property), however, the generated
  /// source code must fall back to having a nullable type.
  /// Consider adding a "default:" property in the specification file to hide this note.
  ///
  String? backchannelLogoutUri;

  /// ID  is the id for this client.
  ///
  /// Please note: This property should have been non-nullable! Since the specification file
  /// does not include a default value (using the "default:" property), however, the generated
  /// source code must fall back to having a nullable type.
  /// Consider adding a "default:" property in the specification file to hide this note.
  ///
  String? clientId;

  /// Name is the human-readable string name of the client to be presented to the end-user during authorization.
  ///
  /// Please note: This property should have been non-nullable! Since the specification file
  /// does not include a default value (using the "default:" property), however, the generated
  /// source code must fall back to having a nullable type.
  /// Consider adding a "default:" property in the specification file to hide this note.
  ///
  String? clientName;

  /// Secret is the client's secret. The secret will be included in the create request as cleartext, and then never again. The secret is stored using BCrypt so it is impossible to recover it. Tell your users that they need to write the secret down as it will not be made available again.
  ///
  /// Please note: This property should have been non-nullable! Since the specification file
  /// does not include a default value (using the "default:" property), however, the generated
  /// source code must fall back to having a nullable type.
  /// Consider adding a "default:" property in the specification file to hide this note.
  ///
  String? clientSecret;

  /// SecretExpiresAt is an integer holding the time at which the client secret will expire or 0 if it will not expire. The time is represented as the number of seconds from 1970-01-01T00:00:00Z as measured in UTC until the date/time of expiration.  This feature is currently not supported and it's value will always be set to 0.
  ///
  /// Please note: This property should have been non-nullable! Since the specification file
  /// does not include a default value (using the "default:" property), however, the generated
  /// source code must fall back to having a nullable type.
  /// Consider adding a "default:" property in the specification file to hide this note.
  ///
  int? clientSecretExpiresAt;

  /// ClientURI is an URL string of a web page providing information about the client. If present, the server SHOULD display this URL to the end-user in a clickable fashion.
  ///
  /// Please note: This property should have been non-nullable! Since the specification file
  /// does not include a default value (using the "default:" property), however, the generated
  /// source code must fall back to having a nullable type.
  /// Consider adding a "default:" property in the specification file to hide this note.
  ///
  String? clientUri;

  List<String> contacts;

  /// CreatedAt returns the timestamp of the client's creation.
  ///
  /// Please note: This property should have been non-nullable! Since the specification file
  /// does not include a default value (using the "default:" property), however, the generated
  /// source code must fall back to having a nullable type.
  /// Consider adding a "default:" property in the specification file to hide this note.
  ///
  DateTime? createdAt;

  /// Boolean value specifying whether the RP requires that iss (issuer) and sid (session ID) query parameters be included to identify the RP session with the OP when the frontchannel_logout_uri is used. If omitted, the default value is false.
  ///
  /// Please note: This property should have been non-nullable! Since the specification file
  /// does not include a default value (using the "default:" property), however, the generated
  /// source code must fall back to having a nullable type.
  /// Consider adding a "default:" property in the specification file to hide this note.
  ///
  bool? frontchannelLogoutSessionRequired;

  /// RP URL that will cause the RP to log itself out when rendered in an iframe by the OP. An iss (issuer) query parameter and a sid (session ID) query parameter MAY be included by the OP to enable the RP to validate the request and to determine which of the potentially multiple sessions is to be logged out; if either is included, both MUST be.
  ///
  /// Please note: This property should have been non-nullable! Since the specification file
  /// does not include a default value (using the "default:" property), however, the generated
  /// source code must fall back to having a nullable type.
  /// Consider adding a "default:" property in the specification file to hide this note.
  ///
  String? frontchannelLogoutUri;

  List<String> grantTypes;

  ///
  /// Please note: This property should have been non-nullable! Since the specification file
  /// does not include a default value (using the "default:" property), however, the generated
  /// source code must fall back to having a nullable type.
  /// Consider adding a "default:" property in the specification file to hide this note.
  ///
  Object? jwks;

  /// URL for the Client's JSON Web Key Set [JWK] document. If the Client signs requests to the Server, it contains the signing key(s) the Server uses to validate signatures from the Client. The JWK Set MAY also contain the Client's encryption keys(s), which are used by the Server to encrypt responses to the Client. When both signing and encryption keys are made available, a use (Key Use) parameter value is REQUIRED for all keys in the referenced JWK Set to indicate each key's intended usage. Although some algorithms allow the same key to be used for both signatures and encryption, doing so is NOT RECOMMENDED, as it is less secure. The JWK x5c parameter MAY be used to provide X.509 representations of keys provided. When used, the bare key values MUST still be present and MUST match those in the certificate.
  ///
  /// Please note: This property should have been non-nullable! Since the specification file
  /// does not include a default value (using the "default:" property), however, the generated
  /// source code must fall back to having a nullable type.
  /// Consider adding a "default:" property in the specification file to hide this note.
  ///
  String? jwksUri;

  /// LogoURI is an URL string that references a logo for the client.
  ///
  /// Please note: This property should have been non-nullable! Since the specification file
  /// does not include a default value (using the "default:" property), however, the generated
  /// source code must fall back to having a nullable type.
  /// Consider adding a "default:" property in the specification file to hide this note.
  ///
  String? logoUri;

  ///
  /// Please note: This property should have been non-nullable! Since the specification file
  /// does not include a default value (using the "default:" property), however, the generated
  /// source code must fall back to having a nullable type.
  /// Consider adding a "default:" property in the specification file to hide this note.
  ///
  Object? metadata;

  /// Owner is a string identifying the owner of the OAuth 2.0 Client.
  ///
  /// Please note: This property should have been non-nullable! Since the specification file
  /// does not include a default value (using the "default:" property), however, the generated
  /// source code must fall back to having a nullable type.
  /// Consider adding a "default:" property in the specification file to hide this note.
  ///
  String? owner;

  /// PolicyURI is a URL string that points to a human-readable privacy policy document that describes how the deployment organization collects, uses, retains, and discloses personal data.
  ///
  /// Please note: This property should have been non-nullable! Since the specification file
  /// does not include a default value (using the "default:" property), however, the generated
  /// source code must fall back to having a nullable type.
  /// Consider adding a "default:" property in the specification file to hide this note.
  ///
  String? policyUri;

  List<String> postLogoutRedirectUris;

  List<String> redirectUris;

  /// RegistrationAccessToken can be used to update, get, or delete the OAuth2 Client.
  ///
  /// Please note: This property should have been non-nullable! Since the specification file
  /// does not include a default value (using the "default:" property), however, the generated
  /// source code must fall back to having a nullable type.
  /// Consider adding a "default:" property in the specification file to hide this note.
  ///
  String? registrationAccessToken;

  /// RegistrationClientURI is the URL used to update, get, or delete the OAuth2 Client.
  ///
  /// Please note: This property should have been non-nullable! Since the specification file
  /// does not include a default value (using the "default:" property), however, the generated
  /// source code must fall back to having a nullable type.
  /// Consider adding a "default:" property in the specification file to hide this note.
  ///
  String? registrationClientUri;

  /// JWS [JWS] alg algorithm [JWA] that MUST be used for signing Request Objects sent to the OP. All Request Objects from this Client MUST be rejected, if not signed with this algorithm.
  ///
  /// Please note: This property should have been non-nullable! Since the specification file
  /// does not include a default value (using the "default:" property), however, the generated
  /// source code must fall back to having a nullable type.
  /// Consider adding a "default:" property in the specification file to hide this note.
  ///
  String? requestObjectSigningAlg;

  List<String> requestUris;

  List<String> responseTypes;

  /// Scope is a string containing a space-separated list of scope values (as described in Section 3.3 of OAuth 2.0 [RFC6749]) that the client can use when requesting access tokens.
  ///
  /// Please note: This property should have been non-nullable! Since the specification file
  /// does not include a default value (using the "default:" property), however, the generated
  /// source code must fall back to having a nullable type.
  /// Consider adding a "default:" property in the specification file to hide this note.
  ///
  String? scope;

  /// URL using the https scheme to be used in calculating Pseudonymous Identifiers by the OP. The URL references a file with a single JSON array of redirect_uri values.
  ///
  /// Please note: This property should have been non-nullable! Since the specification file
  /// does not include a default value (using the "default:" property), however, the generated
  /// source code must fall back to having a nullable type.
  /// Consider adding a "default:" property in the specification file to hide this note.
  ///
  String? sectorIdentifierUri;

  /// SubjectType requested for responses to this Client. The subject_types_supported Discovery parameter contains a list of the supported subject_type values for this server. Valid types include `pairwise` and `public`.
  ///
  /// Please note: This property should have been non-nullable! Since the specification file
  /// does not include a default value (using the "default:" property), however, the generated
  /// source code must fall back to having a nullable type.
  /// Consider adding a "default:" property in the specification file to hide this note.
  ///
  String? subjectType;

  /// Requested Client Authentication method for the Token Endpoint. The options are client_secret_post, client_secret_basic, private_key_jwt, and none.
  ///
  /// Please note: This property should have been non-nullable! Since the specification file
  /// does not include a default value (using the "default:" property), however, the generated
  /// source code must fall back to having a nullable type.
  /// Consider adding a "default:" property in the specification file to hide this note.
  ///
  String? tokenEndpointAuthMethod;

  /// Requested Client Authentication signing algorithm for the Token Endpoint.
  ///
  /// Please note: This property should have been non-nullable! Since the specification file
  /// does not include a default value (using the "default:" property), however, the generated
  /// source code must fall back to having a nullable type.
  /// Consider adding a "default:" property in the specification file to hide this note.
  ///
  String? tokenEndpointAuthSigningAlg;

  /// TermsOfServiceURI is a URL string that points to a human-readable terms of service document for the client that describes a contractual relationship between the end-user and the client that the end-user accepts when authorizing the client.
  ///
  /// Please note: This property should have been non-nullable! Since the specification file
  /// does not include a default value (using the "default:" property), however, the generated
  /// source code must fall back to having a nullable type.
  /// Consider adding a "default:" property in the specification file to hide this note.
  ///
  String? tosUri;

  /// UpdatedAt returns the timestamp of the last update.
  ///
  /// Please note: This property should have been non-nullable! Since the specification file
  /// does not include a default value (using the "default:" property), however, the generated
  /// source code must fall back to having a nullable type.
  /// Consider adding a "default:" property in the specification file to hide this note.
  ///
  DateTime? updatedAt;

  /// JWS alg algorithm [JWA] REQUIRED for signing UserInfo Responses. If this is specified, the response will be JWT [JWT] serialized, and signed using JWS. The default, if omitted, is for the UserInfo Response to return the Claims as a UTF-8 encoded JSON object using the application/json content-type.
  ///
  /// Please note: This property should have been non-nullable! Since the specification file
  /// does not include a default value (using the "default:" property), however, the generated
  /// source code must fall back to having a nullable type.
  /// Consider adding a "default:" property in the specification file to hide this note.
  ///
  String? userinfoSignedResponseAlg;

  @override
  bool operator ==(Object other) => identical(this, other) || other is OAuth2Client &&
     other.allowedCorsOrigins == allowedCorsOrigins &&
     other.audience == audience &&
     other.backchannelLogoutSessionRequired == backchannelLogoutSessionRequired &&
     other.backchannelLogoutUri == backchannelLogoutUri &&
     other.clientId == clientId &&
     other.clientName == clientName &&
     other.clientSecret == clientSecret &&
     other.clientSecretExpiresAt == clientSecretExpiresAt &&
     other.clientUri == clientUri &&
     other.contacts == contacts &&
     other.createdAt == createdAt &&
     other.frontchannelLogoutSessionRequired == frontchannelLogoutSessionRequired &&
     other.frontchannelLogoutUri == frontchannelLogoutUri &&
     other.grantTypes == grantTypes &&
     other.jwks == jwks &&
     other.jwksUri == jwksUri &&
     other.logoUri == logoUri &&
     other.metadata == metadata &&
     other.owner == owner &&
     other.policyUri == policyUri &&
     other.postLogoutRedirectUris == postLogoutRedirectUris &&
     other.redirectUris == redirectUris &&
     other.registrationAccessToken == registrationAccessToken &&
     other.registrationClientUri == registrationClientUri &&
     other.requestObjectSigningAlg == requestObjectSigningAlg &&
     other.requestUris == requestUris &&
     other.responseTypes == responseTypes &&
     other.scope == scope &&
     other.sectorIdentifierUri == sectorIdentifierUri &&
     other.subjectType == subjectType &&
     other.tokenEndpointAuthMethod == tokenEndpointAuthMethod &&
     other.tokenEndpointAuthSigningAlg == tokenEndpointAuthSigningAlg &&
     other.tosUri == tosUri &&
     other.updatedAt == updatedAt &&
     other.userinfoSignedResponseAlg == userinfoSignedResponseAlg;

  @override
  int get hashCode =>
    // ignore: unnecessary_parenthesis
    (allowedCorsOrigins.hashCode) +
    (audience.hashCode) +
    (backchannelLogoutSessionRequired == null ? 0 : backchannelLogoutSessionRequired!.hashCode) +
    (backchannelLogoutUri == null ? 0 : backchannelLogoutUri!.hashCode) +
    (clientId == null ? 0 : clientId!.hashCode) +
    (clientName == null ? 0 : clientName!.hashCode) +
    (clientSecret == null ? 0 : clientSecret!.hashCode) +
    (clientSecretExpiresAt == null ? 0 : clientSecretExpiresAt!.hashCode) +
    (clientUri == null ? 0 : clientUri!.hashCode) +
    (contacts.hashCode) +
    (createdAt == null ? 0 : createdAt!.hashCode) +
    (frontchannelLogoutSessionRequired == null ? 0 : frontchannelLogoutSessionRequired!.hashCode) +
    (frontchannelLogoutUri == null ? 0 : frontchannelLogoutUri!.hashCode) +
    (grantTypes.hashCode) +
    (jwks == null ? 0 : jwks!.hashCode) +
    (jwksUri == null ? 0 : jwksUri!.hashCode) +
    (logoUri == null ? 0 : logoUri!.hashCode) +
    (metadata == null ? 0 : metadata!.hashCode) +
    (owner == null ? 0 : owner!.hashCode) +
    (policyUri == null ? 0 : policyUri!.hashCode) +
    (postLogoutRedirectUris.hashCode) +
    (redirectUris.hashCode) +
    (registrationAccessToken == null ? 0 : registrationAccessToken!.hashCode) +
    (registrationClientUri == null ? 0 : registrationClientUri!.hashCode) +
    (requestObjectSigningAlg == null ? 0 : requestObjectSigningAlg!.hashCode) +
    (requestUris.hashCode) +
    (responseTypes.hashCode) +
    (scope == null ? 0 : scope!.hashCode) +
    (sectorIdentifierUri == null ? 0 : sectorIdentifierUri!.hashCode) +
    (subjectType == null ? 0 : subjectType!.hashCode) +
    (tokenEndpointAuthMethod == null ? 0 : tokenEndpointAuthMethod!.hashCode) +
    (tokenEndpointAuthSigningAlg == null ? 0 : tokenEndpointAuthSigningAlg!.hashCode) +
    (tosUri == null ? 0 : tosUri!.hashCode) +
    (updatedAt == null ? 0 : updatedAt!.hashCode) +
    (userinfoSignedResponseAlg == null ? 0 : userinfoSignedResponseAlg!.hashCode);

  @override
  String toString() => 'OAuth2Client[allowedCorsOrigins=$allowedCorsOrigins, audience=$audience, backchannelLogoutSessionRequired=$backchannelLogoutSessionRequired, backchannelLogoutUri=$backchannelLogoutUri, clientId=$clientId, clientName=$clientName, clientSecret=$clientSecret, clientSecretExpiresAt=$clientSecretExpiresAt, clientUri=$clientUri, contacts=$contacts, createdAt=$createdAt, frontchannelLogoutSessionRequired=$frontchannelLogoutSessionRequired, frontchannelLogoutUri=$frontchannelLogoutUri, grantTypes=$grantTypes, jwks=$jwks, jwksUri=$jwksUri, logoUri=$logoUri, metadata=$metadata, owner=$owner, policyUri=$policyUri, postLogoutRedirectUris=$postLogoutRedirectUris, redirectUris=$redirectUris, registrationAccessToken=$registrationAccessToken, registrationClientUri=$registrationClientUri, requestObjectSigningAlg=$requestObjectSigningAlg, requestUris=$requestUris, responseTypes=$responseTypes, scope=$scope, sectorIdentifierUri=$sectorIdentifierUri, subjectType=$subjectType, tokenEndpointAuthMethod=$tokenEndpointAuthMethod, tokenEndpointAuthSigningAlg=$tokenEndpointAuthSigningAlg, tosUri=$tosUri, updatedAt=$updatedAt, userinfoSignedResponseAlg=$userinfoSignedResponseAlg]';

  Map<String, dynamic> toJson() {
    final _json = <String, dynamic>{};
      _json[r'allowed_cors_origins'] = allowedCorsOrigins;
      _json[r'audience'] = audience;
    if (backchannelLogoutSessionRequired != null) {
      _json[r'backchannel_logout_session_required'] = backchannelLogoutSessionRequired;
    } else {
      _json[r'backchannel_logout_session_required'] = null;
    }
    if (backchannelLogoutUri != null) {
      _json[r'backchannel_logout_uri'] = backchannelLogoutUri;
    } else {
      _json[r'backchannel_logout_uri'] = null;
    }
    if (clientId != null) {
      _json[r'client_id'] = clientId;
    } else {
      _json[r'client_id'] = null;
    }
    if (clientName != null) {
      _json[r'client_name'] = clientName;
    } else {
      _json[r'client_name'] = null;
    }
    if (clientSecret != null) {
      _json[r'client_secret'] = clientSecret;
    } else {
      _json[r'client_secret'] = null;
    }
    if (clientSecretExpiresAt != null) {
      _json[r'client_secret_expires_at'] = clientSecretExpiresAt;
    } else {
      _json[r'client_secret_expires_at'] = null;
    }
    if (clientUri != null) {
      _json[r'client_uri'] = clientUri;
    } else {
      _json[r'client_uri'] = null;
    }
      _json[r'contacts'] = contacts;
    if (createdAt != null) {
      _json[r'created_at'] = createdAt!.toUtc().toIso8601String();
    } else {
      _json[r'created_at'] = null;
    }
    if (frontchannelLogoutSessionRequired != null) {
      _json[r'frontchannel_logout_session_required'] = frontchannelLogoutSessionRequired;
    } else {
      _json[r'frontchannel_logout_session_required'] = null;
    }
    if (frontchannelLogoutUri != null) {
      _json[r'frontchannel_logout_uri'] = frontchannelLogoutUri;
    } else {
      _json[r'frontchannel_logout_uri'] = null;
    }
      _json[r'grant_types'] = grantTypes;
    if (jwks != null) {
      _json[r'jwks'] = jwks;
    } else {
      _json[r'jwks'] = null;
    }
    if (jwksUri != null) {
      _json[r'jwks_uri'] = jwksUri;
    } else {
      _json[r'jwks_uri'] = null;
    }
    if (logoUri != null) {
      _json[r'logo_uri'] = logoUri;
    } else {
      _json[r'logo_uri'] = null;
    }
    if (metadata != null) {
      _json[r'metadata'] = metadata;
    } else {
      _json[r'metadata'] = null;
    }
    if (owner != null) {
      _json[r'owner'] = owner;
    } else {
      _json[r'owner'] = null;
    }
    if (policyUri != null) {
      _json[r'policy_uri'] = policyUri;
    } else {
      _json[r'policy_uri'] = null;
    }
      _json[r'post_logout_redirect_uris'] = postLogoutRedirectUris;
      _json[r'redirect_uris'] = redirectUris;
    if (registrationAccessToken != null) {
      _json[r'registration_access_token'] = registrationAccessToken;
    } else {
      _json[r'registration_access_token'] = null;
    }
    if (registrationClientUri != null) {
      _json[r'registration_client_uri'] = registrationClientUri;
    } else {
      _json[r'registration_client_uri'] = null;
    }
    if (requestObjectSigningAlg != null) {
      _json[r'request_object_signing_alg'] = requestObjectSigningAlg;
    } else {
      _json[r'request_object_signing_alg'] = null;
    }
      _json[r'request_uris'] = requestUris;
      _json[r'response_types'] = responseTypes;
    if (scope != null) {
      _json[r'scope'] = scope;
    } else {
      _json[r'scope'] = null;
    }
    if (sectorIdentifierUri != null) {
      _json[r'sector_identifier_uri'] = sectorIdentifierUri;
    } else {
      _json[r'sector_identifier_uri'] = null;
    }
    if (subjectType != null) {
      _json[r'subject_type'] = subjectType;
    } else {
      _json[r'subject_type'] = null;
    }
    if (tokenEndpointAuthMethod != null) {
      _json[r'token_endpoint_auth_method'] = tokenEndpointAuthMethod;
    } else {
      _json[r'token_endpoint_auth_method'] = null;
    }
    if (tokenEndpointAuthSigningAlg != null) {
      _json[r'token_endpoint_auth_signing_alg'] = tokenEndpointAuthSigningAlg;
    } else {
      _json[r'token_endpoint_auth_signing_alg'] = null;
    }
    if (tosUri != null) {
      _json[r'tos_uri'] = tosUri;
    } else {
      _json[r'tos_uri'] = null;
    }
    if (updatedAt != null) {
      _json[r'updated_at'] = updatedAt!.toUtc().toIso8601String();
    } else {
      _json[r'updated_at'] = null;
    }
    if (userinfoSignedResponseAlg != null) {
      _json[r'userinfo_signed_response_alg'] = userinfoSignedResponseAlg;
    } else {
      _json[r'userinfo_signed_response_alg'] = null;
    }
    return _json;
  }

  /// Returns a new [OAuth2Client] instance and imports its values from
  /// [value] if it's a [Map], null otherwise.
  // ignore: prefer_constructors_over_static_methods
  static OAuth2Client? fromJson(dynamic value) {
    if (value is Map) {
      final json = value.cast<String, dynamic>();

      // Ensure that the map contains the required keys.
      // Note 1: the values aren't checked for validity beyond being non-null.
      // Note 2: this code is stripped in release mode!
      assert(() {
        requiredKeys.forEach((key) {
          assert(json.containsKey(key), 'Required key "OAuth2Client[$key]" is missing from JSON.');
          assert(json[key] != null, 'Required key "OAuth2Client[$key]" has a null value in JSON.');
        });
        return true;
      }());

      return OAuth2Client(
        allowedCorsOrigins: json[r'allowed_cors_origins'] is List
            ? (json[r'allowed_cors_origins'] as List).cast<String>()
            : const [],
        audience: json[r'audience'] is List
            ? (json[r'audience'] as List).cast<String>()
            : const [],
        backchannelLogoutSessionRequired: mapValueOfType<bool>(json, r'backchannel_logout_session_required'),
        backchannelLogoutUri: mapValueOfType<String>(json, r'backchannel_logout_uri'),
        clientId: mapValueOfType<String>(json, r'client_id'),
        clientName: mapValueOfType<String>(json, r'client_name'),
        clientSecret: mapValueOfType<String>(json, r'client_secret'),
        clientSecretExpiresAt: mapValueOfType<int>(json, r'client_secret_expires_at'),
        clientUri: mapValueOfType<String>(json, r'client_uri'),
        contacts: json[r'contacts'] is List
            ? (json[r'contacts'] as List).cast<String>()
            : const [],
        createdAt: mapDateTime(json, r'created_at', ''),
        frontchannelLogoutSessionRequired: mapValueOfType<bool>(json, r'frontchannel_logout_session_required'),
        frontchannelLogoutUri: mapValueOfType<String>(json, r'frontchannel_logout_uri'),
        grantTypes: json[r'grant_types'] is List
            ? (json[r'grant_types'] as List).cast<String>()
            : const [],
        jwks: mapValueOfType<Object>(json, r'jwks'),
        jwksUri: mapValueOfType<String>(json, r'jwks_uri'),
        logoUri: mapValueOfType<String>(json, r'logo_uri'),
        metadata: mapValueOfType<Object>(json, r'metadata'),
        owner: mapValueOfType<String>(json, r'owner'),
        policyUri: mapValueOfType<String>(json, r'policy_uri'),
        postLogoutRedirectUris: json[r'post_logout_redirect_uris'] is List
            ? (json[r'post_logout_redirect_uris'] as List).cast<String>()
            : const [],
        redirectUris: json[r'redirect_uris'] is List
            ? (json[r'redirect_uris'] as List).cast<String>()
            : const [],
        registrationAccessToken: mapValueOfType<String>(json, r'registration_access_token'),
        registrationClientUri: mapValueOfType<String>(json, r'registration_client_uri'),
        requestObjectSigningAlg: mapValueOfType<String>(json, r'request_object_signing_alg'),
        requestUris: json[r'request_uris'] is List
            ? (json[r'request_uris'] as List).cast<String>()
            : const [],
        responseTypes: json[r'response_types'] is List
            ? (json[r'response_types'] as List).cast<String>()
            : const [],
        scope: mapValueOfType<String>(json, r'scope'),
        sectorIdentifierUri: mapValueOfType<String>(json, r'sector_identifier_uri'),
        subjectType: mapValueOfType<String>(json, r'subject_type'),
        tokenEndpointAuthMethod: mapValueOfType<String>(json, r'token_endpoint_auth_method'),
        tokenEndpointAuthSigningAlg: mapValueOfType<String>(json, r'token_endpoint_auth_signing_alg'),
        tosUri: mapValueOfType<String>(json, r'tos_uri'),
        updatedAt: mapDateTime(json, r'updated_at', ''),
        userinfoSignedResponseAlg: mapValueOfType<String>(json, r'userinfo_signed_response_alg'),
      );
    }
    return null;
  }

  static List<OAuth2Client>? listFromJson(dynamic json, {bool growable = false,}) {
    final result = <OAuth2Client>[];
    if (json is List && json.isNotEmpty) {
      for (final row in json) {
        final value = OAuth2Client.fromJson(row);
        if (value != null) {
          result.add(value);
        }
      }
    }
    return result.toList(growable: growable);
  }

  static Map<String, OAuth2Client> mapFromJson(dynamic json) {
    final map = <String, OAuth2Client>{};
    if (json is Map && json.isNotEmpty) {
      json = json.cast<String, dynamic>(); // ignore: parameter_assignments
      for (final entry in json.entries) {
        final value = OAuth2Client.fromJson(entry.value);
        if (value != null) {
          map[entry.key] = value;
        }
      }
    }
    return map;
  }

  // maps a json object with a list of OAuth2Client-objects as value to a dart map
  static Map<String, List<OAuth2Client>> mapListFromJson(dynamic json, {bool growable = false,}) {
    final map = <String, List<OAuth2Client>>{};
    if (json is Map && json.isNotEmpty) {
      json = json.cast<String, dynamic>(); // ignore: parameter_assignments
      for (final entry in json.entries) {
        final value = OAuth2Client.listFromJson(entry.value, growable: growable,);
        if (value != null) {
          map[entry.key] = value;
        }
      }
    }
    return map;
  }

  /// The list of required keys that must be present in a JSON.
  static const requiredKeys = <String>{
  };
}

