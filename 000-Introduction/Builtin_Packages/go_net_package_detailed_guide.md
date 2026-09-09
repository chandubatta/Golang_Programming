# Go `net` Package — Detailed Guide

## 1. What is the `net` package?

The Go `net` package is part of the standard library and provides portable networking support.

```go
import "net"
```

It allows Go programs to:

- Create TCP clients and servers.
- Create UDP clients and servers.
- Resolve domain names and perform reverse DNS lookups.
- Work with IPv4 and IPv6 addresses.
- Inspect network interfaces.
- Work with Unix domain sockets.
- Configure network connection deadlines.
- Build lower-level networking tools and protocols.

The core abstractions include `net.Conn`, `net.Listener`, `net.PacketConn`, `net.TCPConn`, `net.UDPConn`, `net.IP`, `net.IPNet`, and `net.Resolver`.

---

# 2. Simple example — TCP server

```go
package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("Server listening on port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept error:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}

	fmt.Println("Client:", string(buffer[:n]))

	_, err = conn.Write([]byte("Hello from server!\n"))
	if err != nil {
		fmt.Println("Write error:", err)
	}
}
```

A client can connect using:

```bash
nc localhost 8080
```

The basic flow is:

```text
net.Listen()
      ↓
TCP Listener
      ↓
Accept()
      ↓
net.Conn
      ↓
Read()
      ↓
Write()
      ↓
Close()
```

This is the fundamental pattern behind many network servers.

---

# 3. `net` package architecture

Before learning individual functions, understand these important types.

| Type | Purpose |
|---|---|
| `net.Conn` | Generic network connection |
| `net.Listener` | Accepts incoming stream connections |
| `net.PacketConn` | Packet-oriented networking |
| `net.TCPConn` | TCP-specific connection |
| `net.TCPListener` | TCP-specific listener |
| `net.UDPConn` | UDP connection |
| `net.IP` | IP address |
| `net.IPNet` | IP network/CIDR |
| `net.IPAddr` | IP endpoint |
| `net.TCPAddr` | TCP endpoint |
| `net.UDPAddr` | UDP endpoint |
| `net.Interface` | Network interface |
| `net.Resolver` | DNS resolver |

---

# 4. Important `net` package functions

## `net.Dial`

Establishes a connection to a remote server.

```go
conn, err := net.Dial("tcp", "example.com:80")
```

Common network values include:

```text
tcp
tcp4
tcp6
udp
udp4
udp6
unix
unixgram
unixpacket
```

Example:

```go
conn, err := net.Dial("tcp", "example.com:80")
if err != nil {
	panic(err)
}
defer conn.Close()

fmt.Fprintf(conn, "GET / HTTP/1.0\r\nHost: example.com\r\n\r\n")
```

Conceptually:

```text
Client
   |
   | Dial()
   ↓
Server
```

---

## `net.DialTimeout`

Establishes a connection while limiting how long connection establishment can take.

```go
conn, err := net.DialTimeout(
	"tcp",
	"example.com:80",
	5*time.Second,
)
```

This is useful in production applications where a connection attempt should not wait indefinitely.

---

## `net.Listen`

Creates a network listener.

```go
listener, err := net.Listen("tcp", ":8080")
```

The listener waits for incoming clients:

```text
Server
   |
   ↓
Listen()
   |
   ↓
Waiting for clients
```

The listener is normally used with `Accept()`.

---

## `net.ListenPacket`

Creates a packet-oriented network listener, especially useful for UDP.

```go
conn, err := net.ListenPacket("udp", ":9000")
```

Unlike a stream listener, packet-oriented networking receives individual packets.

---

## `net.ListenIP`

Creates an IP-level listener.

```go
conn, err := net.ListenIP(
	"ip4:icmp",
	&net.IPAddr{},
)
```

This is lower-level than ordinary TCP/UDP programming and is useful for specialized networking applications.

---

## `net.DialIP`

Creates an IP-level connection.

```go
conn, err := net.DialIP(
	"ip4:icmp",
	nil,
	&net.IPAddr{IP: net.ParseIP("8.8.8.8")},
)
```

This is useful when working directly with IP protocols.

---

## `net.DialTCP`

Creates a TCP connection and returns a `*net.TCPConn`.

```go
conn, err := net.DialTCP(
	"tcp",
	nil,
	&net.TCPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 8080,
	},
)
```

This is useful when TCP-specific functionality is needed.

---

## `net.DialUDP`

Creates a UDP connection.

```go
conn, err := net.DialUDP(
	"udp",
	nil,
	&net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 9000,
	},
)
```

---

## `net.ListenTCP`

Creates a TCP-specific listener.

```go
listener, err := net.ListenTCP(
	"tcp",
	&net.TCPAddr{
		Port: 8080,
	},
)
```

This is useful when TCP-specific listener functionality is required.

---

## `net.ListenUDP`

Creates a UDP listener.

```go
conn, err := net.ListenUDP(
	"udp",
	&net.UDPAddr{
		Port: 9000,
	},
)
```

---

## `net.ListenMulticastUDP`

Creates a UDP listener for multicast traffic.

Conceptually:

```text
          Multicast Sender
                 |
        +--------+--------+
        |        |        |
       PC1      PC2      PC3
```

It is useful when multiple receivers need to receive the same UDP traffic.

---

# 5. DNS functions

## `net.LookupHost`

Resolves a hostname to host addresses.

```go
addresses, err := net.LookupHost("google.com")
```

---

## `net.LookupIP`

Looks up IP addresses for a hostname.

```go
ips, err := net.LookupIP("google.com")

for _, ip := range ips {
	fmt.Println(ip)
}
```

---

## `net.LookupAddr`

Performs reverse DNS lookup.

```go
names, err := net.LookupAddr("8.8.8.8")
```

Conceptually:

```text
IP address
    ↓
Reverse DNS
    ↓
Hostname
```

---

## `net.LookupCNAME`

Looks up a canonical name.

```go
name, err := net.LookupCNAME("www.example.com")
```

---

## `net.LookupMX`

Looks up mail-exchange records.

```go
records, err := net.LookupMX("example.com")

for _, record := range records {
	fmt.Println(record.Host, record.Pref)
}
```

Useful for understanding where email for a domain should be delivered.

---

## `net.LookupNS`

Looks up DNS name-server records.

```go
records, err := net.LookupNS("example.com")
```

---

## `net.LookupSRV`

Looks up SRV records.

```go
cname, records, err := net.LookupSRV(
	"_http",
	"tcp",
	"example.com",
)
```

SRV records are commonly used for service discovery.

---

## `net.LookupTXT`

Looks up TXT records.

```go
records, err := net.LookupTXT("example.com")
```

TXT records are often used for domain verification and email-related DNS configuration.

---

## `net.LookupPort`

Resolves a service name to a port number.

```go
port, err := net.LookupPort("tcp", "http")
```

Conceptually:

```text
"http"
  ↓
TCP service lookup
  ↓
80
```

---

# 6. Address functions

## `net.JoinHostPort`

Combines a host and port safely.

```go
address := net.JoinHostPort("127.0.0.1", "8080")
fmt.Println(address)
```

Output:

```text
127.0.0.1:8080
```

For IPv6:

```go
address := net.JoinHostPort("::1", "8080")
```

Result:

```text
[::1]:8080
```

This is safer than manually doing:

```go
host + ":" + port
```

because IPv6 addresses need brackets when combined with ports.

---

## `net.SplitHostPort`

Separates a host and port.

```go
host, port, err := net.SplitHostPort("127.0.0.1:8080")
```

Result:

```text
host = 127.0.0.1
port = 8080
```

It also correctly handles IPv6 addresses such as:

```text
[::1]:8080
```

---

# 7. IP address functions

## `net.ParseIP`

Converts a string into an IP address.

```go
ip := net.ParseIP("192.168.1.10")

if ip == nil {
	fmt.Println("Invalid IP")
	return
}

fmt.Println(ip)
```

It supports IPv4 and IPv6.

---

## `net.IPv4`

Creates an IPv4 address.

```go
ip := net.IPv4(192, 168, 1, 10)
```

---

## `net.ParseCIDR`

Parses CIDR notation.

```go
ip, network, err := net.ParseCIDR("192.168.1.0/24")
```

You receive:

```text
ip
network
```

For:

```text
192.168.1.0/24
```

the prefix length is 24 bits.

---

## `net.CIDRMask`

Creates an IP mask.

```go
mask := net.CIDRMask(24, 32)
```

Conceptually this represents:

```text
255.255.255.0
```

---

## `net.IPv4Mask`

Creates an IPv4 mask from four bytes.

```go
mask := net.IPv4Mask(255, 255, 255, 0)
```

---

# 8. Important `net.IP` methods

`net.IP` represents an IP address.

## `ip.String()`

Converts an IP to readable text.

```go
ip := net.ParseIP("192.168.1.10")
fmt.Println(ip.String())
```

---

## `ip.To4()`

Returns the IPv4 representation if the address can be represented as IPv4.

```go
v4 := ip.To4()
```

---

## `ip.To16()`

Returns a 16-byte representation.

```go
v6 := ip.To16()
```

---

## `ip.Equal(other)`

Compares two IP addresses.

```go
if ip1.Equal(ip2) {
	fmt.Println("Same IP")
}
```

Using `Equal` is preferable to assuming that two IP values have identical underlying byte representations.

---

## `ip.IsLoopback()`

Checks whether an address is a loopback address.

Examples include:

```text
127.0.0.1
::1
```

---

## `ip.IsPrivate()`

Checks whether the address belongs to a private IP range.

Common IPv4 private ranges include:

```text
10.0.0.0/8
172.16.0.0/12
192.168.0.0/16
```

---

## `ip.IsMulticast()`

Checks whether the IP address is multicast.

---

## `ip.IsUnspecified()`

Checks for unspecified addresses such as:

```text
0.0.0.0
::
```

---

## `ip.IsGlobalUnicast()`

Checks whether an address is classified as a global unicast address.

---

## `ip.IsLinkLocalUnicast()`

Checks for link-local unicast addresses.

For IPv4 this includes:

```text
169.254.x.x
```

---

## `ip.IsLinkLocalMulticast()`

Checks whether an address is link-local multicast.

---

## `ip.IsInterfaceLocalMulticast()`

Checks whether an address is interface-local multicast.

---

## `ip.Mask(mask)`

Applies a network mask.

```go
ip := net.ParseIP("192.168.1.100")
mask := net.CIDRMask(24, 32)

network := ip.Mask(mask)

fmt.Println(network)
```

---

## `ip.DefaultMask()`

Returns the default mask associated with an IPv4 address.

This is mainly relevant to traditional classful IPv4 behavior and is less important in modern CIDR-based networking.

---

# 9. `IPMask` methods

## `mask.Size()`

Returns:

```text
ones
bits
```

For a `/24` IPv4 mask:

```text
24
32
```

## `mask.String()`

Returns the textual representation of the mask.

---

# 10. `IPNet` methods

An `IPNet` represents an IP network.

Example:

```go
_, network, _ := net.ParseCIDR("192.168.1.0/24")
```

## `network.Contains(ip)`

Checks whether an IP belongs to the network.

```go
if network.Contains(net.ParseIP("192.168.1.50")) {
	fmt.Println("IP belongs to network")
}
```

Useful for:

- Access control.
- Firewall rules.
- Private-network detection.
- Routing logic.
- IP allowlists.

## `network.Network()`

Returns the network type name.

## `network.String()`

Returns a representation such as:

```text
192.168.1.0/24
```

---

# 11. TCP types

## `net.TCPAddr`

Represents a TCP endpoint.

```go
addr := &net.TCPAddr{
	IP:   net.ParseIP("127.0.0.1"),
	Port: 8080,
}
```

Important methods include:

```text
AddrPort()
Network()
String()
```

---

## `net.TCPConn`

Represents an established TCP connection.

Important methods include:

```text
Read()
Write()
Close()
LocalAddr()
RemoteAddr()
SetDeadline()
SetReadDeadline()
SetWriteDeadline()
SetReadBuffer()
SetWriteBuffer()
SetKeepAlive()
SetKeepAlivePeriod()
SetKeepAliveConfig()
SetNoDelay()
SetLinger()
CloseRead()
CloseWrite()
File()
SyscallConn()
MultipathTCP()
```

### `CloseRead()`

Closes the read side of a TCP connection.

### `CloseWrite()`

Closes the write side.

This allows half-close behavior:

```text
Client ---- data ----> Server
Client <---- response ---- Server

Client can finish sending while
still receiving data.
```

### `SetNoDelay(true)`

Controls TCP_NODELAY.

```go
conn.SetNoDelay(true)
```

This can reduce latency for applications that send small messages where waiting for packet aggregation is undesirable.

### `SetKeepAlive`

Controls TCP keep-alive behavior.

### `SetKeepAlivePeriod`

Controls keep-alive timing.

### `SetLinger`

Controls how a connection behaves around `Close()` and unsent data.

---

# 12. TCP listener

## `net.TCPListener`

A TCP-specific listener.

Important methods include:

```text
Accept()
AcceptTCP()
Addr()
Close()
SetDeadline()
File()
SyscallConn()
```

### `Accept()`

Accepts the next incoming connection.

### `AcceptTCP()`

Accepts and returns a `*TCPConn`.

---

# 13. UDP

UDP is fundamentally different from TCP.

TCP:

```text
Connection-oriented
Reliable
Ordered
Stream-based
```

UDP:

```text
Connectionless
No guaranteed delivery
No guaranteed ordering
Datagram-based
```

## `net.UDPAddr`

Represents a UDP endpoint.

Important methods:

```text
AddrPort()
Network()
String()
```

---

## `net.UDPConn`

Important methods include:

```text
Read()
Write()
ReadFrom()
ReadFromUDP()
ReadFromUDPAddrPort()
WriteTo()
WriteToUDP()
WriteToUDPAddrPort()
ReadMsgUDP()
ReadMsgUDPAddrPort()
WriteMsgUDP()
WriteMsgUDPAddrPort()
SetDeadline()
SetReadDeadline()
SetReadBuffer()
SetWriteDeadline()
File()
SyscallConn()
```

### `ReadFromUDP`

Reads a UDP datagram and gives you the sender's address.

```go
buffer := make([]byte, 1024)

n, addr, err := conn.ReadFromUDP(buffer)
```

### `WriteToUDP`

Sends data to a specific UDP address.

```go
_, err := conn.WriteToUDP(
	[]byte("hello"),
	addr,
)
```

This is fundamental for UDP servers.

---

# 14. `net.Conn`

`net.Conn` is one of the most important abstractions in the package.

It provides operations conceptually equivalent to:

```go
type Conn interface {
	Read([]byte) (int, error)
	Write([]byte) (int, error)
	Close() error
	LocalAddr() Addr
	RemoteAddr() Addr
	SetDeadline(time.Time) error
	SetReadDeadline(time.Time) error
	SetWriteDeadline(time.Time) error
}
```

Multiple goroutines may invoke methods on a `Conn` simultaneously.

For example:

```go
go readFromClient(conn)
go writeToClient(conn)
```

---

# 15. Deadlines — extremely important

A beginner may write:

```go
conn.Read(buffer)
```

and assume it will eventually return.

Network operations can block.

Use:

```go
conn.SetReadDeadline(
	time.Now().Add(5 * time.Second),
)
```

You can also use:

```go
conn.SetWriteDeadline(...)
```

or:

```go
conn.SetDeadline(...)
```

A zero `time.Time` removes the deadline.

Deadlines are particularly important for production network services.

---

# 16. `net.Interface`

Represents a machine's network interface.

Examples include:

```text
eth0
en0
Wi-Fi
Loopback
```

You can retrieve interfaces with:

```go
interfaces, err := net.Interfaces()
```

Then:

```go
for _, iface := range interfaces {
	fmt.Println(iface.Name)
	fmt.Println(iface.HardwareAddr)
	fmt.Println(iface.Flags)
}
```

Important fields include:

```text
Index
MTU
Name
HardwareAddr
Flags
```

---

## `net.InterfaceByName`

```go
iface, err := net.InterfaceByName("eth0")
```

Finds a network interface by name.

---

## `net.InterfaceByIndex`

```go
iface, err := net.InterfaceByIndex(1)
```

Finds a network interface by index.

---

## `iface.Addrs()`

Gets addresses assigned to the interface.

```go
addrs, err := iface.Addrs()
```

---

## `iface.MulticastAddrs()`

Gets multicast addresses associated with the interface.

---

## `net.InterfaceAddrs`

Gets system network addresses.

```go
addrs, err := net.InterfaceAddrs()
```

Unlike `Interface.Addrs`, this does not tell you which interface each address belongs to.

---

# 17. Unix domain sockets

The `net` package also supports Unix domain sockets.

Example:

```go
conn, err := net.Dial(
	"unix",
	"/tmp/myapp.sock",
)
```

Unix sockets are useful for communication between processes on the same machine.

Conceptually:

```text
Process A
    |
    | Unix socket
    |
Process B
```

Important types include:

```text
UnixAddr
UnixConn
```

and functions such as:

```text
DialUnix
ListenUnixgram
ResolveUnixAddr
```

---

# 18. `net.Pipe`

`net.Pipe()` creates two connected in-memory network connections.

```go
client, server := net.Pipe()

go func() {
	client.Write([]byte("hello"))
}()

buffer := make([]byte, 100)

n, _ := server.Read(buffer)

fmt.Println(string(buffer[:n]))
```

Conceptually:

```text
Connection A
     ↕
in-memory pipe
     ↕
Connection B
```

It is particularly useful for testing code that expects a `net.Conn`.

---

# 19. `net.Resolver`

`net.Resolver` provides DNS resolution using `context.Context`.

```go
ctx := context.Background()

resolver := net.Resolver{}

ips, err := resolver.LookupHost(
	ctx,
	"example.com",
)
```

Resolver methods include:

```text
LookupAddr
LookupCNAME
LookupHost
LookupIP
LookupIPAddr
LookupMX
LookupNS
LookupNetIP
LookupPort
LookupSRV
LookupTXT
```

Context-aware resolution is useful when you need cancellation or time limits.

Example:

```go
ctx, cancel := context.WithTimeout(
	context.Background(),
	2*time.Second,
)
defer cancel()

ips, err := net.DefaultResolver.LookupHost(
	ctx,
	"example.com",
)
```

---

# 20. `net.ListenConfig`

`ListenConfig` provides more control over creating listeners.

```go
var lc net.ListenConfig

listener, err := lc.Listen(
	context.Background(),
	"tcp",
	":8080",
)
```

Important methods include:

```text
Listen()
ListenPacket()
MultipathTCP()
SetMultipathTCP()
```

This becomes useful when building more configurable network servers.

---

# 21. `net.PacketConn`

`PacketConn` is the generic interface for packet-oriented communication.

Its operations conceptually include:

```text
ReadFrom
WriteTo
Close
LocalAddr
SetDeadline
SetReadDeadline
SetWriteDeadline
```

UDP is the most common implementation.

---

# 22. Network errors

Networking can produce many types of errors.

Important types include:

```text
AddrError
DNSConfigError
DNSError
InvalidAddrError
OpError
ParseError
```

## `net.OpError`

`net.OpError` provides information about a failed network operation, such as:

```text
Operation
Network
Source
Destination
Underlying error
```

For example:

```text
dial tcp 127.0.0.1:8080: connection refused
```

The underlying error can be inspected with standard Go error handling:

```go
if errors.Is(err, net.ErrClosed) {
	// connection was closed
}
```

`net.ErrClosed` represents network I/O on an already-closed connection or a connection closed concurrently.

---

# 23. `net.Buffers`

`net.Buffers` represents multiple byte slices.

```go
buffers := net.Buffers{
	[]byte("Hello "),
	[]byte("world"),
	[]byte("!"),
}
```

It can be useful for efficient batch writing on supported systems.

This becomes particularly relevant in performance-oriented networking code.

---

# 24. Address types

The package contains several important address types:

```text
Addr
IPAddr
TCPAddr
UDPAddr
UnixAddr
```

They represent different kinds of network endpoints.

For example:

```text
TCPAddr
127.0.0.1:8080
```

```text
UDPAddr
127.0.0.1:9000
```

```text
UnixAddr
/tmp/app.sock
```

The common `Addr` abstraction provides:

```go
type Addr interface {
	Network() string
	String() string
}
```

---

# 25. Three common beginner mistakes

## Mistake 1: Assuming TCP preserves message boundaries

Suppose you send:

```go
conn.Write([]byte("Hello"))
conn.Write([]byte("World"))
```

A beginner may assume the receiver will perform:

```text
Read → "Hello"
Read → "World"
```

That is not guaranteed.

TCP is a byte stream, not a message protocol.

You might receive:

```text
"HelloWorld"
```

or:

```text
"Hel"
"loWor"
"ld"
```

### How to avoid it

Design an application-level protocol.

Examples:

```text
[length][message]
```

or:

```text
message\n
```

or use a structured protocol such as HTTP.

---

## Mistake 2: Forgetting deadlines

This can be dangerous:

```go
conn.Read(buffer)
```

A network operation can remain blocked.

### Better

```go
conn.SetReadDeadline(
	time.Now().Add(5 * time.Second),
)
```

Use deadlines or context-aware higher-level APIs where appropriate.

---

## Mistake 3: Assuming UDP behaves like TCP

UDP does not guarantee:

```text
delivery
ordering
duplicate prevention
connection reliability
```

So:

```go
conn.WriteToUDP(...)
```

does not mean that the receiver definitely received the data.

### How to avoid it

If reliability matters:

- Use TCP.
- Build an application-level acknowledgment/retry mechanism.
- Use an appropriate higher-level protocol.

---

# 26. Real-world application #1 — TCP application server

Imagine an internal microservice architecture:

```text
             API Gateway
                  |
                  | TCP
                  ↓
          +---------------+
          | User Service  |
          +---------------+
                  |
                  | TCP
                  ↓
          +---------------+
          | Order Service |
          +---------------+
```

The services can communicate through TCP connections.

The `net` package provides low-level primitives such as:

```go
net.Listen()
net.Dial()
net.Conn
```

A higher-level protocol can then be built on top:

```text
TCP
 ↓
Length-prefixed messages
 ↓
JSON / Protobuf
 ↓
Application logic
```

---

# 27. Real-world application #2 — UDP monitoring/discovery

Imagine thousands of machines inside a data center.

Each machine periodically sends:

```text
"I am alive"
```

using UDP.

```text
Machine A ──┐
Machine B ──┤
Machine C ──┼──> Monitoring Server
Machine D ──┤
Machine E ──┘
```

UDP can be useful because:

- Messages are small.
- Low latency is useful.
- Losing an occasional heartbeat may be acceptable.
- You may not need thousands of persistent TCP connections.

`net.UDPConn` provides the packet operations needed to build this kind of system.

---

# 28. Three progressively challenging exercises

## Exercise 1 — TCP Echo Server

Build a TCP server using `net.Listen`.

Requirements:

1. Listen on port `8080`.
2. Accept multiple clients.
3. For every client, launch a goroutine.
4. Read data from the client.
5. Send exactly the received data back.
6. Close the connection when the client disconnects.
7. Print the client's remote address.
8. Do not use `net/http`.

---

## Exercise 2 — UDP Message Server

Build a UDP server.

Requirements:

1. Listen on UDP port `9000`.
2. Receive datagrams from multiple clients.
3. Print:
   - sender IP
   - sender port
   - received message
4. Send an acknowledgment back to the sender.
5. Add a read deadline.
6. Correctly handle timeout errors.

The goal is to understand the difference between TCP streams and UDP datagrams.

---

## Exercise 3 — Mini TCP Protocol

Build a small TCP-based key/value server.

A client should be able to send commands such as:

```text
SET name Chandu
GET name
DELETE name
GET name
```

Requirements:

1. Accept multiple clients concurrently.
2. Maintain an in-memory key/value store.
3. Define your own message framing protocol.
4. Handle multiple commands over the same TCP connection.
5. Add connection deadlines.
6. Handle clients disconnecting unexpectedly.
7. Prevent concurrent access problems to the shared map.
8. Return meaningful error responses.
9. Log the client's remote address.
10. Gracefully shut down the server.

The important challenge is not just using `net.Listen` and `net.Conn. The real challenge is designing an application protocol on top of TCP.

---

# 29. Recommended learning order

Although the `net` package has a large API surface, do not try to memorize everything immediately.

## Beginner

```text
net.Listen
net.Accept
net.Dial
net.Conn.Read
net.Conn.Write
net.Conn.Close
```

## Then UDP

```text
net.ListenPacket
net.ListenUDP
net.DialUDP
net.UDPConn.ReadFromUDP
net.UDPConn.WriteToUDP
```

## Then DNS

```text
net.LookupHost
net.LookupIP
net.LookupAddr
net.LookupCNAME
net.LookupMX
net.LookupTXT
```

## Then IP networking

```text
net.ParseIP
net.ParseCIDR
net.CIDRMask
net.IP.To4
net.IP.To16
net.IP.Equal
net.IPNet.Contains
```

## Then production networking

```text
SetDeadline
SetReadDeadline
SetWriteDeadline
net.Resolver
net.ListenConfig
net.TCPConn.SetKeepAlive
net.TCPConn.SetNoDelay
```

## Finally advanced networking

```text
IPConn
PacketConn
UnixConn
multicast
SyscallConn
Buffers
multipath TCP
```

---

# 30. A useful mental model

Think about networking in layers:

```text
                    Your Application
                           │
                    Application Protocol
                  JSON / Protobuf / HTTP
                           │
                    TCP or UDP
                    │           │
                   TCP         UDP
                    │           │
                net.Conn    PacketConn
                    │           │
                 net package
                    │
             Operating System
                    │
             Network Interface
                    │
                Ethernet/Wi-Fi
                    │
                 Internet
```

The `net` package gives you building blocks around the TCP/IP and socket layer.

For example, higher-level packages such as `net/http` use networking abstractions to implement HTTP.

---

# 31. `net` vs `net/http`

Do not confuse:

```text
net
```

with:

```text
net/http
```

`net` provides relatively low-level networking.

For example:

```go
conn, _ := net.Dial("tcp", "example.com:80")
```

You are responsible for speaking the protocol.

With:

```go
http.Get("https://example.com")
```

the HTTP package handles HTTP concepts for you.

Conceptually:

```text
net
 ↓
TCP/UDP/DNS/Sockets
```

while:

```text
net/http
 ↓
HTTP
 ↓
net
 ↓
TCP
```

Understanding `net` first makes it easier to understand what higher-level networking packages are doing.

---

# 32. Thought-provoking question

Suppose you have **10,000 clients connecting to your Go TCP server simultaneously**.

You could create one goroutine per connection, but imagine that some clients stay connected for hours without sending any data.

**What problems could this create in terms of memory, file descriptors, goroutine management, connection timeouts, and server capacity — and how would you design your `net`-based server to remain reliable under that workload?**

This question moves you from simply knowing `net` functions to thinking about real-world network-server architecture.

---

## Official documentation

The official Go `net` package documentation is available at:

https://pkg.go.dev/net
