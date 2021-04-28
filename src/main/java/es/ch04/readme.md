### 查询


#### URL search
```shell

GET /m/_search?q=xx&df=xx&sort=year:desc&from=0&size=1&timeout=1s
{
  "profile": true
}
```
q:指定查询语句，使用query string syntax
df默认字段，不指定时，会对所有字典进行查询
sort:排序 from/size用于分页
profile 可以查看查询是如何被执行的


```shell
GET /m/_search?q=2021&df=title

GET /m/_search?q=2021
所有字段

GET /m/_search?q=title:2021

```


### Query String syntax(1)

* 指定字段 Vs 泛查询

    - q=title:2021 / q=2012 
    
* Term vs Phrase
    * Beautiful Mind 等效于Beautiful OR Mind
    * "Beautiful Mind" 等效于Beautiful And Mind Phrase查询，还要要求顺序保持一致
    
* 分组与引号
    * title:(Beautiful And Mind)
    * title="Beautiful Mind"
