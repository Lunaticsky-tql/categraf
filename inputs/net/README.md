# net

网络流量监控插件，比如各个网卡的流量、包量、错包情况等

## Configuration

通常可以维持默认配置，不过有的时候，我们有些网卡不想采集，只想采集指定的网卡，可以通过 interfaces 这个配置来指定。

## 监控大盘

该插件没有单独的监控大盘，OS 的监控大盘统一放到 system 下面了