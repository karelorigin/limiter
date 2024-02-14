# Limiter

A small command-line utility to artificially limit the input rate to STDIN.

Limiter is superior over a command such as `sleep` because it will not wait for slow calls to finish before 'ticking', making up for the time that would have otherwise been lost.

A common use case is making HTTP requests to rate-limited REST API endpoints.

## Installation

```bash
go install github.com/karelorigin/limiter@latest
```

## Usage

```
Usage of limiter:
  -d duration
        The time to wait after each processed batch. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'. (default 1s)
  -r int
        The max processing rate per unit of time. (default 1)
```

## Examples

```bash
echo -e 'dogs\ncats' | limiter -d 1s -r 1 | xargs -I {} -P 5 curl 'https://myapi.com?search={}'
```

## Pitfalls
It is possible that some programs may need minor tweaking to function correctly.

`xargs`, for example, will do input buffering if it becomes too slow, causing it to possibly make multiple calls in a shorter-than-intended timeframe. This can be solved by upping the parallelism count via the `-P` flag.

[httpx](https://github.com/projectdiscovery/httpx), will by default, attempt to read the entire STDIN before finally processing URLs. This can be resolved using the `--stream` flag. Though it's worth noting that httpx has its own rate-limiting functionality.