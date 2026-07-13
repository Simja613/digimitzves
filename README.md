# DigiMitzves

DigiMitzves is an autonomous home automation system designed specifically for Shabbat and Yom Tov observance.

The project is built as a dedicated appliance rather than a general-purpose home automation platform. Its primary goal is deterministic, reliable and recoverable execution of preplanned actions, even after unexpected power loss or system restart.

## Design Goals

* Autonomous operation
* No cloud dependency
* Deterministic behavior
* State-based architecture
* Automatic recovery after reboot
* Hardware abstraction
* Long-term reliability

## High-Level Architecture

```
Configuration
      │
      ▼
 Scheduler
      │
      ▼
  Event Detection
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
      ▼
  Reconcile
      │
      ▼
   Executor
      │
      ▼
 Hardware Devices
```

Device discovery is handled independently through the Registry subsystem.

```
Discovery
     │
     ▼
 Registry
     │
     ▼
  Engine
```

## Main Components

* **Engine** — central control loop and decision making.
* **Scheduler** — determines active and upcoming events.
* **Compiler** — converts events and configuration into executable jobs.
* **Registry** — maintains discovered devices and their operational state.
* **Executor** — performs hardware actions without making decisions.
* **Recovery** — restores system consistency after restart.

## Documentation

* `ARCHITECTURE.md` — overall system architecture
* `ENGINE.md` — Engine lifecycle
* `COMMAND_MODEL.md` — Job and Command model
* `REGISTRY.md` — device registry
* `ROADMAP.md` — project roadmap
* `DESIGN_PRINCIPLES.md` — architectural principles

## Current Status

The project is under active development.

The current implementation already includes:

* Engine lifecycle
* Scheduler
* Compiler
* Job model
* Device Registry
* Synchronization framework
* Recovery framework
* Unit tests for core domain logic

The next major milestone is the complete Job Lifecycle implementation.
