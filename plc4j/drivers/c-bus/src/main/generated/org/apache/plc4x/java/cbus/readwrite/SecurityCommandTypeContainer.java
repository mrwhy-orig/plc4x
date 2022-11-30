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
package org.apache.plc4x.java.cbus.readwrite;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

// Code generated by code-generation. DO NOT EDIT.

public enum SecurityCommandTypeContainer {
  SecurityCommandOff_0Bytes((short) 0x00, (short) 0, SecurityCommandType.OFF),
  SecurityCommandOff_1Bytes((short) 0x01, (short) 1, SecurityCommandType.OFF),
  SecurityCommandOff_2Bytes((short) 0x02, (short) 2, SecurityCommandType.OFF),
  SecurityCommandOff_3Bytes((short) 0x03, (short) 3, SecurityCommandType.OFF),
  SecurityCommandOff_4Bytes((short) 0x04, (short) 4, SecurityCommandType.OFF),
  SecurityCommandOff_5Bytes((short) 0x05, (short) 5, SecurityCommandType.OFF),
  SecurityCommandOff_6Bytes((short) 0x06, (short) 6, SecurityCommandType.OFF),
  SecurityCommandOff_7Bytes((short) 0x07, (short) 7, SecurityCommandType.OFF),
  SecurityCommandEvent_0Bytes((short) 0x08, (short) 0, SecurityCommandType.EVENT),
  SecurityCommandEvent_1Bytes((short) 0x09, (short) 1, SecurityCommandType.EVENT),
  SecurityCommandEvent_2Bytes((short) 0x0A, (short) 2, SecurityCommandType.EVENT),
  SecurityCommandEvent_3Bytes((short) 0x0B, (short) 3, SecurityCommandType.EVENT),
  SecurityCommandEvent_4Bytes((short) 0x0C, (short) 4, SecurityCommandType.EVENT),
  SecurityCommandEvent_5Bytes((short) 0x0D, (short) 5, SecurityCommandType.EVENT),
  SecurityCommandEvent_6Bytes((short) 0x0E, (short) 6, SecurityCommandType.EVENT),
  SecurityCommandEvent_7Bytes((short) 0x0F, (short) 7, SecurityCommandType.EVENT),
  SecurityCommandOn_0Bytes((short) 0x78, (short) 0, SecurityCommandType.ON),
  SecurityCommandOn_1Bytes((short) 0x79, (short) 1, SecurityCommandType.ON),
  SecurityCommandOn_2Bytes((short) 0x7A, (short) 2, SecurityCommandType.ON),
  SecurityCommandOn_3Bytes((short) 0x7B, (short) 3, SecurityCommandType.ON),
  SecurityCommandOn_4Bytes((short) 0x7C, (short) 4, SecurityCommandType.ON),
  SecurityCommandOn_5Bytes((short) 0x7D, (short) 5, SecurityCommandType.ON),
  SecurityCommandOn_6Bytes((short) 0x7E, (short) 6, SecurityCommandType.ON),
  SecurityCommandOn_7Bytes((short) 0x7F, (short) 7, SecurityCommandType.ON),
  SecurityCommandLongOff_0Bytes((short) 0x80, (short) 8, SecurityCommandType.OFF),
  SecurityCommandLongOff_1Bytes((short) 0x81, (short) 1, SecurityCommandType.OFF),
  SecurityCommandLongOff_2Bytes((short) 0x82, (short) 2, SecurityCommandType.OFF),
  SecurityCommandLongOff_3Bytes((short) 0x83, (short) 3, SecurityCommandType.OFF),
  SecurityCommandLongOff_4Bytes((short) 0x84, (short) 4, SecurityCommandType.OFF),
  SecurityCommandLongOff_5Bytes((short) 0x85, (short) 5, SecurityCommandType.OFF),
  SecurityCommandLongOff_6Bytes((short) 0x86, (short) 6, SecurityCommandType.OFF),
  SecurityCommandLongOff_7Bytes((short) 0x87, (short) 7, SecurityCommandType.OFF),
  SecurityCommandLongOff_8Bytes((short) 0x88, (short) 8, SecurityCommandType.OFF),
  SecurityCommandLongOff_9Bytes((short) 0x89, (short) 9, SecurityCommandType.OFF),
  SecurityCommandLongOff_10Bytes((short) 0x8A, (short) 10, SecurityCommandType.OFF),
  SecurityCommandLongOff_11Bytes((short) 0x8B, (short) 11, SecurityCommandType.OFF),
  SecurityCommandLongOff_12Bytes((short) 0x8C, (short) 12, SecurityCommandType.OFF),
  SecurityCommandLongOff_13Bytes((short) 0x8D, (short) 13, SecurityCommandType.OFF),
  SecurityCommandLongOff_14Bytes((short) 0x8E, (short) 14, SecurityCommandType.OFF),
  SecurityCommandLongOff_15Bytes((short) 0x8F, (short) 15, SecurityCommandType.OFF),
  SecurityCommandLongOff_16Bytes((short) 0x90, (short) 16, SecurityCommandType.OFF),
  SecurityCommandLongOff_17Bytes((short) 0x91, (short) 17, SecurityCommandType.OFF),
  SecurityCommandLongOff_18Bytes((short) 0x92, (short) 18, SecurityCommandType.OFF),
  SecurityCommandLongOff_19Bytes((short) 0x93, (short) 19, SecurityCommandType.OFF),
  SecurityCommandLongOff_20Bytes((short) 0x94, (short) 20, SecurityCommandType.OFF),
  SecurityCommandLongOff_21Bytes((short) 0x95, (short) 21, SecurityCommandType.OFF),
  SecurityCommandLongOff_22Bytes((short) 0x96, (short) 22, SecurityCommandType.OFF),
  SecurityCommandLongOff_23Bytes((short) 0x97, (short) 23, SecurityCommandType.OFF),
  SecurityCommandLongOff_24Bytes((short) 0x98, (short) 24, SecurityCommandType.OFF),
  SecurityCommandLongOff_25Bytes((short) 0x99, (short) 25, SecurityCommandType.OFF),
  SecurityCommandLongOff_26Bytes((short) 0x9A, (short) 26, SecurityCommandType.OFF),
  SecurityCommandLongOff_27Bytes((short) 0x9B, (short) 27, SecurityCommandType.OFF),
  SecurityCommandLongOff_28Bytes((short) 0x9C, (short) 28, SecurityCommandType.OFF),
  SecurityCommandLongOff_29Bytes((short) 0x9D, (short) 29, SecurityCommandType.OFF),
  SecurityCommandLongOff_30Bytes((short) 0x9E, (short) 30, SecurityCommandType.OFF),
  SecurityCommandLongOff_31Bytes((short) 0x9F, (short) 31, SecurityCommandType.OFF),
  SecurityCommandLongEvent_0Bytes((short) 0xA0, (short) 0, SecurityCommandType.EVENT),
  SecurityCommandLongEvent_1Bytes((short) 0xA1, (short) 1, SecurityCommandType.EVENT),
  SecurityCommandLongEvent_2Bytes((short) 0xA2, (short) 2, SecurityCommandType.EVENT),
  SecurityCommandLongEvent_3Bytes((short) 0xA3, (short) 3, SecurityCommandType.EVENT),
  SecurityCommandLongEvent_4Bytes((short) 0xA4, (short) 4, SecurityCommandType.EVENT),
  SecurityCommandLongEvent_5Bytes((short) 0xA5, (short) 5, SecurityCommandType.EVENT),
  SecurityCommandLongEvent_6Bytes((short) 0xA6, (short) 6, SecurityCommandType.EVENT),
  SecurityCommandLongEvent_7Bytes((short) 0xA7, (short) 7, SecurityCommandType.EVENT),
  SecurityCommandLongEvent_8Bytes((short) 0xA8, (short) 8, SecurityCommandType.EVENT),
  SecurityCommandLongEvent_9Bytes((short) 0xA9, (short) 9, SecurityCommandType.EVENT),
  SecurityCommandLongEvent_10Bytes((short) 0xAA, (short) 10, SecurityCommandType.EVENT),
  SecurityCommandLongEvent_11Bytes((short) 0xAB, (short) 11, SecurityCommandType.EVENT),
  SecurityCommandLongEvent_12Bytes((short) 0xAC, (short) 12, SecurityCommandType.EVENT),
  SecurityCommandLongEvent_13Bytes((short) 0xAD, (short) 13, SecurityCommandType.EVENT),
  SecurityCommandLongEvent_14Bytes((short) 0xAE, (short) 14, SecurityCommandType.EVENT),
  SecurityCommandLongEvent_15Bytes((short) 0xAF, (short) 15, SecurityCommandType.EVENT),
  SecurityCommandLongEvent_16Bytes((short) 0xB0, (short) 16, SecurityCommandType.EVENT),
  SecurityCommandLongEvent_17Bytes((short) 0xB1, (short) 17, SecurityCommandType.EVENT),
  SecurityCommandLongEvent_18Bytes((short) 0xB2, (short) 18, SecurityCommandType.EVENT),
  SecurityCommandLongEvent_19Bytes((short) 0xB3, (short) 19, SecurityCommandType.EVENT),
  SecurityCommandLongEvent_20Bytes((short) 0xB4, (short) 20, SecurityCommandType.EVENT),
  SecurityCommandLongEvent_21Bytes((short) 0xB5, (short) 21, SecurityCommandType.EVENT),
  SecurityCommandLongEvent_22Bytes((short) 0xB6, (short) 22, SecurityCommandType.EVENT),
  SecurityCommandLongEvent_23Bytes((short) 0xB7, (short) 23, SecurityCommandType.EVENT),
  SecurityCommandLongEvent_24Bytes((short) 0xB8, (short) 24, SecurityCommandType.EVENT),
  SecurityCommandLongEvent_25Bytes((short) 0xB9, (short) 25, SecurityCommandType.EVENT),
  SecurityCommandLongEvent_26Bytes((short) 0xBA, (short) 26, SecurityCommandType.EVENT),
  SecurityCommandLongEvent_27Bytes((short) 0xBB, (short) 27, SecurityCommandType.EVENT),
  SecurityCommandLongEvent_28Bytes((short) 0xBC, (short) 28, SecurityCommandType.EVENT),
  SecurityCommandLongEvent_29Bytes((short) 0xBD, (short) 29, SecurityCommandType.EVENT),
  SecurityCommandLongEvent_30Bytes((short) 0xBE, (short) 30, SecurityCommandType.EVENT),
  SecurityCommandLongEvent_31Bytes((short) 0xBF, (short) 31, SecurityCommandType.EVENT),
  SecurityCommandLongOn_0Bytes((short) 0xE0, (short) 0, SecurityCommandType.ON),
  SecurityCommandLongOn_1Bytes((short) 0xE1, (short) 1, SecurityCommandType.ON),
  SecurityCommandLongOn_2Bytes((short) 0xE2, (short) 2, SecurityCommandType.ON),
  SecurityCommandLongOn_3Bytes((short) 0xE3, (short) 3, SecurityCommandType.ON),
  SecurityCommandLongOn_4Bytes((short) 0xE4, (short) 4, SecurityCommandType.ON),
  SecurityCommandLongOn_5Bytes((short) 0xE5, (short) 5, SecurityCommandType.ON),
  SecurityCommandLongOn_6Bytes((short) 0xE6, (short) 6, SecurityCommandType.ON),
  SecurityCommandLongOn_7Bytes((short) 0xE7, (short) 7, SecurityCommandType.ON),
  SecurityCommandLongOn_8Bytes((short) 0xE8, (short) 8, SecurityCommandType.ON),
  SecurityCommandLongOn_9Bytes((short) 0xE9, (short) 9, SecurityCommandType.ON),
  SecurityCommandLongOn_10Bytes((short) 0xEA, (short) 10, SecurityCommandType.ON),
  SecurityCommandLongOn_11Bytes((short) 0xEB, (short) 11, SecurityCommandType.ON),
  SecurityCommandLongOn_12Bytes((short) 0xEC, (short) 12, SecurityCommandType.ON),
  SecurityCommandLongOn_13Bytes((short) 0xED, (short) 13, SecurityCommandType.ON),
  SecurityCommandLongOn_14Bytes((short) 0xEE, (short) 14, SecurityCommandType.ON),
  SecurityCommandLongOn_15Bytes((short) 0xEF, (short) 15, SecurityCommandType.ON),
  SecurityCommandLongOn_16Bytes((short) 0xF0, (short) 16, SecurityCommandType.ON),
  SecurityCommandLongOn_17Bytes((short) 0xF1, (short) 17, SecurityCommandType.ON),
  SecurityCommandLongOn_18Bytes((short) 0xF2, (short) 18, SecurityCommandType.ON),
  SecurityCommandLongOn_19Bytes((short) 0xF3, (short) 19, SecurityCommandType.ON),
  SecurityCommandLongOn_20Bytes((short) 0xF4, (short) 20, SecurityCommandType.ON),
  SecurityCommandLongOn_21Bytes((short) 0xF5, (short) 21, SecurityCommandType.ON),
  SecurityCommandLongOn_22Bytes((short) 0xF6, (short) 22, SecurityCommandType.ON),
  SecurityCommandLongOn_23Bytes((short) 0xF7, (short) 23, SecurityCommandType.ON),
  SecurityCommandLongOn_24Bytes((short) 0xF8, (short) 24, SecurityCommandType.ON),
  SecurityCommandLongOn_25Bytes((short) 0xF9, (short) 25, SecurityCommandType.ON),
  SecurityCommandLongOn_26Bytes((short) 0xFA, (short) 26, SecurityCommandType.ON),
  SecurityCommandLongOn_27Bytes((short) 0xFB, (short) 27, SecurityCommandType.ON),
  SecurityCommandLongOn_28Bytes((short) 0xFC, (short) 28, SecurityCommandType.ON),
  SecurityCommandLongOn_29Bytes((short) 0xFD, (short) 29, SecurityCommandType.ON),
  SecurityCommandLongOn_30Bytes((short) 0xFE, (short) 30, SecurityCommandType.ON),
  SecurityCommandLongOn_31Bytes((short) 0xFF, (short) 31, SecurityCommandType.ON);
  private static final Map<Short, SecurityCommandTypeContainer> map;

  static {
    map = new HashMap<>();
    for (SecurityCommandTypeContainer value : SecurityCommandTypeContainer.values()) {
      map.put((short) value.getValue(), value);
    }
  }

  private short value;
  private short numBytes;
  private SecurityCommandType commandType;

  SecurityCommandTypeContainer(short value, short numBytes, SecurityCommandType commandType) {
    this.value = value;
    this.numBytes = numBytes;
    this.commandType = commandType;
  }

  public short getValue() {
    return value;
  }

  public short getNumBytes() {
    return numBytes;
  }

  public static SecurityCommandTypeContainer firstEnumForFieldNumBytes(short fieldValue) {
    for (SecurityCommandTypeContainer _val : SecurityCommandTypeContainer.values()) {
      if (_val.getNumBytes() == fieldValue) {
        return _val;
      }
    }
    return null;
  }

  public static List<SecurityCommandTypeContainer> enumsForFieldNumBytes(short fieldValue) {
    List<SecurityCommandTypeContainer> _values = new ArrayList();
    for (SecurityCommandTypeContainer _val : SecurityCommandTypeContainer.values()) {
      if (_val.getNumBytes() == fieldValue) {
        _values.add(_val);
      }
    }
    return _values;
  }

  public SecurityCommandType getCommandType() {
    return commandType;
  }

  public static SecurityCommandTypeContainer firstEnumForFieldCommandType(
      SecurityCommandType fieldValue) {
    for (SecurityCommandTypeContainer _val : SecurityCommandTypeContainer.values()) {
      if (_val.getCommandType() == fieldValue) {
        return _val;
      }
    }
    return null;
  }

  public static List<SecurityCommandTypeContainer> enumsForFieldCommandType(
      SecurityCommandType fieldValue) {
    List<SecurityCommandTypeContainer> _values = new ArrayList();
    for (SecurityCommandTypeContainer _val : SecurityCommandTypeContainer.values()) {
      if (_val.getCommandType() == fieldValue) {
        _values.add(_val);
      }
    }
    return _values;
  }

  public static SecurityCommandTypeContainer enumForValue(short value) {
    return map.get(value);
  }

  public static Boolean isDefined(short value) {
    return map.containsKey(value);
  }
}