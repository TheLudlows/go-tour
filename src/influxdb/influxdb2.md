#### ShardGroup & Shard

ShardGroup是用来InfluxDB中时间分区的，一个ShardGroup覆盖了一段时间，该时间段内的数据只会在对应的ShardGroup中。不同的ShardGroup不会重叠。ShardGroup是一个逻辑容器，它内部组织的Shard才是真正用来存储组织数据。Shard的实现是TSM(Time Sort Merge Tree) 引擎，TSM负责数据的编码存储、读写服务等。TSM类似于LSM。单机版的InfluxDB中一个ShardGroup只会包含一个Shard。