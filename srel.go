package srel

import (
	"regexp"
	"strings"
	"errors"
)

// Breaks a query into a slice of commands and arguments
func parse(query string) []string {
	tokens := strings.Split(query, " ")
	stack := make([]string, 0)

	toCnt := false
	cnt := ""

	for _, e := range tokens {
		// Since we split on the space char I need to allow some way 
		// to have space in a query, so you can use a space if you wrap the query quotes
		if e[0] == '"' && (e[len(e)-1] != '"' || len(e) <= 1) && !toCnt {
			toCnt = true
			cnt += e + " " // We need to insert a string that should be there
			continue
		} else if toCnt {
			if e[len(e)-1] == '"' {
				e = cnt + e
				// We dont want those quotes as part of the arg
				e = strings.TrimLeft(e, `"`)
				e = strings.TrimRight(e, `"`)

				cnt = ""
				toCnt = false
			} else {
				cnt += e + " "
				continue
			}
		}

		// push the arg on the stack to be processed
		stack = append(stack, e)
	}

	return stack
}

// Builds a Regexp with a given SREL query
// TODO: Make this shit accually work!
func Build(query string) (r *regexp.Regexp, err error) {
	stack := parse(query)

	b, err := run(stack)
	
	if err != nil {
		return
	}

	r, err = b.compile()
	return
}

func run(stack []string) (b *builder, err error) {
	b = &builder{mod:"m"}

	stMap, varMap := commandMap()

	for i := 0; i < len(stack); i++ {
		cmd := stack[i]
		if stFunc, found := stMap[cmd]; found {
			stFunc(b)
			continue
		} 

		if varFunc, found := varMap[cmd]; found {
			i++
			if i < len(stack) {
				varFunc(b, stack[i])
			} else {
				err = errors.New("Must end in a valid query")
				return
			}
			continue
		} else {
			err = errors.New("Unknown command: " + cmd)
			return
		}

	}

	return 
}

type cmdMap map[string]builderFunc
type cmdVarMap map[string]builderFuncVar

func commandMap() (cmdMap, cmdVarMap) {
	cm := cmdMap {
		"STARTOFLINE": startOfLine,
		"STARTLINE": startOfLine,
		"SOL": startOfLine,
		"ENDOFLINE": endOfLine,
		"ENDLINE": endOfLine,
		"EOL": endOfLine,
		"ANYTHING": anything,
		"LINEBREAK": lineBreak,
		"BR": lineBreak,
		"TAB": tab,
		"WORD": word,
		"OR": or,
	}

	cvm := cmdVarMap {
		"ANY": any,
		"ANYOF": any,
		"ANYTHINGBUT": anythingBut,
		"ALLBUT": anythingBut,
		"MAYBE": maybe,
		"FIND": find,
		"THEN": find,
	}

	return cm, cvm
}



