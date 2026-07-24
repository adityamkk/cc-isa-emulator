# cc-isa-emulator
"Concurrency Computer" Custom ISA Emulator

(Ongoing) Implementation of a Multicore-focused ISA designed to introduce concurrency
concepts in a straightforward 16-bit architecture inspired by the LC3 (Little Computer 3) ISA.

This ISA aims to be educational. As concurrency becomes more and more critical, it is imperative
that the underlying mechanisms of concurrency can be explored in a straightforward way.

## Installation

Requires [Go](https://go.dev/dl/) 1.22 or higher. Check your version with `go version`.

1. Clone this repository:
   ```
   git clone https://github.com/adityamkk/cc-isa-emulator.git
   cd cc-isa-emulator
   ```
2. Build the emulator:
   ```
   make build
   ```
   The compiled binary will be located at `bin/cc-isa-emulator`.

Alternatively, run `go run . --bin=<filepath>` directly without a separate build step.

## Usage

Run the emulator against a compiled binary file:

```
bin/cc-isa-emulator --bin=<filepath> [--cores=N]
```

### Flags

| Flag       | Required | Default | Description                                                                 |
|------------|----------|---------|-------------------------------------------------------------------------------|
| `--bin`    | Yes      | -       | Path to the binary file to load into memory and execute.                    |
| `--cores`  | No       | `1`     | Number of cores to run concurrently. Must be between 1 and 16 (inclusive).   |

Example, running a binary across 4 cores:

```
bin/cc-isa-emulator --bin=./programs/example.bin --cores=4
```

## Makefile targets

| Target       | Description                                            |
|--------------|---------------------------------------------------------|
| `make build` | Compiles the emulator to `bin/cc-isa-emulator`.        |
| `make run`   | Runs the emulator via `go run` against `README.md` as a sample binary, using 15 cores. |
| `make test`  | Runs the Go test suite (`go test ./...`).              |
| `make vet`   | Runs `go vet ./...` for static analysis.               |
| `make clean` | Removes the `bin/` build output directory.             |

