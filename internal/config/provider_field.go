package config

type providerField string

func (f providerField) Validate() error {
	switch string(f) {
	case "nix", "docker", "homebrew":
	default:
		return ErrProviderNotSupported
	}
	return nil
}
