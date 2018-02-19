package main

import (
	"io/ioutil"
	"log"

	"github.com/frustra/bbcode"
	"github.com/namreg/bbcode/phpgo"
)

const extName = "gobbcode"

func main() {
	panic("should not be called")
}

func init() {
	log.SetOutput(ioutil.Discard)
	phpgo.InitExtension(extName, "")
	phpgo.AddFunc("gobbcode_parse", parse)
}

func parse(str string) string {
	compiler := bbcode.NewCompiler(true, true)
	return compiler.Compile(str)
}
