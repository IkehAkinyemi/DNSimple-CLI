# Zone Records CLI
This is a command line interface (CLI) for managing zone records via the DNSimple API. It allows a developer to list and create zone records, with the option to filter the list of records by name and type and the ability to create SRV records. Validation errors are also handled.

## Prerequisites
- Go version 1.13 or higher
- DNSimple access token
- Zone ID
- Account ID

## Installation
- Clone the repository: `git clone https://github.com/dnsimple-hiring/dnsimple-seint-code-test-IkehAkinyemi.git zone-records-cli`
- Navigate to the project directory: `cd zone-records-cli`
- Install dependencies: `go mod tidy`
- Build the project: `go build -o=./zone-records-cli ./cmd`

## Usage
The CLI has the following commands:
- `list`: lists all zone records, with the option to filter by `name`, `type` and/or `name_like` flags
- `create`: creates a new zone record, with the option to specify the `name`, `type`, and `content` of the record and the `priority`, and `TTL` for SRV records

To use the CLI, you need to provide your DNSimple access token, environment (dedicated the use of `dev` for testing purposes), the name of the zone you want to manage and your account ID. These can be updated using the `config.yaml` file in the `./config` directory.


```
Usage:
  zone-records-cli [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  create      create a new zone record
  help        Help about any command
  list        List all zone records

Flags:
      --cfg-dir string   Directory path to configuration file
  -h, --help             help for zone-records-cli
```

### List command
To list all zone records, use the `list` command. You can specify and load the configuration file using the `--cfg-dir` flag.

```
Usage:
  zone-records-cli list [flags]

Flags:
  -h, --help               help for list
      --name string        filter records by name
      --name_like string   filter records by name like a given string
      --type string        filter records by type
      --ttl int            Time to Live (TTL)


Global Flags:
      --cfg-dir string   Directory path to configuration file
```

#### Example
List all zone records:
```
./zone-records-cli --cfg-dir './config/config.yaml' list 
```

Filter zone records by name:
```
./zone-records-cli --cfg-dir './config/config.yaml' list --name test --type AAA
```

### Create command
To create a new zone record, use the `create` command. You can specify and load the configuration file using the `--cfg-dir` flag.

```
Usage:
  zone-records-cli create [flags]

Flags:
      --content string   content of the record (required)
  -h, --help             help for create
      --name string      Name of the resource (required)
      --priority int     priority of the record (required for SRV records)
      --ttl int          Time to Live (TTL)
      --type string      type of the record (A, AAAA, or SRV) (required)

Global Flags:
      --cfg-dir string   Directory path to configuration file
```

#### Example
Create a new record:
```
./zone-records-cli --cfg-dir './config/config.yaml' create --name test --type A --content 'xxx.example.com'
```

Create a new SRV record:
```
./zone-records-cli --cfg-dir './config/config.yaml' create --name test --type SRV --content 'tcp.test.com' --ttl 0 --priority 0
```

## Testing
To run the unit tests, use the `go test` command:
```
go test ./... -cfg-dir '../../config/config.yaml'
```

Note: Please change `env` field in the YAML file to `dev` and load the configuration file using `--cfg-dir` global flag before running above test command.

## API Documentation
The API documentation for the DNSimple Zone Records endpoints can be found here: https://developer.dnsimple.com/v2/zones/records/.

## Sandbox Environment
It is recommended to use the DNSimple sandbox environment for testing. To use the sandbox environment, set the `env` field in the YAML file to `dev` to switch API endpoint to https://api.sandbox.dnsimple.com.

