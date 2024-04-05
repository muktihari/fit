// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package builder

import "text/template"

// Builder is template data builder
type Builder interface {
	// Build returns slice of template data and an error
	Build() ([]Data, error)
}

// Data wraps template and data for generating proto.
type Data struct {
	Template     *template.Template
	TemplateExec string
	Path         string
	Filename     string
	Data         any
}
