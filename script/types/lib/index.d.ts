
export {}

type obj = { [key: string]: any }
type bytearray = Array<number> | ArrayBuffer | Uint8Array

interface _Event<T extends obj> {
	readonly name: string
	readonly data: T
	readonly cancelable: boolean
	readonly canceled: boolean
	cancel(): void
}

declare global {
	const $: {
		ID: string
		VERSION: string
	} & EventEmitter

	const console: {
		trace(...args: any[]): void
		debug(...args: any[]): void
		log(...args: any[]): void
		info(...args: any[]): void
		warn(...args: any[]): void
		error(...args: any[]): void
	}

	function setInterval(handler: Function, timeout?: number, ...arguments: any[]): Readonly<Object>
	function setTimeout(handler: Function, timeout?: number, ...arguments: any[]): Readonly<Object>
	function setImmediate(handler: Function, ...arguments: any[]): Readonly<Object>
	function clearInterval(id: Readonly<Object> | undefined): void
	function clearTimeout(id: Readonly<Object> | undefined): void
	function clearImmediate(id: Readonly<Object> | undefined): void

	type Event<T extends obj = Object> = _Event<T> & T

	type EventListener<T extends obj = Object> = (event: Event<T>) => void

	interface EventEmitter {
		on<T extends obj = Object>(name: string, listener: EventListener<T>): this
		addListener<T extends obj = Object>(name: string, listener: EventListener<T>): this
		off<T extends obj = Object>(name: string, listener: EventListener<T>): this
		removeListeners<T extends obj = Object>(name: string, listener: EventListener<T>): this
		emit(name: string, cancelable?: boolean): boolean
		emit(name: string, data: obj, cancelable?: boolean): boolean
		removeAllListeners(name: string): void
		eventNames(): string[]
		listeners<T extends obj = Object>(name: string): EventListener<T>[]
	}

	interface Conn extends EventEmitter {
		readonly protocol: number
		readonly localAddr: string
		readonly remoteAddr: string
		close(): void
		newPacket(id: number): PacketBuilder
	}

	interface PacketBuilder {
		readonly protocol: number
		readonly id: number
		send(): Promise<void>
		bytearray(v: bytearray): this
		bool(v: boolean): this
		byte(v: number): this
		ubyte(v: number): this
		short(v: number): this
		ushort(v: number): this
		int(v: number): this
		long(v: number): this
		varInt(v: number): this
		varLong(v: number): this
		float(v: number): this
		double(v: number): this
		string(v: string): this
		uuid(v: string): this
		json(v: any): this
	}

	interface PacketReader {
		readonly protocol: number
		readonly id: number
		send(): Promise<void>
		bytearray(): Uint8Array
		bool(): boolean
		byte(): number
		ubyte(): number
		short(): number
		ushort(): number
		int(): number
		long(): number
		varInt(): number
		varLong(): number
		float(): number
		double(): number
		string(): string
		uuid(): string
		json<T>(): T
	}

	interface PlayerInfo {
		id: string
		name: string
	}

	interface HandshakePkt {
		protocol: number
		addr: string
		addition: string
		port: number
		nextState: number
	}

	type HandshakeEvent = Event<{
		client: Conn
		handshake: Readonly<HandshakePkt>
		target: string
	}>

	type ServeEvent = Event<{
		player?: PlayerInfo // undefined means serve for ping connection
		client: Conn
		server: Conn
		handshake: Readonly<HandshakePkt>
	}>

	type PacketEvent = Event<{
		conn: Conn
		packet: PacketReader
	}>

	type CloseEvent = Event<{
		conn: Conn
	}>

	type BeforeCloseEvent = Event<{
		conn: Conn
		error: string
	}>
}
