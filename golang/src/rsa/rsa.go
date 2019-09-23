package main

import (
	"bytes"
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"fmt"
	"math/big"
)

var privateKey = "MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBALHHnteMDckJNlPxZ7eMzxSDmH4lTMeD6I3EvpVKad8Sw4apQvIG9ZY7ZHeUKKTOlsU0yBOp432BheP74EdU1aljnMqXpFNn+bEgTXpXCzaIdJlij9H4y/2m//mGE9l1OX2EVHZKSmeMY/GihZlMD6tP3yJ8QdolBZI/3CgH7BLDAgMBAAECgYAvkQioBXoeww89MIcerlct1vPzNImxjFKps+2GRk3DeOLF4f3eggwtsSB1ejfRuNDQXQn3cOpER2aKlHbyvvkXkNhrd/lCjpk6wtDYQsq/eeQ7wC8Am6hQ2d8cKySCl5LrpHHzkGkTv1DHw7rNKrMR03ahJWXsyPcqrbhvBMwrMQJBAPlh95E8wPSsqqYA/74o7Iqxa7nq9osXT6t5xrJc2CpI2go4OK1Da1zOI+mCbNpnuA7PnWu9xam2cCmNAsTHGskCQQC2f0L3no9mtGmuB7M7xN4Me5pUlZqVRWzLKDUK3IPEHzUZs7WDQ77RqOJBrvdHxFpY3ZS+bDFYouUbck39vHsrAkEAiIgCKhnA6jO+GbRiT5HILwaDm/3vjKbuj0rUZcI+9qd7+CxfmzxWAzE4qBcn0UsHkdRIszvqg8fGEHmLEoCPQQJAZWT3lBRooCuEu8hTcNXEeTMDYBNuu5jDBWzla49xNjoQiqMqKjAtiNdIPi4z/Y++krkpt1LtZ825dTJg2qUp2QJAMdlZbhYPL99fjhsbUS+xNisTczoi9Y+PEh+expEvfnTIj/YqKHtVdCdIPxktews831vU14GF+UwWFEZQYLt65w=="
var publicKey = "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCxx57XjA3JCTZT8We3jM8Ug5h+JUzHg+iNxL6VSmnfEsOGqULyBvWWO2R3lCikzpbFNMgTqeN9gYXj++BHVNWpY5zKl6RTZ/mxIE16Vws2iHSZYo/R+Mv9pv/5hhPZdTl9hFR2SkpnjGPxooWZTA+rT98ifEHaJQWSP9woB+wSwwIDAQAB"

var (
	ErrInputSize  = errors.New("input size too large")
	ErrEncryption = errors.New("encryption error")
)

func PrivateEncrypt(priv *rsa.PrivateKey, data []byte) (enc []byte, err error) {

	k := (priv.N.BitLen() + 7) / 8
	tLen := len(data)
	if tLen > k-11 {
		err = ErrInputSize
		panic(err)
	}
	em := make([]byte, k)
	em[1] = 1
	for i := 2; i < k-tLen-1; i++ {
		em[i] = 0xff
	}
	copy(em[k-tLen:k], data)
	c := new(big.Int).SetBytes(em)
	if c.Cmp(priv.N) > 0 {
		err = ErrEncryption
		return
	}
	var m *big.Int
	var ir *big.Int
	if priv.Precomputed.Dp == nil {
		m = new(big.Int).Exp(c, priv.D, priv.N)
	} else {
		m = new(big.Int).Exp(c, priv.Precomputed.Dp, priv.Primes[0])
		m2 := new(big.Int).Exp(c, priv.Precomputed.Dq, priv.Primes[1])
		m.Sub(m, m2)
		if m.Sign() < 0 {
			m.Add(m, priv.Primes[0])
		}
		m.Mul(m, priv.Precomputed.Qinv)
		m.Mod(m, priv.Primes[0])
		m.Mul(m, priv.Primes[1])
		m.Add(m, m2)

		for i, values := range priv.Precomputed.CRTValues {
			prime := priv.Primes[2+i]
			m2.Exp(c, values.Exp, prime)
			m2.Sub(m2, m)
			m2.Mul(m2, values.Coeff)
			m2.Mod(m2, prime)
			if m2.Sign() < 0 {
				m2.Add(m2, prime)
			}
			m2.Mul(m2, values.R)
			m.Add(m, m2)
		}
	}

	if ir != nil {
		// Unblind.
		m.Mul(m, ir)
		m.Mod(m, priv.N)
	}
	enc = m.Bytes()
	return
}

func sign(data []byte) string {
	k, _ := base64.StdEncoding.DecodeString(privateKey)
	// 用pkcs8
	privkey, _ := x509.ParsePKCS8PrivateKey(k)
	md5 := md5.New()
	md5.Write(data)
	h := md5.Sum(nil)
	// fmt.Println(base64.StdEncoding.EncodeToString((h)))
	// md5 with rsa
	// opts := &rsa.PSSOptions{SaltLength: rsa.PSSSaltLengthAuto, Hash: crypto.MD5}
	buf, _ := rsa.SignPKCS1v15(rand.Reader, privkey.(*rsa.PrivateKey), crypto.MD5, h)
	return base64.StdEncoding.EncodeToString(buf)
}

func encryptData(data []byte) string {
	k, _ := base64.StdEncoding.DecodeString(privateKey)

	// 用pkcs8
	privkey, _ := x509.ParsePKCS8PrivateKey(k)
	pk := privkey.(*rsa.PrivateKey)
	buf := bytes.NewBuffer(nil)
	inputlen := len(data)
	maxsize := 117
	offset := 0
	for inputlen-offset > 0 {
		var encdata []byte
		if inputlen-offset > maxsize {
			encdata, _ = PrivateEncrypt(pk, data[offset:offset+maxsize])
		} else {
			encdata, _ = PrivateEncrypt(pk, data[offset:offset+inputlen-offset])
		}
		buf.Write(encdata)
		offset += maxsize
	}

	// 输出base64一下 要不然是乱码
	// fmt.Println(base64.StdEncoding.EncodeToString(buf.Bytes()))
	return base64.StdEncoding.EncodeToString(buf.Bytes())
}

func RSAEncrypt(data []byte) string {
	//keyStr := beego.AppConfig.String("publicKey")

	k, _ := base64.StdEncoding.DecodeString(publicKey)

	publicKey, _ := x509.ParsePKIXPublicKey(k)

	pk := publicKey.(*rsa.PublicKey)
	buf := bytes.NewBuffer(nil)
	inputlen := len(data)
	maxsize := 117
	offset := 0

	for inputlen-offset > 0 {
		var encdata []byte
		if inputlen-offset > maxsize {
			encdata, _ = rsa.EncryptPKCS1v15(rand.Reader, pk, data[offset:offset+maxsize])
		} else {
			encdata, _ = rsa.EncryptPKCS1v15(rand.Reader, pk, data[offset:offset+inputlen-offset])
		}
		buf.Write(encdata)
		offset += maxsize
	}

	// 输出base64一下 要不然是乱码
	// fmt.Println(base64.StdEncoding.EncodeToString(buf.Bytes()))
	return base64.StdEncoding.EncodeToString(buf.Bytes())

}

// 私钥解密
func RSADecrypt(data []byte) string {
	k, _ := base64.StdEncoding.DecodeString(privateKey)

	// 用pkcs8
	privkey, _ := x509.ParsePKCS8PrivateKey(k)
	pk := privkey.(*rsa.PrivateKey)
	buf := bytes.NewBuffer(nil)
	inputlen := len(data)
	maxsize := 117
	offset := 0
	for inputlen-offset > 0 {
		var encdata []byte
		if inputlen-offset > maxsize {
			encdata, _ = rsa.DecryptPKCS1v15(rand.Reader, pk, data[offset:offset+maxsize])
		} else {
			encdata, _ = rsa.DecryptPKCS1v15(rand.Reader, pk, data[offset:offset+inputlen-offset])
		}
		buf.Write(encdata)
		offset += maxsize
	}

	// 输出base64一下 要不然是乱码
	// fmt.Println(base64.StdEncoding.EncodeToString(buf.Bytes()))
	return base64.StdEncoding.EncodeToString(buf.Bytes())
}

// copy from crypt/rsa/pkcs1v5.go
var hashPrefixes = map[crypto.Hash][]byte{
	crypto.MD5:       {0x30, 0x20, 0x30, 0x0c, 0x06, 0x08, 0x2a, 0x86, 0x48, 0x86, 0xf7, 0x0d, 0x02, 0x05, 0x05, 0x00, 0x04, 0x10},
	crypto.SHA1:      {0x30, 0x21, 0x30, 0x09, 0x06, 0x05, 0x2b, 0x0e, 0x03, 0x02, 0x1a, 0x05, 0x00, 0x04, 0x14},
	crypto.SHA224:    {0x30, 0x2d, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x04, 0x05, 0x00, 0x04, 0x1c},
	crypto.SHA256:    {0x30, 0x31, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x01, 0x05, 0x00, 0x04, 0x20},
	crypto.SHA384:    {0x30, 0x41, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x02, 0x05, 0x00, 0x04, 0x30},
	crypto.SHA512:    {0x30, 0x51, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x03, 0x05, 0x00, 0x04, 0x40},
	crypto.MD5SHA1:   {}, // A special TLS case which doesn't use an ASN1 prefix.
	crypto.RIPEMD160: {0x30, 0x20, 0x30, 0x08, 0x06, 0x06, 0x28, 0xcf, 0x06, 0x03, 0x00, 0x31, 0x04, 0x14},
}

// copy from crypt/rsa/pkcs1v5.go
func encrypt(c *big.Int, pub *rsa.PublicKey, m *big.Int) *big.Int {
	e := big.NewInt(int64(pub.E))
	c.Exp(m, e, pub.N)
	return c
}

// copy from crypt/rsa/pkcs1v5.go
func pkcs1v15HashInfo(hash crypto.Hash, inLen int) (hashLen int, prefix []byte, err error) {
	// Special case: crypto.Hash(0) is used to indicate that the data is
	// signed directly.
	if hash == 0 {
		return inLen, nil, nil
	}

	hashLen = hash.Size()
	if inLen != hashLen {
		return 0, nil, errors.New("crypto/rsa: input must be hashed message")
	}
	prefix, ok := hashPrefixes[hash]
	if !ok {
		return 0, nil, errors.New("crypto/rsa: unsupported hash function")
	}
	return
}

// copy from crypt/rsa/pkcs1v5.go
func leftPad(input []byte, size int) (out []byte) {
	n := len(input)
	if n > size {
		n = size
	}
	out = make([]byte, size)
	copy(out[len(out)-n:], input)
	return
}
func unLeftPad(input []byte) (out []byte) {
	n := len(input)
	t := 2
	for i := 2; i < n; i++ {
		if input[i] == 0xff {
			t = t + 1
		} else {
			if input[i] == input[0] {
				t = t + int(input[1])
			}
			break
		}
	}
	out = make([]byte, n-t)
	copy(out, input[t:])
	return
}

// copy&modified from crypt/rsa/pkcs1v5.go
func publicDecrypt(pub *rsa.PublicKey, hash crypto.Hash, hashed []byte, sig []byte) (out []byte, err error) {
	hashLen, prefix, err := pkcs1v15HashInfo(hash, len(hashed))
	if err != nil {
		return nil, err
	}

	tLen := len(prefix) + hashLen
	k := (pub.N.BitLen() + 7) / 8
	if k < tLen+11 {
		return nil, fmt.Errorf("length illegal")
	}

	c := new(big.Int).SetBytes(sig)
	m := encrypt(new(big.Int), pub, c)
	em := leftPad(m.Bytes(), k)
	out = unLeftPad(em)

	err = nil
	return
}

/*func PublicDecrypt(pub *rsa.PublicKey, data []byte) ([]byte, error) {
	decData, err := publicDecrypt(pub, crypto.Hash(0), nil, data)
	if err != nil {
		return nil, err
	}
	return decData, nil
}*/

// 公钥解密
func PublicDecrypt(data []byte) string {
	k, _ := base64.StdEncoding.DecodeString(publicKey)

	// 用pkcs8
	public, _ := x509.ParsePKIXPublicKey(k)
	pk := public.(*rsa.PublicKey)

	buf := bytes.NewBuffer(nil)
	inputlen := len(data)
	maxsize := 117
	offset := 0
	for inputlen-offset > 0 {
		var encdata []byte
		if inputlen-offset > maxsize {
			encdata, _ = publicDecrypt(pk, crypto.Hash(0), nil, data[offset:offset+maxsize])
		} else {
			encdata, _ = publicDecrypt(pk, crypto.Hash(0), nil, data[offset:offset+inputlen-offset])
		}
		buf.Write(encdata)
		offset += maxsize
	}

	// 输出base64一下 要不然是乱码
	// fmt.Println(base64.StdEncoding.EncodeToString(buf.Bytes()))
	return base64.StdEncoding.EncodeToString(buf.Bytes())
}

func PrivateEncrypt2(data []byte) string {
	k, _ := base64.StdEncoding.DecodeString(privateKey)

	// 用pkcs8
	privkey, _ := x509.ParsePKCS8PrivateKey(k)
	pk := privkey.(*rsa.PrivateKey)
	buf := bytes.NewBuffer(nil)
	inputlen := len(data)
	maxsize := 117
	offset := 0
	for inputlen-offset > 0 {
		var encdata []byte
		if inputlen-offset > maxsize {
			encdata, _ = rsa.SignPKCS1v15(nil, pk, crypto.Hash(0), data[offset:offset+maxsize])
		} else {
			encdata, _ = rsa.SignPKCS1v15(nil, pk, crypto.Hash(0), data[offset:offset+inputlen-offset])
		}
		buf.Write(encdata)
		offset += maxsize
	}

	// 输出base64一下 要不然是乱码
	// fmt.Println(base64.StdEncoding.EncodeToString(buf.Bytes()))
	return base64.StdEncoding.EncodeToString(buf.Bytes())
}

func main() {

	enD := encryptData([]byte("hello"))
	fmt.Println(enD)

	fmt.Println(PrivateEncrypt2([]byte("hello")))

	rsa.DecryptOAEP(crypto.Hash(0), nil)
	b, _ := base64.StdEncoding.DecodeString(enD)
	fmt.Println(PublicDecrypt(b))
}
