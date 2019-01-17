# Sensu Go Pushover Handler Plugin
TravisCI: [![TravisCI Build Status](https://travis-ci.org/cwjohnston/sensu-pushover-handler.svg?branch=master)](https://travis-ci.org/cwjohnston/sensu-pushover-handler)

A Sensu handler plugin for sending notifications via [Pushover]().

## Installation

Download the latest version of the sensu-pushover-handler from [releases][1],
or create an executable script from this source.

From the local path of the sensu-pushover-handler repository:

```
go build -o /usr/local/bin/sensu-pushover-handler main.go
```

## Configuration

Example Sensu Go definition:

```json
{
    "api_version": "core/v2",
    "type": "Handler",
    "metadata": {
        "namespace": "default",
        "name": "pushover"
    },
    "spec": {
        "...": "..."
    }
}
```

## Usage Examples

Help:

```
The Sensu Go handler for Pushover 

Usage:
  sensu-pushover-handler [flags]

Flags:
  -f, --foo string   example
  -h, --help         help for sensu-pushover-handler
```

## Contributing

See https://github.com/sensu/sensu-go/blob/master/CONTRIBUTING.md

[1]: https://github.com/cwjohnston/sensu-pushover-handlers/releases
