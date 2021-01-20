package bench_secp256k1

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/btcsuite/btcd/btcec"
	"github.com/diodechain/diode_go_client/crypto/secp256k1"
)

const (
	testHexPublicKey  = "04a673638cb9587cb68ea08dbef685c6f2d2a751a8b3c6f2a7e9a4999e6e4bfaf5ca1d22fe57c6103dbaac10cf15d15c0791cab8bb9a04f800e4d215276cb3e008"
	testHexPrivateKey = "22a47fa09a223f2aa079edf85a7c2d4f8720ee63e502ee2869afab7de234b80c"
	// keccak(taiwanhowbonbon)
	testHexMessage = "c532e1d3ed50eb82e34517cf9b72c9233b410a5069865a0125168549fa57af5e"
)

var (
	pubBytes  = decodeHexForce(testHexPublicKey)
	privBytes = decodeHexForce(testHexPrivateKey)
	msgBytes  = decodeHexForce(testHexMessage)
)

func decodeHexForce(src string) (dst []byte) {
	if len(src)%2 != 0 {
		return
	}
	dst, _ = hex.DecodeString(src)
	return
}

// Benchmark sign message
func BenchmarkBTCECSignMessage(b *testing.B) {
	b.ResetTimer()
	priv, _ := btcec.PrivKeyFromBytes(btcec.S256(), privBytes)
	for i := 0; i < b.N; i++ {
		_, err := priv.Sign(msgBytes)
		if err != nil {
			b.Logf("[btcec] sign message failed: %+v\n", err)
		}
	}
}

func BenchmarkCGOSignMessage(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := secp256k1.Sign(msgBytes, privBytes)
		if err != nil {
			b.Logf("[cgo] sign message failed: %+v\n", err)
		}
	}
}

// Benchmark recover public key from message and signature
func BenchmarkBTCECRecoverPublicKey(b *testing.B) {
	b.ResetTimer()
	priv, _ := btcec.PrivKeyFromBytes(btcec.S256(), privBytes)
	sigBytes, err := btcec.SignCompact(btcec.S256(), priv, msgBytes, true)
	if err != nil {
		b.Logf("[btcec] sign message failed: %+v\n", err)
	}
	for i := 0; i < b.N; i++ {
		pub, compressed, err := btcec.RecoverCompact(btcec.S256(), sigBytes, msgBytes)
		if err != nil {
			b.Logf("[btcec] recover public key failed: %+v\n", err)
		}
		if !compressed {
			b.Log("[btcec] signature was not compressed")
		}
		if pub.X.Cmp(priv.X) != 0 || pub.Y.Cmp(priv.Y) != 0 {
			b.Log("[btcec] recovered public key didn't match the original private key")
		}
	}
}

func BenchmarkCGORecoverPublicKey(b *testing.B) {
	b.ResetTimer()
	sig, err := secp256k1.Sign(msgBytes, privBytes)
	if err != nil {
		b.Logf("[cgo] sign message failed: %+v\n", err)
	}
	for i := 0; i < b.N; i++ {
		pub, err := secp256k1.RecoverPubkey(msgBytes, sig)
		if err != nil {
			b.Logf("[cgo] recover public key failed: %+v\n", err)
		}
		if !bytes.Equal(pubBytes, pub) {
			b.Log("[cgo] recovered public key didn't match")
		}
	}
}

// Benchmark verify message
func BenchmarkBTCECVerifyMessage(b *testing.B) {
	b.ResetTimer()
	priv, pub := btcec.PrivKeyFromBytes(btcec.S256(), privBytes)
	sig, err := priv.Sign(msgBytes)
	if err != nil {
		b.Logf("[btcec] sign message failed: %+v\n", err)
	}
	for i := 0; i < b.N; i++ {
		verified := sig.Verify(msgBytes, pub)
		if !verified {
			b.Log("[btcec] signature was not verified")
		}
	}
}

func BenchmarkCGOVerifyMessage(b *testing.B) {
	b.ResetTimer()
	sig, err := secp256k1.Sign(msgBytes, privBytes)
	if err != nil {
		b.Logf("[cgo] sign message failed: %+v\n", err)
	}
	rsig := sig[1:]
	for i := 0; i < b.N; i++ {
		verified := secp256k1.VerifySignature(pubBytes, msgBytes, rsig)
		if !verified {
			b.Log("[cgo] signature was not verified")
		}
	}
}
