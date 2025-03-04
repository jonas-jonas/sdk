/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v0.2.0-alpha.48
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */




#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct RejectRequest {
    /// The error should follow the OAuth2 error format (e.g. `invalid_request`, `login_required`).  Defaults to `request_denied`.
    #[serde(rename = "error", skip_serializing_if = "Option::is_none")]
    pub error: Option<String>,
    /// Debug contains information to help resolve the problem as a developer. Usually not exposed to the public but only in the server logs.
    #[serde(rename = "error_debug", skip_serializing_if = "Option::is_none")]
    pub error_debug: Option<String>,
    /// Description of the error in a human readable format.
    #[serde(rename = "error_description", skip_serializing_if = "Option::is_none")]
    pub error_description: Option<String>,
    /// Hint to help resolve the error.
    #[serde(rename = "error_hint", skip_serializing_if = "Option::is_none")]
    pub error_hint: Option<String>,
    /// Represents the HTTP status code of the error (e.g. 401 or 403)  Defaults to 400
    #[serde(rename = "status_code", skip_serializing_if = "Option::is_none")]
    pub status_code: Option<i64>,
}

impl Default for RejectRequest {
    fn default() -> Self {
        Self::new()
    }
}

impl RejectRequest {
    pub fn new() -> RejectRequest {
        RejectRequest {
                error: None,
                error_debug: None,
                error_description: None,
                error_hint: None,
                status_code: None,
        }
    }
}


