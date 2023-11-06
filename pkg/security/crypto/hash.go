package crypto

import (
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"golang.org/x/crypto/argon2"
	"strings"
)

var (
	ErrInvalidHash         = errors.New("the encoded hash is not in the correct format")
	ErrIncompatibleVersion = errors.New("incompatible argon2 version")
)

type params struct {
	iterations uint32
	memory     uint32
	threads    uint8
	keyLength  uint32
	saltLength uint32
}

const TIME = 1
const MEMORY = 64 * 1024
const THREADS = 4
const KEYLENGTH = 32
const SALTLENGTH = 16

// Params could be used to configure the hashing from outside
var Params = &params{
	iterations: TIME,
	memory:     MEMORY,
	threads:    THREADS,
	keyLength:  KEYLENGTH,
	saltLength: SALTLENGTH,
}

func SecureUrlSafe(nbytes uint32) (string, error) {
	secureBytes, err := SecureRandomBytes(nbytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(secureBytes), nil
}

func argon2Hash(secret string, salt *[]byte) (key []byte, err error) {
	if *salt == nil {
		*salt, err = SecureRandomBytes(Params.saltLength)
	}
	key = argon2.IDKey(
		[]byte(secret),
		*salt,
		Params.iterations,
		Params.memory,
		Params.threads,
		Params.keyLength,
	)
	return
}

func decodeHash(encodedHash string) (p *params, salt, hash []byte, err error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, ErrInvalidHash
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, ErrIncompatibleVersion
	}

	// extract parameters from hash
	p = &params{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &p.memory, &p.iterations, &p.threads)
	if err != nil {
		return nil, nil, nil, err
	}

	// extract hashed value
	hash, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}

	// extract salt
	salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	p.saltLength = uint32(len(salt))
	p.keyLength = uint32(len(hash))

	return p, salt, hash, nil
}

func PasswordHash(password string) (string, error) {
	var salt []byte
	hash, err := argon2Hash(password, &salt)
	if err != nil {
		return "", err
	}
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)

	encodedHash := fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version, Params.memory, Params.iterations, Params.threads, b64Hash, b64Salt)
	return encodedHash, nil
}

func ValidatePassword(plain string, hashed string) (valid bool, err error) {
	_, salt, hash, err := decodeHash(hashed)
	valid = false
	if err != nil {
		return
	}
	otherHashed, err := argon2Hash(plain, &salt)
	if err != nil {
		return
	}
	valid = subtle.ConstantTimeCompare(hash, otherHashed) == 1
	return
}
