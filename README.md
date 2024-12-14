<!--
parent:
  order: false
-->

<div align="center">
  <h1> Hetu Hub </h1>
</div>

<div align="center">
  <a href="https://github.com/hetu-project/hetu-hub/releases/latest">
    <img alt="Version" src="https://img.shields.io/github/tag/hetu-project/hetu-hub.svg" />
  </a>
  <a href="https://github.com/hetu-project/hetu-hub/blob/main/LICENSE">
    <img alt="License: Apache-2.0" src="https://img.shields.io/github/license/hetu-project/hetu-hub.svg" />
  </a>
  <a href="https://pkg.go.dev/github.com/hetu-project/hetu-hub">
    <img alt="GoDoc" src="https://godoc.org/github.com/hetu-project/hetu-hub?status.svg" />
  </a>
  <a href="https://goreportcard.com/report/github.com/hetu-project/hetu-hub">
    <img alt="Go report card" src="https://goreportcard.com/badge/github.com/hetu-project/hetu-hub"/>
  </a>
</div>
<div align="center">
  <a href="https://discord.gg/hetu">
    <img alt="Discord" src="https://img.shields.io/discord/809048090249134080.svg" />
  </a>
  <a href="https://github.com/hetu-project/hetu-hub/actions?query=branch%3Amain+workflow%3ALint">
    <img alt="Lint Status" src="https://github.com/hetu-project/hetu-hub/actions/workflows/lint.yml/badge.svg?branch=main" />
  </a>
  <a href="https://codecov.io/gh/hetu-project/hetu-hub">
    <img alt="Code Coverage" src="https://codecov.io/gh/hetu-project/hetu-hub/branch/main/graph/badge.svg" />
  </a>
</div>

Hetu Hub is a scalable, high-throughput blockchain that is fully compatible and interoperable with Ethereum.
It's built using the [Cosmos SDK](https://github.com/cosmos/cosmos-sdk/) and implements
[Narwhal](https://github.com/MystenLabs/narwhal) + [Bullshark](https://github.com/MystenLabs/sui/tree/main/consensus/bullshark) consensus mechanism
for improved performance and scalability.

## Documentation

Our documentation is hosted in a [separate repository](https://github.com/hetu-project/docs) and can be found at [docs.hetu.org](https://docs.hetu.org).
Head over there and check it out.

**Note**: Requires [Go 1.20+](https://golang.org/dl/)

## Installation

For prerequisites and detailed build instructions
please read the [Installation](https://docs.hetu.org/protocol/hetu-cli) instructions.
Once the dependencies are installed, run:


make install

Or check out the latest [release](https://github.com/hetu-project/hetu-hub/releases).

## Quick Start

To learn how the Hetu Hub works from a high-level perspective,
go to the [Protocol Overview](https://docs.hetu.org/protocol) section from the documentation.
You can also check the instructions to [Run a Node](https://docs.hetu.org/protocol/hetu-cli#run-a-hetu-node).

## Community

The following chat channels and forums are a great spot to ask questions about Hetu Hub:

- [Open an Issue](https://github.com/hetu-project/hetu-hub/issues)
- [Hetu Protocol](https://github.com/hetu-project#hetu-key-research)
- [Follow us on Twitter](https://x.com/hetu_protocol)

## Contributing

We welcome all contributions! There are many ways to contribute to the project, including but not limited to:

- Cloning code repo and opening a [PR](https://github.com/hetu-project/hetu-hub/pulls).
- Submitting feature requests or [bugs](https://github.com/hetu-project/hetu-hub/issues).
- Improving our product or contribution [documentation](./docs/).
- Contributing [use cases](./demos/) to a feature request.

For additional instructions, standards and style guides, please refer to the [Contributing](./CONTRIBUTING.md) document.

