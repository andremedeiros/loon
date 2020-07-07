#{{ define "derivation" }}
loon-go-1_14_4 = pkgs.go.overrideAttrs(attrs: {
  version = "1.14.4";
  src = pkgs.fetchurl {
    url = "https://dl.google.com/go/go1.14.4.src.tar.gz";
    sha256 = "7011af3bbc2ac108d1b82ea8abb87b2e63f78844f0259be20cde4d42c5c40584";
  };
});
#{{ end }}

#{{ define "packages" }}
loon-go-1_14_4
#{{ end }}
