/*
   Copyright 2011-2013 gtalent2@gmail.com

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/
package lex

import "strconv"

type Token struct {
	TokType  int
	TokValue interface{}
}

func (me *Token) Type() int {
	return me.TokType
}

func (me *Token) Value() interface{} {
	return me.TokValue
}

func (me *Token) Set(t int, val interface{}) {
	me.TokType = t
	me.TokValue = val
}

func (me *Token) SetInt(val int) {
	me.TokType = IntLiteral
	me.TokValue = val
}

func (me *Token) Int() int {
	return me.TokValue.(int)
}

func (me *Token) SetString(val string) {
	me.TokType = StringLiteral
	me.TokValue = val
}

func (me *Token) String() string {
	switch me.TokValue.(type) {
	case int:
		return strconv.Itoa(me.TokValue.(int))
	case bool:
		return strconv.FormatBool(me.TokValue.(bool))
	case string:
		return me.TokValue.(string)
	}
	return ""
}

func (me *Token) SetBool(val bool) {
	me.TokType = BoolLiteral
	me.TokValue = val
}

func (me *Token) Bool() bool {
	return me.TokValue.(bool)
}
