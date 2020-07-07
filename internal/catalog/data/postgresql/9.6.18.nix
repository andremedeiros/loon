#{{ define "derivation" }}
loon-postgresql-9_6_18 = pkgs.postgresql_9_6.overrideAttrs(attrs: {
  version = "9.6.18";
  src = pkgs.fetchurl {
    url = "https://ftp.postgresql.org/pub/source/v9.6.18/postgresql-9.6.18.tar.gz";
    sha256 = "5e00bd4ee9e92473de90f5ac05d7889699e8f3175a6f5b90514616741682f4c6";
  };
});
#{{ end }}

#{{ define "packages" }}
loon-postgresql-9_6_18
#{{ end }}
