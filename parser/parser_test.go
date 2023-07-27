package parser

import (
	"Zeon/Programming-Language/ast"
	"Zeon/Programming-Language/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input :=
		`let x = 5;
	let y = 10;
	let foobar = 838383;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3{
		t.Fatalf("program.Statements doesnot contains 3 statements. got = %d\n",len(program.Statements))
	}

	tests := []struct{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests{
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt,tt.expectedIdentifier){
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool{
	if s.TokenLiteral() != "let"{
		t.Errorf("s.TokenLiteral not: 'let'. got=%q",s.TokenLiteral())
		return false
	}

	if letStmt.Name.Value != name{
		t.Errorf("letStmt.Name.Value not %s. got=%s",name,letStmt.Name.Value)
	}
}
