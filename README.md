# log4go
Go log module based on Uber's zap and lumberjack

[中文文档](README_ZH.md)



#### Usage

```go
package main

import (
  "github.com/wanghuiyt/log4go"
)

func main() {
  logger.Info("This is a message")
  logger.Infow("failed to fetch URL",
    "url", "example.com",
    "attempt", 3,
    "backoff", 1,
  )
  logger.Warn("This is a warning message")
  logger.Error("This is an error message")
}
```



#### Configuration

Make sure `log4go.yml` exists in the project directory, If it doesn't exist, the program will throw an exception `The system cannot find the file specified`.

>  The value of the configuration file is not case sensitive.

This is the explanation of the configuration:

log4go.yml

```yaml
LOG4GO:
    # Log record format, can be text mode or Json.
    # The value can only be `Text` or `Json`.
    FORMAT: Json
    # If it is `contain`, the error information will also be included in the info file
    # The value can only be `contain` or `independent`.
    LEVEL_MODE: independent
    # true, flase
    LEVEL_COLOR: true
    INFO:
        # 文件路径
        FILE_PATH_NAME: info.log
        # 滚动的最大数据量(MB)
        MAXSIZE: 50
        # 最大备份数量
        MAXBACKUP_COUNT: 10
        # 最大天数(day)
        MAXAGE: 28
        # 是否启用压缩
        COMPRESS: true
    ERROR:
        # 文件路径
        FILE_PATH_NAME: error.log
        # 滚动的最大数据量(MB)
        MAXSIZE: 50
        # 最大备份数量
        MAXBACKUP_COUNT: 10
        # 最大天数(day)
        MAXAGE: 28
        # 是否启用压缩
        COMPRESS: true
```



