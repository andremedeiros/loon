package nix

type Package struct {
	Name    string
	Inherit string
	Version string
	URL     string
	SHA256  string
}
