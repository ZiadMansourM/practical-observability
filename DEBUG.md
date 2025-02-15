```json title="initial TransformContext before executing StatementSequence"
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
        "log.file.path": "/log.log",
        "source": "docker"
      },
      "body": "{\"log\": \"level=info ts=2025-02-14T11:46:07.432755042Z caller=metrics.go:292 component=frontend org_id=fake traceID=556fafb62b54d438 latency=fast query_type=labels splits=0 start=2025-02-14T10:46:07.397Z end=2025-02-14T11:46:07.397Z start_delta=1h0m0.035751875s end_delta=35.75225ms length=1h0m0s duration=18.571917ms status=200 label= query= query_hash=2166136261 total_entries=2 cache_label_results_req=0 cache_label_results_hit=0 cache_label_results_stored=0 cache_label_results_download_time=0s cache_label_results_query_length_served=0s\\n\", \"stream\": \"stderr\", \"attrs\": {\"tag\": \"lgtm/loki\"}, \"time\": \"2025-02-14T11:46:07.471663667Z\"}",
      "dropped_attribute_count": 0,
      "flags": 0,
      "observed_time_unix_nano": 1739567809660598000,
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

```json title="TransformContext after executing StatementSequence"
{
  "kind": "processor",
  "name": "transform/service-name",
  "pipeline": "logs/filelog",
  "statement": "set(attributes[\"log.record.original\"], body)",
  "condition matched": true,
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
        "log.file.path": "/log.log",
        "source": "docker",
        "log.record.original": "{\"log\": \"level=info ts=2025-02-14T11:46:07.432755042Z caller=metrics.go:292 component=frontend org_id=fake traceID=556fafb62b54d438 latency=fast query_type=labels splits=0 start=2025-02-14T10:46:07.397Z end=2025-02-14T11:46:07.397Z start_delta=1h0m0.035751875s end_delta=35.75225ms length=1h0m0s duration=18.571917ms status=200 label= query= query_hash=2166136261 total_entries=2 cache_label_results_req=0 cache_label_results_hit=0 cache_label_results_stored=0 cache_label_results_download_time=0s cache_label_results_query_length_served=0s\\n\", \"stream\": \"stderr\", \"attrs\": {\"tag\": \"lgtm/loki\"}, \"time\": \"2025-02-14T11:46:07.471663667Z\"}"
      },
      "body": "{\"log\": \"level=info ts=2025-02-14T11:46:07.432755042Z caller=metrics.go:292 component=frontend org_id=fake traceID=556fafb62b54d438 latency=fast query_type=labels splits=0 start=2025-02-14T10:46:07.397Z end=2025-02-14T11:46:07.397Z start_delta=1h0m0.035751875s end_delta=35.75225ms length=1h0m0s duration=18.571917ms status=200 label= query= query_hash=2166136261 total_entries=2 cache_label_results_req=0 cache_label_results_hit=0 cache_label_results_stored=0 cache_label_results_download_time=0s cache_label_results_query_length_served=0s\\n\", \"stream\": \"stderr\", \"attrs\": {\"tag\": \"lgtm/loki\"}, \"time\": \"2025-02-14T11:46:07.471663667Z\"}",
      "dropped_attribute_count": 0,
      "flags": 0,
      "observed_time_unix_nano": 1739567809660598000,
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

```json title="TransformContext after executing StatementSequence Two"
{
  "kind": "processor",
  "name": "transform/service-name",
  "pipeline": "logs/filelog",
  "statement": "merge_maps(cache, ParseJSON(body), \"upsert\")",
  "condition matched": true,
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
        "log.file.path": "/log.log",
        "source": "docker",
        "log.record.original": "{\"log\": \"level=info ts=2025-02-14T11:46:07.432755042Z caller=metrics.go:292 component=frontend org_id=fake traceID=556fafb62b54d438 latency=fast query_type=labels splits=0 start=2025-02-14T10:46:07.397Z end=2025-02-14T11:46:07.397Z start_delta=1h0m0.035751875s end_delta=35.75225ms length=1h0m0s duration=18.571917ms status=200 label= query= query_hash=2166136261 total_entries=2 cache_label_results_req=0 cache_label_results_hit=0 cache_label_results_stored=0 cache_label_results_download_time=0s cache_label_results_query_length_served=0s\\n\", \"stream\": \"stderr\", \"attrs\": {\"tag\": \"lgtm/loki\"}, \"time\": \"2025-02-14T11:46:07.471663667Z\"}"
      },
      "body": "{\"log\": \"level=info ts=2025-02-14T11:46:07.432755042Z caller=metrics.go:292 component=frontend org_id=fake traceID=556fafb62b54d438 latency=fast query_type=labels splits=0 start=2025-02-14T10:46:07.397Z end=2025-02-14T11:46:07.397Z start_delta=1h0m0.035751875s end_delta=35.75225ms length=1h0m0s duration=18.571917ms status=200 label= query= query_hash=2166136261 total_entries=2 cache_label_results_req=0 cache_label_results_hit=0 cache_label_results_stored=0 cache_label_results_download_time=0s cache_label_results_query_length_served=0s\\n\", \"stream\": \"stderr\", \"attrs\": {\"tag\": \"lgtm/loki\"}, \"time\": \"2025-02-14T11:46:07.471663667Z\"}",
      "dropped_attribute_count": 0,
      "flags": 0,
      "observed_time_unix_nano": 1739567809660598000,
      "severity_number": 0,
      "severity_text": "",
      "span_id": "0000000000000000",
      "time_unix_nano": 0,
      "trace_id": "00000000000000000000000000000000"
    },
    "cache": {
      "attrs": {
        "tag": "lgtm/loki"
      },
      "time": "2025-02-14T11:46:07.471663667Z",
      "log": "level=info ts=2025-02-14T11:46:07.432755042Z caller=metrics.go:292 component=frontend org_id=fake traceID=556fafb62b54d438 latency=fast query_type=labels splits=0 start=2025-02-14T10:46:07.397Z end=2025-02-14T11:46:07.397Z start_delta=1h0m0.035751875s end_delta=35.75225ms length=1h0m0s duration=18.571917ms status=200 label= query= query_hash=2166136261 total_entries=2 cache_label_results_req=0 cache_label_results_hit=0 cache_label_results_stored=0 cache_label_results_download_time=0s cache_label_results_query_length_served=0s\n",
      "stream": "stderr"
    }
  }
}
```

```json title="TransformContext after executing StatementSequence Three"
{
  "kind": "processor",
  "name": "transform/service-name",
  "pipeline": "logs/filelog",
  "statement": "set(resource.attributes[\"service.name\"], cache[\"attrs\"][\"tag\"])",
  "condition matched": true,
  "TransformContext": {
    "resource": {
      "attributes": {
        "service.name": "lgtm/loki"
      },
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
        "log.file.path": "/log.log",
        "source": "docker",
        "log.record.original": "{\"log\": \"level=info ts=2025-02-14T11:46:07.432755042Z caller=metrics.go:292 component=frontend org_id=fake traceID=556fafb62b54d438 latency=fast query_type=labels splits=0 start=2025-02-14T10:46:07.397Z end=2025-02-14T11:46:07.397Z start_delta=1h0m0.035751875s end_delta=35.75225ms length=1h0m0s duration=18.571917ms status=200 label= query= query_hash=2166136261 total_entries=2 cache_label_results_req=0 cache_label_results_hit=0 cache_label_results_stored=0 cache_label_results_download_time=0s cache_label_results_query_length_served=0s\\n\", \"stream\": \"stderr\", \"attrs\": {\"tag\": \"lgtm/loki\"}, \"time\": \"2025-02-14T11:46:07.471663667Z\"}"
      },
      "body": "{\"log\": \"level=info ts=2025-02-14T11:46:07.432755042Z caller=metrics.go:292 component=frontend org_id=fake traceID=556fafb62b54d438 latency=fast query_type=labels splits=0 start=2025-02-14T10:46:07.397Z end=2025-02-14T11:46:07.397Z start_delta=1h0m0.035751875s end_delta=35.75225ms length=1h0m0s duration=18.571917ms status=200 label= query= query_hash=2166136261 total_entries=2 cache_label_results_req=0 cache_label_results_hit=0 cache_label_results_stored=0 cache_label_results_download_time=0s cache_label_results_query_length_served=0s\\n\", \"stream\": \"stderr\", \"attrs\": {\"tag\": \"lgtm/loki\"}, \"time\": \"2025-02-14T11:46:07.471663667Z\"}",
      "dropped_attribute_count": 0,
      "flags": 0,
      "observed_time_unix_nano": 1739567809660598000,
      "severity_number": 0,
      "severity_text": "",
      "span_id": "0000000000000000",
      "time_unix_nano": 0,
      "trace_id": "00000000000000000000000000000000"
    },
    "cache": {
      "attrs": {
        "tag": "lgtm/loki"
      },
      "time": "2025-02-14T11:46:07.471663667Z",
      "log": "level=info ts=2025-02-14T11:46:07.432755042Z caller=metrics.go:292 component=frontend org_id=fake traceID=556fafb62b54d438 latency=fast query_type=labels splits=0 start=2025-02-14T10:46:07.397Z end=2025-02-14T11:46:07.397Z start_delta=1h0m0.035751875s end_delta=35.75225ms length=1h0m0s duration=18.571917ms status=200 label= query= query_hash=2166136261 total_entries=2 cache_label_results_req=0 cache_label_results_hit=0 cache_label_results_stored=0 cache_label_results_download_time=0s cache_label_results_query_length_served=0s\n",
      "stream": "stderr"
    }
  }
}
```

```json title="TransformContext after executing StatementSequence Four"
{
  "kind": "processor",
  "name": "transform/service-name",
  "pipeline": "logs/filelog",
  "statement": "set(body, cache[\"log\"])",
  "condition matched": true,
  "TransformContext": {
    "resource": {
      "attributes": {
        "service.name": "lgtm/loki"
      },
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
        "log.file.path": "/log.log",
        "source": "docker",
        "log.record.original": "{\"log\": \"level=info ts=2025-02-14T11:46:07.432755042Z caller=metrics.go:292 component=frontend org_id=fake traceID=556fafb62b54d438 latency=fast query_type=labels splits=0 start=2025-02-14T10:46:07.397Z end=2025-02-14T11:46:07.397Z start_delta=1h0m0.035751875s end_delta=35.75225ms length=1h0m0s duration=18.571917ms status=200 label= query= query_hash=2166136261 total_entries=2 cache_label_results_req=0 cache_label_results_hit=0 cache_label_results_stored=0 cache_label_results_download_time=0s cache_label_results_query_length_served=0s\\n\", \"stream\": \"stderr\", \"attrs\": {\"tag\": \"lgtm/loki\"}, \"time\": \"2025-02-14T11:46:07.471663667Z\"}"
      },
      "body": "level=info ts=2025-02-14T11:46:07.432755042Z caller=metrics.go:292 component=frontend org_id=fake traceID=556fafb62b54d438 latency=fast query_type=labels splits=0 start=2025-02-14T10:46:07.397Z end=2025-02-14T11:46:07.397Z start_delta=1h0m0.035751875s end_delta=35.75225ms length=1h0m0s duration=18.571917ms status=200 label= query= query_hash=2166136261 total_entries=2 cache_label_results_req=0 cache_label_results_hit=0 cache_label_results_stored=0 cache_label_results_download_time=0s cache_label_results_query_length_served=0s\n",
      "dropped_attribute_count": 0,
      "flags": 0,
      "observed_time_unix_nano": 1739567809660598000,
      "severity_number": 0,
      "severity_text": "",
      "span_id": "0000000000000000",
      "time_unix_nano": 0,
      "trace_id": "00000000000000000000000000000000"
    },
    "cache": {
      "attrs": {
        "tag": "lgtm/loki"
      },
      "time": "2025-02-14T11:46:07.471663667Z",
      "log": "level=info ts=2025-02-14T11:46:07.432755042Z caller=metrics.go:292 component=frontend org_id=fake traceID=556fafb62b54d438 latency=fast query_type=labels splits=0 start=2025-02-14T10:46:07.397Z end=2025-02-14T11:46:07.397Z start_delta=1h0m0.035751875s end_delta=35.75225ms length=1h0m0s duration=18.571917ms status=200 label= query= query_hash=2166136261 total_entries=2 cache_label_results_req=0 cache_label_results_hit=0 cache_label_results_stored=0 cache_label_results_download_time=0s cache_label_results_query_length_served=0s\n",
      "stream": "stderr"
    }
  }
}
```

```json title="TransformContext after executing StatementSequence Five"

```