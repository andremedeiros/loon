#{{ define "derivation" }}
loon-go-1_13_12 = pkgs.go_1_13.overrideAttrs(attrs: {
  version = "1.13.12";
  src = pkgs.fetchurl {
    url = "https://dl.google.com/go/go1.13.12.src.tar.gz";
    sha256 = "17ba2c4de4d78793a21cc659d9907f4356cd9c8de8b7d0899cdedcef712eba34";
  };
});
#{{ end }}

#{{ define "packages" }}
loon-go-1_13_12
#{{ end }}
