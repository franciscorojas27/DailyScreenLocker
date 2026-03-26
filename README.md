# DailyScreenLocker

> A compact Windows utility that automatically locks the workstation at a configured time each day.

## Overview

`DailyScreenLocker` is a small, focused tool written in Go that locks a Windows session at a scheduled hour every day. It is intended to help enforce breaks, preserve device security, or automate end-of-day locking.

## Features

- Simple, single-binary Go program
- Schedules a daily lock at the configured hour
- Minimal dependencies (Windows API via golang.org/x/sys/windows)
- Clear console status and next-run countdown

## Requirements

- Windows (desktop) — uses `user32.dll` LockWorkStation API
- Go 1.18+ (to build)

## Build

Clone the repository and build with Go:

```bash
go build -o dailyscreenlocker ./
```

## Usage

Run the binary on a Windows machine. The program schedules a daily lock for 17:00 (5pm) by default. To change the scheduled time, edit the `main.go` source and adjust the `goal` time construction.

Example run:

```powershell
.\dailyscreenlocker.exe
```

The program prints a live countdown to the next scheduled lock and will call the Windows `LockWorkStation` API at the scheduled time.

## Configuration

Currently, the schedule is defined in `main.go` as a hard-coded time. If you want runtime configuration, consider one of these quick improvements:

- Read the target hour/minute from environment variables
- Accept a CLI flag (e.g., `--hour=17 --minute=0`)
- Add a JSON/TOML config file

## Contributing

Contributions are welcome. Keep changes focused and documented. Please open issues for feature requests or bug reports.

## License

This project is provided under the Non-Commercial Source-Available License (see `LICENSE`). You may use, modify and distribute the code for non-commercial purposes, with attribution.

## Contact

Open an issue or create a pull request in this repository for questions or improvements.
