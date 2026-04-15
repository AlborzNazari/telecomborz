package crypto

type RatchetState struct {
    // TODO: root key, chain keys, counters, etc.
}

func NewRatchetState() *RatchetState {
    return &RatchetState{}
}

func (r *RatchetState) Encrypt(plaintext []byte) (ciphertext []byte, header []byte, err error) {
    // TODO: implement Double Ratchet + AEAD
    return plaintext, []byte("hdr"), nil
}

func (r *RatchetState) Decrypt(ciphertext, header []byte) ([]byte, error) {
    // TODO: implement Double Ratchet + AEAD
    return ciphertext, nil
}
