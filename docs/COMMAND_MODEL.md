# COMMAND MODEL

## Purpose

The Command Model describes how DigiMitzves transforms an event into executable actions.

The model separates planning from execution.

Planning determines **what should happen**.

Execution determines **when it actually happens**.

---

# Overview

```text
Configuration

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

   Commands

      │

      ▼

   Executor

      │

      ▼

   Hardware
```

---

# Event

An Event represents a time interval.

Examples:

* Shabbat
* Yom Tov

Events contain:

* start time
* end time

Events never contain device actions.

---

# Configuration

Configuration represents user intent.

Examples:

* enabled intervals
* preferred states
* schedules
* device settings

Configuration changes rarely.

---

# Compiler

The Compiler transforms:

* Event
* Configuration
* Current State

into

one immutable Job.

Compiler never executes commands.

---

# Job

A Job represents the complete execution plan for one Event.

```text
Job

├── Event

├── Created

├── Status

└── Commands[]
```

A Job is immutable.

If assumptions change, a new Job is generated.

---

# Command

A Command represents exactly one planned action.

Example:

```text
18:42

Kitchen Relay

ON
```

Each Command contains:

* execution time
* target device
* desired action
* execution status

A Command never contains business logic.

---

# Command Status

Typical lifecycle:

```text
Planned

↓

Running

↓

Completed
```

Possible future states:

```text
Failed

Skipped

Cancelled
```

---

# Job Lifecycle

```text
No Job

↓

Compile

↓

Planned

↓

Running

↓

Completed

↓

Archived
```

Whenever assumptions change:

```text
Configuration Changed

or

Event Changed

↓

Recompile

↓

New Job
```

The previous Job becomes obsolete.

---

# Reconciliation

The Engine never assumes that a Job remains valid.

Every cycle asks:

```text
Does the current Job still represent reality?
```

Possible outcomes:

* Job is valid
* Job must be regenerated
* Job is completed
* No Job exists

Only then does execution continue.

---

# Execution

Execution follows the Job exactly.

The Executor does not make decisions.

It simply performs Commands.

```text
Job

↓

Command

↓

Executor

↓

Hardware
```

---

# Separation of Responsibilities

| Component | Responsibility            |
| --------- | ------------------------- |
| Scheduler | Determines time           |
| Compiler  | Builds the execution plan |
| Job       | Stores the execution plan |
| Engine    | Decides what to do next   |
| Executor  | Executes commands         |
| Hardware  | Performs physical actions |

No component may perform another component's responsibility.

---

# Design Principles

A Job represents intent.

A Command represents one action.

Execution never modifies planning.

Planning never communicates with hardware.

Every physical action must originate from a Command.

Every Command must belong to exactly one Job.

Business decisions are made before execution begins.

Execution must remain deterministic.
