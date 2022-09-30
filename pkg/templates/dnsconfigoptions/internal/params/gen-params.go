// Code generated by kube-linter template codegen. DO NOT EDIT.
//go:build !templatecodegen
// +build !templatecodegen

package params

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"golang.stackrox.io/kube-linter/pkg/check"
	"golang.stackrox.io/kube-linter/pkg/templates/util"
)

var (
	// Use some imports in case they don't get used otherwise.
	_ = util.MustParseParameterDesc
	_ = fmt.Sprintf

	keyParamDesc = util.MustParseParameterDesc(`{
	"Name": "key",
	"Type": "string",
	"Description": "Key of the dnsConfig option.",
	"Examples": null,
	"Enum": null,
	"SubParameters": null,
	"ArrayElemType": "",
	"Required": false,
	"NoRegex": false,
	"NotNegatable": false,
	"XXXStructFieldName": "Key",
	"XXXIsPointer": false
}
`)

	valueParamDesc = util.MustParseParameterDesc(`{
	"Name": "value",
	"Type": "string",
	"Description": "Value of the dnsConfig option.",
	"Examples": null,
	"Enum": null,
	"SubParameters": null,
	"ArrayElemType": "",
	"Required": false,
	"NoRegex": false,
	"NotNegatable": false,
	"XXXStructFieldName": "Value",
	"XXXIsPointer": false
}
`)

	ParamDescs = []check.ParameterDesc{
		keyParamDesc,
		valueParamDesc,
	}
)

func (p *Params) Validate() error {
	var validationErrors []string
	if len(validationErrors) > 0 {
		return errors.Errorf("invalid parameters: %s", strings.Join(validationErrors, ", "))
	}
	return nil
}

// ParseAndValidate instantiates a Params object out of the passed map[string]interface{},
// validates it, and returns it.
// The return type is interface{} to satisfy the type in the Template struct.
func ParseAndValidate(m map[string]interface{}) (interface{}, error) {
	var p Params
	if err := util.DecodeMapStructure(m, &p); err != nil {
		return nil, err
	}
	if err := p.Validate(); err != nil {
		return nil, err
	}
	return p, nil
}

// WrapInstantiateFunc is a convenience wrapper that wraps an untyped instantiate function
// into a typed one.
func WrapInstantiateFunc(f func(p Params) (check.Func, error)) func(interface{}) (check.Func, error) {
	return func(paramsInt interface{}) (check.Func, error) {
		return f(paramsInt.(Params))
	}
}
