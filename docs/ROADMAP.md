# ROADMAP

## Vision

DigiMitzves is being developed as a reliable appliance dedicated to Shabbat and Yom Tov automation.

The roadmap follows architectural maturity rather than feature count.

Each phase should leave the system stable, testable and recoverable.

---

# Phase 1 — Core Foundation

Status: **Completed**

Objectives:

* Project structure
* Engine lifecycle
* Scheduler
* Compiler
* Job model
* Registry
* Synchronization framework
* Recovery framework
* Initial unit tests

Result:

A deterministic application core capable of observing the system and making consistent decisions.

---

# Phase 2 — Job Lifecycle

Status: **Current**

Objectives:

* Job validation
* Job generation
* Job replacement
* Job persistence
* Job execution lifecycle
* Command lifecycle
* Reconciliation improvements

Result:

The Engine manages a complete execution plan instead of isolated commands.

---

# Phase 3 — Device Integration

Objectives:

* Discovery improvements
* Missing device detection
* Ignored devices
* Configuration workflow
* Registry persistence
* Device recovery

Result:

Reliable long-term device management independent of transport technology.

---

# Phase 4 — Executor

Objectives:

* MQTT adapter
* Zigbee2MQTT integration
* Hardware abstraction
* Command acknowledgements
* Execution reliability

Result:

Business logic becomes completely independent from physical hardware.

---

# Phase 5 — Web Interface

Objectives:

* Device configuration
* Schedule configuration
* Registry inspection
* Job inspection
* System status
* Recovery interface

Result:

Complete appliance management without exposing implementation details.

---

# Phase 6 — Installation Mode

Objectives:

* Automatic discovery
* Friendly name assignment
* Initial configuration
* Ignore workflow
* Device replacement
* Configuration validation

Result:

A guided installation experience suitable for non-technical users.

---

# Phase 7 — Appliance

Objectives:

* Raspberry Pi image
* systemd services
* automatic startup
* persistent storage
* logging
* update strategy

Result:

A standalone appliance requiring minimal administration.

---

# Phase 8 — Production

Objectives:

* Long-term stability
* Extensive testing
* Documentation completion
* Deployment automation
* Release process

Result:

A production-ready appliance for everyday operation.

---

# Long-Term Direction

The long-term goal is not to become a general home automation platform.

The goal is to become a dependable appliance that performs one specific task exceptionally well.

Future features will be accepted only if they improve:

* reliability
* recoverability
* simplicity
* maintainability

Features that increase complexity without improving these qualities should be rejected.

---

# Guiding Principle

Every completed phase should improve the architecture without requiring a major redesign of previous phases.
