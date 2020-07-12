#{{ define "derivation" }}
loon-redis-6_0_5 = pkgs.redis.overrideAttrs(attrs: {
  version = "6.0.5";
  src = pkgs.fetchurl {
    url = "http://download.redis.io/releases/redis-6.0.5.tar.gz";
    sha256 = "15pmk3w3cjhnv40jibdavfkn446hsjn9dnpwk2w5396j2jhqdks2";
  };
});
#{{ end }}

#{{ define "packages" }}
loon-redis-6_0_5
#{{ end }}
