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

package object

import (
	. "github.com/apache/plc4x/plc4go/internal/bacnetip/bacgopes/basetypes"
	. "github.com/apache/plc4x/plc4go/internal/bacnetip/bacgopes/comp"
	. "github.com/apache/plc4x/plc4go/internal/bacnetip/bacgopes/primitivedata"
)

type LoopObject struct {
	Object
	objectType           string // TODO: migrateme
	properties           []Property
	_object_supports_cov bool
}

func NewLoopObject(arg Arg) (*LoopObject, error) {
	o := &LoopObject{
		objectType:           "loop",
		_object_supports_cov: true,
		properties: []Property{
			NewReadableProperty("presentValue", V2P(NewReal)),
			NewReadableProperty("statusFlags", V2P(NewStatusFlags)),
			NewReadableProperty("eventState", V2P(NewEventState)),
			NewOptionalProperty("reliability", V2P(NewReliability)),
			NewReadableProperty("outOfService", V2P(NewBoolean)),
			NewReadableProperty("updateInterval", V2P(NewUnsigned)),
			NewReadableProperty("outputUnits", V2P(NewEngineeringUnits)),
			NewReadableProperty("manipulatedVariableReference", V2P(NewObjectPropertyReference)),
			NewReadableProperty("controlledVariableReference", V2P(NewObjectPropertyReference)),
			NewReadableProperty("controlledVariableValue", V2P(NewReal)),
			NewReadableProperty("controlledVariableUnits", V2P(NewEngineeringUnits)),
			NewReadableProperty("setpointReference", V2P(NewSetpointReference)),
			NewReadableProperty("setpoint", V2P(NewReal)),
			NewReadableProperty("action", V2P(NewAction)),
			NewOptionalProperty("proportionalConstant", V2P(NewReal)),
			NewOptionalProperty("proportionalConstantUnits", V2P(NewEngineeringUnits)),
			NewOptionalProperty("integralConstant", V2P(NewReal)),
			NewOptionalProperty("integralConstantUnits", V2P(NewEngineeringUnits)),
			NewOptionalProperty("derivativeConstant", V2P(NewReal)),
			NewOptionalProperty("derivativeConstantUnits", V2P(NewEngineeringUnits)),
			NewOptionalProperty("bias", V2P(NewReal)),
			NewOptionalProperty("maximumOutput", V2P(NewReal)),
			NewOptionalProperty("minimumOutput", V2P(NewReal)),
			NewReadableProperty("priorityForWriting", V2P(NewUnsigned)),
			NewOptionalProperty("covIncrement", V2P(NewReal)),
			NewOptionalProperty("timeDelay", V2P(NewUnsigned)),
			NewOptionalProperty("notificationClass", V2P(NewUnsigned)),
			NewOptionalProperty("errorLimit", V2P(NewReal)),
			NewOptionalProperty("deadband", V2P(NewReal)),
			NewOptionalProperty("eventEnable", V2P(NewEventTransitionBits)),
			NewOptionalProperty("ackedTransitions", V2P(NewEventTransitionBits)),
			NewOptionalProperty("notifyType", V2P(NewNotifyType)),
			NewOptionalProperty("eventTimeStamps", ArrayOfP(NewTimeStamp, 3, 0)),
			NewOptionalProperty("eventMessageTexts", ArrayOfP(NewCharacterString, 3, 0)),
			NewOptionalProperty("eventMessageTextsConfig", ArrayOfP(NewCharacterString, 3, 0)),
			NewOptionalProperty("eventDetectionEnable", V2P(NewBoolean)),
			NewOptionalProperty("eventAlgorithmInhibitRef", V2P(NewObjectPropertyReference)),
			NewOptionalProperty("eventAlgorithmInhibit", V2P(NewBoolean)),
			NewOptionalProperty("timeDelayNormal", V2P(NewUnsigned)),
			NewOptionalProperty("reliabilityEvaluationInhibit", V2P(NewBoolean)),
			NewOptionalProperty("lowDiffLimit", V2P(NewOptionalReal)),
		},
	}
	// TODO: @register_object_type
	return o, nil
}