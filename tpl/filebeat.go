package template
var(
  Filebeat=`
# 输入
# filebeat 下载地址
# https://www.elastic.co/cn/downloads/past-releases/filebeat-7-9-3/
filebeat.inputs:
  - type: log
    enabled: true
    paths:
      - ./*.log
#输出
output.logstash:
  hosts: ["localhost:5044"]
`
)
