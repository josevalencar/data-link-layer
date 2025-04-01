# Data link layer implementation in Go

This project implements a basic data link layer protocol in Go, demonstrating the principles of framing, synchronization, and error detection/correction.

## Overview

The implementation consists of two main components:

1. **Sender**: Encodes data using Hamming code for error correction, applies bit stuffing to prevent frame synchronization issues, and adds header/terminator for framing.

2. **Receiver**: Detects frame boundaries, removes bit stuffing, checks for and corrects errors using Hamming code, and returns the original data.

## Features

- Frame structure with header and terminator (01111110)
- Bit stuffing to prevent header/terminator patterns in the payload
- Hamming code implementation for single-bit error correction
- Detection of missing headers or terminators

## Protocol documentation

The frame structure is as follows:

```
+------------------+--------------------+------------------+
|     Header       |      Payload       |    Terminator    |
|    (8 bits)      |  (variable + bits  |     (8 bits)     |
|                  |  of Hamming)       |                  |
+------------------+--------------------+------------------+
```

- **Header**: 01111110 (Flag for start of frame)
- **Payload**: Original data encoded with Hamming code and bit stuffing
- **Terminator**: 01111110 (Flag for end of frame)

### Error correction

The implementation uses Hamming code, which allows:
- Detection of up to 2-bit errors
- Correction of single-bit errors

## Limitations

- Cannot correct multi-bit errors
- No flow control mechanisms
- No addressing (assumes point-to-point communication)

## Building and running

### Prerequisites

- Go 1.13 or higher

### Compilation

```bash
go build -o sender sender.go
go build -o receiver receiver.go
```

### Usage

Basic usage:
```bash
./sender "01101001" | ./receiver
```

Simulate a single-bit error:
```bash
./sender "01101001" | sed 's/\(.\{15\}\)./\10/' | ./receiver
```