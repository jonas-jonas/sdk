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
import io.swagger.annotations.ApiModel;
import com.google.gson.annotations.SerializedName;

import java.io.IOException;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;

/**
 * show_form: No user data has been collected, or it is invalid, and thus the form should be shown. success: Indicates that the settings flow has been updated successfully with the provided data. Done will stay true when repeatedly checking. If set to true, done will revert back to false only when a flow with invalid (e.g. \&quot;please use a valid phone number\&quot;) data was sent.
 */
@JsonAdapter(SelfServiceSettingsFlowState.Adapter.class)
public enum SelfServiceSettingsFlowState {
  
  SHOW_FORM("show_form"),
  
  SUCCESS("success");

  private String value;

  SelfServiceSettingsFlowState(String value) {
    this.value = value;
  }

  public String getValue() {
    return value;
  }

  @Override
  public String toString() {
    return String.valueOf(value);
  }

  public static SelfServiceSettingsFlowState fromValue(String value) {
    for (SelfServiceSettingsFlowState b : SelfServiceSettingsFlowState.values()) {
      if (b.value.equals(value)) {
        return b;
      }
    }
    throw new IllegalArgumentException("Unexpected value '" + value + "'");
  }

  public static class Adapter extends TypeAdapter<SelfServiceSettingsFlowState> {
    @Override
    public void write(final JsonWriter jsonWriter, final SelfServiceSettingsFlowState enumeration) throws IOException {
      jsonWriter.value(enumeration.getValue());
    }

    @Override
    public SelfServiceSettingsFlowState read(final JsonReader jsonReader) throws IOException {
      String value = jsonReader.nextString();
      return SelfServiceSettingsFlowState.fromValue(value);
    }
  }
}

