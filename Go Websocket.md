

# 🚀 Go WebSocket Server: Technical Concepts & Code

---

## 1. WebSockets: Real-Time Communication 🗣️


WebSockets provide a persistent, full-duplex communication channel between client and server over a single TCP connection. This enables real-time data exchange, making them ideal for chat, live feeds, collaborative tools, and financial dashboards.


**Key Features:**
- 🔄 Bi-directional: Both client and server can send/receive messages at any time.
- ⚡ Low latency: No HTTP request/response overhead after handshake.
- 📢 Suitable for broadcasting, subscriptions, and real-time updates.

---


## 2. Go WebSocket Server Structure 🏗️


This Go server uses the `golang.org/x/net/websocket` package to manage multiple WebSocket clients, broadcast messages, and send periodic data feeds. It is designed for real-time applications where multiple clients may need to receive the same or different data streams.


### Server State
```go
// Server manages all active WebSocket client connections.
type Server struct {
    Connection map[*websocket.Conn]bool // Map of active WebSocket connections.
}

// NewServer initializes and returns a new Server instance with an empty connection map.
func NewServer() *Server {
    return &Server{
        Connection: make(map[*websocket.Conn]bool),
    }
}
```

**Technical Notes:**
- The server maintains a map of all active WebSocket connections for broadcasting and management.
- ⚠️ Go maps are not safe for concurrent writes. Use `sync.Mutex` or `sync.RWMutex` for thread safety in production.

---


## 3. WebSocket Endpoints and Handlers 🎯


### WebSocket Endpoints

- `/ws` — General message exchange and broadcasting.
- `/orderbookfeed` — Periodic data feed (e.g., for order book updates).

```go
http.Handle("/ws", websocket.Handler(server.handleWSConnection)) // WebSocket handler for /ws
http.Handle("/orderbookfeed", websocket.Handler(server.handleWSOrderBook)) // WebSocket handler for /orderbookfeed
```

**Technical Notes:**
- `websocket.Handler` wraps a function to handle WebSocket connections, similar to `http.HandleFunc` for HTTP.
- Each endpoint can have custom logic for different data streams or client groups.

---


## 4. Broadcasting & Subscription Feeds 📡


### Broadcasting Messages 📢
```go
// Broadcast sends a message to all connected WebSocket clients.
func (s *Server) Broadcast(b []byte) {
    for ws := range s.Connection {
        go func(ws *websocket.Conn) {
            // Use a goroutine to avoid blocking if a client is slow or unresponsive.
            if _, err := ws.Write(b); err != nil {
                fmt.Println("write error: ", err)
            }
        }(ws)
    }
}
```

**Technical Notes:**
- Broadcasting is essential for chat, notifications, and real-time dashboards.
- Each send is done in a goroutine to ensure non-blocking operation.

---


### Periodic Data Feeds (Subscription) ⏰
```go
// handleWSOrderBook sends periodic data to the client every 2 seconds.
func (s *Server) handleWSOrderBook(ws *websocket.Conn) {
    fmt.Println("New connection to orderbook feed: ", ws.RemoteAddr().String())
    for {
        payload := fmt.Sprintf("orderbook data: %v\n", time.Now().Local())
        ws.Write([]byte(payload))
        time.Sleep(time.Second * 2)
    }
}
```

**Technical Notes:**
- Useful for live market data, telemetry, or any time-based updates.
- All clients connected to `/orderbookfeed` receive the same data at the same interval.

---


## 5. Reading & Responding: The Conversation Loop 🔄


### Reading and Responding to Client Messages 🔄
```go
// readLoop reads messages from a client, broadcasts them, and replies.
func (s *Server) readLoop(ws *websocket.Conn) {
    buffer := make([]byte, 1024)
    for {
        n, err := ws.Read(buffer) // Read data from the WebSocket connection
        if err != nil {
            if err == io.EOF {
                break // Client closed the connection.
            }
            fmt.Println("error in reading websocket connection: ", err)
            continue // Continue to keep the connection alive for recoverable errors.
        }
        msg := buffer[:n]
        fmt.Println(string(msg))
        s.Broadcast(msg) // Broadcast the received message to all connected clients
        ws.Write([]byte("Thank you for the message!!!")) // Reply to the sender
    }
}
```

**Technical Notes:**
- Handles incoming messages, logs them, broadcasts to all clients, and sends an acknowledgment.
- If the client disconnects, the loop exits for that connection.

---


## 6. Main Function: The Party Starts Here! 🎈


### Main Function: Server Startup 🏁
```go
func main() {

    server := NewServer()

    http.Handle("/ws", websocket.Handler(server.handleWSConnection))
    http.Handle("/orderbookfeed", websocket.Handler(server.handleWSOrderBook))

    http.ListenAndServe(":1337", nil)
}
```

**Technical Notes:**
- The server listens on port 1337 and registers WebSocket handlers for `/ws` and `/orderbookfeed`.
- You can add more HTTP or WebSocket endpoints as needed for your application.

---



## 7. Best Practices & Technical Tips 🧠

- 🦺 **Go maps are not safe for concurrent writes!** Use `sync.Mutex` or `sync.RWMutex` for thread safety.
- 🕹️ **WebSocket handlers** maintain persistent connections for real-time communication.
- 🏃 **Broadcasting** in goroutines ensures non-blocking, scalable message delivery.
- ⏰ **Periodic feeds** are ideal for live data (e.g., stock tickers, telemetry, order books).
- 🧹 **Resource management:** Always close connections and clean up resources when clients disconnect.
- 🔒 **Security:** Consider authentication, authorization, and input validation for production systems.

---





**Technical Notes:**
- Add HTTP endpoints for status, health checks, or documentation as needed.

---



## 8. Compatibility Matrix 🧩

- 🟢 Works with browser WebSocket clients and browser-based API tools (e.g., Hoppscotch).
- 🔴 Not compatible with Node.js `ws` clients (protocol mismatch). For universal compatibility, use `github.com/gorilla/websocket` in Go.

---



## 9. Summary & Next Steps 🚦

This Go WebSocket server demonstrates real-time, bidirectional communication, message broadcasting, and periodic data feeds. The code and technical notes above provide a complete, self-contained reference for building scalable, real-time systems in Go.

**Next steps:**
- Add authentication and authorization for secure communication 🔐
- Use `gorilla/websocket` for broader client compatibility 🌐
- Implement resource cleanup and error handling for production readiness 🧹
- Extend with custom endpoints for your application's needs 🚀