[toc]

#### 1. 快速使用

```shell    
docker pull influxdbdocker run -d -p 8086:8086 --name myinfluxdb influxdb
docker exec -it [myinfluxdb] bash
```

#### 2. 基本概念
类比于传统数据库

| influxDB中的名词 | 传统数据库中的概念 |
| ---------------- | ------------------ |
| database         | 数据库             |
| measurement      | 数据库中的表       |
| points           | 表里面的一行数据   |

Point由时间戳（time）、数据（field）、标签（tags）组成。Point相当于传统数据库里的一行数据，如下表所示：

| Point属性 |                                                              |
| --------- | ------------------------------------------------------------ |
| time      | 每个数据记录时间，是数据库中的主索引(会自动生成)             |
| fields    | 各种记录值，没有索引的属性，值可以为64位整型，64位浮点型，字符串以及布尔型 |
| tags      | 各种有索引的属性，值只能为字符串                             |



#### 3. InfluxDB操作

   - 客户端命令行方式
   - HTTP API方式
   - 各种语言库

客户端方式

 ```sql
influx -- 进入client
show databases -- 查看库
create database MyDB -- 创建数据库
drop database MyDB -- 删除数据库
use MyDB --使用数据库
 ```

Measurement 类似于table，但不需要提前创建结构，第一次Insert操作会自动创建，在插入新数据时，tag、field和timestamp之间用空格分隔，

```sql
show measurements	-- 查看表
insert log,rt=3,method="get" path="/login" -- rt、method为tag，path为faild
show tag keys from log -- 查看log表的tag key
show tag values from log with key="method" [where 筛选field条件]-- 查看log表中tag为method的values
show field keys from log -- 查看log表的field key
```

InfluxDB没有提供修改和删除数据的方法,删除可以通过InfluxDB的数据保存策略来实现。

#### 4. 保留策略(Retention Policies)
 InfluxDB本身不提供数据的删除操作，因此用来控制数据量的方式就是定义数据保留策略。定义数据保留策略的目的是让InfluxDB能够知道可以丢弃哪些数据，从而更高效的处理数据。

InfluxDB会比较服务器本地的时间戳和请求数据里的时间戳，并删除比在RPS里面用`DURATION`设置的更老的数据。一个数据库中可以有多个RP。

```sql
show retention policies on [DB name] -- 查看库的保存策略
name      duration shardGroupDuration replicaN default
----      -------- ------------------ -------- -------
autogen   0s       168h0m0s           1        false

--新增/删除/修改
create retention policy "my_policy" on "logprocess" duration 2h replication 1 [shard duration 1h] default
alter retention policy "my_policy" on logprocess duration 4h
drop retention policy "my_policy" on "logprocess"
```

默认的策略如上所示，

- name：名称

- duration：持续时间，0代表无限制

- shardGroupDuration: 简单的说InfluxDB为了有利于查询删除等操作，持久化是以时间来划分数据，比如一个小时的数据存为一个分片，shardGroupDuration指的是一个分片组覆盖特定的时间间隔，如果没有指定 shard duration <duration>，InfluxDB通过查看相关保留策略的持续时间来确定该时间间隔。 下表概述了RP的DURATION和分片组的时间间隔之间的默认关系：

  | RP duration               | Shard group interval |
  | ------------------------- | -------------------- |
  | < 2 days                  | 1 hour               |
  | >= 2 days and <= 6 months | 1 day                |
  | > 6 months                | 7 days               |

- replicaN ：副本个数，单机版只能为1

- default：是否是默认策略


写入数据是可以指定使用的RP，若不指定则使用默认的。指定方式为如下,注意和普通的多了`into`关键字
````sql
insert into my_ploicy log,rt=3,method="get" path="/login"
```


查询也可以指定RP，若不指定使用默认的

```sql
select * from my_ploicy.log
```

#### 5. 连续查询(Continuous Query)

如果我们不想完全将这些数据删除掉，就需要连续查询（Continuous Queries）的帮助了。连续查询主要用在将数据归档，以降低系统空间的占用率，主要是以降低精度为代价。

Continuous Query其实就是InfluxDB内部周期执行的一个查询，将结果放入另一个表中。这个查询有如下限制：CQS需要在`SELECT`语句中使用一个函数，并且一定包括一个`GROUP BY time()`语句。

```sql
create Continuous Query "my_cq" on "logprocess"  
begin 
  select mean("rt") AS "mean_rt"
  into "my_policy"."avg_rt"
  from "log"
  group by time(30m)
end
```

上面创建了一个`my_cq`的 CQ，作用于logprocess库的log表，每过30分钟对rt字段求平均值，然后插入avg_rt表中，使用my_rp保留策略。

查看/删除

```sql
show Continuous Queries
drop Continuous Query [cq_name] on [database_name]
```

#### 6. 补充概念

- Filed key，以`insert into my_ploicy log,rt=3,method="get" path="/login"为例，Filed key为path
- Filed value为/login
- Tag keys为rt和method
- Tag values为3和“get”，这里3存储将会转化为字符串类型
- Field Set为path="/login",注意value相同为同一个Field Set
- Tag Set为rt=3,method="get"，value相同为同一个Tag Set
- Series是一些具有相同RP、Measurement和Tag Set的的Points集合
- Series key 标识一个Series，由Measurement和Tag set组成
- Point表示一行记录，Series和time决定一个Points，比如插入多次相同的RP、time和tag value的记录，但是只会产生一条记录。

#### 7. ShardGroup & Shard & Sharding

ShardGroup是用来做InfluxDB中时间分区，一个ShardGroup覆盖了一段时间，该时间段内的数据只会在对应的ShardGroup中。不同的ShardGroup不会重叠。ShardGroup是一个逻辑容器，它内部组织的Shard才是真正的存储引擎。Shard的实现是TSM(Time Sort Merge Tree) 引擎，TSM负责数据的编码存储、读写服务等。ShardGroup内部可能包含多个Shard，首先数据根据时间选择落在哪个ShardGroup，然后根据Series进行Hash再进行一次分区，决定进入哪个Shard。第一层Rang Sharding，第二层Hash Sharding。双层Sharding设计主要是为了解决热点写入问题。单机版的的Shard个数固定为1，而集群版的Shard个数取决于副本数和节点数。

#### 附1. 函数

##### 聚合函数

- count

  返回一个field中的非空值的数量（log表中提前填充了几条记录）

  ```sql
  > select count(path) from log
  name: log
  time count
  ---- -----
  0    5
  ```

  函数如果没有指定时间的话，会默认以 0开始

- Distinct

  返回一个field的唯一值。

  ```sql
  > select distinct(path) from log
  name: log
  time distinct
  ---- --------
  0    /login
  0    /user
  ```

- Mean

  返回一个field中的值的算术平均值。字段类型必须是长整型或float64

-  Median 

  从单个field中的排序值返回中间值。字段值的类型必须是长整型或float64格式

- Spread

  返回字段的最小值和最大值之间的差值。数据的类型必须是长整型或float64。

- Sum

  返回一个字段中的所有值的和。字段的类型必须是长整型或float64。

##### 选择函数、变换函数

比如Top函数返回一个字段中最大的N个值，字段类型必须是长整型或float64类型。还有Bottom、First、Last、Max、Min等，使用都比较简单。

#### 附2. Go Client的使用

参考Demo : https://github.com/TheLudlows/go-tour/blob/master/src/influxdb/influx/InfluxDB.go

