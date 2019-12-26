[toc]

#### LSM Tree

LSM Tree (Log-Structured Merge Tree) 即日志合并树，被用于大量的数据库引擎中，如Hbase、LevelDB等。适用于海量数据的写入，而查询少的情况。主要思想是随机写转化为顺序写。基本流程为，最新的数据驻留在磁盘中，等到积累到足够多之后，再内存中有序的数据合并追加到磁盘队尾。为了解决防止数据丢失，使用WAL(Write Ahead Log)方式，写入内存的同时写入文件，用来恢复内存中的数据。以 LevelDB 为例，内存 中的数据达到指定阀值后会在写入一个新的文件。当某一层的文件数超过一定值后，就会将该层下的一个文件和更高一级的文件合并，由于 文件数据都是有序的，相当于是一个多路归并排序，所以合并操作相当快速，最终生成一个新的 文件，将旧的文件删除，这样就完成了一次合并过程。这也大概是为什么叫Level的的原因吧。

这里推荐几篇优秀的LSM Tree文章

[1. LSM-Tree VS B-Tree](https://blog.bcmeng.com/post/lsm-tree-vs-b-tree.html#5-sstables-and-lsm-trees)

[2. 野猪书读书笔记第三章](https://xieyuanpeng.com/2018/10/18/野猪书读书笔记第三章/)

#### InfluxDB 存储架构

InfluxDB在经历了LSM Tree、B+Tree等集中尝试后，最终自研TSM，TTSM全称是Time-Structured Merge Tree，思想类似LSM。我们先看它的整体架构：
![1](./1.png)

1. Shard
上一篇文章中提到过这个概念，InfluxDB 中按照数据的时间戳所在的范围，会去创建不同的Shard Group，而Shard Group中会包含一个至多个Shard，单机版本中只有一个Shard。每一个 shard 都有自己的 cache、wal、tsm files 以及 compactor。
2. Cache
  内存中暂存数据的地方，其实是一个map，key 为 seriesKey + FiledName，value为entry,具体实现为List<fieldkey,values>,values根据时间来排序。插入数据时，同时往 cache 与 wal 中写入数据，当Cache中的数据达到25M(默认)全部写入 tsm 文件。
3. WAL
  wal 文件的内容与内存中的 cache 相同，其作用就是为了防止系统崩溃导致的数据丢失。由于数据是顺序写入文件中，所以写入效率非常高。
4. TSM
5. Compactor

#### 文件目录介绍

Linux版本的数据文件位于`/var/lib/influxdb/`,Windows位于`/用户/.influxd/`下。主要有三个目录，分别是 meta, wal 以及 data 。meta 用于存储数据库的一些元数据，meta.db 文件。wal 目录存放预写日志文件，以 .wal 结尾。data 目录存放实际存储的数据文件，以 .tsm 结尾。整体结构如下所示：

```tree
influxdb/
  ├── data/
  |    └── [DB Name]/
  |        └── autogen/
  |            └── 33/ #shard group id
  |                ├── 000000001-0000000001.tsm
  |                └── fields.idx
  |.   		 └── _series
  ├── meta/
  |    └── meta.db
  └── wal/
      └── [DB Name]/
          └── autogen/ #retention policy
              └── 33/ # shard group id
                  └── _00001.wal
```

#### WAL

#### TSM File 

#### 

#### 

