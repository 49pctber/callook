# Callook

A simple Go API for finding information about FCC callsigns from [callook.info](https://callook.info/).

Disclaimer: I am in no way related to this project. See their [API reference](https://callook.info/api_reference.php) for more information.

## Installation

```
go install github.com/ae6nr/callook/cmd/...@latest
```

## Usage

```
callook <callsign>
```

If the callsign is valid, the available information will be printed to your console.
