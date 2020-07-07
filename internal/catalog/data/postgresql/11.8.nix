#{{ define "derivation" }}
loon-postgresql-11_8 = pkgs.postgresql_11.overrideAttrs(attrs: {
  version = "11.8";
  src = pkgs.fetchurl {
    url = "https://ftp.postgresql.org/pub/source/v11.8/postgresql-11.8.tar.gz";
    sha256 = "1w0vaal5d2wl30vygarbbahkhj170x5g780nnpg67nqnzgd6nq4j";
  };
});
#{{ end }}

#{{ define "packages" }}
loon-postgresql-11_8
#{{ end }}
