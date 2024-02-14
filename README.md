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
echo -e 'foo\nbar' | limiter
```

```bash
echo -e 'dogs\ncats' | limiter -d 1s -r 1 | xargs -I {} curl 'https://myapi.com?search={}'
```
