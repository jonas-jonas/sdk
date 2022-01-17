/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v0.0.1-alpha.48
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */




#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct ProjectRevision {
    /// The Project API URL  The URL where the Project's APIs are available.
    #[serde(rename = "api_url")]
    pub api_url: String,
    /// Your Application URL  The URL where your application is available. Your users will be redirected to this URL when they successfully complete a login, logout, verification, recovery, or registration flow.  More fine-grained redirection patterns are available for the individual flows.
    #[serde(rename = "application_url")]
    pub application_url: String,
    /// The Project's Revision Creation Date
    #[serde(rename = "created_at")]
    pub created_at: String,
    /// Default Identity Schema URL  This represents your Ory Kratos Default Identity Schema. It is your identity's default schema. This allows setting custom fields such as \"address\", specifying whether you want to log in using email or a username, and more. For more information on this topic, please check out the identity documentation.  The value of this field can be either any \"https://\" URL - for example a file hosted at GitHub, or a `preset://`-prefixed string. Available profiles are:  profile://email profile://username
    #[serde(rename = "default_identity_schema_url")]
    pub default_identity_schema_url: String,
    /// Self-Service Error UI URL  Sets the UI URL for the error UI. If left empty, this will use Ory's hosted pages.
    #[serde(rename = "error_ui_url", skip_serializing_if = "Option::is_none")]
    pub error_ui_url: Option<String>,
    #[serde(rename = "hosts")]
    pub hosts: Vec<String>,
    #[serde(rename = "id")]
    pub id: String,
    /// The Project's Kratos Config Version
    #[serde(rename = "kratos_config_version", skip_serializing_if = "Option::is_none")]
    pub kratos_config_version: Option<String>,
    #[serde(rename = "kratos_custom_schema_id", skip_serializing_if = "Option::is_none")]
    pub kratos_custom_schema_id: Option<String>,
    /// Self-Service Login UI URL  Sets the UI URL for the login UI. If left empty, this will use Ory's hosted pages.
    #[serde(rename = "login_ui_url", skip_serializing_if = "Option::is_none")]
    pub login_ui_url: Option<String>,
    #[serde(rename = "lookup_secret", skip_serializing_if = "Option::is_none")]
    pub lookup_secret: Option<Box<crate::models::ProjectLookupSecretConfig>>,
    /// The project's name.
    #[serde(rename = "name")]
    pub name: String,
    #[serde(rename = "oidc_providers", skip_serializing_if = "Option::is_none")]
    pub oidc_providers: Option<Vec<crate::models::ProjectOidcConfig>>,
    #[serde(rename = "password", skip_serializing_if = "Option::is_none")]
    pub password: Option<Box<crate::models::ProjectPasswordConfig>>,
    #[serde(rename = "project_id")]
    pub project_id: String,
    #[serde(rename = "recovery", skip_serializing_if = "Option::is_none")]
    pub recovery: Option<Box<crate::models::ProjectRecoveryConfig>>,
    /// Self-Service Login UI URL  Sets the UI URL for the recovery UI. If left empty, this will use Ory's hosted pages.
    #[serde(rename = "recovery_ui_url", skip_serializing_if = "Option::is_none")]
    pub recovery_ui_url: Option<String>,
    #[serde(rename = "redirection_config", skip_serializing_if = "Option::is_none")]
    pub redirection_config: Option<Box<crate::models::RedirectionConfig>>,
    /// Self-Service Login UI URL  Sets the UI URL for the registration UI. If left empty, this will use Ory's hosted pages.
    #[serde(rename = "registration_ui_url", skip_serializing_if = "Option::is_none")]
    pub registration_ui_url: Option<String>,
    /// Issue Session after Sign Up  If set to true, users will receive a session after they successfully signed up. Enabling this option allows account enumeration during registration flows. Read more: https://www.ory.sh/kratos/docs/self-service/flows/user-registration#successful-registration
    #[serde(rename = "session_after_sign_up")]
    pub session_after_sign_up: bool,
    /// Enable Soft 2FA for Login Sessions
    #[serde(rename = "session_soft_2fa", skip_serializing_if = "Option::is_none")]
    pub session_soft_2fa: Option<bool>,
    /// Duration in Seconds of how long a Session is Privileged  Defines how long a session is considered privileged in seconds. If the session's authenticated_at is older than the value specified here, the user needs to re-authenticate to perform certain actions (e.g. password change).
    #[serde(rename = "settings_privileged_session_max_age_seconds", skip_serializing_if = "Option::is_none")]
    pub settings_privileged_session_max_age_seconds: Option<i64>,
    /// Enable Soft 2FA for Self-Service Settings Flows
    #[serde(rename = "settings_soft_2fa", skip_serializing_if = "Option::is_none")]
    pub settings_soft_2fa: Option<bool>,
    /// Self-Service Settings UI URL  Sets the UI URL for the settings UI. If left empty, this will use Ory's hosted pages.
    #[serde(rename = "settings_ui_url", skip_serializing_if = "Option::is_none")]
    pub settings_ui_url: Option<String>,
    #[serde(rename = "totp", skip_serializing_if = "Option::is_none")]
    pub totp: Option<Box<crate::models::ProjectTotpConfig>>,
    /// Last Time Project's Revision was Updated
    #[serde(rename = "updated_at")]
    pub updated_at: String,
    #[serde(rename = "verification", skip_serializing_if = "Option::is_none")]
    pub verification: Option<Box<crate::models::ProjectVerificationConfig>>,
    /// Self-Service Login UI URL  Sets the UI URL for the verification UI. If left empty, this will use Ory's hosted pages.
    #[serde(rename = "verification_ui_url", skip_serializing_if = "Option::is_none")]
    pub verification_ui_url: Option<String>,
    #[serde(rename = "webauthn", skip_serializing_if = "Option::is_none")]
    pub webauthn: Option<Box<crate::models::ProjectWebAuthnConfig>>,
}

impl ProjectRevision {
    pub fn new(api_url: String, application_url: String, created_at: String, default_identity_schema_url: String, hosts: Vec<String>, id: String, name: String, project_id: String, session_after_sign_up: bool, updated_at: String) -> ProjectRevision {
        ProjectRevision {
            api_url,
            application_url,
            created_at,
            default_identity_schema_url,
            error_ui_url: None,
            hosts,
            id,
            kratos_config_version: None,
            kratos_custom_schema_id: None,
            login_ui_url: None,
            lookup_secret: None,
            name,
            oidc_providers: None,
            password: None,
            project_id,
            recovery: None,
            recovery_ui_url: None,
            redirection_config: None,
            registration_ui_url: None,
            session_after_sign_up,
            session_soft_2fa: None,
            settings_privileged_session_max_age_seconds: None,
            settings_soft_2fa: None,
            settings_ui_url: None,
            totp: None,
            updated_at,
            verification: None,
            verification_ui_url: None,
            webauthn: None,
        }
    }
}


