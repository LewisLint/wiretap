# wiretap

`wiretap` is a small Go command-line utility for inspecting the headers at the
start of a classic PCAP capture. It prints the global PCAP header and the
metadata for the first packet record in both hexadecimal and numeric form.

This is an exploratory parser, not a full packet analyzer: it does not decode
packet contents or iterate over subsequent records.

## Requirements

- Go (the module declares Go `1.27.2`)
- A classic `.pcap` file containing a global header and at least one packet
  record

## Usage

Run directly with a capture file:

```sh
go run . path/to/capture.pcap
```

Or build and run the executable:

```sh
go build -o wiretap .
./wiretap path/to/capture.pcap
```

For example, to run against the sample capture in this repository:

```sh
go run . test.pcap
```

The output includes the global header fields (magic number, version, time
zone offset, timestamp accuracy, snapshot length, and link-layer type), plus
the first record's timestamp, timestamp fraction, captured length, and original
length.

## Current limitations

- Only the first packet record header is displayed; packet data and later
  records are not decoded.
- Numeric fields are currently interpreted as little-endian. Other PCAP byte
  orders are not detected or handled.
- The input is expected to be a valid classic PCAP capture with a complete
  first record. PCAPNG files and malformed or truncated captures are not
  supported.
- A capture file path is required as the first argument.