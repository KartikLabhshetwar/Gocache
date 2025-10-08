
<img width="1024" height="1024" alt="Gemini_Generated_Image_4jdlh44jdlh44jdl" src="https://github.com/user-attachments/assets/3e35a35f-98bf-41be-8710-22acab2b481e" />

# gocache

gocache is a small, educational Redis-like in-memory cache server implemented in Go. It implements a subset of the Redis protocol (RESP) and supports a few basic commands (PING, SET, GET, HSET, HGET, HGETALL). The project also includes an append-only file (AOF) persistence helper and a RESP parser/writer.

This project is intended as a learning implementation to demonstrate how a simple RESP server, command handlers, and persistent AOF logging can be structured in Go.

## Features

- RESP (REdis Serialization Protocol) parser and writer
- Basic commands: `PING`, `SET`, `GET`, `HSET`, `HGET`, `HGETALL`
- In-memory stores for simple keys and hash maps
- Append-only file (AOF) writer and reader for simple persistence
- Concurrency-safe access to the in-memory stores using RWMutex

## Components

- `resp.go` - RESP parser and writer (reading and writing RESP values, marshalling/unmarshalling). It provides `Value`, `Resp`, and `Writer` types.
- `handler.go` - Command handlers for supported commands, and in-memory maps with mutexes.
- `aof.go` - Append-only file helper that writes marshalled RESP entries to disk and can replay them.
- `main.go` - (entry point) expected to wire up the TCP listener, accept connections and use RESP, handlers and AOF. (See code for exact behavior.)

## Architecture


<img width="1955" height="612" alt="Mermaid Chart - Create complex, visual diagrams with text -2025-10-08-100226" src="https://github.com/user-attachments/assets/2bb90ef0-e024-4c1f-a4e9-14dc09d923f1" />



## RESP basics used by gocache

gocache implements a subset of RESP and models values with a `Value` type. The important RESP types used are:

- Bulk strings (prefixed by `$`) — used for keys and values
- Arrays (prefixed by `*`) — used for commands and multi-part responses

The repository includes helper methods to Marshal and Read these types from streams.

Example of the RAW RESP for the command `SET mykey myvalue`:

```text
*3
$3
SET
$5
mykey
$7
myvalue
```

## How to build and run

Requirements: Go 1.20+ (any modern Go should work).

Build the binary:

```bash
go build -o gocache
```

Run it (if `main.go` listens on a port, run the binary; otherwise use `go run` during development):

```bash
go run .
# or
./gocache
```

Check `main.go` for the port and listen logic used by this repository. If `main.go` binds a port (for example :6379), you can connect with `redis-cli -p 6379` or with `nc`/`telnet` and speak RESP manually.

## Quick examples 


https://github.com/user-attachments/assets/d46c0bc7-e48d-43c8-9030-e7f91499db4c



## Persistence (AOF)

The `aof.go` file provides an `Aof` struct which opens a file, writes marshalled `Value` entries, and periodically syncs the file to disk. The `Aof` writer uses a mutex to make writes safe across goroutines. The `Read` method allows replaying the file by invoking a callback for each recorded `Value`.

This AOF implementation is intentionally simple:

- It writes the RESP-marshalled bytes directly.
- It periodically (every second) calls `file.Sync()` from a background goroutine.
- It uses a single mutex for file operations.

On startup you can replay the AOF to restore the state by reading each recorded command and re-executing it against the in-memory handlers.

## Concurrency and safety

- The in-memory maps are protected with `sync.RWMutex` for safe concurrent reads and writes.
- The AOF uses a `sync.Mutex` for serializing file writes and Sync calls.

## Limitations

- Protocol: The RESP implementation is minimal (supports arrays and bulk strings primarily). Some RESP types and error handling are simplified.
- Commands: Only a small subset of Redis commands are implemented.
- Persistence: The AOF is naive and appends entire RESP frames. There is no rewriting/compaction.
- Robustness: Error handling and malformed-resp handling can be improved.
- Tests: Add unit tests for RESP parsing, handlers, and AOF replay.

## Contributing

Feel free to open issues or PRs. This project is intended for learning and exploration: improvements, refactors and test coverage are welcome.

## License

This repository includes a [LICENSE](#LICENSE) file — respect its terms when using or contributing.
