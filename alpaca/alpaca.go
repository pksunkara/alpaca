package alpaca

type Module struct {
	Name    string
	Module  string
	Version string

	Official bool

	License string

	Author struct {
		Name  string
		Email string
		Url   string
	}

	Git struct {
		Site string
		User string
		Name string
	}
}
