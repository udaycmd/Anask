// Copyright 2026 Uday Tiwari. All rights reserved.
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/udaycmd/Anask/priory/config"
	. "github.com/udaycmd/Anask/priory/internal"
)

// A [Handler] object manages all the dependencies for a gin router
type Handler struct {
	C *config.Config
}

func NewHandlerWithConfig(config *config.Config) *Handler {
	return &Handler{
		C: config,
	}
}

func (h *Handler) Signature(c *gin.Context) {
	var req *signatureReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	sign, err := DoSignature(h.C, req.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.AsciiJSON(http.StatusOK, gin.H{
		"signature": sign,
	})
}

type signatureReq struct {
	Id string `json:"id" binding:"required"`
}
