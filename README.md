# dzlog

`dzlog` is a lightweight and extensible logger for Go applications, designed for
microservices and backend systems that require structured logs, context propagation,
and correlation IDs.

It provides an Apache-style log format, configurable log levels, message splitting
for large payloads, and seamless `context.Context` integration.

---

## ✨ Features

- Log levels: `DEBUG`, `INFO`, `WARN`, `ERROR`
- Environment-based log level (`LOG_LEVEL`)
- Context-aware logging (`context.Context`)
- Correlation ID support
- Automatic caller info (file + line)
- Large message splitting
- Apache-style log format
- Zero configuration by default

---

## 📦 Installation

```bash
go get github.com/dmytrozilnyk/dzlog
```

---

## 🚀 Quick start

```go
import "github.com/dmytrozilnyk/dzlog"

log := dzlog.New(
    dzlog.LogModuleName("my-service"),
    dzlog.LogLevel(dzlog.InfoLevel),
)

log.Info("Service started")
log.Warning("Cache miss")
log.Error("Database error")
```

---

## ⚙️ Configuration

### Log level

You can configure the log level in two ways:

#### 1️⃣ Programmatically

```go
dzlog.LogLevel(dzlog.DebugLevel)
```

#### 2️⃣ Environment variable (overrides code)

```bash
export LOG_LEVEL=DEBUG
```

Supported values:
- `DEBUG`
- `INFO`
- `WARN`
- `ERROR`

---

### Module name

Used to identify the service or component emitting logs.

```go
dzlog.LogModuleName("users-api")
```

---

### Max message length

Messages longer than the configured length will be split automatically.

```go
dzlog.LogMaxLength(3000)
```

If the value is `<= 0`, the default (3000) is used.

---

## 🧠 Context & Correlation ID

`dzlog` supports `context.Context` to propagate metadata such as correlation IDs.

The correlation ID key used internally is:

```go
"correlationId"
```

### Example

```go
ctx := context.WithValue(context.Background(), "correlationId", "abc-123")

log.InfoCtx(ctx, "Processing request")
log.ErrorCtx(ctx, "Request failed")
```

---

## ⏱️ Time measurement (Debug only)

Useful for performance tracing.

```go
start := log.GetTimeNow()

// do work...

elapsed := log.CalculateDifference(start)
log.Debug("Elapsed time:", elapsed)
```

> Time tracking is only active when `DEBUG` level is enabled.

---

## 🔁 Returning context

Some methods return the context for chaining or propagation:

```go
ctx = log.InfoReturnCtx(ctx, "Step completed")
```

Available methods:
- `DebugReturnCtx`
- `InfoReturnCtx`
- `WarningReturnCtx`

---

## 📄 Log format

Example output:

```
[2026-02-04 15:04:05.123][users-api:INFO][pid:1234 tid:0][correlationId: abc-123][entries: 1/1][handler.go:42] User created
```

Each log entry includes:
- Timestamp
- Module name
- Log level
- Process ID (PID)
- Thread ID (TID)
- Correlation ID (if present)
- Source file and line number
- Message content

---

## 📄 License

MIT License