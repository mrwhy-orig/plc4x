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

package org.apache.plc4x.java.ctrlx.readwrite.rest.datalayer.api;

import com.fasterxml.jackson.core.type.TypeReference;
import org.apache.plc4x.java.ctrlx.readwrite.rest.datalayer.ApiClient;
import org.apache.plc4x.java.ctrlx.readwrite.rest.datalayer.ApiException;
import org.apache.plc4x.java.ctrlx.readwrite.rest.datalayer.Configuration;
import org.apache.plc4x.java.ctrlx.readwrite.rest.datalayer.Pair;
import org.apache.plc4x.java.ctrlx.readwrite.rest.datalayer.model.Reflection;

import java.util.*;

@jakarta.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-11-18T13:34:36.056861+01:00[Europe/Berlin]")
public class TypesApi {
  private ApiClient apiClient;

  public TypesApi() {
    this(Configuration.getDefaultApiClient());
  }

  public TypesApi(ApiClient apiClient) {
    this.apiClient = apiClient;
  }

  public ApiClient getApiClient() {
    return apiClient;
  }

  public void setApiClient(ApiClient apiClient) {
    this.apiClient = apiClient;
  }

  /**
   * Generic type access
   * Returns a type definition. Types can be referenced by nodes. A type describes the layout and structure of a complex node value.
   * @param typeName The TypeName is a string. It can be separated by &#39;/&#39; (required)
   * @return Reflection
   * @throws ApiException if fails to make API call
   */
  public Reflection readTypeNode(String typeName) throws ApiException {
    Object localVarPostBody = null;
    
    // verify the required parameter 'typeName' is set
    if (typeName == null) {
      throw new ApiException(400, "Missing the required parameter 'typeName' when calling readTypeNode");
    }
    
    // create path and map variables
    String localVarPath = "/types/{TypeName}"
      .replaceAll("\\{" + "TypeName" + "\\}", apiClient.escapeString(typeName.toString()));

    StringJoiner localVarQueryStringJoiner = new StringJoiner("&");
    String localVarQueryParameterBaseName;
    List<Pair> localVarQueryParams = new ArrayList<Pair>();
    List<Pair> localVarCollectionQueryParams = new ArrayList<Pair>();
    Map<String, String> localVarHeaderParams = new HashMap<String, String>();
    Map<String, String> localVarCookieParams = new HashMap<String, String>();
    Map<String, Object> localVarFormParams = new HashMap<String, Object>();

    
    
    
    final String[] localVarAccepts = {
      "application/json", "application/octet-stream", "plain/text"
    };
    final String localVarAccept = apiClient.selectHeaderAccept(localVarAccepts);

    final String[] localVarContentTypes = {
      
    };
    final String localVarContentType = apiClient.selectHeaderContentType(localVarContentTypes);

    String[] localVarAuthNames = new String[] { "Bearer", "UsernamePassword" };

    TypeReference<Reflection> localVarReturnType = new TypeReference<Reflection>() {};
    return apiClient.invokeAPI(
        localVarPath,
        "GET",
        localVarQueryParams,
        localVarCollectionQueryParams,
        localVarQueryStringJoiner.toString(),
        localVarPostBody,
        localVarHeaderParams,
        localVarCookieParams,
        localVarFormParams,
        localVarAccept,
        localVarContentType,
        localVarAuthNames,
        localVarReturnType
    );
  }
}
