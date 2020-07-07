#{{ define "derivation" }}
loon-memcached-1_6_6 = pkgs.memcached.overrideAttrs(attrs: {
  version = "1.6.6";
  src = pkgs.fetchurl {
    url = "https://memcached.org/files/memcached-1.6.6.tar.gz";
    sha256 = "908f0eecfa559129c9e44edc46f02e73afe8faca355b4efc5c86d902fc3e32f7";
  };
});
#{{ end }}

#{{ define "packages" }}
loon-memcached-1_6_6
#{{ end }}
