package parser

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
)

var (
	errEmptyLine = errors.New("Empty Line")
)

type CommandType int

//go:generate go run golang.org/x/tools/cmd/stringer -type=CommandType
const (
	C_ARITHMETIC CommandType = iota
	C_PUSH
	C_POP
	C_LABEL
	C_GOTO
	C_IF
	C_FUNCTION
	C_RETURN
	C_CALL
)

var commandTypeMapping = map[string]CommandType{
	"add":      C_ARITHMETIC,
	"sub":      C_ARITHMETIC,
	"neq":      C_ARITHMETIC,
	"eq":       C_ARITHMETIC,
	"gt":       C_ARITHMETIC,
	"lt":       C_ARITHMETIC,
	"and":      C_ARITHMETIC,
	"or":       C_ARITHMETIC,
	"not":      C_ARITHMETIC,
	"push":     C_PUSH,
	"pop":      C_POP,
	"label":    C_LABEL,
	"goto":     C_GOTO,
	"if-goto":  C_IF,
	"function": C_FUNCTION,
	"return":   C_RETURN,
	"call":     C_CALL,
}

type Parser struct {
	scanner     *bufio.Scanner
	commandType CommandType
	arg1        string
	arg2        int64
}

func NewParser(r io.Reader) *Parser {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	return &Parser{
		scanner: scanner,
	}
}

func (p *Parser) Advance() error {
	err := p.read()
	for err == errEmptyLine {
		err = p.read()
	}

	return err
}

// Read one line, parse it, and set command into parser's fields.
// If the line is empty or a comment, this method returns errEmptyLine.
func (p *Parser) read() error {
	if ok := p.scanner.Scan(); !ok {
		if p.scanner.Err() == nil {
			return io.EOF
		} else {
			return p.scanner.Err()
		}
	}

	line := p.scanner.Text()
	parts := splitIntoParts(line)
	if len(parts) == 0 {
		return errEmptyLine
	}
	if len(parts) > 3 {
		return fmt.Errorf("invalid command (too many arguments): %s", line)
	}

	commandType, err := getCommandType(getOrEmpty(parts, 0))
	if err != nil {
		return err
	}
	p.commandType = commandType

	if commandType == C_ARITHMETIC {
		p.arg1 = getOrEmpty(parts, 0)
	} else {
		p.arg1 = getOrEmpty(parts, 1)
	}

	arg2 := getOrEmpty(parts, 2)
	if len(arg2) > 0 {
		arg2int, err := strconv.ParseInt(arg2, 10, 64)
		if err != nil {
			return fmt.Errorf("failed to parse arg2 as int64: %s, %+v", arg2, err)
		}
		p.arg2 = arg2int
	}

	return nil
}

func getCommandType(op string) (CommandType, error) {
	typ, ok := commandTypeMapping[op]
	if !ok {
		return C_ARITHMETIC, fmt.Errorf("unknown command type: %s", op)
	}

	return typ, nil
}

func getOrEmpty(arr []string, i int) string {
	if len(arr) <= i {
		return ""
	}

	return arr[i]
}

func splitIntoParts(str string) []string {
	ret := make([]string, 0, 3)

	buf := ""
	startCommentOut := false
	for _, r := range str {
		if r == ' ' {
			if len(buf) > 0 {
				ret = append(ret, buf)
				buf = ""
				continue
			} else {
				// skip space
				continue
			}
		}

		if r == '/' {
			if startCommentOut {
				break
			}
			startCommentOut = true
			continue
		}

		buf += string(r)
	}

	if len(buf) > 0 {
		ret = append(ret, buf)
	}

	return ret
}

func (p *Parser) CommandType() CommandType {
	return p.commandType
}

func (p *Parser) Arg1() string {
	return p.arg1
}

func (p *Parser) Arg2() int64 {
	return p.arg2
}
