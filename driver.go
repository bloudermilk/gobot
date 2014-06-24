package gobot

import (
	"fmt"
	"time"
)

type Driver struct {
	Adaptor  AdaptorInterface
	Interval time.Duration
	Pin      string
	Name     string
	Commands map[string]func(map[string]interface{}) interface{}
	Events   map[string]*Event
	Type     string
}

type DriverInterface interface {
	Start() bool
	Halt() bool
	adaptor() AdaptorInterface
	setInterval(time.Duration)
	interval() time.Duration
	setName(string)
	name() string
	commands() map[string]func(map[string]interface{}) interface{}
	ToJSON() *JSONDevice
}

func (d *Driver) adaptor() AdaptorInterface {
	return d.Adaptor
}

func (d *Driver) setInterval(t time.Duration) {
	d.Interval = t
}

func (d *Driver) interval() time.Duration {
	return d.Interval
}

func (d *Driver) setName(s string) {
	d.Name = s
}

func (d *Driver) name() string {
	return d.Name
}

func (d *Driver) commands() map[string]func(map[string]interface{}) interface{} {
	return d.Commands
}

func (d *Driver) AddCommand(name string, f func(map[string]interface{}) interface{}) {
	d.Commands[name] = f
}

func NewDriver(name string, t string, commands Commands, a AdaptorInterface) *Driver {
	if name == "" {
		name = fmt.Sprintf("%X", Rand(int(^uint(0)>>1)))
	}
	return &Driver{
		Type:     t,
		Name:     name,
		Interval: 10 * time.Millisecond,
		Commands: commands,
		Adaptor:  a,
	}
}

func (d *Driver) ToJSON() *JSONDevice {
	jsonDevice := &JSONDevice{
		Name:       d.Name,
		Driver:     d.Type,
		Commands:   []string{},
		Connection: nil,
	}

	if d.adaptor() != nil {
		//jsonDevice.Connection = d.Robot.Connection(d.adaptor().name()).ToJSON()
	}

	commands := d.commands()
	for command := range commands {
		jsonDevice.Commands = append(jsonDevice.Commands, command)
	}

	return jsonDevice
}
