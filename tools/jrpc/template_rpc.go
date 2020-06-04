package main

import (
	"fmt"
	"path"
	"text/template"

	"github.com/jakewright/home-automation/tools/jrpc/imports"
)

type rpcDataEndpoint struct {
	InputType  string
	OutputType string
	HTTPMethod string
	URL        string
}

type rpcData struct {
	PackageName string
	Imports     []*imports.Imp
	Endpoints   []*rpcDataEndpoint
}

const rpcTemplateText = `// Code generated by jrpc. DO NOT EDIT.

package {{ .PackageName }}

{{ if .Imports }}
	import (
		{{- range .Imports }}
			{{ .Alias }} "{{ .Path }}"
		{{- end}}
	)
{{- end }}

{{ range .Endpoints }}
	// Request builds an RPC request
	func (m *{{ .InputType }}) Request() *rpc.Request {
		return &rpc.Request{
			Method: "{{ .HTTPMethod }}",
			URL: "{{ .URL }}",
			Body: m,
		}
	}

	// Do performs the request
	func (m *{{ .InputType }}) Do(ctx context.Context) (*{{ .OutputType }}, error) {
		rsp := &{{ .OutputType }}{}
		_, err := rpc.Do(ctx, m.Request(), rsp)
		return rsp, err
	}
{{ end }}
`

type rpcGenerator struct {
	baseGenerator
}

func (g *rpcGenerator) Template() (*template.Template, error) {
	return template.New("rpc_template").Parse(rpcTemplateText)
}

func (g *rpcGenerator) PackageDir() string {
	return packageDirExternal
}

func (g *rpcGenerator) Data(im *imports.Manager) (interface{}, error) {
	im.Add("github.com/jakewright/home-automation/libraries/go/rpc")
	im.Add("context")

	if g.file.Service == nil {
		return nil, nil
	}

	if len(g.file.Service.RPCs) == 0 {
		return nil, nil
	}

	routerPath, ok := g.file.Service.Options["path"].(string)
	if !ok {
		return nil, fmt.Errorf("path not set on service %q", g.file.Service.Name)
	}

	endpoints := make([]*rpcDataEndpoint, len(g.file.Service.RPCs))
	for i, r := range g.file.Service.RPCs {
		method, err := getMethod(r)
		if err != nil {
			return nil, fmt.Errorf("failed to get RPC %q method: %w", r.Name, err)
		}

		rpcPath, err := getPath(r)
		if err != nil {
			return nil, fmt.Errorf("failed to get RPC %q path: %w", r.Name, err)
		}

		inType, err := resolveTypeName(r.InputType, g.file, im)
		if err != nil {
			return nil, fmt.Errorf("failed to resolve RPC %q input type: %w", r.Name, err)
		}

		outType, err := resolveTypeName(r.OutputType, g.file, im)
		if err != nil {
			return nil, fmt.Errorf("failed to resolve RPC %q output type: %w", r.Name, err)
		}

		endpoints[i] = &rpcDataEndpoint{
			InputType:  inType.TypeName,
			OutputType: outType.TypeName,
			HTTPMethod: method,
			URL:        path.Join(routerPath, rpcPath),
		}
	}

	return &rpcData{
		PackageName: externalPackageName(g.options),
		Imports:     im.Get(),
		Endpoints:   endpoints,
	}, nil
}

func (g *rpcGenerator) Filename() string {
	return "rpc.go"
}
