# ARCHITECTURE

## Purpose

DigiMitzves is designed as a deterministic appliance for Shabbat and Yom Tov automation.

Unlike general-purpose home automation systems, DigiMitzves does not continuously make decisions based on external events. Instead, it continuously verifies that the system state matches a previously calculated execution plan.

The architecture is centered around deterministic state transitions, recovery, and reliability.

---

# System Overview

```text
                   Configuration
                         │
                         ▼
                  Event Scheduler
                         │
                         ▼
                 Current Event Detection
                         │
                         ▼
                     Compiler
                         │
                         ▼
                       Job
                         │
                         ▼
                     Engine
                         │
          ┌──────────────┼──────────────┐
          ▼              ▼              ▼
      Recovery     Synchronize     Reconcile
                         │
                         ▼
                     Execute
                         │
                         ▼
                    Hardware
```

Device discovery follows an independent path.

```text
Devices

    │

Discovery

    │

Registry

    │

Engine
```

---

# Architectural Layers

## Configuration Layer

Contains:

* user settings
* schedules
* device preferences
* Shabbat configuration

Persistent.

---

## Scheduler

Responsible for:

* locating the current event
* locating the next event
* determining event boundaries

Scheduler never creates commands.

---

## Compiler

Input:

* Event
* Configuration
* Current State

Output:

* Job

Compiler creates a complete execution plan.

Compiler never executes anything.

---

## Job

A Job is a complete execution plan for one event.

```
Job

├── Event

├── Status

├── Commands[]
```

A Job is immutable after creation.

If conditions change, a new Job is created.

---

## Engine

The Engine is the central decision maker.

Engine executes a continuous control loop.

```
Validate

↓

DetectEvent

↓

DetectMode

↓

Recover

↓

Synchronize

↓

Reconcile

↓

Execute

↓

Finalize
```

Each stage has exactly one responsibility.

Stages communicate only through:

* Engine state
* Registry
* Job
* Context

Stages never invoke each other directly.

---

## Context

Context represents the observations and decisions of a single Engine cycle.

It is recreated every cycle.

Typical flags include:

* RecoveryRequired
* SynchronizationRequired
* ExecutionRequired
* SaveRequired
* ConfigurationRequired
* NewDevices
* MissingDevices

Context is not persistent.

---

## Registry

Registry represents the currently known devices.

Each device has one unique record.

Registry stores:

* identity
* configuration state
* presence
* ignore state

Registry never stores schedules.

Registry never executes commands.

---

## Executor

Executor performs physical actions.

Executor knows:

* hardware protocol
* MQTT
* GPIO
* Zigbee

Executor never makes decisions.

---

## Hardware

Hardware is the final execution layer.

Possible implementations include:

* Zigbee2MQTT
* GPIO
* RS-485
* future adapters

The Engine must remain completely independent from hardware implementation.

---

# Control Philosophy

The system follows a repeating observation cycle.

```
Observe

↓

Context

↓

Decision

↓

Execution

↓

Persistence
```

Every cycle begins with observing the world.

No stage assumes that previous assumptions are still valid.

---

# Recovery Philosophy

Recovery has higher priority than execution.

Before executing any command, the Engine must first verify:

* current state
* active event
* device availability
* Job validity

Only after the system is consistent may execution continue.

---

# Design Objectives

The architecture is designed to provide:

* deterministic behavior
* repeatable execution
* hardware independence
* autonomous operation
* power-loss recovery
* simple testing
* long-term maintainability

These goals take precedence over execution speed or implementation convenience.
