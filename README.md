# Go Port Scanner

This project is a simple port scanner written in Go. It allows users to scan a specified range of ports on a target host to determine which ports are open.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Examples](#examples)
- [License](#license)

## Installation

To get started, clone the repository and navigate into the project directory:

```bash
git clone https://github.com/vp5h/go-port-scanner.git
cd go-port-scanner
```

Then, install the necessary dependencies:

```bash
go mod tidy
```

## Usage

You can run the port scanner by executing the following command:

```bash
go run cmd/main.go <target_host>
```

Replace `<target_host>` with the IP address or hostname of the target

## Examples

To scan ports 80 to 100 on the target host `192.168.1.1`, you would run:

```bash
go run cmd/main.go 192.168.1.1 80 100
```

## License

This project is licensed under the MIT License. See the LICENSE file for more details.
