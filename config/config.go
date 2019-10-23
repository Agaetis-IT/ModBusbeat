// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import "time"

type Config struct {
	Period  time.Duration  `config:"period"`
	Devices []DeviceConfig `config:"devices"`
}

var DefaultConfig = Config{
	Period: 5 * time.Second,
}

type DeviceConfig struct {
	Address   string           `config:"address"`
	Registers []RegisterConfig `config:"registers"`
}

type RegisterConfig struct {
	Type      string   `config:"type"`
	Addresses []uint16 `config:"addresses"`
}

/**
 * Example config
 */
//modbusbeat:
//	period: xs
//	devices:
//		- address: 127.0.0.1
//		  registers:
//			- type: "Holding"
//				address:
//					- 101
//					- 0x64
//			- type: "Input"
//				address:
//					- 101
//					- 0x64
//			- type: "Coil"
//				address:
//					- [101,2] // ???
