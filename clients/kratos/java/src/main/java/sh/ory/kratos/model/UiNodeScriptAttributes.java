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
 * UiNodeScriptAttributes
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2022-06-01T11:19:13.300472677Z[Etc/UTC]")
public class UiNodeScriptAttributes {
  public static final String SERIALIZED_NAME_ASYNC = "async";
  @SerializedName(SERIALIZED_NAME_ASYNC)
  private Boolean async;

  public static final String SERIALIZED_NAME_CROSSORIGIN = "crossorigin";
  @SerializedName(SERIALIZED_NAME_CROSSORIGIN)
  private String crossorigin;

  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private String id;

  public static final String SERIALIZED_NAME_INTEGRITY = "integrity";
  @SerializedName(SERIALIZED_NAME_INTEGRITY)
  private String integrity;

  public static final String SERIALIZED_NAME_NODE_TYPE = "node_type";
  @SerializedName(SERIALIZED_NAME_NODE_TYPE)
  private String nodeType;

  public static final String SERIALIZED_NAME_NONCE = "nonce";
  @SerializedName(SERIALIZED_NAME_NONCE)
  private String nonce;

  public static final String SERIALIZED_NAME_REFERRERPOLICY = "referrerpolicy";
  @SerializedName(SERIALIZED_NAME_REFERRERPOLICY)
  private String referrerpolicy;

  public static final String SERIALIZED_NAME_SRC = "src";
  @SerializedName(SERIALIZED_NAME_SRC)
  private String src;

  public static final String SERIALIZED_NAME_TYPE = "type";
  @SerializedName(SERIALIZED_NAME_TYPE)
  private String type;

  public UiNodeScriptAttributes() { 
  }

  public UiNodeScriptAttributes async(Boolean async) {
    
    this.async = async;
    return this;
  }

   /**
   * The script async type
   * @return async
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "The script async type")

  public Boolean getAsync() {
    return async;
  }


  public void setAsync(Boolean async) {
    this.async = async;
  }


  public UiNodeScriptAttributes crossorigin(String crossorigin) {
    
    this.crossorigin = crossorigin;
    return this;
  }

   /**
   * The script cross origin policy
   * @return crossorigin
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "The script cross origin policy")

  public String getCrossorigin() {
    return crossorigin;
  }


  public void setCrossorigin(String crossorigin) {
    this.crossorigin = crossorigin;
  }


  public UiNodeScriptAttributes id(String id) {
    
    this.id = id;
    return this;
  }

   /**
   * A unique identifier
   * @return id
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "A unique identifier")

  public String getId() {
    return id;
  }


  public void setId(String id) {
    this.id = id;
  }


  public UiNodeScriptAttributes integrity(String integrity) {
    
    this.integrity = integrity;
    return this;
  }

   /**
   * The script&#39;s integrity hash
   * @return integrity
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "The script's integrity hash")

  public String getIntegrity() {
    return integrity;
  }


  public void setIntegrity(String integrity) {
    this.integrity = integrity;
  }


  public UiNodeScriptAttributes nodeType(String nodeType) {
    
    this.nodeType = nodeType;
    return this;
  }

   /**
   * NodeType represents this node&#39;s types. It is a mirror of &#x60;node.type&#x60; and is primarily used to allow compatibility with OpenAPI 3.0. In this struct it technically always is \&quot;script\&quot;.
   * @return nodeType
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "NodeType represents this node's types. It is a mirror of `node.type` and is primarily used to allow compatibility with OpenAPI 3.0. In this struct it technically always is \"script\".")

  public String getNodeType() {
    return nodeType;
  }


  public void setNodeType(String nodeType) {
    this.nodeType = nodeType;
  }


  public UiNodeScriptAttributes nonce(String nonce) {
    
    this.nonce = nonce;
    return this;
  }

   /**
   * Nonce for CSP  A nonce you may want to use to improve your Content Security Policy. You do not have to use this value but if you want to improve your CSP policies you may use it. You can also choose to use your own nonce value!
   * @return nonce
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "Nonce for CSP  A nonce you may want to use to improve your Content Security Policy. You do not have to use this value but if you want to improve your CSP policies you may use it. You can also choose to use your own nonce value!")

  public String getNonce() {
    return nonce;
  }


  public void setNonce(String nonce) {
    this.nonce = nonce;
  }


  public UiNodeScriptAttributes referrerpolicy(String referrerpolicy) {
    
    this.referrerpolicy = referrerpolicy;
    return this;
  }

   /**
   * The script referrer policy
   * @return referrerpolicy
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "The script referrer policy")

  public String getReferrerpolicy() {
    return referrerpolicy;
  }


  public void setReferrerpolicy(String referrerpolicy) {
    this.referrerpolicy = referrerpolicy;
  }


  public UiNodeScriptAttributes src(String src) {
    
    this.src = src;
    return this;
  }

   /**
   * The script source
   * @return src
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "The script source")

  public String getSrc() {
    return src;
  }


  public void setSrc(String src) {
    this.src = src;
  }


  public UiNodeScriptAttributes type(String type) {
    
    this.type = type;
    return this;
  }

   /**
   * The script MIME type
   * @return type
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "The script MIME type")

  public String getType() {
    return type;
  }


  public void setType(String type) {
    this.type = type;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    UiNodeScriptAttributes uiNodeScriptAttributes = (UiNodeScriptAttributes) o;
    return Objects.equals(this.async, uiNodeScriptAttributes.async) &&
        Objects.equals(this.crossorigin, uiNodeScriptAttributes.crossorigin) &&
        Objects.equals(this.id, uiNodeScriptAttributes.id) &&
        Objects.equals(this.integrity, uiNodeScriptAttributes.integrity) &&
        Objects.equals(this.nodeType, uiNodeScriptAttributes.nodeType) &&
        Objects.equals(this.nonce, uiNodeScriptAttributes.nonce) &&
        Objects.equals(this.referrerpolicy, uiNodeScriptAttributes.referrerpolicy) &&
        Objects.equals(this.src, uiNodeScriptAttributes.src) &&
        Objects.equals(this.type, uiNodeScriptAttributes.type);
  }

  @Override
  public int hashCode() {
    return Objects.hash(async, crossorigin, id, integrity, nodeType, nonce, referrerpolicy, src, type);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class UiNodeScriptAttributes {\n");
    sb.append("    async: ").append(toIndentedString(async)).append("\n");
    sb.append("    crossorigin: ").append(toIndentedString(crossorigin)).append("\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    integrity: ").append(toIndentedString(integrity)).append("\n");
    sb.append("    nodeType: ").append(toIndentedString(nodeType)).append("\n");
    sb.append("    nonce: ").append(toIndentedString(nonce)).append("\n");
    sb.append("    referrerpolicy: ").append(toIndentedString(referrerpolicy)).append("\n");
    sb.append("    src: ").append(toIndentedString(src)).append("\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
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

