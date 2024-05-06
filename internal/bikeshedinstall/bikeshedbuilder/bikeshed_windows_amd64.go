package bikeshedbuilder

import _ "embed"

//go:generate go run ../../../task.go wget-template "https://github.com/jcbhmr/bikeshed-builder/releases/download/v{{.BikeshedBuilderVersion}}/bikeshed-x86_64-pc-windows-msvc.zip"
//go:embed bikeshed-x86_64-pc-windows-msvc.zip
var BikeshedZip []byte
