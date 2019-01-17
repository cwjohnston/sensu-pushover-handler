# Sensu Go Pushover Handler Plugin
TravisCI: [![TravisCI Build Status](https://travis-ci.org/cwjohnston/sensu-pushover-handler.svg?branch=master)](https://travis-ci.org/cwjohnston/sensu-pushover-handler)

A Sensu handler plugin for sending notifications via [Pushover]().

## Installation

Download the latest version of the sensu-pushover-handler from [releases][2],
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
        "type": "pipe",
        "command": "sensu-pushover-handler",
        "timeout": 10,
        "env_vars": [
            "PUSHOVER_APP_TOKEN=your_app_token_here",
            "PUSHOVER_USER_KEY=your_user_key_here"
        ],
        "filters": [
            "is_incident",
            "not_silenced"
        ],
        "runtime_assets": [ "sensu-pushover-handler" ]
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
  -a, --app.token string   Pushover v1 API app token, use default from PUSHOVER_APP_TOKEN env var
  -h, --help               help for sensu-pushover-handler
  -u, --user.key string    Pushover v1 API user key, use default from PUSHOVER_USER_KEY env var
```

## Contributing

See https://github.com/sensu/sensu-go/blob/master/CONTRIBUTING.md

[1]: https://pushover.net
[2]: https://github.com/cwjohnston/sensu-pushover-handlers/releases
