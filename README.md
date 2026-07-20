# Mattix-Agent

> A lightweight system monitoring agent written in Go

Mattix is a simple, fast and linux-plataform monitoring agent designed to collect system metrics and expose them through an HTTP API.

It was created as a learning project inspired by tools like **Node Exporter** and **Zabbix Agent**, focusing on simplicity,  and low resource usage.

## Features

- CPU usage
- Memory usage
- Disk usage
- Network statistics
- Hostname
- System uptime
- JSON HTTP API
- Lightweight
- Zero external dependencies (except x/sys)

## Example

```json
{
  "hostname": "Ubuntu",
  "uptime": 22187,
  "cpu": {
    "usage": 12.73,
    "cores": 12
  },
  "memory": {
    "total": 16402235392,
    "used": 5715677184,
    "free": 10686558208
  },
  "disk": {
    "total": 64424509440,
    "used": 49529016320,
    "free": 14895493120
  },
  "network": {
    "name": "wlp0s20f3",
    "ipv4": "192.168.1.20",
    "mac": "28:0c:50:db:5c:6d",
    "rx_bytes": 872827727,
    "tx_bytes": 44564219,
    "rx_speed": 1745,
    "tx_speed": 231
  }
}
```

## Running

```bash
git clone https://github.com/matesu777/Mattix-agent

cd Mattix-agent

go run .
```


The API will be available at

```
http://localhost:8080/metrics
```

---

## Roadmap

- [x] CPU monitoring
- [x] Memory monitoring
- [x] Disk monitoring
- [x] Network monitoring
- [x] Hostname
- [x] Uptime

### Next

- [ ] Goroutines
- [ ] Temperature
- [ ] Per-core CPU usage
- [ ] Multiple disks
- [ ] Multiple network interfaces
- [ ] Process monitoring
- [ ] HTTP authentication
- [ ] Configuration file

---

## Technologies

- Go
