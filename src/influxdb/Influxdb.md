1. 安装及配置
    `docker pull influxdb`
    `docker run -d -p 8086:8086 --name myinfluxdb influxdb`

    `mkdir -p /Users/liuchao56/influxdb/data /Users/liuchao56/influxdb/conf /Users/liuchao56/influxdb/meta /Users/liuchao56/influxdb/wal`
    

docker exec -it myinfluxdb bash