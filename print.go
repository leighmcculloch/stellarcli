package main

import (
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/stellar/go/xdr"
)

func Dump(w io.Writer, v interface{}) {
	dump(w, reflect.ValueOf(v), 0, reflect.Invalid)
}

const indentStr = "  "

func dump(w io.Writer, v reflect.Value, indent int, parentKind reflect.Kind) {
	if v.Type() == reflect.TypeOf(xdr.AccountId{}) {
		accountId := v.Interface().(xdr.AccountId)
		fmt.Fprintf(w, "%s\n", accountId.Address())
		return
	}
	switch v.Kind() {
	case reflect.Ptr, reflect.Interface:
		vv := v.Elem()
		if vv.IsValid() {
			dump(w, vv, indent, parentKind)
		} else {
			fmt.Fprint(w, "nil\n")
		}
	case reflect.Struct:
		t := v.Type()
		fmt.Fprintf(w, "\n")
		for i := 0; i < v.NumField(); i++ {
			sf := t.Field(i)
			f := v.Field(i)
			if reflect.DeepEqual(f.Interface(), reflect.Zero(f.Type()).Interface()) {
				continue
			}
			fmt.Fprintf(w, "%s%s: ", strings.Repeat(indentStr, indent), sf.Name)
			dump(w, f, indent+1, v.Kind())
		}
	case reflect.Slice:
		fmt.Fprintf(w, "\n")
		for i := 0; i < v.Len(); i++ {
			fmt.Fprintf(w, "%s- ", strings.Repeat(indentStr, indent))
			e := v.Index(i)
			dump(w, e, indent+1, v.Kind())
		}
	case reflect.Map:
		fmt.Fprintf(w, "\n")
		for _, k := range v.MapKeys() {
			fmt.Fprintf(w, "%s%#v: ", strings.Repeat(indentStr, indent), k)
			e := v.MapIndex(k)
			dump(w, e, indent+1, v.Kind())
		}
	default:
		fmt.Fprintf(w, "%v\n", v)
	}
}
