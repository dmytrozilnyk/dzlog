# dzlog

`dzlog` is a lightweight and extensible logger for Go applications, designed for
backend services and microservices that need **structured logs**, **context propagation**
and **correlation (trace) IDs**, without pulling heavy logging frameworks.

It provides an Apache-style log format, configurable log levels, message splitting
for large payloads, and seamless `context.Context` integration.

---

## ✨ Features

- Log levels: `DEBUG`, `INFO`, `WARN`, `ERROR`
- Environment-based log level via `LOG_LEVEL`
- Context-aware logging (`context.Context`)
- Trace / correlation ID support (`tracerId`)
- Automatic caller info (file + line)
- Large message splitting
- Apache-style log format
- Zero configuration by default
- No reflection, no unsafe

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
	dzlog.LogLevel("INFO"),
	dzlog.LogEnvironment("dev"),
)

log.Info("Service started")
log.Warning("Cache miss")
log.Error("Database error")
```

---

## ⚙️ Configuration

`dzlog` uses functional options for configuration.

### Log level

You can configure the log level in two ways:

#### 1️⃣ Programmatically

```go
dzlog.LogLevel("DEBUG")
```

Supported values:
- `DEBUG`
- `INFO`
- `WARN`
- `ERROR`

If an invalid value is provided, it falls back to `INFO`.

#### 2️⃣ Environment variable (overrides code)

```bash
export LOG_LEVEL=DEBUG
```

If `LOG_LEVEL` is set, it **always overrides** the configured value.

---

### Module name

Identifies the service or component emitting logs.

```go
dzlog.LogModuleName("users-api")
```

---

### Environment name

Useful to distinguish logs between environments (`dev`, `staging`, `prod`, etc.).

```go
dzlog.LogEnvironment("prod")
```

---

### Max message length

Messages longer than the configured length are automatically split into multiple
log entries.

```go
dzlog.LogMaxLength(3000)
```

If the value is `<= 0`, the default (`3000`) is used.

Each split entry includes an `entries: x/y` indicator.

---

## 🧠 Context & Trace ID

`dzlog` supports `context.Context` to propagate metadata such as trace / correlation IDs.

The context key used internally is:

```go
"tracerId"
```

### Example

```go
ctx := context.WithValue(context.Background(), "tracerId", "abc-123")

log.InfoCtx(ctx, "Processing request")
log.ErrorCtx(ctx, "Request failed")
```

If the context contains a `tracerId`, it will automatically appear in the logs.

---

## ⏱️ Time measurement (DEBUG only)

Useful for performance tracing. Time measurement is only active when the log level
is `DEBUG`.

```go
start := log.GetTimeNow()

// do work...

elapsed := log.CalculateDifference(start)
log.Debug("Elapsed time:", elapsed)
```

If the log level is not `DEBUG`, `GetTimeNow()` returns a zero `time.Time` and
`CalculateDifference` returns `0`.

---

## 🔁 Returning context

Some methods return the same context to allow chaining or propagation.

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
[2026-02-04 15:04:05.123][users-api:INFO][env:dev][pid:1234 tid:0][tracerId: abc-123][entries: 1/1][handler.go:42] User created
```

Each log entry includes:

- Timestamp
- Module name
- Log level
- Environment
- Process ID (PID)
- Thread ID (TID)
- Trace ID (`tracerId`) if present
- Entry index when message splitting is used
- Source file and line number
- Message content

---

## 📄 License

MIT License
