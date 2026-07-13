# ENGINE

## Overview

The Engine is the central controller of DigiMitzves.

It continuously observes the system, evaluates its current state, makes decisions, and performs the minimum required actions to keep the system consistent.

The Engine never assumes that previous state information is still valid.

Every cycle starts with a fresh observation of the current system state.

---

# Engine Cycle

```text
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

---

# Validate

Purpose:

Verify that the Engine has all required dependencies before any processing begins.

Typical checks:

* Configuration loaded
* Registry available
* State available
* Scheduler available

Validate never modifies the system.

---

# DetectEvent

Purpose:

Determine whether the current time belongs to a scheduled event.

Output:

* Current Event
* No Event

DetectEvent never creates Jobs.

DetectEvent never executes commands.

---

# DetectMode

Purpose:

Determine the current operating mode.

Examples:

* WeekdayMode
* ShabbatMode

Future modes may include:

* InstallationMode
* RecoveryMode
* ConfigurationMode

Mode detection depends only on current facts.

---

# Recover

Purpose:

Restore consistency after reboot or unexpected interruption.

Examples:

* unfinished Job
* partially executed commands
* inconsistent state

Recovery always has higher priority than execution.

---

# Synchronize

Purpose:

Synchronize the Registry with the currently discovered devices.

Synchronize may:

* register new devices
* update presence
* clear recovered devices
* raise configuration requirements

Synchronize never:

* sends MQTT commands
* changes relay state
* creates Jobs
* executes actions

Synchronize is designed to be idempotent.

Running Synchronize repeatedly without changes must not modify the system.

---

# Reconcile

Purpose:

Compare the desired state with the actual state.

Typical responsibilities:

* determine whether a Job exists
* determine whether the Job is still valid
* determine whether a new Job must be created
* determine whether execution is required

Reconcile produces decisions.

It does not execute them.

---

# Execute

Purpose:

Perform actions approved by Reconcile.

Typical actions:

* publish MQTT messages
* activate relays
* update command status

Execute never creates decisions.

---

# Finalize

Purpose:

Persist the results of the current cycle.

Typical operations:

* save Registry
* save State
* save Job
* clear temporary Context flags

Finalize prepares the system for the next cycle.

---

# Context

Context represents the result of one Engine cycle.

Typical flags:

```text
RecoveryRequired

SynchronizationRequired

ExecutionRequired

SaveRequired

ConfigurationRequired

NewDevices

MissingDevices
```

Context is recreated every cycle.

It is never persisted.

---

# Core Responsibilities

The Engine is responsible for:

* observing
* deciding
* coordinating

The Engine is not responsible for:

* MQTT protocol
* GPIO access
* Zigbee implementation
* hardware drivers
* web interface

---

# Engine Invariants

The following rules must always remain true.

1. Every cycle starts from current observations.

2. Recovery always precedes execution.

3. Synchronize never changes device state.

4. Reconcile never communicates with hardware.

5. Execute never makes decisions.

6. Finalize never changes business logic.

7. Context exists for one cycle only.

8. Hardware failures must not corrupt Engine state.

9. Every decision must be deterministic.

10. Every cycle must leave the system in a consistent state.

---

# Design Philosophy

The Engine behaves as a deterministic state machine rather than an event callback processor.

Its purpose is not to react immediately to every change.

Its purpose is to repeatedly verify that the current world matches the expected world and to perform only the minimum actions required to restore consistency.
