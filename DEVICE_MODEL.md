# DEVICE MODEL

## Purpose

This document defines the meaning of a Device within DigiMitzves.

A Device is a logical entity used by the Engine to perform automation.

It is intentionally independent from any specific transport or hardware implementation.

---

# Philosophy

For DigiMitzves, a Device is **not** a physical Zigbee switch.

A Device is an independently controlled logical object.

The Engine controls Devices.

The transport layer controls hardware.

This separation allows the automation engine to remain independent from Zigbee, MQTT, GPIO, RS-485, or any future communication technology.

---

# Logical Device

A Device owns:

* identity
* configuration
* schedule
* operational state

A Device does **not** own:

* IEEE address
* MQTT topics
* Zigbee pairing
* network topology
* transport protocol

Those belong to the transport layer.

---

# Physical Representation

A single physical device may expose one or more independently controlled endpoints.

Example:

```text
Kitchen Relay

├── Channel 1
└── Channel 2
```

Inside DigiMitzves:

```text
Kitchen Channel 1

Kitchen Channel 2
```

Each channel becomes an independent Device.

From the Engine's perspective, there is no difference between:

* one physical relay,
* one channel of a multi-channel relay,
* a virtual actuator,
* or any future controllable endpoint.

Every controllable endpoint is represented as one logical Device.

---

# Ownership

## Transport Layer

The transport layer owns:

* hardware communication
* device discovery
* pairing
* availability detection
* firmware management
* protocol-specific information

Examples include:

* Zigbee2MQTT
* GPIO
* RS-485
* future adapters

---

## DigiMitzves

DigiMitzves owns:

* Device identity
* user configuration
* schedules
* operational state
* Shabbat behavior

The Engine never depends on transport-specific details.

---

# Device Identity

Each Device has one stable internal identifier.

The identifier must never change during the lifetime of the Device.

Changing a friendly name does not create a new Device.

Changing the transport implementation does not create a new Device.

---

# Operational State

Whether a Device is currently available or configured is determined by the Registry.

The Device itself remains a logical object regardless of its current operational state.

---

# Engine Responsibility

The Engine never communicates directly with physical hardware.

It always operates on logical Devices.

The Engine never needs to know:

* IEEE addresses
* MQTT topics
* Zigbee endpoints
* transport implementation

It only knows Devices.

---

# Design Principles

A Device represents **what** can be controlled.

The transport layer determines **how** it is controlled.

Business logic operates exclusively on logical Devices.

Hardware-specific details remain isolated outside the Engine.
