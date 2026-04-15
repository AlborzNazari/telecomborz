package protocol

type WireMessage struct {
    From      string
    To        string
    Payload   []byte
    Header    []byte
    Timestamp int64
}
