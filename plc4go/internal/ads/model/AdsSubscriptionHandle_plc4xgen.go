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

// Code generated by "plc4xgenerator -type=AdsSubscriptionHandle"; DO NOT EDIT.

package model

import (
	"context"
	"encoding/binary"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

var _ = fmt.Printf

func (d *AdsSubscriptionHandle) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := d.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (d *AdsSubscriptionHandle) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	if err := writeBuffer.PushContext("AdsSubscriptionHandle"); err != nil {
		return err
	}

	if d.subscriber != nil {
		if serializableField, ok := d.subscriber.(utils.Serializable); ok {
			if err := writeBuffer.PushContext("subscriber"); err != nil {
				return err
			}
			if err := serializableField.SerializeWithWriteBuffer(ctx, writeBuffer); err != nil {
				return err
			}
			if err := writeBuffer.PopContext("subscriber"); err != nil {
				return err
			}
		} else {
			stringValue := fmt.Sprintf("%v", d.subscriber)
			if err := writeBuffer.WriteString("subscriber", uint32(len(stringValue)*8), "UTF-8", stringValue); err != nil {
				return err
			}
		}
	}

	if err := writeBuffer.WriteString("tagName", uint32(len(d.tagName)*8), "UTF-8", d.tagName); err != nil {
		return err
	}
	{
		_value := fmt.Sprintf("%v", d.directTag)

		if err := writeBuffer.WriteString("directTag", uint32(len(_value)*8), "UTF-8", _value); err != nil {
			return err
		}
	}
	if err := writeBuffer.PushContext("consumers", utils.WithRenderAsList(true)); err != nil {
		return err
	}
	for _, elem := range d.consumers {
		var elem any = elem

		if elem != nil {
			if serializableField, ok := elem.(utils.Serializable); ok {
				if err := writeBuffer.PushContext("value"); err != nil {
					return err
				}
				if err := serializableField.SerializeWithWriteBuffer(ctx, writeBuffer); err != nil {
					return err
				}
				if err := writeBuffer.PopContext("value"); err != nil {
					return err
				}
			} else {
				stringValue := fmt.Sprintf("%v", elem)
				if err := writeBuffer.WriteString("value", uint32(len(stringValue)*8), "UTF-8", stringValue); err != nil {
					return err
				}
			}
		}
	}
	if err := writeBuffer.PopContext("consumers", utils.WithRenderAsList(true)); err != nil {
		return err
	}
	if err := writeBuffer.PushContext("_options", utils.WithRenderAsList(true)); err != nil {
		return err
	}
	for _, elem := range d._options {
		var elem any = elem

		if elem != nil {
			if serializableField, ok := elem.(utils.Serializable); ok {
				if err := writeBuffer.PushContext("value"); err != nil {
					return err
				}
				if err := serializableField.SerializeWithWriteBuffer(ctx, writeBuffer); err != nil {
					return err
				}
				if err := writeBuffer.PopContext("value"); err != nil {
					return err
				}
			} else {
				stringValue := fmt.Sprintf("%v", elem)
				if err := writeBuffer.WriteString("value", uint32(len(stringValue)*8), "UTF-8", stringValue); err != nil {
					return err
				}
			}
		}
	}
	if err := writeBuffer.PopContext("_options", utils.WithRenderAsList(true)); err != nil {
		return err
	}
	if err := writeBuffer.PopContext("AdsSubscriptionHandle"); err != nil {
		return err
	}
	return nil
}

func (d *AdsSubscriptionHandle) String() string {
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), d); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}