# log4go
Go log module based on Uber's zap and lumberjack

[中文文档](README_ZH.md)



#### Install

```sh
go get -u github.com/wanghuiyt/log4go
```



#### Usage

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
    # If it is `contain`, the error information will also be included in the info file.
    # The value can only be `contain` or `independent`.
    LEVEL_MODE: independent
    # If true, the color of the level is displayed in the file.
    # The value can only be `true` or `flase`.
    LEVEL_COLOR: true
    INFO:
        # File path with level greater than INFO. 
        # If LEVEL_MODE is `independent`, this file only records INFO and WARN logs.
        # example: /var/log/info.log
        FILE_PATH_NAME: info.log
        # The maximum size in megabytes of the log file before it gets rotated.
        MAXSIZE: 50
        # The maximum number of old log files to retain.
        MAXBACKUP_COUNT: 10
        # The maximum number of days to retain old log files based on the timestamp encoded in their filename.
        MAXAGE: 28
        # Determines if the rotated log files should be compressed using gzip.
        COMPRESS: true
    ERROR:
    	# The file path with a level greater than or equal to ERROR.
        FILE_PATH_NAME: error.log
        MAXSIZE: 50
        MAXBACKUP_COUNT: 10
        MAXAGE: 28
        COMPRESS: true
```





