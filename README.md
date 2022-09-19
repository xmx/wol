# WoL

Wake on LAN 用 go 实现的网络唤醒功能

## 例子

```go
package main

import (
	"net"
	
	"github.com/xmx/wol"
)

func main() {
	mac, _ := net.ParseMAC("E4-54-E8-A0-3D-67")
	wol.Wake(mac)
}
```

