# Go-Redis

用go实现redis,参考 https://github.com/HDT3213/godis

```
_________            ________     _____________        
__  ____/_____       ___  __ \__________  /__(_)_______
_  / __ _  __ \________  /_/ /  _ \  __  /__  /__  ___/
/ /_/ / / /_/ //_____/  _, _//  __/ /_/ / _  / _(__  ) 
\____/  \____/       /_/ |_| \___/\__,_/  /_/  /____/

v1.0.0
````

## 1. Echo服务器
接受客户端连接并将客户端发送的内原样传回客户端.

### 1.1 修改内容
```shell
tcp
├─ echo.go # client端
├─ echo_test.go # 测试
└─ server.go # server端
```

## 2. Redis协议解析器

### 2.1 Redis 通信协议

Redis 自 2.0 版本起使用了统一的协议 `RESP (REdis Serialization Protocol)`，该协议易于实现， 计算机可以高效的进行解析且易于被人类读懂. RESP 是一个二进制安全的文本协议，工作于 TCP
协议上. RESP 以行作为单位，客户端和服务器发送的命令或数据一律以 `\r\n` (CRLF) 作为换行符.

二进制安全是指允许协议中出现任意字符而不会导致故障. 比如 C 语言的字符串以 `\0` 作为结尾不允许字符串中间出现`\0`, 而 Go 语言的 string 则允许出现 `\0`，我们说 Go 语言的 string 是二进制安全的，而
C 语言字符串不是二进制安全的. RESP 的二进制安全性允许我们在 key 或者 value 中包含 `\r` 或者 `\n` 这样的特殊字符. 在使用 redis 存储 protobuf、msgpack
等二进制数据时，二进制安全性尤为重要.

RESP 定义了5种格式:

1. 简单字符串(Simple String): 服务器用来返回简单的结果，比如"OK". 非二进制安全，且不允许换行.
2. 错误信息(Error):  服务器用来返回简单的错误信息，比如"ERR Invalid Synatx". 非二进制安全，且不允许换行.
3. 整数(Integer): llen、scard 等命令的返回值, 64位有符号整数
4. 字符串(Bulk String): 二进制安全字符串, 比如 get 等命令的返回值
5. 数组(Array, 又称 Multi Bulk Strings): Bulk String 数组，客户端发送指令以及 lrange 等命令响应的格式

RESP 通过第一个字符来表示格式:

- 简单字符串: 以 `+` 开始， 如: `+OK\r\n`
- 错误: 以 `-` 开始，如: `-ERR Invalid Synatx\r\n`
- 整数: 以 `:` 开始，如: `:1\r\n`
- 字符串: 以 `$` 开始
- 数组: 以 `*` 开始

例子:

```shell
> 二进制安全
$4
a\r\nb

> nil
$-1 表示 nil, 比如使用 get 命令查询一个不存在的key时, 响应即为$-1

> Array格式: * + 数组长度
*2
$3
foo
$3
bar

> SET命令: 
*3
$3
SET
$3
key
$5
value
```

### 2.2 修改内容
```shell
go-redis
|── interface
│ └── redis
│     └── reply.go # redis响应体接口定义
├── lib
│ ├── logger
│ │ ├── file.go # 日志文件相关操作
│ │ └── logger.go # 日志配置类
│ └── utils
│     └── utils.go # 工具类, 包含BytesEquals判断方法
└── redis
  ├── parser
  │ ├── parser.go # 协议转换
  │ └── parser_test.go # 协议转换测试类
  └── reply
      ├── consts.go # 常用的响应体定义
      └── reply.go # 响应体信息
```