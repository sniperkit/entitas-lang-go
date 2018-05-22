package elang

import (
	"fmt"
	"io"
	"strings"
)

// Parser represents a parser.
type Parser struct {
	s   *Scanner
	buf struct {
		tok         Token  // last read token
		lit         string // last read literal
		isUnscanned bool   // true if you should read buf first
	}
}

// NewParser returns a new instance of Parser.
func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

// scan returns the next token from the underlying scanner.
// If a token has been unscanned then read that instead.
func (p *Parser) scan() (tok Token, lit string) {
	// If we have a token on the buffer, then return it.
	if p.buf.isUnscanned {
		p.buf.isUnscanned = false
		return p.buf.tok, p.buf.lit
	}

	// Otherwise read the next token from the scanner.
	tok, lit = p.s.Scan()

	// Save it to the buffer in case we unscan later.
	p.buf.tok, p.buf.lit = tok, lit

	return
}

// unscan pushes the previously read token back onto the buffer.
func (p *Parser) unscan() { p.buf.isUnscanned = true }

// scanIgnoreWhitespace scans the next non-whitespace token.
func (p *Parser) scanIgnoreWhitespace() (tok Token, lit string) {
	tok, lit = p.scan()
	if tok == WHITESPACE {
		tok, lit = p.scan()
	}
	return
}

func (p *Parser) parseParameter() (kv KeyValue, err error) {
	tok, _ := p.scanIgnoreWhitespace()
	if tok != LPAREN {
		p.unscan()
		return nil, nil
	}
	kv = make(KeyValue, 0)
	for {
		k, v, err := p.parseKeyValue()
		if err != nil {
			return nil, err
		}
		kv[k] = v
		tok, lit := p.scan()
		if tok == RPAREN {
			return kv, nil
		} else if tok != COMMA {
			return nil, fmt.Errorf("Parse parameter failed. Found %q, expected ','", lit)
		}
	}
}

func (p *Parser) parseString() (str string, err error) {
	tok, lit := p.scanIgnoreWhitespace()
	if tok != QUOTE {
		return "", fmt.Errorf("Parse string failed. Found %q, expected '\"'", lit)
	}
	s := ""
	for {
		tok, lit := p.scan()
		if tok == NEWLINE {
			return "", fmt.Errorf("Parse string failed. Found newline, expected '\"', word or character")
		} else if tok == QUOTE {
			break
		}
		s += lit
	}
	return s, nil
}

func (p *Parser) parseAlias() (*Alias, error) {
	a := NewAlias()
	id, err := p.parseIdentifier()
	if err != nil {
		return nil, err
	}
	a.AliasName = id
	tok, lit := p.scanIgnoreWhitespace()
	if tok != COLON {
		return nil, fmt.Errorf("Parse alias failed. Found %q, expected ':'", lit)
	}
	str, err := p.parseString()
	a.AliasValue = str
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (p *Parser) parseAliasDecl() (*AliasDecl, error) {
	ad := NewAliasDecl()
	tok, lit := p.scanIgnoreWhitespace()
	if tok != KW_ALIAS {
		return nil, fmt.Errorf("Parse alias failed. Found %q, expected 'alias'", lit)
	}
	for {
		a, err := p.parseAlias()
		ad.AddAlias(a)
		if err != nil {
			return nil, err
		}
		tok, _ := p.scan()
		if tok == NEWLINE || tok == EOF {
			return ad, nil
		}
		p.unscan()
	}
}

func (p *Parser) parseContextDecl() (*ContextDecl, error) {
	tok, lit := p.scanIgnoreWhitespace()
	if tok != KW_CONTEXT {
		return nil, fmt.Errorf("Parse context failed. Found %q, expected 'context'", lit)
	}
	cd := NewContextDecl()
	for {
		tok, lit = p.scanIgnoreWhitespace()
		if tok != WORD {
			return nil, fmt.Errorf("Parse context failed. Found %q, expected identifier", lit)
		}
		p.unscan()
		id, err := p.parseIdentifier()
		if err != nil {
			return nil, err
		}
		c := NewContext()
		c.ContextName = id
		cd.AddContext(c)
		kv, err := p.parseParameter()
		if err != nil {
			return nil, err
		}
		c.ContextParameter = kv
		tok, lit = p.scan()
		if tok == NEWLINE || tok == EOF {
			break
		} else if tok != COMMA {
			return nil, fmt.Errorf("Parse context failed. Found %q, expected ','", lit)
		}
	}
	return cd, nil
}

func (p *Parser) parseKeyValue() (key string, value string, err error) {
	k, err := p.parseIdentifier()
	if err != nil {
		return "", "", err
	}
	tok, _ := p.scanIgnoreWhitespace()
	if tok == COLON {
		v, err := p.parseIdentifier()
		if err != nil {
			return "", "", err
		}
		return k, v, nil
	}
	p.unscan()
	return k, "", nil
}

func (p *Parser) parseNamespaceDecl() (*NamespaceDecl, error) {
	tok, lit := p.scanIgnoreWhitespace()
	if tok != KW_NAMESPACE {
		return nil, fmt.Errorf("Parse namespace failed. Found %q, expected 'namespace'", lit)
	}
	ns := NewNamespaceDecl()
	nsv := ""
	for {
		id, err := p.parseIdentifier()
		if err != nil {
			return nil, err
		}
		nsv += id
		tok, lit := p.scan()
		if tok != PERIOD {
			if tok != NEWLINE {
				return nil, fmt.Errorf("Parse namespace failed. Found %q, exspected '.' or newline", lit)
			}
			ns.Namespace = nsv
			return ns, nil
		}
		nsv += "."
	}
}

func (p *Parser) parseIdentifier() (string, error) {
	s := ""
	for {
		tok, lit := p.scanIgnoreWhitespace()
		if tok != UNDERSCORE {
			p.unscan()
			break
		}
		s += lit
	}
	tok, lit := p.scan()
	if isKeyword(tok) {
		s += strings.ToLower(lit)
	} else if tok == WORD {
		s += lit
	} else {
		p.unscan()
		return "", fmt.Errorf("Parse identifier failed. Found %q, expected word", lit)
	}
	for {
		tok, lit := p.scan()
		if isKeyword(tok) {
			s += strings.ToLower(lit)
			continue
		} else if tok == WORD || tok == UNDERSCORE {
			s += lit
			continue
		}
		p.unscan()
		break
	}
	return s, nil
}

func (p *Parser) parseTargetDecl() (*TargetDecl, error) {
	t := NewTargetDecl()
	tok, lit := p.scanIgnoreWhitespace()
	if tok != KW_TARGET {
		return nil, fmt.Errorf("Parse target failed. Found %q, expected 'target'", lit)
	}
	id, err := p.parseIdentifier()
	if err != nil {
		return nil, err
	}
	t.Target = id
	tok, lit = p.scan()
	if tok != NEWLINE {
		return nil, fmt.Errorf("Parse target failed. Found %q, expected newline", lit)
	}

	return t, nil
}

// Parse one entire entitas-lang file.
func (p *Parser) Parse() (*Project, error) {
	f := NewProject()
	t, err := p.parseTargetDecl()
	if err != nil {
		return nil, err
	}
	f.TargetDecl = t
	ns, err := p.parseNamespaceDecl()
	if err != nil {
		return nil, err
	}
	f.NamespaceDecl = ns
	cd, err := p.parseContextDecl()
	if err != nil {
		return nil, err
	}
	f.ContextDecl = cd

	ad, err := p.parseAliasDecl()
	f.AddAliasDecl(ad)
	if err != nil {
		return nil, err
	}
	ad, err = p.parseAliasDecl()
	f.AddAliasDecl(ad)
	if err != nil {
		return nil, err
	}

	for {
		ad, err := p.parseAliasDecl()
		f.AddAliasDecl(ad)
		if err != nil {
			return nil, err
		}
		tok, _ := p.scan()
		p.unscan()
		if tok != KW_ALIAS {
			break
		}
	}

	return f, nil
}

// Parse one entire entitas-lang file.
func Parse(reader io.Reader) (*Project, error) {
	p := NewParser(reader)
	return p.Parse()
}
