[![Build Status](https://travis-ci.com/Agaetis-IT/ModBusbeat.svg?branch=master)](https://travis-ci.com/Agaetis-IT/ModBusbeat)
# Modbusbeat 

Modbusbeat is a lightweight agent that formats and ships data from ModBus devices.


## Getting Started
### Requirements

* [Golang](https://golang.org/dl/) 1.7


### Clone

To clone Modbusbeat from the git repository, run the following commands:

```
mkdir -p ${GOPATH}/src/modbusbeat
git clone https://github.com/Agaetis-IT/ModBusbeat.git ${GOPATH}/src/modbusbeat
```

### Init Project
To get running with Modbusbeat and also install the
dependencies, run the following command:

```
make setup
```

It will create a clean git history for each major step. Note that you can always rewrite the history if you wish before pushing your changes.

### Build

To build the binary for Modbusbeat run the command below. This will generate a binary
in the same directory with the name modbusbeat.

```
make
```


### Run

To run Modbusbeat with debugging output enabled, run:

```
./modbusbeat -c modbusbeat.yml -e -d "*"
```


### Test

To test Modbusbeat, run the following command:

```
make testsuite
```

alternatively:
```
make unit-tests
make system-tests
make integration-tests
make coverage-report
```

The test coverage is reported in the folder `./build/coverage/`

### Update

Each beat has a template for the mapping in elasticsearch and a documentation for the fields
which is automatically generated based on `fields.yml` by running the following command.

```
make update
```


### Cleanup

To clean  Modbusbeat source code, run the following commands:

```
make fmt
make simplify
```

To clean up the build directory and generated artifacts, run:

```
make clean
```



For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).


### How to configure it

> Default configuration file
```yaml
modbusbeat:
  period: 5s
  devices:
    - address: 127.0.0.1
      registers:
        - type: "Holding" # [Holding|Input|Coil|Discret]
          addresses:
            - 101 # Or 0x65
```

## Packaging

The beat frameworks provides tools to crosscompile and package your beat for different platforms. This requires [docker](https://www.docker.com/) and vendoring as described above. To build packages of your beat, run the following command:

```
make release
```

This will fetch and create all images required for the build process. The whole process to finish can take several minutes.
