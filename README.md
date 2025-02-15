```bash
otel-collector-1  | 2025-02-14T20:02:03.614Z    debug   ottl@v0.118.0/parser.go:356     initial TransformContext before executing StatementSequence     {"kind": "processor", "name": "transform/service-name", "pipeline": "logs/filelog", "TransformContext": {"resource": {"attributes": {}, "dropped_attribute_count": 0}, "scope": {"attributes": {}, "dropped_attribute_count": 0, "name": "", "version": ""}, "log_record": {"attributes": {"log.file.path": "/var/lib/docker/containers/47a457e11ccce5fda8bb853cbe46c4eefa9a440bdc548576336dfa6b698b1908/47a457e11ccce5fda8bb853cbe46c4eefa9a440bdc548576336dfa6b698b1908-json.log", "source": "docker"}, "body": "{\"log\":\"2025-02-14T20:02:03.510Z\\u0009debug\\u0009fileconsumer/file.go:125\\u0009matched files\\u0009{\\\"kind\\\": \\\"receiver\\\", \\\"name\\\": \\\"filelog\\\", \\\"data_type\\\": \\\"logs\\\", \\\"component\\\": \\\"fileconsumer\\\", \\\"paths\\\": [\\\"/var/lib/docker/containers/0ffba8bce179d7620b78bc162aeafe5103b9eb07ee1b3226c090b7877f56fb1c/0ffba8bce179d7620b78bc162aeafe5103b9eb07ee1b3226c090b7877f56fb1c-json.log\\\", \\\"/var/lib/docker/containers/160adb257e91b037da0565696f4c0e12c764846aa336d3c9300869ce3ba884fb/160adb257e91b037da0565696f4c0e12c764846aa336d3c9300869ce3ba884fb-json.log\\\", \\\"/var/lib/docker/containers/1b2417fab3554afc6561ae6ba232d32900a3e59be1a07de7a5ece2e3d8857c37/1b2417fab3554afc6561ae6ba232d32900a3e59be1a07de7a5ece2e3d8857c37-json.log\\\", \\\"/var/lib/docker/containers/47a457e11ccce5fda8bb853cbe46c4eefa9a440bdc548576336dfa6b698b1908/47a457e11ccce5fda8bb853cbe46c4eefa9a440bdc548576336dfa6b698b1908-json.log\\\", \\\"/var/lib/docker/containers/92ee77848b0dcd1db20b5a53819b8a84b90ea915b1c6821cb7871f5ae2912220/92ee77848b0dcd1db20b5a53819b8a84b90ea915b1c6821cb7871f5ae2912220-json.log\\\", \\\"/var/lib/docker/containers/94b4fcb6dc940eeb7296a1c273bb63984b7d7631173e96f5a673a27adda31b5c/94b4fcb6dc940eeb7296a1c273bb63984b7d7631173e96f5a673a27adda31b5c-json.log\\\", \\\"/var/lib/docker/containers/c90e10160dc2967b2154d07872b8cfe0c85a3a5ee471e5e7149d0a3680d17a88/c90e10160dc2967b2154d07872b8cfe0c85a3a5ee471e5e7149d0a3680d17a88-json.log\\\", \\\"/var/lib/docker/containers/ed320d7df19b818c7dc618ccb74246af2f33ecde786b9588a2e0179d2ecf769f/ed320d7df19b818c7dc618ccb74246af2f33ecde786b9588a2e0179d2ecf769f-json.log\\\"]}\\n\",\"stream\":\"stderr\",\"attrs\":{\"tag\":\"lgtm/otel-collector\"},\"time\":\"2025-02-14T20:02:03.510686251Z\"}", "dropped_attribute_count": 0, "flags": 0, "observed_time_unix_nano": 1739563323511175918, "severity_number": 0, "severity_text": "", "span_id": "0000000000000000", "time_unix_nano": 0, "trace_id": "00000000000000000000000000000000"}, "cache": {}}}
```

```bash
{
  "kind": "processor",
  "name": "transform/service-name",
  "pipeline": "logs/filelog",
  "TransformContext": {
    "resource": {
      "attributes": {},
      "dropped_attribute_count": 0
    },
    "scope": {
      "attributes": {},
      "dropped_attribute_count": 0,
      "name": "",
      "version": ""
    },
    "log_record": {
      "attributes": {
        "log.file.path": "/var/lib/docker/containers/47a457e11ccce5fda8bb853cbe46c4eefa9a440bdc548576336dfa6b698b1908/47a457e11ccce5fda8bb853cbe46c4eefa9a440bdc548576336dfa6b698b1908-json.log",
        "source": "docker"
      },
      "body": "{\"log\":\"2025-02-14T20:02:03.510Z\\u0009debug\\u0009fileconsumer/file.go:125\\u0009matched files\\u0009{\\\"kind\\\": \\\"receiver\\\", \\\"name\\\": \\\"filelog\\\", \\\"data_type\\\": \\\"logs\\\", \\\"component\\\": \\\"fileconsumer\\\", \\\"paths\\\": [\\\"/var/lib/docker/containers/0ffba8bce179d7620b78bc162aeafe5103b9eb07ee1b3226c090b7877f56fb1c/0ffba8bce179d7620b78bc162aeafe5103b9eb07ee1b3226c090b7877f56fb1c-json.log\\\", \\\"/var/lib/docker/containers/160adb257e91b037da0565696f4c0e12c764846aa336d3c9300869ce3ba884fb/160adb257e91b037da0565696f4c0e12c764846aa336d3c9300869ce3ba884fb-json.log\\\", \\\"/var/lib/docker/containers/1b2417fab3554afc6561ae6ba232d32900a3e59be1a07de7a5ece2e3d8857c37/1b2417fab3554afc6561ae6ba232d32900a3e59be1a07de7a5ece2e3d8857c37-json.log\\\", \\\"/var/lib/docker/containers/47a457e11ccce5fda8bb853cbe46c4eefa9a440bdc548576336dfa6b698b1908/47a457e11ccce5fda8bb853cbe46c4eefa9a440bdc548576336dfa6b698b1908-json.log\\\", \\\"/var/lib/docker/containers/92ee77848b0dcd1db20b5a53819b8a84b90ea915b1c6821cb7871f5ae2912220/92ee77848b0dcd1db20b5a53819b8a84b90ea915b1c6821cb7871f5ae2912220-json.log\\\", \\\"/var/lib/docker/containers/94b4fcb6dc940eeb7296a1c273bb63984b7d7631173e96f5a673a27adda31b5c/94b4fcb6dc940eeb7296a1c273bb63984b7d7631173e96f5a673a27adda31b5c-json.log\\\", \\\"/var/lib/docker/containers/c90e10160dc2967b2154d07872b8cfe0c85a3a5ee471e5e7149d0a3680d17a88/c90e10160dc2967b2154d07872b8cfe0c85a3a5ee471e5e7149d0a3680d17a88-json.log\\\", \\\"/var/lib/docker/containers/ed320d7df19b818c7dc618ccb74246af2f33ecde786b9588a2e0179d2ecf769f/ed320d7df19b818c7dc618ccb74246af2f33ecde786b9588a2e0179d2ecf769f-json.log\\\"]}\\n\",\"stream\":\"stderr\",\"attrs\":{\"tag\":\"lgtm/otel-collector\"},\"time\":\"2025-02-14T20:02:03.510686251Z\"}",
      "dropped_attribute_count": 0,
      "flags": 0,
      "observed_time_unix_nano": 1739563323511176000,
      "severity_number": 0,
      "severity_text": "",
      "span_id": "0000000000000000",
      "time_unix_nano": 0,
      "trace_id": "00000000000000000000000000000000"
    },
    "cache": {}
  }
}
```

```bash
{
  "log": "level=info ts=2025-02-14T11:46:07.432755042Z caller=metrics.go:292 component=frontend org_id=fake traceID=556fafb62b54d438 latency=fast query_type=labels splits=0 start=2025-02-14T10:46:07.397Z end=2025-02-14T11:46:07.397Z start_delta=1h0m0.035751875s end_delta=35.75225ms length=1h0m0s duration=18.571917ms status=200 label= query= query_hash=2166136261 total_entries=2 cache_label_results_req=0 cache_label_results_hit=0 cache_label_results_stored=0 cache_label_results_download_time=0s cache_label_results_query_length_served=0s\n",
  "stream": "stderr",
  "attrs": {
    "tag": "lgtm/loki"
  },
  "time": "2025-02-14T11:46:07.471663667Z"
}
```

![alt text](image.png)

```bash
ziadh@Ziads-MacBook-Air lgtm % docker inspect --format '{{ index .Config.Labels "com.docker.compose.project" }}/{{ index .Config.Labels "com.docker.compose.service" }}' 1e575ea1b0be
lgtm/otel-collector
```

```bash
while true; do curl --fail http://localhost:3100/ready; sleep 1; done
while true; do curl --fail http://localhost:3200/ready; sleep 1; done

wget --quiet --tries=1 --output-document=- http://127.0.0.1:9009/ready

while true; do wget --quiet --tries=1 --output-document=- http://localhost:9009/ready; sleep 1; done
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
