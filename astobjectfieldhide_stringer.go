// Generated by: main
// TypeWriter: stringer
// Directive: +gen on astObjectFieldHide

package jsonnet

import (
	"fmt"
)

const _astObjectFieldHide_name = "astObjectFieldHiddenastObjectFieldInheritastObjectFieldVisible"

var _astObjectFieldHide_index = [...]uint8{0, 20, 41, 62}

func (i ObjectFieldHide) String() string {
	if i < 0 || i+1 >= ObjectFieldHide(len(_astObjectFieldHide_index)) {
		return fmt.Sprintf("astObjectFieldHide(%d)", i)
	}
	return _astObjectFieldHide_name[_astObjectFieldHide_index[i]:_astObjectFieldHide_index[i+1]]
}