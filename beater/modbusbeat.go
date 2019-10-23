package beater

import (
    "encoding/binary"
    "fmt"
    "github.com/elastic/beats/libbeat/beat"
    "github.com/elastic/beats/libbeat/common"
    "github.com/elastic/beats/libbeat/logp"
    "github.com/goburrow/modbus"
    "modbusbeat/config"
    "time"
)

// Modbusbeat configuration.
type Modbusbeat struct {
    done   chan struct{}
    config config.Config
    client beat.Client
    beat   *beat.Beat
}

// New creates an instance of modbusbeat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
    c := config.DefaultConfig
    if err := cfg.Unpack(&c); err != nil {
        return nil, fmt.Errorf("error reading config file: %v", err)
    }

    bt := &Modbusbeat{
        done:   make(chan struct{}),
        config: c,
        beat:   b,
    }

    return bt, nil
}

// Run starts modbusbeat.
func (mb *Modbusbeat) Run(b *beat.Beat) error {
    logp.L().Info("modbusbeat is running! Hit CTRL-C to stop it.")

    var err error
    mb.client, err = b.Publisher.Connect()

    if err != nil {
        return err
    }

    ticker := time.NewTicker(mb.config.Period)
    for {
        select {
        case <-mb.done:
            return nil
        case <-ticker.C:
        }

        err := mb.refreshMetrics()

        if err != nil {
            return err
        }
    }
}

func (mb *Modbusbeat) refreshMetrics() error {
    for _,device := range mb.config.Devices {

        err := mb.fetchMetric(&device)

        if err != nil {
            return err
        }
    }

    return nil
}

func (mb *Modbusbeat) fetchMetric(device *config.DeviceConfig) error {

    handler := modbus.NewTCPClientHandler(device.Address + ":502")
    handler.Timeout = 10 * time.Second
    clientmb := modbus.NewClient(handler)

    for _, register := range device.Registers {
        for _, address := range register.Addresses {
            var res []byte
            var err error

            switch registerType := register.Type; registerType {
            case "Holding":
                res, err = clientmb.ReadHoldingRegisters(address, 1)
            case "Input":
                res, err = clientmb.ReadInputRegisters(address, 1)
            case "Coil":
                res, err = clientmb.ReadCoils(address, 1)
            case "Discrete":
                res, err = clientmb.ReadDiscreteInputs(address, 1)
            }

            if err != nil {
                return err
            }

            data := binary.BigEndian.Uint16(res)

            if err != nil {
                return err
            }

            event := beat.Event{
                Timestamp: time.Now(),
                Fields: common.MapStr{
                    "modbusbeat": common.MapStr{
                        "type":    register.Type,
                        "address": address,
                        "value":   data,
                    },
                },
            }

            mb.client.Publish(event)
            logp.L().Info("Event sent")
        }
    }

    err := handler.Close()

    if err != nil {
        return err
    }

    return nil
}

// Stop stops modbusbeat.
func (mb *Modbusbeat) Stop() {
    err := mb.client.Close()

    if err != nil {
        logp.L().Info("Erreur client close()")
    }

    close(mb.done)
}
