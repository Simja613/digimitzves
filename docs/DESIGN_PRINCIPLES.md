# DESIGN PRINCIPLES

## Purpose

This document defines the architectural principles of DigiMitzves.

These principles are intentionally more stable than the implementation itself.

Code may evolve.

Architecture may evolve.

These principles should change only with great care.

---

# 1. Reliability Before Features

The primary goal of DigiMitzves is reliability.

New functionality must never reduce the predictability or recoverability of the system.

---

# 2. Deterministic Behavior

The same inputs must always produce the same decisions.

The Engine should never depend on hidden state, timing side effects, or execution order.

---

# 3. Observe Before Acting

Every Engine cycle begins by observing the current state of the world.

No decision should be based solely on assumptions from a previous cycle.

---

# 4. Recovery Before Execution

After startup or failure, the system must first restore consistency.

Only then may it execute new actions.

---

# 5. Planning Is Separate From Execution

Planning answers:

> What should happen?

Execution answers:

> Perform it.

Planning never communicates with hardware.

Execution never makes decisions.

---

# 6. One Responsibility Per Component

Each subsystem has one primary responsibility.

Examples:

* Scheduler determines time.
* Compiler builds Jobs.
* Registry tracks devices.
* Engine makes decisions.
* Executor performs actions.

Responsibilities should not overlap.

---

# 7. Hardware Independence

Business logic must never depend on a specific hardware implementation.

Replacing MQTT, Zigbee, GPIO, or RS-485 must not require changes to Engine logic.

---

# 8. State Before Events

The system is driven by state rather than instantaneous events.

The Engine continuously verifies that the current world matches the expected world.

---

# 9. Simplicity Over Generality

DigiMitzves is a dedicated appliance.

The project should not introduce abstractions for hypothetical future requirements.

Every new abstraction must solve an existing problem.

---

# 10. Add Complexity Only When Necessary

New layers, interfaces, or components should appear only when they simplify an existing design.

Avoid designing for imagined future use cases.

---

# 11. Explicit Is Better Than Implicit

State transitions should be visible.

Responsibilities should be obvious.

Hidden behavior should be avoided.

---

# 12. Fail Safely

Unexpected situations should never leave the system in an undefined state.

Whenever possible:

* preserve consistency
* preserve recoverability
* preserve observability

---

# 13. Test Behavior, Not Implementation

Tests should describe system behavior.

They should remain valid even if the internal implementation changes.

---

# 14. Documentation Is Part of the Architecture

Architecture exists not only in code but also in documentation.

Every significant architectural decision should be reflected in the project documentation.

---

# 15. Keep the Core Small

The Engine should remain focused on decision making.

Features that do not belong in the core should remain outside the core.

A small core is easier to understand, test, and maintain.

---

# Final Principle

When in doubt, prefer the solution that makes the system easier to understand five years from now.
