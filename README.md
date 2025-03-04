# Mainflux

[![build][ci-badge]][ci-url]
[![go report card][grc-badge]][grc-url]
[![coverage][cov-badge]][cov-url]
[![license][license]](LICENSE)
[![chat][gitter-badge]][gitter]

![banner][banner]

Mainflux is modern, scalable, secure, open-source, and patent-free IoT cloud platform written in Go.

It accepts user and thing (sensor, actuator, application) connections over various network protocols (i.e. HTTP,
MQTT, WebSocket, CoAP), thus making a seamless bridge between them. It is used as the IoT middleware
for building complex IoT solutions.

For more details, check out the [official documentation][docs].

## Features

- Multi-protocol connectivity and bridging (HTTP, MQTT, WebSocket and CoAP)
- Device management and provisioning (Zero Touch provisioning)
- Mutual TLS Authentication (mTLS) using X.509 Certificates
- Fine-grained access control (policies, ABAC/RBAC)
- Message persistence (Cassandra, InfluxDB, MongoDB and PostgresSQL)
- Platform logging and instrumentation support (Grafana, Prometheus and OpenTracing)
- Event sourcing
- Container-based deployment using [Docker][docker] and [Kubernetes][kubernetes]
- [LoRaWAN][lora] network integration
- [OPC UA](opcua) integration
- Edge [Agent](agent) and [Export](export) services for remote IoT gateway management and edge computing
- SDK
- CLI
- Small memory footprint and fast execution
- Domain-driven design architecture, high-quality code and test coverage

## Prerequisites

The following are needed to run Mainflux:

- [Docker](https://docs.docker.com/install/) (version 20.10)
- [Docker compose](https://docs.docker.com/compose/install/) (version 1.29)

Developing Mainflux will also require:

- [Go](https://golang.org/doc/install) (version 1.19.2)
- [Protobuf](https://github.com/protocolbuffers/protobuf#protocol-compiler-installation) (version 3.6.1)

## Install

Once the prerequisites are installed, execute the following commands from the project's root:

```bash
docker-compose -f docker/docker-compose.yml up
```

This will bring up the Mainflux docker services and interconnect them. This command can also be executed using the project's included Makefile:

```bash
make run
```

If you want to run services from specific release checkout code from github and make sure that
`MF_RELEASE_TAG` in [.env](.env) is being set to match the release version

```bash
git checkout tags/<release_number> -b <release_number>
# e.g. `git checkout tags/0.13.0 -b 0.13.0`
```

Check that `.env` file contains:

```bash
MF_RELEASE_TAG=<release_number>
```

>`docker-compose` should be used for development and testing deployments. For production we suggest using [Kubernetes](https://docs.mainflux.io/kubernetes).

## Usage

The quickest way to start using Mainflux is via the CLI. The latest version can be downloaded from the [official releases page][rel].

It can also be built and used from the project's root directory:

```bash
make cli
./build/mainflux-cli version
```

Additional details on using the CLI can be found in the [CLI documentation](https://docs.mainflux.io/cli).

## Documentation

Official documentation is hosted at [Mainflux official docs page][docs]. Documentation is auto-generated, checkout the instructions on [official docs repository](https://github.com/mainflux/docs):

If you spot an error or a need for corrections, please let us know - or even better: send us a PR.

## Authors

Main architect and BDFL of Mainflux project is [@drasko][drasko].

Additionally, [@nmarcetic][nikola] and [@janko-isidorovic][janko] assured
overall architecture and design, while [@manuio][manu] and [@darkodraskovic][darko]
helped with crafting initial implementation and continuously worked on the project evolutions.

Besides them, Mainflux is constantly improved and actively
developed by [@anovakovic01][alex], [@dusanb94][dusan], [@srados][sava],
[@gsaleh][george], [@blokovi][iva], [@chombium][kole], [@mteodor][mirko] and a large set of contributors.

Maintainers are listed in [MAINTAINERS](MAINTAINERS) file.

The Mainflux team would like to give special thanks to [@mijicd][dejan] for his monumental work
on designing and implementing a highly improved and optimized version of the platform,
and [@malidukica][dusanm] for his effort on implementing the initial user interface.

## Professional Support

There are many companies offering professional support for the Mainflux system.

If you need this kind of support, best is to reach out to [@drasko][drasko] directly, and he will point you out to the best-matching support team.

## Contributing

Thank you for your interest in Mainflux and the desire to contribute!

1. Take a look at our [open issues](https://github.com/mainflux/mainflux/issues). The [good-first-issue](https://github.com/mainflux/mainflux/labels/good-first-issue) label is specifically for issues that are great for getting started.
2. Checkout the [contribution guide](CONTRIBUTING.md) to learn more about our style and conventions.
3. Make your changes compatible to our workflow.

### We're Hiring

You like Mainflux and you would like to make it your day job? We're always looking for talented engineers interested in open-source, IoT and distributed systems. If you recognize yourself, reach out to [@drasko][drasko] - he will contact you back.

>The best way to grab our attention is, of course, by sending PRs :sunglasses:.

## Community

- [Google group][forum]
- [Gitter][gitter]
- [Twitter][twitter]

## License

[Apache-2.0](LICENSE)

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fmainflux%2Fmainflux.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fmainflux%2Fmainflux?ref=badge_large)

[banner]: https://github.com/mainflux/docs/blob/master/docs/img/gopherBanner.jpg
[ci-badge]: https://semaphoreci.com/api/v1/mainflux/mainflux/branches/master/badge.svg
[ci-url]: https://semaphoreci.com/mainflux/mainflux
[docs]: https://docs.mainflux.io
[docker]: https://www.docker.com
[forum]: https://groups.google.com/forum/#!forum/mainflux
[gitter]: https://gitter.im/mainflux/mainflux?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge
[gitter-badge]: https://badges.gitter.im/Join%20Chat.svg
[grc-badge]: https://goreportcard.com/badge/github.com/mainflux/mainflux
[grc-url]: https://goreportcard.com/report/github.com/mainflux/mainflux
[cov-badge]: https://codecov.io/gh/mainflux/mainflux/branch/master/graph/badge.svg
[cov-url]: https://codecov.io/gh/mainflux/mainflux
[license]: https://img.shields.io/badge/license-Apache%20v2.0-blue.svg
[twitter]: https://twitter.com/mainflux
[lora]: https://lora-alliance.org/
[opcua]: https://opcfoundation.org/about/opc-technologies/opc-ua/
[agent]: https://github.com/mainflux/agent
[export]: https://github.com/mainflux/export
[kubernetes]: https://kubernetes.io/
[rel]: https://github.com/mainflux/mainflux/releases
[careers]: https://www.mainflux.com/careers.html
[lf]: https://www.linuxfoundation.org/
[edgex]: https://www.edgexfoundry.org/
[company]: https://www.mainflux.com/
[blog]: https://medium.com/mainflux-iot-platform
[drasko]: https://github.com/drasko
[nikola]: https://github.com/nmarcetic
[dejan]: https://github.com/mijicd
[manu]: https://github.com/manuIO
[darko]: https://github.com/darkodraskovic
[janko]: https://github.com/janko-isidorovic
[alex]: https://github.com/anovakovic01
[dusan]: https://github.com/dusanb94
[sava]: https://github.com/srados
[george]: https://github.com/gesaleh
[iva]: https://github.com/blokovi
[kole]: https://github.com/chombium
[dusanm]: https://github.com/malidukica
[mirko]: https://github.com/mteodor
