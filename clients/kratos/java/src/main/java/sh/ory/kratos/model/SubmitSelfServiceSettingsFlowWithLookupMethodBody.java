/*
 * Ory Kratos API
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * The version of the OpenAPI document: v0.10.1
 * Contact: hi@ory.sh
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package sh.ory.kratos.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;

/**
 * SubmitSelfServiceSettingsFlowWithLookupMethodBody
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2022-06-01T11:19:13.300472677Z[Etc/UTC]")
public class SubmitSelfServiceSettingsFlowWithLookupMethodBody {
  public static final String SERIALIZED_NAME_CSRF_TOKEN = "csrf_token";
  @SerializedName(SERIALIZED_NAME_CSRF_TOKEN)
  private String csrfToken;

  public static final String SERIALIZED_NAME_LOOKUP_SECRET_CONFIRM = "lookup_secret_confirm";
  @SerializedName(SERIALIZED_NAME_LOOKUP_SECRET_CONFIRM)
  private Boolean lookupSecretConfirm;

  public static final String SERIALIZED_NAME_LOOKUP_SECRET_DISABLE = "lookup_secret_disable";
  @SerializedName(SERIALIZED_NAME_LOOKUP_SECRET_DISABLE)
  private Boolean lookupSecretDisable;

  public static final String SERIALIZED_NAME_LOOKUP_SECRET_REGENERATE = "lookup_secret_regenerate";
  @SerializedName(SERIALIZED_NAME_LOOKUP_SECRET_REGENERATE)
  private Boolean lookupSecretRegenerate;

  public static final String SERIALIZED_NAME_LOOKUP_SECRET_REVEAL = "lookup_secret_reveal";
  @SerializedName(SERIALIZED_NAME_LOOKUP_SECRET_REVEAL)
  private Boolean lookupSecretReveal;

  public static final String SERIALIZED_NAME_METHOD = "method";
  @SerializedName(SERIALIZED_NAME_METHOD)
  private String method;

  public SubmitSelfServiceSettingsFlowWithLookupMethodBody() { 
  }

  public SubmitSelfServiceSettingsFlowWithLookupMethodBody csrfToken(String csrfToken) {
    
    this.csrfToken = csrfToken;
    return this;
  }

   /**
   * CSRFToken is the anti-CSRF token
   * @return csrfToken
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "CSRFToken is the anti-CSRF token")

  public String getCsrfToken() {
    return csrfToken;
  }


  public void setCsrfToken(String csrfToken) {
    this.csrfToken = csrfToken;
  }


  public SubmitSelfServiceSettingsFlowWithLookupMethodBody lookupSecretConfirm(Boolean lookupSecretConfirm) {
    
    this.lookupSecretConfirm = lookupSecretConfirm;
    return this;
  }

   /**
   * If set to true will save the regenerated lookup secrets
   * @return lookupSecretConfirm
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "If set to true will save the regenerated lookup secrets")

  public Boolean getLookupSecretConfirm() {
    return lookupSecretConfirm;
  }


  public void setLookupSecretConfirm(Boolean lookupSecretConfirm) {
    this.lookupSecretConfirm = lookupSecretConfirm;
  }


  public SubmitSelfServiceSettingsFlowWithLookupMethodBody lookupSecretDisable(Boolean lookupSecretDisable) {
    
    this.lookupSecretDisable = lookupSecretDisable;
    return this;
  }

   /**
   * Disables this method if true.
   * @return lookupSecretDisable
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Disables this method if true.")

  public Boolean getLookupSecretDisable() {
    return lookupSecretDisable;
  }


  public void setLookupSecretDisable(Boolean lookupSecretDisable) {
    this.lookupSecretDisable = lookupSecretDisable;
  }


  public SubmitSelfServiceSettingsFlowWithLookupMethodBody lookupSecretRegenerate(Boolean lookupSecretRegenerate) {
    
    this.lookupSecretRegenerate = lookupSecretRegenerate;
    return this;
  }

   /**
   * If set to true will regenerate the lookup secrets
   * @return lookupSecretRegenerate
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "If set to true will regenerate the lookup secrets")

  public Boolean getLookupSecretRegenerate() {
    return lookupSecretRegenerate;
  }


  public void setLookupSecretRegenerate(Boolean lookupSecretRegenerate) {
    this.lookupSecretRegenerate = lookupSecretRegenerate;
  }


  public SubmitSelfServiceSettingsFlowWithLookupMethodBody lookupSecretReveal(Boolean lookupSecretReveal) {
    
    this.lookupSecretReveal = lookupSecretReveal;
    return this;
  }

   /**
   * If set to true will reveal the lookup secrets
   * @return lookupSecretReveal
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "If set to true will reveal the lookup secrets")

  public Boolean getLookupSecretReveal() {
    return lookupSecretReveal;
  }


  public void setLookupSecretReveal(Boolean lookupSecretReveal) {
    this.lookupSecretReveal = lookupSecretReveal;
  }


  public SubmitSelfServiceSettingsFlowWithLookupMethodBody method(String method) {
    
    this.method = method;
    return this;
  }

   /**
   * Method  Should be set to \&quot;lookup\&quot; when trying to add, update, or remove a lookup pairing.
   * @return method
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "Method  Should be set to \"lookup\" when trying to add, update, or remove a lookup pairing.")

  public String getMethod() {
    return method;
  }


  public void setMethod(String method) {
    this.method = method;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    SubmitSelfServiceSettingsFlowWithLookupMethodBody submitSelfServiceSettingsFlowWithLookupMethodBody = (SubmitSelfServiceSettingsFlowWithLookupMethodBody) o;
    return Objects.equals(this.csrfToken, submitSelfServiceSettingsFlowWithLookupMethodBody.csrfToken) &&
        Objects.equals(this.lookupSecretConfirm, submitSelfServiceSettingsFlowWithLookupMethodBody.lookupSecretConfirm) &&
        Objects.equals(this.lookupSecretDisable, submitSelfServiceSettingsFlowWithLookupMethodBody.lookupSecretDisable) &&
        Objects.equals(this.lookupSecretRegenerate, submitSelfServiceSettingsFlowWithLookupMethodBody.lookupSecretRegenerate) &&
        Objects.equals(this.lookupSecretReveal, submitSelfServiceSettingsFlowWithLookupMethodBody.lookupSecretReveal) &&
        Objects.equals(this.method, submitSelfServiceSettingsFlowWithLookupMethodBody.method);
  }

  @Override
  public int hashCode() {
    return Objects.hash(csrfToken, lookupSecretConfirm, lookupSecretDisable, lookupSecretRegenerate, lookupSecretReveal, method);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class SubmitSelfServiceSettingsFlowWithLookupMethodBody {\n");
    sb.append("    csrfToken: ").append(toIndentedString(csrfToken)).append("\n");
    sb.append("    lookupSecretConfirm: ").append(toIndentedString(lookupSecretConfirm)).append("\n");
    sb.append("    lookupSecretDisable: ").append(toIndentedString(lookupSecretDisable)).append("\n");
    sb.append("    lookupSecretRegenerate: ").append(toIndentedString(lookupSecretRegenerate)).append("\n");
    sb.append("    lookupSecretReveal: ").append(toIndentedString(lookupSecretReveal)).append("\n");
    sb.append("    method: ").append(toIndentedString(method)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}

