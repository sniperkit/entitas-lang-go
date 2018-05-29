package elang

import (
	"fmt"
	"io"
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
	if p.buf.isUnscanned {
		p.buf.isUnscanned = false
		return p.buf.tok, p.buf.lit
	}
	tok, lit = p.s.Scan()
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

// parseKeyValue `Key:Value` ...
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

// parseKeyValueList `Key_One:Value, Key_Two:Value, Key_Three:Value` ...
func (p *Parser) parseKeyValueList() (kv KeyValue, err error) {
	kv = NewKeyValue()
	for {
		k, v, err := p.parseKeyValue()
		if err != nil {
			return nil, err
		}
		kv[k] = v
		tok, _ := p.scan()
		if tok == COMMA {
			continue
		}
		p.unscan()
		return kv, nil
	}
}

// parseParameter `(Key_One:Value, Key_Two:Value, Key_Three:Value)` ...
func (p *Parser) parseParameter() (kv KeyValue, err error) {
	tok, lit := p.scanIgnoreWhitespace()
	if tok != LPAREN {
		p.unscan()
		return nil, nil
	}
	kv, err = p.parseKeyValueList()
	if err != nil {
		return nil, err
	}
	tok, lit = p.scan()
	if tok != RPAREN {
		p.unscan()
		return nil, fmt.Errorf("EntitasLang: Parse parameter failed. Found '%s', expected ')'", lit)
	}
	return kv, nil
}

// parseString `"this is a string"` ...
func (p *Parser) parseString() (str string, err error) {
	tok, lit := p.scanIgnoreWhitespace()
	if tok != QUOTE {
		return "", fmt.Errorf("EntitasLang: Parse string failed. Found '%s', expected '\"'", lit)
	}
	s := ""
	for {
		tok, lit := p.scan()
		if tok == NEWLINE {
			return "", fmt.Errorf("EntitasLang: Parse string failed. Found newline, expected '\"', word or character")
		} else if tok == QUOTE {
			break
		}
		s += lit
	}
	return s, nil
}

// parseIdentifier `my_id` ...
func (p *Parser) parseIdentifier() (string, error) {
	s := ""
	tok, lit := p.scanIgnoreWhitespace()
	if isKeyword(tok) || tok == WORD || tok == UNDERSCORE {
		s += lit
	} else {
		return "", fmt.Errorf("EntitasLang: Parse identifier failed. Found '%s', expected keyword, word or underscore", lit)
	}
	for {
		tok, lit := p.scan()
		if isKeyword(tok) || tok == WORD || tok == UNDERSCORE {
			s += lit
			continue
		}
		p.unscan()
		if containsOnly(s, '_') {
			return "", fmt.Errorf("EntitasLang: Parse identifier failed. Identifier cannot consist only of \"_\"")
		}
		return s, nil
	}
}

// parseIdentifierList `Key_One, Key_Two, Key_Three` ...
func (p *Parser) parseIdentifierList() (l []string, err error) {
	l = make([]string, 0)
	for {
		id, err := p.parseIdentifier()
		if err != nil {
			return nil, err
		}
		l = append(l, id)
		tok, _ := p.scan()
		if tok == COMMA {
			continue
		}
		p.unscan()
		return l, nil
	}
}

// parseTargetDecl `target entitas_csharp` ...
func (p *Parser) parseTargetDecl() (*TargetDecl, error) {
	t := NewTargetDecl()
	tok, lit := p.scanIgnoreWhitespace()
	if tok != KW_TARGET {
		return nil, fmt.Errorf("EntitasLang: Parse target failed. Found '%s', expected 'target'", lit)
	}
	id, err := p.parseIdentifier()
	if err != nil {
		return nil, err
	}
	t.Target = id
	tok, lit = p.scan()
	return t, nil
}

// parseNamespace `my.game.name_space` ...
func (p *Parser) parseNamespace() (string, error) {
	nsv := ""
	for {
		id, err := p.parseIdentifier()
		if err != nil {
			return "", err
		}
		nsv += id
		tok, _ := p.scan()
		if tok != PERIOD {
			return nsv, nil
		}
		nsv += "."
	}
}

// parseNamespaceDecl `namespace my.game.name_space` ...
func (p *Parser) parseNamespaceDecl() (*NamespaceDecl, error) {
	ns := NewNamespaceDecl()
	tok, lit := p.scanIgnoreWhitespace()
	if tok != KW_NAMESPACE {
		return nil, fmt.Errorf("EntitasLang: Parse namespace failed. Found '%s', expected 'namespace'", lit)
	}
	str, err := p.parseNamespace()
	if err != nil {
		return nil, err
	}
	ns.Namespace = str
	return ns, nil
}

// parseContext `my_game` `my_game (key_one:value)` `my_game (key_one:value, key_two:value)` ...
func (p *Parser) parseContext() (*Context, error) {
	c := NewContext()
	id, err := p.parseIdentifier()
	if err != nil {
		return nil, err
	}
	c.Name = id
	kv, err := p.parseParameter()
	if err != nil {
		return nil, err
	}
	if kv != nil {
		c.Parameter = kv
	}
	return c, nil
}

// parseContextDecl `context my_game` `context my_game (key:value), second_context(key:value)` ...
func (p *Parser) parseContextDecl() (*ContextDecl, error) {
	cd := NewContextDecl()
	tok, lit := p.scanIgnoreWhitespace()
	if tok != KW_CONTEXT {
		return nil, fmt.Errorf("EntitasLang: Parse context failed. Found '%s', expected 'context'", lit)
	}
	for {
		c, err := p.parseContext()
		if err != nil {
			return nil, err
		}
		cd.AddContext(c)
		tok, lit = p.scan()
		if tok == NEWLINE || tok == EOF {
			break
		} else if tok != COMMA {
			return nil, fmt.Errorf("EntitasLang: Parse context failed. Found '%s', expected ','", lit)
		}
	}
	return cd, nil
}

// parseAlias `my_int : "int"` ...
func (p *Parser) parseAlias() (*Alias, error) {
	a := NewAlias()
	id, err := p.parseIdentifier()
	if err != nil {
		return nil, err
	}
	a.Name = id
	tok, lit := p.scanIgnoreWhitespace()
	if tok != COLON {
		return nil, fmt.Errorf("EntitasLang: Parse alias failed. Found '%s', expected ':'", lit)
	}
	str, err := p.parseString()
	a.Value = str
	if err != nil {
		return nil, err
	}
	return a, nil
}

// parseAlias `alias my_int : "int"` `alias my_int : "int" my_string : "string"` ...
func (p *Parser) parseAliasDecl() (*AliasDecl, error) {
	ad := NewAliasDecl()
	tok, lit := p.scanIgnoreWhitespace()
	if tok != KW_ALIAS {
		return nil, fmt.Errorf("EntitasLang: Parse alias failed. Found '%s', expected 'alias'", lit)
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

// parseComponentDecl `comp my_component (key:value) in game` ...
func (p *Parser) parseComponentDecl() (*ComponentDecl, []string, error) {
	comp := NewComponentDecl()
	list := []string{}
	tok, lit := p.scanIgnoreWhitespace()
	if tok != KW_COMP {
		return nil, nil, fmt.Errorf("EntitasLang: Parse component failed. Found '%s', expected 'comp'", lit)
	}
	id, err := p.parseIdentifier()
	if err != nil {
		return nil, nil, err
	}
	comp.Name = id
	kv, err := p.parseParameter()
	if err != nil {
		return nil, nil, err
	}
	if kv != nil {
		comp.Parameter = kv
	}
	tok, _ = p.scanIgnoreWhitespace()
	if tok == KW_IN {
		list, err = p.parseIdentifierList()
		if err != nil {
			return nil, nil, err
		}
		tok, _ = p.scan()
		if tok != NEWLINE {
			p.unscan()
		}
	}
	tok, _ = p.scan()
	if tok != NEWLINE {
		p.unscan()
	}
	return comp, list, nil
}

// Parse one entire entitas-lang file.
func (p *Parser) Parse() (*Project, error) {
	project := NewProject()
	t, err := p.parseTargetDecl()
	if err != nil {
		return nil, err
	}
	project.TargetDecl = t
	ns, err := p.parseNamespaceDecl()
	if err != nil {
		return nil, err
	}
	project.NamespaceDecl = ns
	cd, err := p.parseContextDecl()
	if err != nil {
		return nil, err
	}
	project.ContextDecl = cd

	for {
		tok, _ := p.scan()
		p.unscan()
		if tok != KW_ALIAS {
			break
		}
		ad, err := p.parseAliasDecl()
		if err != nil {
			return nil, err
		}
		project.AddAliasDecl(ad)
	}

	for {
		tok, _ := p.scan()
		p.unscan()
		if tok != KW_COMP {
			break
		}
		Component, ContextList, err := p.parseComponentDecl()
		if err != nil {
			return nil, err
		}
		for _, Context := range ContextList {
			if project.ContextDecl.GetContextWithName(Context) == nil {
				return nil, fmt.Errorf("EntitasLang: Component(%s) Context '%s' is not defined!", Component.Name, Context)
			}
		}
		Component.Context = project.ContextDecl.GetContextSliceWithName(ContextList...)
		project.AddComponentDecl(Component)
	}

	for _, Context := range project.ContextDecl.Context {
		if len(project.ContextDecl.GetContextSliceWithName(Context.Name)) > 1 {
			return nil, fmt.Errorf("EntitasLang: Context '%s' is already defined!", Context.Name)
		}
	}

	return project, nil
}

// Parse one entire entitas-lang file.
func Parse(reader io.Reader) (*Project, error) {
	p := NewParser(reader)
	return p.Parse()
}
