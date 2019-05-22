# NFS-Ganesha Exporter
Prometheus exporter for [NFS-Ganesha](https://github.com/nfs-ganesha/nfs-ganesha/)

This exporter uses Dbus interface to get the metrics.

## Installation
You can build the latest version using Go v1.11+ via `go get`:
```
go get -u github.com/Gandi/ganesha_exporter
```

## Usage
```
usage: ganesha_exporter [<flags>]

Flags:
  -h, --help                     Show context-sensitive help (also try --help-long and --help-man).
      --web.listen-address=":9587"
                                 Address on which to expose metrics and web interface.
      --web.telemetry-path="/metrics"
                                 Path under which to expose metrics.
      --gandi                    Activate Gandi specific fields
      --collector.exports        Activate exports collector
      --collector.exports.nfsv3  Activate NFSv3 stats
      --collector.exports.nfsv40
                                 Activate NFSv4.0 stats
      --collector.exports.nfsv41
                                 Activate NFSv4.1 stats
      --collector.exports.pnfsv41
                                 Activate pNFSv4.1 stats
      --collector.clients        Activate clients collector
      --collector.clients.nfsv3  Activate NFSv3 stats
      --collector.clients.nfsv40
                                 Activate NFSv4.0 stats
      --collector.clients.nfsv41
                                 Activate NFSv4.1 stats
      --collector.clients.pnfsv41
                                 Activate pNFSv4.1 stats
      --log.level="info"         Only log messages with the given severity or above. Valid levels: [debug,
                                 info, warn, error, fatal]
      --log.format="logger:stderr"
                                 Set the log target and format. Example:
                                 "logger:syslog?appname=bob&local=7" or "logger:stdout?json=true"
      --version                  Show application version.
```

All collectors are activated by default, they can be de-activated using `--no-collector.XXX`

The additional statistics retrieved by the `--gandi` flag are part of an internal WIP to get more
comprehensive statistics and will be proposed upstream as soon as they are fully done.


