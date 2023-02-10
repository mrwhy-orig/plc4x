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
package org.apache.plc4x.java.eip.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class CipReadResponse extends CipService implements Message {

  // Accessors for discriminator values.
  public Short getService() {
    return (short) 0xCC;
  }

  // Properties.
  protected final short status;
  protected final short extStatus;
  protected final CIPDataTypeCode dataType;
  protected final byte[] data;

  public CipReadResponse(short status, short extStatus, CIPDataTypeCode dataType, byte[] data) {
    super();
    this.status = status;
    this.extStatus = extStatus;
    this.dataType = dataType;
    this.data = data;
  }

  public short getStatus() {
    return status;
  }

  public short getExtStatus() {
    return extStatus;
  }

  public CIPDataTypeCode getDataType() {
    return dataType;
  }

  public byte[] getData() {
    return data;
  }

  @Override
  protected void serializeCipServiceChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("CipReadResponse");

    // Reserved Field (reserved)
    writeReservedField("reserved", (short) 0x00, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (status)
    writeSimpleField("status", status, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (extStatus)
    writeSimpleField("extStatus", extStatus, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (dataType)
    writeSimpleEnumField(
        "dataType",
        "CIPDataTypeCode",
        dataType,
        new DataWriterEnumDefault<>(
            CIPDataTypeCode::getValue, CIPDataTypeCode::name, writeUnsignedInt(writeBuffer, 16)));

    // Array Field (data)
    writeByteArrayField("data", data, writeByteArray(writeBuffer, 8));

    writeBuffer.popContext("CipReadResponse");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    CipReadResponse _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Reserved Field (reserved)
    lengthInBits += 8;

    // Simple field (status)
    lengthInBits += 8;

    // Simple field (extStatus)
    lengthInBits += 8;

    // Simple field (dataType)
    lengthInBits += 16;

    // Array field
    if (data != null) {
      lengthInBits += 8 * data.length;
    }

    return lengthInBits;
  }

  public static CipServiceBuilder staticParseCipServiceBuilder(
      ReadBuffer readBuffer, Integer serviceLen) throws ParseException {
    readBuffer.pullContext("CipReadResponse");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    Short reservedField0 =
        readReservedField("reserved", readUnsignedShort(readBuffer, 8), (short) 0x00);

    short status = readSimpleField("status", readUnsignedShort(readBuffer, 8));

    short extStatus = readSimpleField("extStatus", readUnsignedShort(readBuffer, 8));

    CIPDataTypeCode dataType =
        readEnumField(
            "dataType",
            "CIPDataTypeCode",
            new DataReaderEnumDefault<>(
                CIPDataTypeCode::enumForValue, readUnsignedInt(readBuffer, 16)));

    byte[] data = readBuffer.readByteArray("data", Math.toIntExact((serviceLen) - (6)));

    readBuffer.closeContext("CipReadResponse");
    // Create the instance
    return new CipReadResponseBuilderImpl(status, extStatus, dataType, data);
  }

  public static class CipReadResponseBuilderImpl implements CipService.CipServiceBuilder {
    private final short status;
    private final short extStatus;
    private final CIPDataTypeCode dataType;
    private final byte[] data;

    public CipReadResponseBuilderImpl(
        short status, short extStatus, CIPDataTypeCode dataType, byte[] data) {
      this.status = status;
      this.extStatus = extStatus;
      this.dataType = dataType;
      this.data = data;
    }

    public CipReadResponse build() {
      CipReadResponse cipReadResponse = new CipReadResponse(status, extStatus, dataType, data);
      return cipReadResponse;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof CipReadResponse)) {
      return false;
    }
    CipReadResponse that = (CipReadResponse) o;
    return (getStatus() == that.getStatus())
        && (getExtStatus() == that.getExtStatus())
        && (getDataType() == that.getDataType())
        && (getData() == that.getData())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getStatus(), getExtStatus(), getDataType(), getData());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}