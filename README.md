### Kubevars

Kubernetes Pod has environment variables set via Docker to point to services running on the same cluster:

    "Env": [
        "KUBERNETES_SERVICE_HOST=10.84.0.1",
        "KUBERNETES_SERVICE_PORT=443",
        "KUBERNETES_PORT=tcp://10.84.0.1:443",
        "KUBERNETES_PORT_443_TCP=tcp://10.84.0.1:443",
        "KUBERNETES_PORT_443_TCP_PROTO=tcp",
        "KUBERNETES_PORT_443_TCP_PORT=443",
        "KUBERNETES_PORT_443_TCP_ADDR=10.84.0.1",
        "KUBERNETES_RO_SERVICE_HOST=10.84.0.2",
        "KUBERNETES_RO_SERVICE_PORT=80",
        "KUBERNETES_RO_PORT=tcp://10.84.0.2:80",
        "KUBERNETES_RO_PORT_80_TCP=tcp://10.84.0.2:80",
        "KUBERNETES_RO_PORT_80_TCP_PROTO=tcp",
        "KUBERNETES_RO_PORT_80_TCP_PORT=80",
        "KUBERNETES_RO_PORT_80_TCP_ADDR=10.84.0.2",
        "INFLUX_MASTER_SERVICE_HOST=10.84.0.3",
        "INFLUX_MASTER_SERVICE_PORT=8085",
        "INFLUX_MASTER_PORT=tcp://10.84.0.3:8085",
        "INFLUX_MASTER_PORT_8085_TCP=tcp://10.84.0.3:8085",
        "INFLUX_MASTER_PORT_8085_TCP_PROTO=tcp",
        "INFLUX_MASTER_PORT_8085_TCP_PORT=8085",
        "INFLUX_MASTER_PORT_8085_TCP_ADDR=10.84.0.3",
        "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
    ],

I want to get the same information for my applications that are launched by other means, like [CoreOS] Fleet or [Deis].

`kubevars` come to rescue:

    $ ./bin/kubevars.amd64 -h
    Usage: kubevars [-etcd http://host[:port]] [-format <docker|json|plain>]
    Flags:
      -etcd="http://localhost:4001": The ETCD endpoint
      -format="docker": One of: docker, json, plain
    $ ./bin/kubevars.amd64
    -e INFLUX_MASTER_SERVICE_HOST="10.84.0.3"
    -e INFLUX_MASTER_SERVICE_PORT="8085"
    -e INFLUX_MASTER_PORT="tcp://10.84.0.3:8085"
    -e INFLUX_MASTER_PORT_8085_TCP="tcp://10.84.0.3:8085"
    -e INFLUX_MASTER_PORT_8085_TCP_PROTO="tcp"
    -e INFLUX_MASTER_PORT_8085_TCP_PORT="8085"
    -e INFLUX_MASTER_PORT_8085_TCP_ADDR="10.84.0.3"
    -e KUBERNETES_SERVICE_HOST="10.84.0.1"
    -e KUBERNETES_SERVICE_PORT="443"
    -e KUBERNETES_PORT="tcp://10.84.0.1:443"
    -e KUBERNETES_PORT_443_TCP="tcp://10.84.0.1:443"
    -e KUBERNETES_PORT_443_TCP_PROTO="tcp"
    -e KUBERNETES_PORT_443_TCP_PORT="443"
    -e KUBERNETES_PORT_443_TCP_ADDR="10.84.0.1"
    -e KUBERNETES_RO_SERVICE_HOST="10.84.0.2"
    -e KUBERNETES_RO_SERVICE_PORT="80"
    -e KUBERNETES_RO_PORT="tcp://10.84.0.2:80"
    -e KUBERNETES_RO_PORT_80_TCP="tcp://10.84.0.2:80"
    -e KUBERNETES_RO_PORT_80_TCP_PROTO="tcp"
    -e KUBERNETES_RO_PORT_80_TCP_PORT="80"
    -e KUBERNETES_RO_PORT_80_TCP_ADDR="10.84.0.2"


[CoreOS]: https://coreos.com/
[Deis]: http://deis.io
