package tree

type XMLDocument struct {
	Prolog     *Prolog           `json:"prolog,omitempty"`
	Element    *XMLNode          `json:"element,omitempty"`
	Namespaces map[string]string `json:"namespaces,omitempty"`
}

func NewXMLDocument(prolog *Prolog, rootElement *XMLNode) *XMLDocument {
	return &XMLDocument{
		Prolog:  prolog,
		Element: rootElement,
	}
}

func (d *XMLDocument) WithNamespaces(nsMap map[string]string) *XMLDocument {
	d.Namespaces = nsMap
	return d
}

type XMLNode struct {
	Parent     *XMLNode          `json:"-"`
	Children   []*XMLNode        `json:"children,omitempty"`
	Name       string            `json:"name,omitempty"`
	Content    string            `json:"content,omitempty"`
	Namespace  string            `json:"namespace,omitempty"`
	Attributes map[string]string `json:"attributes,omitempty"`
	Line       int               `json:"line,omitempty"`
	Column     int               `json:"column,omitempty"`
}

func NewXMLRootNode() *XMLNode {
	return &XMLNode{}
}

func NewXMLChildNode(parent *XMLNode) *XMLNode {
	return &XMLNode{
		Parent: parent,
	}
}

func (n *XMLNode) WithPosition(line, column int) *XMLNode {
	n.Line = line
	n.Column = column
	return n
}

func (n *XMLNode) WithNSPrefix(prefix string) *XMLNode {
	n.Namespace = prefix
	return n
}

func (n *XMLNode) WithAttribute(name, value string) *XMLNode {
	if n.Attributes == nil {
		n.Attributes = make(map[string]string, 0)
	}
	n.Attributes[name] = value
	return n
}
