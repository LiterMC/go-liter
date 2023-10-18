
- [English](./README.MD)
- **中文**

# Script

该包实现了脚本管理器等功能, 目前暂不支持 TS

每个插件都应放在工作目录下的`plugins`文件夹内, 文件名以`.js`结尾并符合如下格式:
```regexp
^([a-z_][0-9a-z_]{0,31})(?:@(\d+(?:\.\d+)*))?(?:-.+)?\..+$
```
简单来说, 文件名分为3个部分:
1. 插件ID: 该部分必须使用小写字母或下划线开头, 仅能包含数字、小写字母或下划线. 长度在1~32之间
2. 插件版本 (可选): 该部分必须以 `@` 开头, 仅包含数字和小数点
3. 附加信息: 该部分以小数点(`.`)或横杠(`-`)开头, 直到文件名末尾

每个脚本拥有独立的魔法变量 `$`, 该变量提供脚本的基本信息, 如 `$.ID`、`$.VERSION`.  
`$.storage` 实现了 [Storage API](https://developer.mozilla.org/zh-CN/docs/Web/API/Storage) 目前仅在运行时有效, 暂不推荐使用.  
脚本可以使用 `$.on`、`$.emit` 等方法收发事件

### 日志

所有脚本日志都应该通过 `console` 对象输出.  
`trace`、`debug`、`info`/`log`、`warn`、`error` 6个方法代表了日志的5个级别, 级别过低的日志在没有明确配置的情况下不会被记录.

日志输出时会自动加上脚本ID作为前缀.

### 事件系统

`EventEmitter` 对象实现了 [NodeJS EventEmitter](https://nodejs.dev/en/learn/the-nodejs-event-emitter/) 的大部分功能.

`EventEmitter` 的监听器有且仅传入一个参数 `event`, `Event` 对象的定义如下:
```ts
class Event {
	name: string
	data: Object
	get cancelable(): bool
	get canceled(): bool
	cancel()
	// 也会将 data 对象的键值对展开到 Event 对象中
}
```

发送事件时, `emit(<name>, [<data>], [<cancelable>])` 会自动打包成 `Event` 对象

大部分对象都实现了 `EventEmitter`, 常见事件有:

| 所在对象 | 事件名       | 描述          | 撤销后果         |
|---------|-------------|--------------|-----------------|
| `$`     | `unload`    | 插件卸载事件   | 不可撤销         |
| `$`     | `handshake` | 客户端握手     | 阻止连接到服务端, 允许解析数据包并将触发 `packet` 事件 |
| `$`     | `serve`     | 服务端连接建立 | 阻止底层转发, 允许解析数据包并将触发 `packet` 事件 |
| `Conn`  | `packet`    | 连接收到数据包 | 阻止该数据包的转发 |
| `Conn`  | `error`     | 连接产生错误   | 不可撤销         |
| `Conn`  | `close`     | 连接断开      | 对等连接将不会自动断开 |

如果您想处理连接, 则必须在 `handshake` 或 `serve` 上调用 `cancel()` 方法, 否则 `packet` 事件不会触发
