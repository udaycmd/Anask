// Copyright 2026 Uday Tiwari. All rights reserved.
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package internal

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"github.com/udaycmd/Anask/priory/config"
)

// Creates a signature of an identifer string.
func DoSignature(cfg *config.Config, identifier string) (string, error) {
	salt := cfg.Salt.Value
	if salt == "" {
		return "", errors.New("no secret salt found in config")
	}
	return sign(salt, identifier), nil
}

// This creates the actual signature by combining the
// message with a secret salt via [HMAC_SHA256 digest].
//
// [HMAC_SHA256 digest]: https://en.wikipedia.org/wiki/HMAC
func sign(salt, msg string) string {
	hash := hmac.New(sha256.New, []byte(salt))
	hash.Write([]byte(msg))
	return hex.EncodeToString(hash.Sum(nil))
}
