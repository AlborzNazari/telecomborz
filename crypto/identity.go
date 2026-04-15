package crypto

type Identity struct {
    Public  []byte
    Private []byte
}

func GenerateIdentity() (*Identity, error) {
    // TODO: use Ed25519/X25519 from a real crypto library
    return &Identity{
        Public:  []byte("dummy-public"),
        Private: []byte("dummy-private"),
    }, nil
}
