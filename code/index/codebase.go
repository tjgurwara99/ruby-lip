package index

type NodeType int

const (
	NodeModule NodeType = iota
	NodeClass
	NodeMethod
)

type Node interface {
	Type() NodeType
	Range() *Range
}

type ModuleDecl struct {
	Name    string
	r       *Range
	Classes []*ClassDecl
	Modules []*ModuleDecl
	Methods []*MethodDecl
}

// Range implements Node.
func (m *ModuleDecl) Range() *Range {
	return m.r
}

// Type implements Node.
func (m *ModuleDecl) Type() NodeType {
	return NodeModule
}

type Location struct {
	Line      int
	Character int
	FileURI   string
}

type Range struct {
	Start *Location
	End   *Location
}

type ClassDecl struct {
	Name    string
	r       *Range
	Methods []*MethodDecl
	Classes []*ClassDecl
	Modules []*ModuleDecl
}

// Range implements Node.
func (c *ClassDecl) Range() *Range {
	return c.r
}

// Type implements Node.
func (c *ClassDecl) Type() NodeType {
	return NodeClass
}

type MethodDecl struct {
	Name string
	r    *Range
	Args []string
}

// Range implements Node.
func (m *MethodDecl) Range() *Range {
	return m.r
}

// Type implements Node.
func (m *MethodDecl) Type() NodeType {
	return NodeMethod
}
