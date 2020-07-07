#{{ define "derivation" }}
loon-memcached-1_6_5 = pkgs.memcached.overrideAttrs(attrs: {
  version = "1.6.5";
  src = pkgs.fetchurl {
    url = "https://memcached.org/files/memcached-1.6.5.tar.gz";
    sha256 = "1f4da3706fc13c33be9df97b2c1c8d7b0891d5f0dc88aebc603cb178e68b27df";
  };
});
#{{ end }}

#{{ define "packages" }}
loon-memcached-1_6_5
#{{ end }}
