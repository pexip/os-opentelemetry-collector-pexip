# Pexip FIPS checker module

This is a very simple Go module which can be used as follows in your `main` routine:

```
import (
    ...
    _ "github.com/pexip/gofipschecker"
    ...
)
```

The module init will then check for the FIPS enabled flag used in Pexip and
ensure that FIPS is enabled in the Go process if required. If not, it will look
for a binary with the same name with `.fips` appended.  If this exists, it will
exec it and the check will re-run.  If all of this fails, the binary will
panic.

This means that there must be two binary builds for each binary: `binary` and
`binary.fips`.  The second *must* be built using `GOFIPS140=certified go build`

Any Go-based binary running on PexOS doing crypto is required to include this
module.
