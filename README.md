# Log Parser

This Go program parses an access log and reports:
- Number of unique IP addresses
- Top 3 most visited URLs
- Top 3 most active IPs

## How to run the app

1. Ensure Go is installed.
2. Clone the repo and the log file is placed at `log-parser/access.log`
3. Run:

```
make build
make run
```

## How to run the tests

```
make test
```

## How to run the linting

```
make lint
```

## Code considerations
1. Zerolog library is used for structured logging
2. When determining the top 3 most visited URLs and top 3 most active IP addresses, 
the current implementation do not apply any tie-breaking rule.
3. The resulting order for tied entries may vary between runs and left as-is for simplicity.

