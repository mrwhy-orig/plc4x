/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/*
 * ctrlX CORE - Data Layer API
 * This is the base API for the ctrlX Data Layer access on ctrlX CORE devices <ul> <li>Click 'Authorize' to open the 'Available authorizations' dialog.</li> <li>Enter 'username' and 'password'. The 'Client credentials location' selector together with the 'client_id' and 'client_secret' fields as well as the 'Bearer' section can be ignored.</li> <li>Click 'Authorize' and then 'Close' to close the 'Available authorizations' dialog.</li> <li>Try out those GET, PUT, ... operations you're interested in.</li> </ul>
 *
 * The version of the OpenAPI document: 2.1.0
 * Contact: support@boschrexroth.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.apache.plc4x.java.ctrlx.readwrite.rest.datalayer.model;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonPropertyOrder;
import com.fasterxml.jackson.annotation.JsonTypeName;

import java.io.UnsupportedEncodingException;
import java.math.BigDecimal;
import java.net.URLEncoder;
import java.util.Objects;
import java.util.StringJoiner;

/**
 * Event data
 */
@JsonPropertyOrder({
  SSEEventData.JSON_PROPERTY_NODE,
  SSEEventData.JSON_PROPERTY_TYPE,
  SSEEventData.JSON_PROPERTY_SCHEMA,
  SSEEventData.JSON_PROPERTY_VALUE,
  SSEEventData.JSON_PROPERTY_TIMESTAMP
})
@JsonTypeName("SSE_EventData")
@jakarta.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-11-18T13:34:36.056861+01:00[Europe/Berlin]")
public class SSEEventData {
  public static final String JSON_PROPERTY_NODE = "node";
  private String node;

  public static final String JSON_PROPERTY_TYPE = "type";
  private String type;

  public static final String JSON_PROPERTY_SCHEMA = "schema";
  private String schema;

  public static final String JSON_PROPERTY_VALUE = "value";
  private String value;

  public static final String JSON_PROPERTY_TIMESTAMP = "timestamp";
  private BigDecimal timestamp;

  public SSEEventData() {
  }

  public SSEEventData node(String node) {
    
    this.node = node;
    return this;
  }

   /**
   * Node address of event item
   * @return node
  **/
  @jakarta.annotation.Nonnull
  @JsonProperty(JSON_PROPERTY_NODE)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)

  public String getNode() {
    return node;
  }


  @JsonProperty(JSON_PROPERTY_NODE)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)
  public void setNode(String node) {
    this.node = node;
  }


  public SSEEventData type(String type) {
    
    this.type = type;
    return this;
  }

   /**
   * type of value (basetype or \&quot;object\&quot;)
   * @return type
  **/
  @jakarta.annotation.Nullable
  @JsonProperty(JSON_PROPERTY_TYPE)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public String getType() {
    return type;
  }


  @JsonProperty(JSON_PROPERTY_TYPE)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setType(String type) {
    this.type = type;
  }


  public SSEEventData schema(String schema) {
    
    this.schema = schema;
    return this;
  }

   /**
   * reference to flatbuffer type in datalayer (only if type is \&quot;object\&quot;)
   * @return schema
  **/
  @jakarta.annotation.Nullable
  @JsonProperty(JSON_PROPERTY_SCHEMA)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public String getSchema() {
    return schema;
  }


  @JsonProperty(JSON_PROPERTY_SCHEMA)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setSchema(String schema) {
    this.schema = schema;
  }


  public SSEEventData value(String value) {
    
    this.value = value;
    return this;
  }

   /**
   * value of the node
   * @return value
  **/
  @jakarta.annotation.Nonnull
  @JsonProperty(JSON_PROPERTY_VALUE)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)

  public String getValue() {
    return value;
  }


  @JsonProperty(JSON_PROPERTY_VALUE)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)
  public void setValue(String value) {
    this.value = value;
  }


  public SSEEventData timestamp(BigDecimal timestamp) {
    
    this.timestamp = timestamp;
    return this;
  }

   /**
   * Timestamp as Filetime
   * @return timestamp
  **/
  @jakarta.annotation.Nonnull
  @JsonProperty(JSON_PROPERTY_TIMESTAMP)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)

  public BigDecimal getTimestamp() {
    return timestamp;
  }


  @JsonProperty(JSON_PROPERTY_TIMESTAMP)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)
  public void setTimestamp(BigDecimal timestamp) {
    this.timestamp = timestamp;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    SSEEventData ssEEventData = (SSEEventData) o;
    return Objects.equals(this.node, ssEEventData.node) &&
        Objects.equals(this.type, ssEEventData.type) &&
        Objects.equals(this.schema, ssEEventData.schema) &&
        Objects.equals(this.value, ssEEventData.value) &&
        Objects.equals(this.timestamp, ssEEventData.timestamp);
  }

  @Override
  public int hashCode() {
    return Objects.hash(node, type, schema, value, timestamp);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class SSEEventData {\n");
    sb.append("    node: ").append(toIndentedString(node)).append("\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("    schema: ").append(toIndentedString(schema)).append("\n");
    sb.append("    value: ").append(toIndentedString(value)).append("\n");
    sb.append("    timestamp: ").append(toIndentedString(timestamp)).append("\n");
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

  /**
   * Convert the instance into URL query string.
   *
   * @return URL query string
   */
  public String toUrlQueryString() {
    return toUrlQueryString(null);
  }

  /**
   * Convert the instance into URL query string.
   *
   * @param prefix prefix of the query string
   * @return URL query string
   */
  public String toUrlQueryString(String prefix) {
    String suffix = "";
    String containerSuffix = "";
    String containerPrefix = "";
    if (prefix == null) {
      // style=form, explode=true, e.g. /pet?name=cat&type=manx
      prefix = "";
    } else {
      // deepObject style e.g. /pet?id[name]=cat&id[type]=manx
      prefix = prefix + "[";
      suffix = "]";
      containerSuffix = "]";
      containerPrefix = "[";
    }

    StringJoiner joiner = new StringJoiner("&");

    // add `node` to the URL query string
    if (getNode() != null) {
      try {
        joiner.add(String.format("%snode%s=%s", prefix, suffix, URLEncoder.encode(String.valueOf(getNode()), "UTF-8").replaceAll("\\+", "%20")));
      } catch (UnsupportedEncodingException e) {
        // Should never happen, UTF-8 is always supported
        throw new RuntimeException(e);
      }
    }

    // add `type` to the URL query string
    if (getType() != null) {
      try {
        joiner.add(String.format("%stype%s=%s", prefix, suffix, URLEncoder.encode(String.valueOf(getType()), "UTF-8").replaceAll("\\+", "%20")));
      } catch (UnsupportedEncodingException e) {
        // Should never happen, UTF-8 is always supported
        throw new RuntimeException(e);
      }
    }

    // add `schema` to the URL query string
    if (getSchema() != null) {
      try {
        joiner.add(String.format("%sschema%s=%s", prefix, suffix, URLEncoder.encode(String.valueOf(getSchema()), "UTF-8").replaceAll("\\+", "%20")));
      } catch (UnsupportedEncodingException e) {
        // Should never happen, UTF-8 is always supported
        throw new RuntimeException(e);
      }
    }

    // add `value` to the URL query string
    if (getValue() != null) {
      try {
        joiner.add(String.format("%svalue%s=%s", prefix, suffix, URLEncoder.encode(String.valueOf(getValue()), "UTF-8").replaceAll("\\+", "%20")));
      } catch (UnsupportedEncodingException e) {
        // Should never happen, UTF-8 is always supported
        throw new RuntimeException(e);
      }
    }

    // add `timestamp` to the URL query string
    if (getTimestamp() != null) {
      try {
        joiner.add(String.format("%stimestamp%s=%s", prefix, suffix, URLEncoder.encode(String.valueOf(getTimestamp()), "UTF-8").replaceAll("\\+", "%20")));
      } catch (UnsupportedEncodingException e) {
        // Should never happen, UTF-8 is always supported
        throw new RuntimeException(e);
      }
    }

    return joiner.toString();
  }

}

