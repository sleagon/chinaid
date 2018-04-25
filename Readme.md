# CHINA ID(中国大陆身份证)

## 身份证号校验

```go
package main
import (
    "log"
    "github.com/sleagon/chinaid"
)

func main() {
    id := chinaid.IDCard("420683199006041237")
    result := id.Check()
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
    "cityCode":    420683,
    "city":        "枣阳市",
    "province":    "湖北省",
    "sex": "男",
    "birthday":    "1990-06-04 00:00:00 +0000 UTC",
}
```