# pG - Your Terminal's Network Heartbeat 🛰️

Ever wondered if you're *really* connected? Is it your app, your DNS, or the whole network? Stop guessing! 🕵️‍♂️

`pG` is a sleek, terminal-based utility that gives you a simple, at-a-glance status of your internet connection. It's like a heartbeat monitor for your connectivity, living right in your terminal. No more switching to a browser to check.

### ✨ Visual Status

**✅ CONNECTED**
```
┌──────────────────────────────────────────┐
│                                          │
│                                          │
│                 ┌█████┐                  │
│                 │█████│ (Green)          │
│                 └█████┘                  │
│                                          │
│                                          │
└──────────────────────────────────────────┘
```

**❌ DISCONNECTED**
```
┌──────────────────────────────────────────┐
│                                          │
│                                          │
│                 ┌█████┐                  │
│                 │█████│ (Red)            │
│                 └█████┘                  │
│                                          │
│                                          │
└──────────────────────────────────────────┘
```

---

## 🌟 Features

*   **🚦 Instant Visual Feedback:** A simple colored block tells you what you need to know. Green for GO, Red for NO.
*   **📡 Real-time Monitoring:** Continuously checks your connection at sub-second intervals so you know the moment it drops or comes back.
*   **📜 Detailed Logs:** Press `l` to toggle a log panel. See every check, its latency, and the HTTP status code. Perfect for troubleshooting flaky connections.
*   **⌨️ Keyboard First:** Designed for power users. Control everything with simple, intuitive key presses.
*   **⚡ Lightweight & Fast:** Written in pure Go, it's incredibly efficient and won't hog your system resources.

## 🚀 Getting Started

It couldn't be simpler.

1.  **Clone the repository:**
    ```bash
    git clone <your-repo-url>
    cd pG
    ```

2.  **Build the binary:**
    ```bash
    go build
    ```

3.  **Run it!**
    ```bash
    ./pG
    ```

## 🕹️ Controls

*   `q` or `Esc` or `Ctrl+C`: Quit the application.
*   `l`: Toggle the detailed log view on and off.

## 🔧 Make It Your Own

Want to ping a different server or change the check interval? It's easy!

Just tweak the `const` values at the top of `main.go` and recompile.

```go
const (
	checkURL        = "http://clients3.google.com/generate_204" // Change this URL
	checkInterval   = 500 * time.Millisecond                   // Change how often it checks
	checkTimeout    = 1000 * time.Millisecond                  // Change the timeout
	// ...
)
```

## 🛠️ Built With

*   **[Go](https://golang.org/)** - For the core application logic.
*   **[tview](https://github.com/rivo/tview)** - For the awesome TUI components.
*   **[tcell](https://github.com/gdamore/tcell)** - For low-level terminal control.

---
## 👤 Creator

Made with ❤️ by **b4sh0ne**.

Feel free to contribute!