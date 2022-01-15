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

Build using `docker buildx build  --push(?) --tag solsson(?)/prometheus-deconz-exporter:$(git rev-parse HEAD) --platform=linux/amd64,linux/arm64 .`

## References

 * https://phoscon.de/en/conbee2/install#docker
 * https://github.com/kubernetes/kubernetes/issues/5607#issuecomment-274942938
   - https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.17/#podspec-v1-core

 * https://github.com/marthoc/docker-deconz#configuring-raspbian-for-raspbee


 * https://github.com/dresden-elektronik/deconz-rest-plugin/issues/2273
 * https://github.com/dresden-elektronik/deconz-rest-plugin/wiki/Network-lost-issues
 * https://github.com/jurgen-kluft/go-conbee

## Fresh start 2022

Still on a Raspberry Pi 1. Used the [imager](https://www.raspberrypi.com/software/) and selected 32bit Bullseye.

Deconz version [2.13.04](https://deconz.dresden-elektronik.de/raspbian/stable/deconz-2.13.04-qt5.deb).

Firmware [0x26720700](http://deconz.dresden-elektronik.de/deconz-firmware/deCONZ_RaspBeeII_0x26720700.bin.GCF)
[md5](http://deconz.dresden-elektronik.de/deconz-firmware/deCONZ_RaspBeeII_0x26720700.bin.GCF.md5).

Deconz installation

```
sudo dpkg -i ...
sudo apt --fix-broken install
sudo systemctl enable deconz
```

To get the API key, in the gateway click "Authenticate app" then run something like `curl -X POST -H 'Content-Type: application/json' --data '{"devicetype":"local-kubernetes"}' http://192.168.3.45/api`. Save the KEY.
