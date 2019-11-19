# logrus_fluent_example
example fluentd hooks for logrus

:octopus:

## Install Fluent

```Bash
docker pull fluent/fluentd
vim /tmp/fluentd.conf
docker run -it -p 9880:9880 -v /tmp:/fluentd/etc -e FLUENTD_CONF=fluentd.conf -v /data:/fluentd/log  fluent/fluentd
```
```fluentd.conf
<source>
    @type forward
    port 9880
  </source>
  <match **>
    @type stdout
  </match>
```

## Run main.go

```bash
go run main.go
```
