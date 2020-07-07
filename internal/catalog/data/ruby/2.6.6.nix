#{{ define "derivation" }}
loon-ruby-2_6_6 = pkgs.ruby.overrideAttrs(attrs: {
  version = "2.6.6";
  src = pkgs.fetchurl {
    url = "https://cache.ruby-lang.org/pub/ruby/2.6/ruby-2.6.6.tar.gz";
    sha256 = "364b143def360bac1b74eb56ed60b1a0dca6439b00157ae11ff77d5cd2e92291";
  };
  buildInputs = [pkgs.libiconv pkgs.openssl pkgs.zlib];
});
#{{ end }}

#{{ define "packages" }}
loon-ruby-2_6_6
#{{ end }}
