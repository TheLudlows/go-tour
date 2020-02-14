#### TSI 

默认配置下索引信息存放于内存中，`index-version="tsi1"`配置作用是将索引信息持久化至硬盘，TSI索引主要包含tsl、tsi文件，其实Series 文件也算。tsl、tsi文件位于在每个shard中`index`目录下。

##### TSL

##### TSI

当TSL文件大小达到配置的compaction阈值时（由配置文件中的max-index-log-file-size指定，默认为1M），TSL文件会compaction成TSI文件，TSI文件算是存储格式最为复杂的。

| Magic  | Tag Set Blocks | Measurement Block | SeriesID Set | TombstoneSeriesIDSet | SeriesSketch | TombstoneSketch | Trailer |
| :----: | :------------: | :---------------: | :----------: | :------------------: | ------------ | --------------- | ------- |
| 4 Byte |                |                   |              |                      |              |                 | 82 Byte |



