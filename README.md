# Sensors to Prometheus via ZigBee

This repo is for exploring how to operate ZigBee sensors in the long run using Kubernetes.
I liked the "without cloud" approach of [Conbee II](https://phoscon.de/en/conbee2) so that's my starting point.
The goal isn't primarily automation, but only a neat Grafana dashboard.
Also I've learned to appreciate how many self-hosting problems Kubernetes solves,
so I'm not very interested in existing high level frameworks. My approach is:

```
Hardware  -> Driver -> Vendor's gateway/API -> a prometheus metrics endpoint -> Prometheus -> Grafana
```

The current metrics container assumes access to a [Deconz REST API](https://github.com/dresden-elektronik/deconz-rest-plugin)
and is therefore built as the image [solsson/prometheus-deconz-exporter](https://hub.docker.com/r/solsson/prometheus-deconz-exporter).

## References

 * https://phoscon.de/en/conbee2/install#docker
 * https://github.com/kubernetes/kubernetes/issues/5607#issuecomment-274942938
   - https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.17/#podspec-v1-core

 * https://github.com/marthoc/docker-deconz#configuring-raspbian-for-raspbee


 * https://github.com/dresden-elektronik/deconz-rest-plugin/issues/2273
 * https://github.com/dresden-elektronik/deconz-rest-plugin/wiki/Network-lost-issues
 * https://github.com/jurgen-kluft/go-conbee
