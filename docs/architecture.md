# Architecture

- Core crypto: identity keys, sessions, ratchet.
- Protocol: sessions, wire messages.
- Transport: pluggable (IP for v1).
- Storage: encrypted keystore and message store.
- App: CLI that wires everything together.
