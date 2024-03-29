
- [English](./README.MD)
- **中文**

# Script

该包实现了脚本管理器等功能.  
插件基于 [goja](https://github.com/dop251/goja) 运行, 然而 goja 并不支持全部的 ES6 特性, 请开发者注意.

每个脚本拥有独立的魔法变量 `$`, 该变量提供脚本的基本信息, 如 `$.ID`、`$.VERSION`.  
`$.storage` 实现了 [Storage API](https://developer.mozilla.org/zh-CN/docs/Web/API/Storage) 目前仅在运行时有效, 暂不推荐使用.  
脚本可以使用 `$.on`、`$.emit` 等方法收发事件

# 初始化

```sh
npm init glp@latest "<plugin-name>"
```

注: 同时会自动配置 Typescript

# 日志

所有脚本日志都应该通过 `console` 对象输出.  
`trace`、`debug`、`info`/`log`、`warn`、`error` 6个方法代表了日志的5个级别, 级别过低的日志在没有明确配置的情况下不会被记录.

日志输出时会自动加上脚本ID作为前缀.

# 常见事件

| 所在对象 | 事件名       | 描述          | 撤销后果         |
|---------|-------------|--------------|-----------------|
| `$`     | `unload`    | 插件卸载事件   | 不可撤销         |
| `$`     | `handshake` | 客户端握手     | 阻止连接到服务端, 允许解析数据包并将触发 `packet` 事件 |
| `$`     | `serve`     | 服务端连接建立 | 阻止底层转发, 允许解析数据包并将触发 `packet` 事件 |
| `Conn`  | `packet`    | 连接收到数据包 | 阻止该数据包的转发 |
| `Conn`  | `error`     | 连接产生错误   | 不可撤销         |
| `Conn`  | `close`     | 连接断开      | 不可撤销          |
| `Conn`  | `before_close` | 由于对等方连接断开导致该连接准备断开 | 该连接将不会断开 |

如果您想处理连接, 则必须在 `handshake` 或 `serve` 上调用 `cancel()` 方法, 否则 `packet` 事件不会触发
