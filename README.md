# easyLogs

## Description

This repository is an alternate of golang default logger.Purpose is to provide a different experience with logging in golang with multiple customization options.

- **Vizualization** Provides multiple visualization style, suitable for both local and production envs. You can switch between colorful logs and json formatted logs.
- **Descriptive** Logs can be easily identified different logs based on severity. Caller function and line number will be helpful for debugging.
- **TraceEnabled** In case of error it is possible to turn on stack trace based on requirement.
- **Memstats** Printing memory stats and no of go routines can be turned on along with logs if required.

## Import

This package can be imported by `go get github.com/KIVUOS1999/easyLogs`

## Usage

Once imported it can be used directly use it with default options.

``` go
package main

import (
 "github.com/KIVUOS1999/easyLogs/pkg/log"
)

func main() {
    log.Debug("Debug", 2)
    log.Info("Info", 2)
 log.Warn("Warn", 2)
 log.Error("Error", 2)
    
    log.Debugf("Debug %d", 2)
    log.Infof("Info %+v", 2)
 log.Warnf("Warn %s ... ", "test")
 log.Errorf("Error %d", 2)

    log.ErrorWithTrace("Error this is")
}
```

### Default output

[Image]("https://1drv.ms/i/s!Aho-_IuMswcjgqtrBbixKEkH97N_Vw?embed=1&width=2528&height=516")

### Available levels

This package supports level for `Info`, `Debug`, `Warn`, `Error`. You can use a format specifier with `Infof`, `Debugf`, `Warnf`, `Errorf`. You can control the level of output by passing some params that is discussed here. There is another type log print `ErrorWithTrace` which will generate the stack trace.


## Initalization

If you want to trigger customization you can do it like below

### Viewing only upto certain log level

`log.Init(log_level)`

- For log_level `configs.Error` - Error logs will only come.
- For log_level `configs.Warn` - Error, Warn logs will come.
- For log_level `configs.Info` - Error, Warn, Info log will come.
- For log_level `configs.Debug` - Error, Warn, Info, Debug logs will come.

### Code Example

``` go
package main

import(
    "github.com/KIVUOS1999/easyLogs/pkg/configs"
 "github.com/KIVUOS1999/easyLogs/pkg/log"
)

func main(){
    log.Init(configs.Info)

    log.Debug("Debug", 2)
 log.Info("Info", 2)
 log.Warn("Warn", 2)
 log.Error("Error", 2)
}
```

``` log
2024-11-07 21:40:33 [ INFO  ] main.main()#4    Info 2
2024-11-07 21:40:33 [ WARN  ] main.main()#5    Warn 2
2024-11-07 21:40:33 [ ERROR ] main.main()#6    Error 2
```

As you have set the log level to info Debug logs will not be printed.

### Changing format of log

`log.Init("", configs.JsonLogs)` - You can either pass the 1st param of log level or it can be kept as empty. It will take default log level of Debug.

JsonFormat is usually preferred in production envs or envs where colored logs will be of no use such as logs of cloud watch.

```go
package main

import(
    "github.com/KIVUOS1999/easyLogs/pkg/configs"
 "github.com/KIVUOS1999/easyLogs/pkg/log"
)

func main(){
    log.Init("", configs.JsonLogs)

    log.Debug("Debug", 2)
 log.Info("Info", 2)
 log.Warn("Warn", 2)
 log.Error("Error", 2)
}
```

```json
{"type":"Debug","time":"2024-11-07 21:46:13","message":"Debug 2","caller_file":"main.go","caller_func_name":"main.main","line_number":21}
{"type":"Info","time":"2024-11-07 21:46:13","message":"Info 2","caller_file":"main.go","caller_func_name":"main.main","line_number":22}
{"type":"Warn","time":"2024-11-07 21:46:13","message":"Warn 2","caller_file":"main.go","caller_func_name":"main.main","line_number":23}
{"type":"Error","time":"2024-11-07 21:46:13","message":"Error 2","caller_file":"main.go","caller_func_name":"main.main","line_number":24}
```

### System related information

With this logger it is possible to get more system related information like memory usage, no of goroutines etc... Which is useful in dev like envs which can detect early memleaks. 

**Available System Stats**

- current_heap_alloc_MB - Current heap memory allocation.
- total_alloc_MB - Total memory allocated from start of time.
- current_stack_alloc_MB - Current stack memory allocation.
- sys_alloc_MB - Total memory allocated by system.
- total_garbage_collected - Total number of time garbage collector has ran.
- current_go_routine - current go routines used by program.

> **Note**: This is only available with json format logs.

```go
package main

import(
    "github.com/KIVUOS1999/easyLogs/pkg/configs"
 "github.com/KIVUOS1999/easyLogs/pkg/log"
)

func main(){
    log.Init("", configs.JsonLogs, true)

    log.Debug("Debug", 2)
 log.Info("Info", 2)
 log.Warn("Warn", 2)
 log.Error("Error", 2)
}
```

```json
{"type":"Debug","time":"2024-11-07 22:02:29","message":"Debug 2","caller_file":"main.go","caller_func_name":"main.main","line_number":21,"current_heap_alloc_MB":0.16,"current_stack_alloc_MB":0.28,"total_alloc_MB":0.16,"sys_alloc_MB":6.64,"total_garbage_collected":0,"current_go_routine":1}
{"type":"Info","time":"2024-11-07 22:02:29","message":"Info 2","caller_file":"main.go","caller_func_name":"main.main","line_number":22,"current_heap_alloc_MB":0.17,"current_stack_alloc_MB":0.28,"total_alloc_MB":0.17,"sys_alloc_MB":6.64,"total_garbage_collected":0,"current_go_routine":1}
{"type":"Warn","time":"2024-11-07 22:02:29","message":"Warn 2","caller_file":"main.go","caller_func_name":"main.main","line_number":23,"current_heap_alloc_MB":0.17,"current_stack_alloc_MB":0.28,"total_alloc_MB":0.17,"sys_alloc_MB":6.64,"total_garbage_collected":0,"current_go_routine":1}
{"type":"Error","time":"2024-11-07 22:02:29","message":"Error 2","caller_file":"main.go","caller_func_name":"main.main","line_number":24,"current_heap_alloc_MB":0.17,"current_stack_alloc_MB":0.28,"total_alloc_MB":0.17,"sys_alloc_MB":6.64,"total_garbage_collected":0,"current_go_routine":1}
```
