package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"testing"
)

func TestParse(t *testing.T) {
	//input, _ := antlr.NewFileStream("/Users/zhangzhen/go/src/github.com/YoKoa/sea/database/grammars/examples/bitrix_queries_cut.sql")
	input := antlr.NewInputStream("SELECT * FROM `b_forum_user` LIMIT 0;")
	lexer := NewMySqlLexer(input)

	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := NewMySqlParser(stream)
	parser.BuildParseTrees = true
	tree := parser.Root()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)


}
type TreeShapeListener struct {
	*BaseMySqlParserListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	//fmt.Println("test : " , ctx.)
}