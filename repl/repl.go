package repl

import (
	"book/monkeyLang/evaluator"
	"book/monkeyLang/lexer"
	"book/monkeyLang/object"
	"book/monkeyLang/parser"
	"bufio"
	"fmt"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
    scanner := bufio.NewScanner(in)
    env := object.NewEnvironment()
    for {
        fmt.Fprintf(out,PROMPT)
        scanned := scanner.Scan()
        
        if !scanned {
            return
        
        }
        line := scanner.Text()
        if line == "exit" || line == "exit()" { 
            return 
        }
        l := lexer.New(line)
        p := parser.New(l)

        program := p.ParseProgram()
        if(len(p.Errors())) != 0 {
            printParserErrors(out, p.Errors())
        }
        evaluated := evaluator.Eval(program,env)
        if evaluated != nil {
            io.WriteString(out, evaluated.Inspect())
            io.WriteString(out, "\n")
        }
    }
}

func printParserErrors(out io.Writer, errors []string) {
    for _, msg := range errors {
        io.WriteString(out, "\t"+msg+"\n")
    }
}
