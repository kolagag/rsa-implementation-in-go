# RSA Key Generator (Go)

A minimal implementation of RSA key generation using big integers.

### Features
- Random prime generation using crypto/rand
- Big integer modulus (1024/2048/4096 bits)
- Extended Euclidean modular inverse
- Generates public & private key pairs
- Saves keys to files

### Usage

```bash
go run .           # generates 1024-bit keys
go run . 2048      # generates 2048-bit keys

```

### Whats Next

- [ ] Add RSA Encryption & Decryption functions
- [ ] Implement PKCS#1 v1.5 Padding
- [ ] Build `Encrypt()` and `Decrypt()` APIs
- [ ] Store & Load Keys securely (PEM / custom format)
- [ ] Build CLI Tool for keygen/encrypt/decrypt
- [ ] Create Server–Client encrypted chat system
- [ ] Add Message Signing + Diffie-Hellman for forward secrecy
- [ ] OAuth-based authentication
- [ ] Encrypted Web/GUI Chat Application
