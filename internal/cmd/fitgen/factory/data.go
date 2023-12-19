// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package factory

// Data represent factory.tmpl
type Data struct {
	SDKVersion string // Fit SDK Version
	Package    string // Package name
	Messages   string // string formated struct []*proto.Message{...}
	Mesgnums   []string
}
