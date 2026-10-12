# Pexip Opentelemetry Collector build

This is a repo which contains the code for building the Pexip distribution of
the Opentelemetry collector.  The distribution contains the minimum number of
modules needed to implement our required functionality.

# Updating the Pexip opentelemetry collector

The key file is `manifest.yaml`.  Everything else should be treat as "throw-away".
For examples of manifest files, see the upstream releases repository:
https://github.com/open-telemetry/opentelemetry-collector-releases/tree/main/distributions

You will need the `opentelemetry-collector-builder` package installed.  For
newer versions of the collector, you may need to update the builder package
first.

After editing `manifest.yaml`, run the prepare source script which will rebuild
the `_build` directory, apply our patches and vendor the relevant modules:

```
./prepare_source
```

Watch for any warnings or errors.

Bump the Debian changelog, commit the new source tree and build in CI.
