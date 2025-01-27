```bash
while true; do curl --fail http://localhost:3100/ready; sleep 1; done
while true; do curl --fail http://localhost:3200/ready; sleep 1; done
```

## ToDo:
- [X] Grafana Volumes to save dashboards.
- [X] Loki to MinIO "Why not working!!!".
    - Showed folder named fake!
- [X] Loki and Tempo Volumes.
- [X] Tempo to MinIO.
    - DID NOT show folder named fake!
- [ ] derived fields.
- [ ] Get Mimir Up and running:
    - Instrument the code.
    - Modify Collector to send metrics to Mimir.
- [ ] Correlate Metrics with Traces.
- [ ] OTel Collector.

## Later:
- [ ] Kubernetes.
