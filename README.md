# ids

一个golang的id生成器

### 安装

```shell
go get github.com/zituocn/ids
```

### 使用

```go
package main

import (
    "github.com/zituocn/ids"
)

func main(){
    id := ids.New()
    fmt.Println(id)
}
 
```


### 使用的第三方库
* github.com/rs/xid - 生成uuid