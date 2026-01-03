# Go Select & Timeout Quest — Goroutine Timing

This quest focuses on **goroutine timing** and how the `select` statement reacts to **channel readiness**, not code order.

All goroutines start at the same time.
The **only reason** messages arrive in a specific order must be **different timeouts** inside each goroutine.

---

## Objective

Implement `FunctionOrdered` such that messages are received in this exact order:

1. `c1`
2. `c4`
3. `c2`
4. `c5`
5. `c3`

The receiving logic **must not** enforce this order manually.
The order must emerge **naturally** from timing.

---

## Core Rules

- Use **5 goroutines**
- Use **5 unbuffered channels**: `c1` → `c5`
- Every goroutine must:

  - Include a **timeout**
  - Send **exactly one message** to its channel

- Use **one `select` loop** to receive
- Do **not** use:

  - `time.Sleep`
  - `sync.WaitGroup`
  - buffered channels
  - conditional ordering logic

---

## Channel Message Contract (Strict)

Each channel must send **only** the following string:

- `c1` → `"from c1"`
- `c2` → `"from c2"`
- `c3` → `"from c3"`
- `c4` → `"from c4"`
- `c5` → `"from c5"`

No extra text. No formatting changes.

---

## Single Example (Pattern Only)

Each goroutine should follow this **pattern**:

```go
go func() {
    // wait using a timeout
    // send exactly one string to its channel
}()
```

The **timeout duration** determines **when** the message becomes available — and therefore **when `select` receives it**.

---

## What This Quest Tests

- Understanding that `select` reacts to **readiness**, not order
- Correct use of **timeouts inside goroutines**
- Ability to reason about **concurrent timing**
- Letting channels, not logic, control flow

---

## Success Condition

If implemented correctly, the output will **always** be:

```
from c1
from c4
from c2
from c5
from c3
```

If the order changes, the timing logic is incorrect.

---

## Running the Tests

Once implemented, run:

```bash
go test ./quests/17.select_timeout -v
```
