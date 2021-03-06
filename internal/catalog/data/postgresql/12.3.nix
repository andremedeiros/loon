#{{ define "derivation" }}
loon-postgresql-12_3 = pkgs.postgresql_12.overrideAttrs(attrs: {
  version = "12.3";
  src = pkgs.fetchurl {
    url = "https://ftp.postgresql.org/pub/source/v12.3/postgresql-12.3.tar.gz";
    sha256 = "708fd5b32a97577679d3c13824c633936f886a733fc55ab5a9240b615a105f50";
  };
});
#{{ end }}

#{{ define "packages" }}
loon-postgresql-12_3
#{{ end }}
