<a href="https://www.streamingfast.io/">
	<img width="100%" src="assets/firehose-banner.png" alt="StreamingFast Firehose Banner" />
</a>

# Firehose for Vara

[![reference](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://pkg.go.dev/github.com/streamingfast/firehose-vara)

Quick start with Firehose for Vara can be found in the official Firehose docs. Here some quick links to it:

- [Firehose Overview](https://firehose.streamingfast.io/introduction/firehose-overview)
- [Concepts & Architectures](https://firehose.streamingfast.io/concepts-and-architeceture)
  - [Components](https://firehose.streamingfast.io/concepts-and-architeceture/components)
  - [Data Flow](https://firehose.streamingfast.io/concepts-and-architeceture/data-flow)
  - [Data Storage](https://firehose.streamingfast.io/concepts-and-architeceture/data-storage)
  - [Design Principles](https://firehose.streamingfast.io/concepts-and-architeceture/design-principles)

# Running the Firehose poller

The below command with start streaming Firehose Vara blocks, check `proto/sg/gear/type/v1/block.proto` for more information.

```bash
firevara fetch {FIRST_STREAMABLE_BLOCK} --endpoints {VARA_RPC_ENDPOINT} --state-dir {STATE_DIR}
```

## Contributing

Report any protocol-specific issues in their
[respective repositories](https://github.com/streamingfast/streamingfast#protocols)

**Please first refer to the general
[StreamingFast contribution guide](https://github.com/streamingfast/streamingfast/blob/master/CONTRIBUTING.md)**,
if you wish to contribute to this code base.

This codebase uses unit tests extensively, please write and run tests.
