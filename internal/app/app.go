package app

import (
    "fmt"

    "github.com/yourname/telecomborz/internal/crypto"
    "github.com/yourname/telecomborz/internal/protocol"
    "github.com/yourname/telecomborz/internal/transport/ip"
)

func Run() error {
    fmt.Println("telecomborz starting...")

    // 1. Load or generate identity
    id, err := crypto.GenerateIdentity()
    if err != nil {
        return err
    }
    fmt.Println("Identity public key:", id.Public)

    // 2. Start IP listener
    listener := ip.Listener{Addr: ":9000"}
    go func() {
        if err := listener.Start(func(msg protocol.WireMessage) {
            fmt.Printf("Received from %s: %x\n", msg.From, msg.Payload)
        }); err != nil {
            fmt.Println("listener error:", err)
        }
    }()

    // 3. For now, just block (later: REPL / UI)
    select {}
}
