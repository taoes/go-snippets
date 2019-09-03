# 基础语法


## 逻辑判断

## 常用结构体 

## CLI 参数解析

1. 官方提供的方式(暂未完成) [src/command-line/BaseCommandLine.go](src/command-line/BaseCommandLine.go)

## 函数

1. 延迟函数 [src/function/DelayFunction.go](src/function/DelayFunction.go)
2. 命名函数返回值 [src/function/funcAndRename.go](src/function/funcAndRename.go)
3. 返回函数 [src/function/modifyArgFunction.go](src/function/modifyArgFunction.go)

### 闭包(匿名函数)

+ 基础用法，返回一个匿名函数 [src/functional/closure.go](src/functional/closure.go)
+ 创建一个复杂的闭包示例 [src/functional/Sum.go](src/functional/Sum.go)
+ 闭包参数的记忆效应 [src/functional/Accumulate.go](src/functional/Accumulate.go)



# 并发
1. [线程间数据传输实例](src/concurrent/ChanExample.go)
2. 通道示例 [src/concurrent/ChanExample.go](src/concurrent/ChanExample.go)
3. CPU核数优化 [src/concurrent/CoreOptimizer.go](src/concurrent/CoreOptimizer.go)
4. Go 提供的锁机制 [src/concurrent/Matux.go](src/concurrent/Matux.go)
5. 性能更好地读写锁 [src/concurrent/RWMatux.go](src/concurrent/RWMatux.go)
6. IO 的多路复用简单示例 [src/concurrent/Select.go](src/concurrent/Select.go)
7. 等待组 阻塞线程 [src/concurrent/WaitGroup.go](src/concurrent/WaitGroup.go)



# IO

## 读写文件

+ 通过字节读文件 [src/io/file/ReadFile.go](src/io/file/ReadFile.go)
+ 通过字节写文件 [src/io/file/WriteFile.go](src/io/file/WriteFile.go)
+ 100个goroutine 并发写入文件 [src/io/file/ConcurrentWriteFile.go](src/io/file/ConcurrentWriteFile.go)

## JSON

### 官方提供的JSON工具

+ 序列化  [src/io/json/conversionObjectToJson.go](src/io/json/conversionObjectToJson.go)
+ 反序列化  [src/io/json/ConversionJsonToObject.go](src/io/json/ConversionJsonToObject.go)


# 网络

## 套接字

## Web 服务器

+ 官方提供的实例 [src/web/WebBase.go](src/web/WebBase.go)


