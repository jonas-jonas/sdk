/*
 * ORY Oathkeeper
 *
 * ORY Oathkeeper is a reverse proxy that checks the HTTP Authorization for validity against a set of rules. This service uses Hydra to validate access tokens and policies.
 *
 * The version of the OpenAPI document: v0.0.0
 * Contact: hi@ory.am
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
using OpenAPIDateConverter = Ory.Oathkeeper.Client.Client.OpenAPIDateConverter;

namespace Ory.Oathkeeper.Client.Model
{
    /// <summary>
    /// OathkeeperRule
    /// </summary>
    [DataContract(Name = "rule")]
    public partial class OathkeeperRule : IEquatable<OathkeeperRule>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="OathkeeperRule" /> class.
        /// </summary>
        /// <param name="authenticators">Authenticators is a list of authentication handlers that will try and authenticate the provided credentials. Authenticators are checked iteratively from index 0 to n and if the first authenticator to return a positive result will be the one used.  If you want the rule to first check a specific authenticator  before \&quot;falling back\&quot; to others, have that authenticator as the first item in the array..</param>
        /// <param name="authorizer">authorizer.</param>
        /// <param name="description">Description is a human readable description of this rule..</param>
        /// <param name="id">ID is the unique id of the rule. It can be at most 190 characters long, but the layout of the ID is up to you. You will need this ID later on to update or delete the rule..</param>
        /// <param name="match">match.</param>
        /// <param name="mutators">Mutators is a list of mutation handlers that transform the HTTP request. A common use case is generating a new set of credentials (e.g. JWT) which then will be forwarded to the upstream server.  Mutations are performed iteratively from index 0 to n and should all succeed in order for the HTTP request to be forwarded..</param>
        /// <param name="upstream">upstream.</param>
        public OathkeeperRule(List<OathkeeperRuleHandler> authenticators = default(List<OathkeeperRuleHandler>), OathkeeperRuleHandler authorizer = default(OathkeeperRuleHandler), string description = default(string), string id = default(string), OathkeeperRuleMatch match = default(OathkeeperRuleMatch), List<OathkeeperRuleHandler> mutators = default(List<OathkeeperRuleHandler>), OathkeeperUpstream upstream = default(OathkeeperUpstream))
        {
            this.Authenticators = authenticators;
            this.Authorizer = authorizer;
            this.Description = description;
            this.Id = id;
            this.Match = match;
            this.Mutators = mutators;
            this.Upstream = upstream;
        }

        /// <summary>
        /// Authenticators is a list of authentication handlers that will try and authenticate the provided credentials. Authenticators are checked iteratively from index 0 to n and if the first authenticator to return a positive result will be the one used.  If you want the rule to first check a specific authenticator  before \&quot;falling back\&quot; to others, have that authenticator as the first item in the array.
        /// </summary>
        /// <value>Authenticators is a list of authentication handlers that will try and authenticate the provided credentials. Authenticators are checked iteratively from index 0 to n and if the first authenticator to return a positive result will be the one used.  If you want the rule to first check a specific authenticator  before \&quot;falling back\&quot; to others, have that authenticator as the first item in the array.</value>
        [DataMember(Name = "authenticators", EmitDefaultValue = false)]
        public List<OathkeeperRuleHandler> Authenticators { get; set; }

        /// <summary>
        /// Gets or Sets Authorizer
        /// </summary>
        [DataMember(Name = "authorizer", EmitDefaultValue = false)]
        public OathkeeperRuleHandler Authorizer { get; set; }

        /// <summary>
        /// Description is a human readable description of this rule.
        /// </summary>
        /// <value>Description is a human readable description of this rule.</value>
        [DataMember(Name = "description", EmitDefaultValue = false)]
        public string Description { get; set; }

        /// <summary>
        /// ID is the unique id of the rule. It can be at most 190 characters long, but the layout of the ID is up to you. You will need this ID later on to update or delete the rule.
        /// </summary>
        /// <value>ID is the unique id of the rule. It can be at most 190 characters long, but the layout of the ID is up to you. You will need this ID later on to update or delete the rule.</value>
        [DataMember(Name = "id", EmitDefaultValue = false)]
        public string Id { get; set; }

        /// <summary>
        /// Gets or Sets Match
        /// </summary>
        [DataMember(Name = "match", EmitDefaultValue = false)]
        public OathkeeperRuleMatch Match { get; set; }

        /// <summary>
        /// Mutators is a list of mutation handlers that transform the HTTP request. A common use case is generating a new set of credentials (e.g. JWT) which then will be forwarded to the upstream server.  Mutations are performed iteratively from index 0 to n and should all succeed in order for the HTTP request to be forwarded.
        /// </summary>
        /// <value>Mutators is a list of mutation handlers that transform the HTTP request. A common use case is generating a new set of credentials (e.g. JWT) which then will be forwarded to the upstream server.  Mutations are performed iteratively from index 0 to n and should all succeed in order for the HTTP request to be forwarded.</value>
        [DataMember(Name = "mutators", EmitDefaultValue = false)]
        public List<OathkeeperRuleHandler> Mutators { get; set; }

        /// <summary>
        /// Gets or Sets Upstream
        /// </summary>
        [DataMember(Name = "upstream", EmitDefaultValue = false)]
        public OathkeeperUpstream Upstream { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            var sb = new StringBuilder();
            sb.Append("class OathkeeperRule {\n");
            sb.Append("  Authenticators: ").Append(Authenticators).Append("\n");
            sb.Append("  Authorizer: ").Append(Authorizer).Append("\n");
            sb.Append("  Description: ").Append(Description).Append("\n");
            sb.Append("  Id: ").Append(Id).Append("\n");
            sb.Append("  Match: ").Append(Match).Append("\n");
            sb.Append("  Mutators: ").Append(Mutators).Append("\n");
            sb.Append("  Upstream: ").Append(Upstream).Append("\n");
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
            return this.Equals(input as OathkeeperRule);
        }

        /// <summary>
        /// Returns true if OathkeeperRule instances are equal
        /// </summary>
        /// <param name="input">Instance of OathkeeperRule to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(OathkeeperRule input)
        {
            if (input == null)
                return false;

            return 
                (
                    this.Authenticators == input.Authenticators ||
                    this.Authenticators != null &&
                    input.Authenticators != null &&
                    this.Authenticators.SequenceEqual(input.Authenticators)
                ) && 
                (
                    this.Authorizer == input.Authorizer ||
                    (this.Authorizer != null &&
                    this.Authorizer.Equals(input.Authorizer))
                ) && 
                (
                    this.Description == input.Description ||
                    (this.Description != null &&
                    this.Description.Equals(input.Description))
                ) && 
                (
                    this.Id == input.Id ||
                    (this.Id != null &&
                    this.Id.Equals(input.Id))
                ) && 
                (
                    this.Match == input.Match ||
                    (this.Match != null &&
                    this.Match.Equals(input.Match))
                ) && 
                (
                    this.Mutators == input.Mutators ||
                    this.Mutators != null &&
                    input.Mutators != null &&
                    this.Mutators.SequenceEqual(input.Mutators)
                ) && 
                (
                    this.Upstream == input.Upstream ||
                    (this.Upstream != null &&
                    this.Upstream.Equals(input.Upstream))
                );
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
                if (this.Authenticators != null)
                    hashCode = hashCode * 59 + this.Authenticators.GetHashCode();
                if (this.Authorizer != null)
                    hashCode = hashCode * 59 + this.Authorizer.GetHashCode();
                if (this.Description != null)
                    hashCode = hashCode * 59 + this.Description.GetHashCode();
                if (this.Id != null)
                    hashCode = hashCode * 59 + this.Id.GetHashCode();
                if (this.Match != null)
                    hashCode = hashCode * 59 + this.Match.GetHashCode();
                if (this.Mutators != null)
                    hashCode = hashCode * 59 + this.Mutators.GetHashCode();
                if (this.Upstream != null)
                    hashCode = hashCode * 59 + this.Upstream.GetHashCode();
                return hashCode;
            }
        }

        /// <summary>
        /// To validate all properties of the instance
        /// </summary>
        /// <param name="validationContext">Validation context</param>
        /// <returns>Validation Result</returns>
        IEnumerable<System.ComponentModel.DataAnnotations.ValidationResult> IValidatableObject.Validate(ValidationContext validationContext)
        {
            yield break;
        }
    }

}
