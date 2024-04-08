/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package org.apache.plc4x.java.eip.base;

import org.apache.plc4x.java.api.PlcConnection;
import org.apache.plc4x.java.api.PlcDriverManager;
import org.apache.plc4x.java.api.messages.PlcReadRequest;

import java.util.concurrent.TimeUnit;

public class ManualEipTest {

    public static void main(String[] args) {
        try (PlcConnection plcConnection = PlcDriverManager.getDefault().getConnectionManager().getConnection("eip:tcp://192.168.23.10:44818")) {
            PlcReadRequest.Builder builder = plcConnection.readRequestBuilder();
            builder.addTagAddress("param", "%out01");
            var readRequest = builder.build();
            var readResponse = readRequest.execute().get(5000, TimeUnit.MILLISECONDS);
            System.out.println(readResponse.getResponseCode("param"));
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
    }

}