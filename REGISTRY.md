# REGISTRY

## Purpose

The Registry represents the Engine's current knowledge about physical devices.

It is the single source of truth describing which devices are known, present, configured, or ignored.

The Registry does not contain schedules, commands, or business logic.

---

# Overview

```text
Physical Devices

      │

      ▼

 Discovery

      │

      ▼

  Registry

      │

      ▼

   Engine
```

The Engine never communicates directly with device discovery.

All information passes through the Registry.

---

# Device Model

Each physical device is represented by exactly one Registry entry.

Example:

```text
Device

ID

Friendly Name

Parent

Channel

Present

Configured

Ignored
```

Future fields may be added without changing Engine logic.

---

# Device Identity

A device is identified by a stable internal ID.

The ID must remain unchanged throughout the lifetime of the device.

Changing a friendly name must never change the device ID.

---

# Device States

A device may be:

## Present

The device is currently visible.

---

## Missing

The device exists in the Registry but was not discovered.

Missing does not automatically mean deleted.

---

## Configured

The user has completed the required configuration.

Configured devices participate in normal operation.

---

## Unconfigured

The device has been discovered but has not yet received user configuration.

This state requires operator attention.

---

## Ignored

The user intentionally removed the device from operation.

Ignored devices remain stored inside the Registry.

If the device appears again during discovery, the Ignored state is automatically cleared and the device is treated as a newly active device requiring configuration.

---

# Synchronization

Synchronization compares:

```text
Discovery

↓

Registry
```

Possible outcomes:

* new device
* known device
* missing device
* recovered device

Synchronize never controls hardware.

Synchronize never creates Jobs.

Synchronize only updates Registry state and Engine Context.

---

# Registry Responsibilities

The Registry is responsible for:

* maintaining device identity
* maintaining operational state
* tracking configuration state
* tracking device presence

The Registry is not responsible for:

* schedules
* Events
* Jobs
* Commands
* MQTT
* Zigbee communication
* hardware control

---

# Registry Invariants

The following rules must always remain true.

1. One physical device corresponds to one Registry entry.

2. Device IDs never change.

3. Registry contains no business logic.

4. Registry never executes hardware actions.

5. Discovery updates the Registry but never bypasses it.

6. Engine always uses the Registry instead of raw discovery data.

7. Ignored devices remain stored until explicitly removed by the user.

---

# Design Philosophy

The Registry is not a device database.

It is the Engine's current view of the physical world.

Every Engine decision concerning devices must be based on the Registry rather than directly on hardware discovery.
