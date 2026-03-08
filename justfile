test:
    #!/usr/bin/env bash
    set -euo pipefail
    jobs="$(getconf _NPROCESSORS_ONLN 2>/dev/null || sysctl -n hw.logicalcpu)"
    go test -tags=unit -p "$jobs" -race ./...
