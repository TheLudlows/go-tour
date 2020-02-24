



#### TSI 

默认配置下索引信息存放于内存中，`index-version="tsi1"`配置作用是将索引信息持久化至硬盘，TSI索引主要包含logFile文件(.tsl结尾)、tsi文件，其实Series 文件也算。tsl、tsi文件位于在每个shard中`index`目录下。

##### LogFile

##### Index File

当TSL文件大小达到配置的compaction阈值时（由配置文件中的max-index-log-file-size指定，默认为1M），LogFile文件会compaction成TSI文件，TSI文件算是存储格式最为复杂的。

| Magic  | Tag Blocks | Measurement Block | SeriesID Set | TombstoneSeriesIDSet | SeriesSketch | TombstoneSketch | Trailer |
| :----: | :--------: | :---------------: | :----------: | :------------------: | ------------ | --------------- | ------- |
| 4 Byte |            |                   |              |                      |              |                 | 82 Byte |

- Magic：占用4Byte，1.79版为`TSI1`
- Tag Blocks：每个Measurement占用拥有一个Tag Block，Tag Block中保存tag value相关信息。
- Measurement Block：Measurement Block保存所有的Measurement信息，通过Measurement可以获取到该Measurement对应Tag Block的地址。
- SeriesID Set：采用Roaring Bitmap记录当前文件中的SeriesID
- TombstoneSeriesIDSet：删除的SeriesID
- SeriesSketch：HyhperLogLog++结构记数
- TombstoneSketch：上同
- Trailer：记录每个Block的的地址和大小

其中最为重要的是Measurement Block和Tag Blocks，因为Measurement Block中有Tag Block的索引。先来看Measurement Block的结构

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

通过Measurement Block拿到了Tag Block的地址，一个Measurement可能存在多个tag，也放在一个Tag block中。接下里看Tag Block的存储格式

| Tag Block Value Iterators | Tag Block key Iterator | Key Hash Index | Trailer |
| :-----------------------: | :--------------------: | :------------: | :-----: |
|                           |                        |                | 58 Byte |

TagBlock保存了某个Measurement的Tagkey和TagValue，前面我们根据Measurement Block获取到该Measurement的TagBlock的地址以及大小(size)，因此可以直接获取TagBlock的Trailer，同样该Trailer保存了各个模块的地址和大小，首先是Tag Block Value Iterators，我们知道一个tag key对应多个value，同样一个Measurement中也可能存在多个tag key。一个key对应的values 保存在一个Tag Block Value Iterator中。Tag Block Value Iterator是TagBlockValueElem的集合，（至于为什么这么命名，是因为源码中就是这么命名的）。接下里是Tag Block key Iterator，这个模块是tag key的集合，一个Tag key对应的结构称为TagBlockKeyElem，TagBlockKeyElem保存了某个key，以及key对应的Tag Block Value Iterator的地址以及大小等信息。Key Hash Index是对TagBlockKeyElem的索引，该模块的结构非常简单，和前面的索引结构相同，即[cout,off1,off2...]，每个值占用8byte。当然是通过key name来定位。拿到了TagBlockKeyElem的地址，接下里看TagBlockKeyElem的结构

| Flag   | Values Offsize | Values Size | Values index offSize | Values index Size | Key size | key Name |
| ------ | :------------: | :---------: | :------------------: | :---------------: | :------: | :------: |
| 1 Byte |     8Byte      |    8Byte    |        8Byte         |       8Byte       |    -     |    -     |

Flag表示是否删除，0表示未删除，Values Offsize表示key对应的Tag Block Value Iterator的地址，对于多个value如何快速的找到某个具体Value，这里就需要对value进行索引，这就是Values index的作用，每个Tag Block Value Iterator都一个ValueIndex。Values index offSize表示该Tag Block Value Iterator的地址。接下里是key size采用变长编码，key name就是 tag key。通过TagBlockKeyElem获取到了Tag Block Value Iterator的地址以及Values index offSize，我们可以通过Index获取到任意一个value的信息，接下里看Tag Block Value Iterator中的TagBlockValueElem结构

| Flag   | Value size | Value | Series count | Series size | Series data |
| ------ | :--------: | :---: | :----------: | ----------- | :---------: |
| 1 Byte |     -      |   -   |      -       | -           |      -      |

Flag表示该value对应的Series data的编码方式，Roaring bitmap或者原生varint编码，Value size、Series count、Series size采用varint编码。在讲述Measurement Block中也提到了Series data ，Measurement Block中的Series data保存了当前文件中该Measurement拥有的所有的Series ID（bitmap存储），而在TagBlockValueElem中的Series data 保存了该Value对应的Series ID。