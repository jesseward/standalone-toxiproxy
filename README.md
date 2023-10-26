## Summary

An example that highlights how we can decouple a custom built Toxic from the repo.

* toxics and tests we build are are placed into `ghtoxics` , right now there is a single `debug_toxic.go` from the upstream _examples
* (re)implement the cmd/server.go command (taken from upstream). Though we ensure we `toxics.Register("debug", new(ghtoxics.DebugToxic))` in the cmd to ensure it's discovered
* The makefile will spit out the server and cli.
* `tools/tools.go` is a placeholder to ensure that `"github.com/Shopify/toxiproxy/v2/cmd/cli` isn't wiped from go.mod

## Build

```sh
make all # and then make build
```

Boot the server and create the proxy and toxics
```sh
$ dist/toxiproxy-server &  # boot the server
$ ./dist/toxiproxy-cli --host "http://localhost:8474" create -l :9999 -u localhost:9998 blah # create proxy
$ ./dist/toxiproxy-cli --host "http://localhost:8474" toxic add --type debug blah # add our debug proxy
$ ./dist/toxiproxy-cli --host "http://localhost:8474" toxic add --type latency blah # add the latency proxy
```

inspect server state to see what is bound.
```sh
# dist/toxiproxy-cli inspect blah # inspect the state.

Name: blah      Listen: [::]:9999       Upstream: localhost:9998
======================================================================
Upstream toxics:
Proxy has no Upstream toxics enabled.

Downstream toxics:
debug_downstream:       type=debug      stream=downstream       toxicity=1.00   attributes=[    ]
latency_downstream:     type=latency    stream=downstream       toxicity=1.00   attributes=[    jitter=0        latency=0       ]
```