# check-file-agedir

## Help

```
$ check-file-agedir -h
Usage:
  check-file-agedir [OPTIONS]

Application Options:
  -b, --base=         the base directory(required)
  -w, --warning-age=  warning if more old than(sec) (default: 21600)
  -c, --critical-age= critical if more old than(sec) (default: 43200)
  -d, --debug         debug print

Help Options:
  -h, --help          Show this help message
```

## Install

```
sudo mkr plugin install itsubaki/check-file-agedir@v0.2
```

```
# /etc/mackerel-agent/mackerel-agent.conf
[plugin.checks.fage_td-agent_buffer]
command = "/opt/mackerel-agent/plugins/bin/check-file-agedir -b /var/log/td-agent/buffer/"
```
