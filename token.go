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

type TokenList []Token

//Returns next Token
func (me *TokenList) Peak() Token {
	return (*me)[0]
}

//Drops the next Token from the list and returns it
func (me *TokenList) Next() Token {
	retval := (*me)[0]
	*me = (*me)[1:]
	return retval
}

func (me *TokenList) HasNext() bool {
	return len(*me) != 0
}

type Token struct {
	Type  int
	Value string
}

func (me *Token) Set(t int, val interface{}) {
	me.Type = t
	switch val.(type) {
	case int:
		me.Value = strconv.Itoa(val.(int))
	case bool:
		me.Value = strconv.FormatBool(val.(bool))
	case string:
		me.Value = val.(string)
	}
}

func (me *Token) SetInt(val int) {
	me.Type = IntLiteral
	me.Value = strconv.Itoa(val)
}

func (me *Token) Int() int {
	v, _ := strconv.Atoi(me.Value)
	return v
}

func (me *Token) SetString(val string) {
	me.Type = StringLiteral
	me.Value = val
}

func (me *Token) String() string {
	return me.Value
}

func (me *Token) SetBool(val bool) {
	me.Type = BoolLiteral
	me.Value = strconv.FormatBool(val)
}

func (me *Token) Bool() bool {
	v, _ := strconv.ParseBool(me.Value)
	return v
}
