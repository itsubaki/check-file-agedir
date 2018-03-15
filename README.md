# check-fileagedir

# Help

```
$ check-fileagedir -h
Usage:
  check-fileagedir [OPTIONS]

Application Options:
  -b, --base=         the base directory(required)
  -w, --warning-age=  warning if more old than(sec) (default: 21600)
  -c, --critical-age= critical if more old than(sec) (default: 43200)
  -d, --debug         debug print

Help Options:
  -h, --help          Show this help message
```

# Example

```
# /etc/mackerel-agent/mackerel-agent.conf
[plugin.checks.fage_td-agent_buffer]
command = "check-fileagedir -b /var/log/td-agent/buffer/"
```