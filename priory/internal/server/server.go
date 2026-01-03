// Copyright 2026 Uday Tiwari. All rights reserved.
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/udaycmd/Anask/priory/config"
)

func NewServer(cfg *config.Config) *http.Server {
	router := NewRouter(NewHandlerWithConfig(cfg))

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Priory.Port),
		Handler:      router,
		IdleTimeout:  30 * time.Second,
		WriteTimeout: 14 * time.Second,
		ReadTimeout:  7 * time.Second,
	}

	return s
}
