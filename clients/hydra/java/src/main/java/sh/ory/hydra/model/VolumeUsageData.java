/*
 * ORY Hydra
 * Welcome to the ORY Hydra HTTP API documentation. You will find documentation for all HTTP APIs here.
 *
 * The version of the OpenAPI document: v1.11.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package sh.ory.hydra.model;

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
 * VolumeUsageData Usage details about the volume. This information is used by the &#x60;GET /system/df&#x60; endpoint, and omitted in other endpoints.
 */
@ApiModel(description = "VolumeUsageData Usage details about the volume. This information is used by the `GET /system/df` endpoint, and omitted in other endpoints.")
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2022-05-27T08:52:01.831221547Z[Etc/UTC]")
public class VolumeUsageData {
  public static final String SERIALIZED_NAME_REF_COUNT = "RefCount";
  @SerializedName(SERIALIZED_NAME_REF_COUNT)
  private Long refCount;

  public static final String SERIALIZED_NAME_SIZE = "Size";
  @SerializedName(SERIALIZED_NAME_SIZE)
  private Long size;

  public VolumeUsageData() { 
  }

  public VolumeUsageData refCount(Long refCount) {
    
    this.refCount = refCount;
    return this;
  }

   /**
   * The number of containers referencing this volume. This field is set to &#x60;-1&#x60; if the reference-count is not available.
   * @return refCount
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "The number of containers referencing this volume. This field is set to `-1` if the reference-count is not available.")

  public Long getRefCount() {
    return refCount;
  }


  public void setRefCount(Long refCount) {
    this.refCount = refCount;
  }


  public VolumeUsageData size(Long size) {
    
    this.size = size;
    return this;
  }

   /**
   * Amount of disk space used by the volume (in bytes). This information is only available for volumes created with the &#x60;\&quot;local\&quot;&#x60; volume driver. For volumes created with other volume drivers, this field is set to &#x60;-1&#x60; (\&quot;not available\&quot;)
   * @return size
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "Amount of disk space used by the volume (in bytes). This information is only available for volumes created with the `\"local\"` volume driver. For volumes created with other volume drivers, this field is set to `-1` (\"not available\")")

  public Long getSize() {
    return size;
  }


  public void setSize(Long size) {
    this.size = size;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    VolumeUsageData volumeUsageData = (VolumeUsageData) o;
    return Objects.equals(this.refCount, volumeUsageData.refCount) &&
        Objects.equals(this.size, volumeUsageData.size);
  }

  @Override
  public int hashCode() {
    return Objects.hash(refCount, size);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class VolumeUsageData {\n");
    sb.append("    refCount: ").append(toIndentedString(refCount)).append("\n");
    sb.append("    size: ").append(toIndentedString(size)).append("\n");
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

