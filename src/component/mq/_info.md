## Redis Stream
也算一种消息队列，但正经人不会在生产环境使用（除了yozo）.

## Consumer
!!!: 建议使用 goroutine池(e.g. ants) 来处理收到的消息，以免造成消息堆积.


