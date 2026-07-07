# DigiMitzves Architecture

## Overview

DigiMitzves is designed as an autonomous embedded appliance rather than a traditional smart home application.

The system operates entirely offline and is intended to continue functioning predictably regardless of Internet connectivity, cloud services or external infrastructure.

The primary architectural goal is deterministic behavior.

---

# Design Principles

The project follows several core principles.

## Deterministic

The same input must always produce the same output.

The system never relies on randomness, cloud state or asynchronous external services.

---

## Autonomous

All decisions are made locally.

No Internet connection is required for normal operation.

---

## Recoverable

Power loss is considered a normal operating condition.

After any reboot the system must recover its state automatically without user intervention.

---

## Observable

The current state of the system must always be derivable from persistent data.

The engine never depends on hidden runtime state.

---

## Idempotent

Every engine cycle should safely repeat the same operations without causing duplicate actions.

Running the control loop twice must produce the same result as running it once.

---

# High-Level Architecture

```
           Events
              │
              ▼
         Scheduler
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
     ┌────────┴────────┐
     ▼                 ▼
 Reconciliation    Execution
     │                 │
     └────────┬────────┘
              ▼
        Physical Devices
```

---

# Engine

The Engine is the central control loop.

It continuously observes the current system state, evaluates required actions and executes only the necessary operations.

The engine itself does not contain hardware-specific logic.

Its responsibility is orchestration.

Current lifecycle:

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

---

# Compiler

The compiler produces a Job.

A Job is not execution.

A Job is a plan describing what should happen.

---

# Job

A Job represents a complete execution plan for one event.

A Job contains:

- metadata
- event boundaries
- execution status
- command list

Commands are immutable plans.

Execution updates their status but never changes their intent.

---

# Reconciliation

Reconciliation compares three independent realities.

```
Current State

Desired State

Current Job
```

Only after comparing all three does the Engine decide whether action is required.

---

# Execution

Execution performs commands that have become due.

Execution never makes business decisions.

All decisions are made earlier by the Engine.

---

# Persistence

The appliance stores persistent information locally.

Examples include:

- configuration
- event schedule
- runtime state
- execution jobs

Persistent storage allows complete recovery after reboot.

---

# Hardware Abstraction

The Engine never communicates directly with hardware.

Future architecture:

```
Engine

   │

Executor Interface

   │

MQTT Adapter

   │

Mosquitto

   │

Zigbee2MQTT

   │

Zigbee Devices
```

This separation allows different transport layers without changing business logic.

---

# State Machines

The system is evolving toward multiple independent state machines.

Examples include:

- Runtime mode
- Shabbat mode
- Device state
- Recovery state
- Installation mode

State machines communicate through events rather than direct calls.

---

# Future Direction

The long-term objective is a reliable embedded appliance capable of operating continuously for years with predictable behavior and automatic recovery after any interruption.
