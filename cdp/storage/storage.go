// Package storage provides the Chrome Debugging Protocol
// commands, types, and events for the Chrome Storage domain.
//
// Generated by the chromedp-gen command.
package storage

// AUTOGENERATED. DO NOT EDIT.

import (
	"context"

	. "github.com/knq/chromedp/cdp"
	"github.com/mailru/easyjson"
)

var (
	_ BackendNode
	_ BackendNodeID
	_ ComputedProperty
	_ ErrorType
	_ Frame
	_ FrameID
	_ LoaderID
	_ Message
	_ MessageError
	_ MethodType
	_ Node
	_ NodeID
	_ NodeType
	_ PseudoType
	_ RGBA
	_ ShadowRootType
	_ Timestamp
)

// ClearDataForOriginParams clears storage for origin.
type ClearDataForOriginParams struct {
	Origin       string `json:"origin"`       // Security origin.
	StorageTypes string `json:"storageTypes"` // Comma separated origin names.
}

// ClearDataForOrigin clears storage for origin.
//
// parameters:
//   origin - Security origin.
//   storageTypes - Comma separated origin names.
func ClearDataForOrigin(origin string, storageTypes string) *ClearDataForOriginParams {
	return &ClearDataForOriginParams{
		Origin:       origin,
		StorageTypes: storageTypes,
	}
}

// Do executes Storage.clearDataForOrigin.
func (p *ClearDataForOriginParams) Do(ctxt context.Context, h FrameHandler) (err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// marshal
	buf, err := easyjson.Marshal(p)
	if err != nil {
		return err
	}

	// execute
	ch := h.Execute(ctxt, CommandStorageClearDataForOrigin, easyjson.RawMessage(buf))

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			return nil

		case error:
			return v
		}

	case <-ctxt.Done():
		return ErrContextDone
	}

	return ErrUnknownResult
}