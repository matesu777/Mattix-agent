# Mattix-Agent

> Um agente de monitoramento de sistema leve escrito em Go

O Mattix é um agente de monitoramento simples, rápido e para plataformas Linux, projetado para coletar métricas do sistema e expô-las através de uma API HTTP.

Ele foi criado como um projeto de aprendizado inspirado em ferramentas como o **Node Exporter** e o **Zabbix Agent**, focando em simplicidade, e baixo uso de recursos.

## Funcionalidades

- Uso de CPU
- Uso de memória
- Uso de disco
- Estatísticas de rede
- Nome do host (Hostname)
- Tempo de atividade do sistema (Uptime)
- API HTTP JSON
- Leve
- Zero dependências externas (exceto x/sys)

## Exemplo

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

## Executando

```bash
git clone https://github.com/matesu777/Mattix-agent

cd Mattix-agent

go run .
```

A API estará disponível em

```
http://localhost:8080/metrics
```

---

## Roteiro (Roadmap)

- [x] Monitoramento de CPU
- [x] Monitoramento de memória
- [x] Monitoramento de disco
- [x] Monitoramento de rede
- [x] Nome do host (Hostname)
- [x] Tempo de atividade (Uptime)

### Próximos passos

- [X] Goroutines
- [ ] Temperatura
- [ ] Uso de CPU por núcleo
- [ ] Múltiplos discos
- [ ] Múltiplas interfaces de rede
- [ ] Monitoramento de processos
- [ ] Autenticação HTTP
- [ ] Arquivo de configuração

---

### Languages 
- [Portuguese](/README.pt-BR.md)
- [English](/README.md)
