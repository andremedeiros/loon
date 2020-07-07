#{{ define "derivation" }}
loon-ruby-2_7_1 = pkgs.ruby_2_7.overrideAttrs(attrs: {
  version = "2.7.1";
  src = pkgs.fetchurl {
    url = "https://cache.ruby-lang.org/pub/ruby/2.7/ruby-2.7.1.tar.gz";
    sha256 = "d418483bdd0000576c1370571121a6eb24582116db0b7bb2005e90e250eae418";
  };
  buildInputs = [pkgs.libiconv pkgs.openssl pkgs.zlib];
});
#{{ end }}

#{{ define "packages" }}
loon-ruby-2_7_1
#{{ end }}
