# CHINA ID(中国大陆身份证)

[![Build Status](https://travis-ci.org/sleagon/chinaid.svg?branch=master)](https://travis-ci.org/sleagon/chinaid)  [![Go Report Card](https://goreportcard.com/badge/github.com/sleagon/chinaid)](https://goreportcard.com/report/github.com/sleagon/chinaid)  [![GoDoc](https://godoc.org/github.com/jinzhu/gorm?status.svg)](https://godoc.org/github.com/sleagon/chinaid)  [![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)


> 校验、解析中国大陆身份证号

## 身份证号校验

```go
package main
import (
    "log"
    "github.com/sleagon/chinaid"
)

func main() {
    id := chinaid.IDCard("420683199006041237")
    result := id.Valid()
    log.Println(">>>>", result)
}
```


## 身份证信息解析

```go
package main
import (
    "log"
    "github.com/sleagon/chinaid"
)

func main() {
    id := chinaid.IDCard("420683199006041237")
    result, err := id.Decode()
    if err != nil {
        log.Println("非法身份证号")
        return
    }
    log.Println(">>>>", result)
}
```

## 结果示例

```json
{
    "sex":       1,
    "code":      420683,
    "district":  "枣阳市",
    "city":      "襄阳市",
    "province":  "湖北省",
    "birthday":  "1990-06-04T00:00:00Z"
}
```

## 地域映射

身份证里的地域码往地域转换的映射表来自[中华人民共和国民政部][1]官网，本项目里目前用的版本是2019年4月更新的[版本](http://www.mca.gov.cn/article/sj/xzqh/2019/201901-06/201905271021.html)，后续会不定期更新。

## 依赖示例

```bash
go get github.com/sleagon/chinaid
```

> dep

```yml
[[constraint]]
   name = "github.com/sleagon/chinaid"
   version = "0.3"
```

[1]: http://www.mca.gov.cn/