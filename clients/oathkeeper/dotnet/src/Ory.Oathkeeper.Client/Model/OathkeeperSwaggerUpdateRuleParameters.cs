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
    /// SwaggerUpdateRuleParameters SwaggerUpdateRuleParameters SwaggerUpdateRuleParameters SwaggerUpdateRuleParameters SwaggerUpdateRuleParameters swagger update rule parameters
    /// </summary>
    [DataContract(Name = "SwaggerUpdateRuleParameters")]
    public partial class OathkeeperSwaggerUpdateRuleParameters : IEquatable<OathkeeperSwaggerUpdateRuleParameters>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="OathkeeperSwaggerUpdateRuleParameters" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected OathkeeperSwaggerUpdateRuleParameters() { }
        /// <summary>
        /// Initializes a new instance of the <see cref="OathkeeperSwaggerUpdateRuleParameters" /> class.
        /// </summary>
        /// <param name="body">body.</param>
        /// <param name="id">in: path (required).</param>
        public OathkeeperSwaggerUpdateRuleParameters(OathkeeperSwaggerRule body = default(OathkeeperSwaggerRule), string id = default(string))
        {
            // to ensure "id" is required (not null)
            this.Id = id ?? throw new ArgumentNullException("id is a required property for OathkeeperSwaggerUpdateRuleParameters and cannot be null");
            this.Body = body;
        }

        /// <summary>
        /// Gets or Sets Body
        /// </summary>
        [DataMember(Name = "Body", EmitDefaultValue = false)]
        public OathkeeperSwaggerRule Body { get; set; }

        /// <summary>
        /// in: path
        /// </summary>
        /// <value>in: path</value>
        [DataMember(Name = "id", IsRequired = true, EmitDefaultValue = false)]
        public string Id { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            var sb = new StringBuilder();
            sb.Append("class OathkeeperSwaggerUpdateRuleParameters {\n");
            sb.Append("  Body: ").Append(Body).Append("\n");
            sb.Append("  Id: ").Append(Id).Append("\n");
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
            return this.Equals(input as OathkeeperSwaggerUpdateRuleParameters);
        }

        /// <summary>
        /// Returns true if OathkeeperSwaggerUpdateRuleParameters instances are equal
        /// </summary>
        /// <param name="input">Instance of OathkeeperSwaggerUpdateRuleParameters to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(OathkeeperSwaggerUpdateRuleParameters input)
        {
            if (input == null)
                return false;

            return 
                (
                    this.Body == input.Body ||
                    (this.Body != null &&
                    this.Body.Equals(input.Body))
                ) && 
                (
                    this.Id == input.Id ||
                    (this.Id != null &&
                    this.Id.Equals(input.Id))
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
                if (this.Body != null)
                    hashCode = hashCode * 59 + this.Body.GetHashCode();
                if (this.Id != null)
                    hashCode = hashCode * 59 + this.Id.GetHashCode();
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
