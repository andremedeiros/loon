#{{ define "derivation" }}
loon-redis-6_0_4 = pkgs.redis.overrideAttrs(attrs: {
  version = "6.0.4";
  src = pkgs.fetchurl {
    url = "http://download.redis.io/releases/redis-6.0.4.tar.gz";
    sha256 = "3337005a1e0c3aa293c87c313467ea8ac11984921fab08807998ba765c9943de";
  };
});
#{{ end }}

#{{ define "packages" }}
loon-redis-6_0_4
#{{ end }}
