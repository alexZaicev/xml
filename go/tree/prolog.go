package tree

type Prolog struct {
	Version    string `json:"version,omitempty"`
	Standalone string `json:"standalone,omitempty"`
	Encoding   string `json:"encoding,omitempty"`
}

func NewProlog() *Prolog {
	return &Prolog{}
}

func (p *Prolog) WithVersion(version string) *Prolog {
	p.Version = version
	return p
}

func (p *Prolog) WithStandalone(standalone string) *Prolog {
	p.Standalone = standalone
	return p
}

func (p *Prolog) WithEncoding(encoding string) *Prolog {
	p.Encoding = encoding
	return p
}
