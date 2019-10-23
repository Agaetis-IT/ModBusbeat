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
