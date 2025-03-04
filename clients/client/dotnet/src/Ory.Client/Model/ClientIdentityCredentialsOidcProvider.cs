/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v0.2.0-alpha.48
 * Contact: support@ory.sh
 * Generated by: https://github.com/openapitools/openapi-generator.git
 */


using System;
using System.Collections;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Linq;
using System.IO;
using System.Runtime.Serialization;
using System.Text;
using System.Text.RegularExpressions;
using Newtonsoft.Json;
using Newtonsoft.Json.Converters;
using Newtonsoft.Json.Linq;
using System.ComponentModel.DataAnnotations;
using OpenAPIDateConverter = Ory.Client.Client.OpenAPIDateConverter;

namespace Ory.Client.Model
{
    /// <summary>
    /// ClientIdentityCredentialsOidcProvider
    /// </summary>
    [DataContract(Name = "identityCredentialsOidcProvider")]
    public partial class ClientIdentityCredentialsOidcProvider : IEquatable<ClientIdentityCredentialsOidcProvider>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="ClientIdentityCredentialsOidcProvider" /> class.
        /// </summary>
        /// <param name="initialAccessToken">initialAccessToken.</param>
        /// <param name="initialIdToken">initialIdToken.</param>
        /// <param name="initialRefreshToken">initialRefreshToken.</param>
        /// <param name="provider">provider.</param>
        /// <param name="subject">subject.</param>
        public ClientIdentityCredentialsOidcProvider(string initialAccessToken = default(string), string initialIdToken = default(string), string initialRefreshToken = default(string), string provider = default(string), string subject = default(string))
        {
            this.InitialAccessToken = initialAccessToken;
            this.InitialIdToken = initialIdToken;
            this.InitialRefreshToken = initialRefreshToken;
            this.Provider = provider;
            this.Subject = subject;
            this.AdditionalProperties = new Dictionary<string, object>();
        }

        /// <summary>
        /// Gets or Sets InitialAccessToken
        /// </summary>
        [DataMember(Name = "initial_access_token", EmitDefaultValue = false)]
        public string InitialAccessToken { get; set; }

        /// <summary>
        /// Gets or Sets InitialIdToken
        /// </summary>
        [DataMember(Name = "initial_id_token", EmitDefaultValue = false)]
        public string InitialIdToken { get; set; }

        /// <summary>
        /// Gets or Sets InitialRefreshToken
        /// </summary>
        [DataMember(Name = "initial_refresh_token", EmitDefaultValue = false)]
        public string InitialRefreshToken { get; set; }

        /// <summary>
        /// Gets or Sets Provider
        /// </summary>
        [DataMember(Name = "provider", EmitDefaultValue = false)]
        public string Provider { get; set; }

        /// <summary>
        /// Gets or Sets Subject
        /// </summary>
        [DataMember(Name = "subject", EmitDefaultValue = false)]
        public string Subject { get; set; }

        /// <summary>
        /// Gets or Sets additional properties
        /// </summary>
        [JsonExtensionData]
        public IDictionary<string, object> AdditionalProperties { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            StringBuilder sb = new StringBuilder();
            sb.Append("class ClientIdentityCredentialsOidcProvider {\n");
            sb.Append("  InitialAccessToken: ").Append(InitialAccessToken).Append("\n");
            sb.Append("  InitialIdToken: ").Append(InitialIdToken).Append("\n");
            sb.Append("  InitialRefreshToken: ").Append(InitialRefreshToken).Append("\n");
            sb.Append("  Provider: ").Append(Provider).Append("\n");
            sb.Append("  Subject: ").Append(Subject).Append("\n");
            sb.Append("  AdditionalProperties: ").Append(AdditionalProperties).Append("\n");
            sb.Append("}\n");
            return sb.ToString();
        }

        /// <summary>
        /// Returns the JSON string presentation of the object
        /// </summary>
        /// <returns>JSON string presentation of the object</returns>
        public virtual string ToJson()
        {
            return Newtonsoft.Json.JsonConvert.SerializeObject(this, Newtonsoft.Json.Formatting.Indented);
        }

        /// <summary>
        /// Returns true if objects are equal
        /// </summary>
        /// <param name="input">Object to be compared</param>
        /// <returns>Boolean</returns>
        public override bool Equals(object input)
        {
            return this.Equals(input as ClientIdentityCredentialsOidcProvider);
        }

        /// <summary>
        /// Returns true if ClientIdentityCredentialsOidcProvider instances are equal
        /// </summary>
        /// <param name="input">Instance of ClientIdentityCredentialsOidcProvider to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(ClientIdentityCredentialsOidcProvider input)
        {
            if (input == null)
            {
                return false;
            }
            return 
                (
                    this.InitialAccessToken == input.InitialAccessToken ||
                    (this.InitialAccessToken != null &&
                    this.InitialAccessToken.Equals(input.InitialAccessToken))
                ) && 
                (
                    this.InitialIdToken == input.InitialIdToken ||
                    (this.InitialIdToken != null &&
                    this.InitialIdToken.Equals(input.InitialIdToken))
                ) && 
                (
                    this.InitialRefreshToken == input.InitialRefreshToken ||
                    (this.InitialRefreshToken != null &&
                    this.InitialRefreshToken.Equals(input.InitialRefreshToken))
                ) && 
                (
                    this.Provider == input.Provider ||
                    (this.Provider != null &&
                    this.Provider.Equals(input.Provider))
                ) && 
                (
                    this.Subject == input.Subject ||
                    (this.Subject != null &&
                    this.Subject.Equals(input.Subject))
                )
                && (this.AdditionalProperties.Count == input.AdditionalProperties.Count && !this.AdditionalProperties.Except(input.AdditionalProperties).Any());
        }

        /// <summary>
        /// Gets the hash code
        /// </summary>
        /// <returns>Hash code</returns>
        public override int GetHashCode()
        {
            unchecked // Overflow is fine, just wrap
            {
                int hashCode = 41;
                if (this.InitialAccessToken != null)
                {
                    hashCode = (hashCode * 59) + this.InitialAccessToken.GetHashCode();
                }
                if (this.InitialIdToken != null)
                {
                    hashCode = (hashCode * 59) + this.InitialIdToken.GetHashCode();
                }
                if (this.InitialRefreshToken != null)
                {
                    hashCode = (hashCode * 59) + this.InitialRefreshToken.GetHashCode();
                }
                if (this.Provider != null)
                {
                    hashCode = (hashCode * 59) + this.Provider.GetHashCode();
                }
                if (this.Subject != null)
                {
                    hashCode = (hashCode * 59) + this.Subject.GetHashCode();
                }
                if (this.AdditionalProperties != null)
                {
                    hashCode = (hashCode * 59) + this.AdditionalProperties.GetHashCode();
                }
                return hashCode;
            }
        }

        /// <summary>
        /// To validate all properties of the instance
        /// </summary>
        /// <param name="validationContext">Validation context</param>
        /// <returns>Validation Result</returns>
        public IEnumerable<System.ComponentModel.DataAnnotations.ValidationResult> Validate(ValidationContext validationContext)
        {
            yield break;
        }
    }

}
