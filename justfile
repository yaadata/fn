dev-dependencies:
    mise install
    mise exec -- pre-commit install --hook-type commit-msg

bootstrap: dev-dependencies

unit-test:
    #!/usr/bin/env bash
    set -euo pipefail
    jobs="$(getconf _NPROCESSORS_ONLN 2>/dev/null || sysctl -n hw.logicalcpu)"
    mise exec -- go test -tags=unit -p "$jobs" -race ./...

test: unit-test

generate-pkl:
    mise exec -- pkl run package://pkg.pkl-lang.org/pkl-go/pkl.golang@0.13.1#/gen.pkl --output-path pkl/golang --base-path github.com/yaadata/fn/pkl/golang pkl/schema.pkl
