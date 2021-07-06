# log4go

基于 Uber 的 zap 和 lumberjack 的 Go 日志模块



### 安装

```sh
go get -u github.com/wanghuiyt/log4go
```



### 使用

example.go

```go
package main

import "github.com/wanghuiyt/log4go"

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

example file structure:

```
exampleProject
├─util
│    xx1.go
│    xx2.go
│
│ main.go
│ log4go.yml  # The configuration file named `log4go.yml` must exist
```



### Configuration

确保 `log4go.yml` 存在于项目目录, 如果不存在, 程序会抛出 `The system cannot find the file specified`异常.

>  YAML配置文件的值不区分大小写

配置文件的说明:

log4go.yml

```yaml
LOG4GO:
    # 日志记录格式, 可以是文本模式或Json.
    # 值只能是`Text`或者`Json`.
    FORMAT: Json
    # 如果为`contain`, 错误日志也会包含在info文件中.
    # 值只能是`contain`或者`independent`.
    LEVEL_MODE: independent
    # 如果为true, 则LEVEL的颜色显示在文件中.
    # 值只能是`true`或者`false`.
    LEVEL_COLOR: true
    INFO:
        # 级别大于 INFO 的日志文件路径. 
        # 如果 LEVEL_MODE 是`independent`，这个文件只记录 INFO 和 WARN 日志。
        # 示例: /var/log/info.log
        FILE_PATH_NAME: info.log
        # 日志文件在滚动之前的最大大小(以兆字节MB为单位)
        MAXSIZE: 50
        # 要保留的最大旧日志文件数
        MAXBACKUP_COUNT: 10
        # 根据文件名中编码的时间戳保留旧日志文件的最大天数
        MAXAGE: 28
        # 是否应使用 gzip 压缩轮换的日志文件
        COMPRESS: true
    ERROR:
    	# 级别大于或等于ERROR的文件路径。
        FILE_PATH_NAME: error.log
        MAXSIZE: 50
        MAXBACKUP_COUNT: 10
        MAXAGE: 28
        COMPRESS: true
```





