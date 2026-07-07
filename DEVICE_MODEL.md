# DigiMitzves Device Model

## Philosophy

For DigiMitzves, a Device is not a physical Zigbee switch.

A Device is an independently controlled object that owns its own configuration,
its own schedule and its own lifecycle.

The system controls Devices.

The transport layer controls hardware.

This separation allows the automation engine to remain independent from Zigbee,
MQTT or any future transport technology.

---

## Physical hierarchy

One physical Zigbee relay may expose multiple independently controlled channels.

Example

Kitchen Relay

    Channel 1
    Channel 2

Inside DigiMitzves these become

Kitchen Channel 1

Kitchen Channel 2

Each channel is an independent Device.

---

## Device ownership

Zigbee2MQTT owns:

- IEEE address
- pairing
- network topology
- availability
- firmware
- LQI
- endpoints

DigiMitzves owns:

- Device identity
- user configuration
- schedules
- operational status
- Shabbat behaviour
- lifecycle
- warnings

---

## Device lifecycle

Discovered

↓

Named

↓

Configuration Required

↓

Operational

↓

Missing

↓

Ignored

↓

Operational

or

↓

Deleted

---

## Ignore

Ignore does not remove the Device.

Ignore only suppresses the warning state.

If the Device reappears on the network:

Ignored flag is cleared automatically

↓

the previous configuration becomes active again.

---

## Operational definition

A Device is considered operational when

Present

AND

Configured

are both true.

---

## Engine responsibility

The Engine never creates or deletes Devices directly.

The Engine only observes their current state.

Administrative actions are performed through the Web interface.

---

## Design principle

The Engine controls logical Devices.

The transport layer controls physical hardware.
