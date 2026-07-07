# DigiMitzves Roadmap

## Vision

DigiMitzves aims to become a fully autonomous appliance for Shabbat and Yom Tov automation.

The long-term objective is to provide a reliable embedded system that continues operating safely and predictably without Internet connectivity or user intervention.

---

# Development Stages

## Stage 1 — Core Foundation

Current stage.

Goals:

- Engine lifecycle
- Scheduler
- Compiler
- Job model
- Command model
- Persistent state
- Recovery pipeline

Status:

🟢 In Progress

---

## Stage 2 — Device Abstraction

Goals:

- Executor interface
- Device registry
- Hardware abstraction
- Transport-independent execution
- State synchronization

Status:

⚪ Planned

---

## Stage 3 — Zigbee Integration

Goals:

- Zigbee2MQTT adapter
- MQTT execution
- Device discovery
- Friendly name management
- Signal quality monitoring

Status:

⚪ Planned

---

## Stage 4 — Installation Workflow

Goals:

- Installation mode
- Device pairing
- Device naming
- Configuration validation
- Diagnostics

Status:

⚪ Planned

---

## Stage 5 — Web Configuration

Goals:

- Local web interface
- Configuration editor
- Schedule editor
- Device management
- Runtime monitoring

Status:

⚪ Planned

---

## Stage 6 — Embedded Appliance

Goals:

- Raspberry Pi deployment
- systemd services
- Automatic startup
- Automatic recovery
- Read-only operating mode where appropriate

Status:

⚪ Planned

---

## Stage 7 — Production Appliance

Goals:

- Installation image
- Complete documentation
- Stable release process
- Long-term maintenance
- Community feedback

Status:

⚪ Planned

---

# Engineering Priorities

The project prioritizes correctness over features.

New functionality is introduced only after the underlying architecture has been validated.

Preferred order:

1. Reliability
2. Recovery
3. Determinism
4. Simplicity
5. Performance
6. Convenience

---

# Non-Goals

The project is intentionally **not** designed to be:

- cloud-dependent;
- mobile-first;
- Internet-connected by default;
- dependent on third-party automation platforms;
- a generic home automation framework.

Its purpose is focused, predictable and reliable Shabbat and Yom Tov automation.

---

# Long-Term Goals

- Autonomous operation for years without maintenance.
- Complete recovery after unexpected power loss.
- Predictable behavior under all operating conditions.
- Hardware independence through well-defined interfaces.
- Professional embedded software architecture.
- A trustworthy appliance suitable for everyday use in observant Jewish homes.
