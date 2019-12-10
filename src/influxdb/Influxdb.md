#### 1. 安装及配置
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

| Point属性 |                                                  |
| --------- | ------------------------------------------------ |
| time      | 每个数据记录时间，是数据库中的主索引(会自动生成) |
| fields    | 各种记录值，没有索引的属性                       |
| tags      | 各种有索引的属性                                 |

#### 3. InfluxDB操作
   - 客户端命令行方式
   - HTTP API方式
   - 各种语言库

客户端方式

 ```sql
// DB相关
show databases // 查看库
create database MyDB // 创建数据库
drop database MyDB // 删除数据库
use MyDB //使用数据库
 ```

Measurement 类似于table，但不需要提前创建结构，第一次Insert操作会自动创建，在插入新数据时，tag、field和timestamp之间用空格分隔，

```sql
show measurements	// 查看表
insert log,rt=3,method="get" path="/login"
```

InfluxDB没有提供修改和删除数据的方法,删除可以通过InfluxDB的数据保存策略来实现。

#### 4. 保留策略(Retention Policies)
 InfluxDB本身不提供数据的删除操作，因此用来控制数据量的方式就是定义数据保留策略。定义数据保留策略的目的是让InfluxDB能够知道可以丢弃哪些数据，从而更高效的处理数据。

InfluxDB会比较服务器本地的时间戳和请求数据里的时间戳，并删除比在RPS里面用`DURATION`设置的更老的数据。一个数据库中可以有多个RPS。

```sql
show retention policies on [DB name] // 查看库的保存策略
 
name      duration shardGroupDuration replicaN default
----      -------- ------------------ -------- -------
autogen   0s       168h0m0s           1        false
```

默认的策略如上所示，

- name：名称

- duration：持续时间，0代表无限制

- shardGroupDuration: 简单的说InfluxDB为了有利于查询删除等操作，持久化是以时间来划分数据，比如一个小时的数据存为一个分片，shardGroupDuration指的是一个分片组覆盖特定的时间间隔， InfluxDB通过查看相关保留策略的持续时间来确定该时间间隔。 下表概述了RP的DURATION和分片组的时间间隔之间的默认关系：

  | RP duration               | Shard group interval |
  | ------------------------- | -------------------- |
  | < 2 days                  | 1 hour               |
  | >= 2 days and <= 6 months | 1 day                |
  | > 6 months                | 7 days               |

- replicaN ：副本个数，单机版只能为1

- default：是否是默认策略

新增/删除/修改

```sql
create retention policy "my_policy" on "logprocess" duration 2h replication 1 default
alter retention policy "my_policy" on logprocess duration 4h
drop retention policy "my_policy" on "logprocess"
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

上面创建了一个`my_cq`de CQ，作用于logprocess库的log表，每过30分钟对rt字段求平均值，然后插入avg_rt表中，使用my_rp保留策略。

查看/删除

```sql
show Continuous Queries
drop Continuous Query [cq_name] on [database_name]
```



#### 6. 聚合函数

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

#### 7. 选择函数

比如Top函数返回一个字段中最大的N个值，字段类型必须是长整型或float64类型。还有Bottom、First、Last、Max、Min等，使用都比较简单。

#### 8. 变换函数

