/*
 * Ory Keto API
 *
 * Documentation for all of Ory Keto's REST APIs. gRPC is documented separately. 
 *
 * The version of the OpenAPI document: v0.10.0-alpha.0
 * Contact: hi@ory.sh
 * Generated by: https://openapi-generator.tech
 */




#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct GetVersion200Response {
    /// The version of Ory Keto.
    #[serde(rename = "version")]
    pub version: String,
}


impl GetVersion200Response {
    pub fn new(version: String) -> GetVersion200Response {
        GetVersion200Response {
                version,
        }
    }
}


