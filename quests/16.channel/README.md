# Go Channels Quest — Explicit Read & Write Responsibilities

This quest validates your understanding of **how channels are used by behavior**, not by type-level direction annotations.

You will work with **plain `chan string` values**, but each function must strictly follow a **single, well-defined responsibility**.

---

## What This Quest Is Testing

You are being tested on your ability to:

- Correctly **receive** from channels when required
- Correctly **send** to channels when required
- Avoid accidental reads or writes outside the defined role
- Understand how **channel blocking** coordinates execution naturally

There is no trick, no race condition, and no requirement to use goroutines inside these functions.

Correctness is purely about **doing exactly what the function contract states**.

---

## Functions to Implement

You must implement **three independent functions**.

Each function receives channels as parameters, but **how those channels are used is fixed**.

---

## 1. `ReadFromBoth` — Receive Only

```go
func ReadFromBoth(ch1 chan string, ch2 chan string) string
```

### Required Behavior

- Receive **value** from `ch1`
- Receive **value** from `ch2`
- Do **not** send to either channel
- Return a formatted string

### Return Format

```text
read: <value_from_ch1> & <value_from_ch2>
```

### Example

If:

- `ch1` sends `"message from ch1"`
- `ch2` sends `"message from ch2"`

The function must return:

```text
read: message from ch1 & message from ch2
```

---

## 2. `WriteToBoth` — Send Only

```go
func WriteToBoth(ch1 chan string, ch2 chan string, msg string)
```

### Required Behavior

- Send **value** to `ch1`
- Send **value** to `ch2`
- Do **not** receive from either channel

### Value to Send

```text
write: <msg>
```

### Example

If `msg` is `"hello"`, both channels must receive:

```text
write: hello
```

This function **does not return anything**.

---

## 3. `ReadThenWrite` — Receive First, Then Send

```go
func ReadThenWrite(input chan string, output chan string)
```

### Required Behavior

- Receive **value** from `input`
- Send **value** to `output`
- Do **not** send to `input`
- Do **not** receive from `output`

### Value to Send

```text
transform: <received_value>
```

### Example

If `input` sends:

```text
msg
```

Then `output` must receive:

```text
transform: msg
```

---

## Mental Model (Important)

Think in terms of **roles**, not helpers:

- `WriteToBoth` → **producer**
- `ReadThenWrite` → **transformer**
- `ReadFromBoth` → **final consumer**

Each function owns its responsibility **fully and exclusively**.

---

## Why This Matters

In real Go codebases:

- Channels are often shared without directional typing
- Correctness depends on **discipline**, not compiler enforcement
- One incorrect send or receive can cause deadlocks or subtle bugs

This quest ensures you can:

- Reason about channel ownership
- Respect read/write boundaries
- Build correct channel-based flows confidently

---

## Running the Tests

Once implemented, run:

```bash
go test ./quests/16.channel -v
```
