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

package bacgopes

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/apache/plc4x/plc4go/protocols/bacnetip/readwrite/model"
	"github.com/apache/plc4x/plc4go/spi"
)

type DistributeBroadcastToNetwork struct {
	*_BVLPDU

	// post construct function
	_postConstruct []func()
}

var _ BVLPDU = (*DistributeBroadcastToNetwork)(nil)

func NewDistributeBroadcastToNetwork(pdu PDU, opts ...func(*DistributeBroadcastToNetwork)) (*DistributeBroadcastToNetwork, error) {
	o := &DistributeBroadcastToNetwork{}
	for _, opt := range opts {
		opt(o)
	}
	switch npdu := pdu.(type) {
	case model.NPDU:
		o._BVLPDU = NewBVLPDU(model.NewBVLCDistributeBroadcastToNetwork(o.produceInnerNPDU(npdu))).(*_BVLPDU)
	case nil:
		o._BVLPDU = NewBVLPDU(nil).(*_BVLPDU)
	default:
		// TODO: re-encode seems expensive... check if there is a better option (e.g. only do it on the message bridge)
		data := pdu.GetPduData()
		parse, err := model.NPDUParse(context.Background(), data, uint16(len(data)))
		if err != nil {
			return nil, errors.Wrap(err, "error re-encoding")
		}
		o._BVLPDU = NewBVLPDU(model.NewBVLCDistributeBroadcastToNetwork(o.produceInnerNPDU(parse))).(*_BVLPDU)
	}
	// Do a post construct for a bit more easy initialization
	for _, f := range o._postConstruct {
		f()
	}
	o._postConstruct = nil
	return o, nil
}

func WithDistributeBroadcastToNetworkDestination(destination *Address) func(*DistributeBroadcastToNetwork) {
	return func(o *DistributeBroadcastToNetwork) {
		o._postConstruct = append(o._postConstruct, func() {
			o.SetPDUDestination(destination)
		})
	}
}

func WithDistributeBroadcastToNetworkUserData(userData spi.Message) func(*DistributeBroadcastToNetwork) {
	return func(o *DistributeBroadcastToNetwork) {
		o._postConstruct = append(o._postConstruct, func() {
			o.SetPDUUserData(userData)
		})
	}
}

func (d *DistributeBroadcastToNetwork) produceInnerNPDU(inNpdu model.NPDU) (npdu model.NPDU, bvlcPayloadLength uint16) {
	npdu = inNpdu
	return
}

func (d *DistributeBroadcastToNetwork) Encode(bvlpdu Arg) error {
	switch bvlpdu := bvlpdu.(type) {
	case BVLCI:
		if err := bvlpdu.getBVLCI().Update(d); err != nil {
			return errors.Wrap(err, "error updating BVLPDU")
		}
	}

	switch bvlpdu := bvlpdu.(type) {
	case PDUData:
		bvlpdu.PutData(d.GetPduData()...)
	default:
		return errors.Errorf("invalid BVLPDU type %T", bvlpdu)
	}
	return nil
}

func (d *DistributeBroadcastToNetwork) Decode(bvlpdu Arg) error {
	if err := d._BVLCI.Update(bvlpdu); err != nil {
		return errors.Wrap(err, "error updating BVLCI")
	}
	switch bvlpdu := bvlpdu.(type) {
	case BVLPDU:
		switch rm := bvlpdu.GetRootMessage().(type) {
		case model.BVLCDistributeBroadcastToNetwork:
			d.rootMessage = rm
		}
	}
	switch bvlpdu := bvlpdu.(type) {
	case PDUData:
		d.SetPduData(bvlpdu.GetPduData())
	}
	return nil
}

func (d *DistributeBroadcastToNetwork) String() string {
	if d == nil {
		return "(*DistributeBroadcastToNetwork)(nil)"
	}
	return fmt.Sprintf("DistributeBroadcastToNetwork{%s}", d._BVLPDU)
}