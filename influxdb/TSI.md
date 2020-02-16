



#### TSI 

默认配置下索引信息存放于内存中，`index-version="tsi1"`配置作用是将索引信息持久化至硬盘，TSI索引主要包含logFile文件(.tsl结尾)、tsi文件，其实Series 文件也算。tsl、tsi文件位于在每个shard中`index`目录下。

##### LogFile

##### Index File

当TSL文件大小达到配置的compaction阈值时（由配置文件中的max-index-log-file-size指定，默认为1M），LogFile文件会compaction成TSI文件，TSI文件算是存储格式最为复杂的。

| Magic  | Tag Blocks | Measurement Block | SeriesID Set | TombstoneSeriesIDSet | SeriesSketch | TombstoneSketch | Trailer |
| :----: | :--------: | :---------------: | :----------: | :------------------: | ------------ | --------------- | ------- |
| 4 Byte |            |                   |              |                      |              |                 | 82 Byte |

- Magic：占用4Byte，1.79版本位`TSI1`
- Tag Blocks：每个Measurement占用拥有一个Tag Block，Tag Block中保存tag value相关信息。
- Measurement Block：Measurement Block保存所有的Measurement信息，通过Measurement可以获取到Tag Block的地址。
- SeriesID Set
- TombstoneSeriesIDSet：
- SeriesSketch：HyhperLogLog++结构记数
- TombstoneSketch：上同
- Trailer：记录每个Block的的地址和大小

其中最为重要的是Measurement Block和Tag Blocks，先从Measurement Block看起，因为Measurement Block中有Tag Block的索引。先来看Measurement Block的结构

| Measurements | HashIndex | Measurements Sketch | Measurements TombStone Sketch | Trailer |
| ------------ | --------- | ------------------- | ----------------------------- | ------- |
|              |           |                     |                               | 66 Byte |

同样Trailer包含各个模块的地址和大小。Measurements中包含多个Measurement，通过HashIndex来确定Measurement的地址，key为Measurement name。Measurements Sketch和Measurements TombStone Sketch用来记数统计使用。先通过Trailer定取到HashIndex的地址和大小，读出HashIndex再根据我们提供的name去取Measurement。先来看HashIndex的格式

| Measurements count | Measurement 1 offsize | Measurement 2 offsize... |
| :----------------: | :-------------------: | :----------------------: |
|       8 Byte       |        8 Byte         |                          |

通过Measurement name经过hash计算就可以得到Measurement的地址，接下里继续看Measurement的结构

| Flag | Tag block offsize | Tag block Size | Measurement name size | Measurement name | Series count | Series Data Size | Series Data |
| :--: | :---------------: | :------------: | :-------------------: | :--------------: | :----------: | :--------------: | ----------- |
|  1   |         8         |       8        |           -           |        -         |      -       |        -         | -           |

Flag表示Series Data 的编码方式，分别是变长varint编码或者Roaring bitmap方式，Tag block offsize指的是该Measurement对应的Tag Block的地址。Measurement name size、Series count、Series Data Size用varint方式编码。Series count指的是当前Measurement在此文件中的Series Id 个数。

