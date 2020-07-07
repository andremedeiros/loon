#{{ define "derivation" }}
loon-postgresql-9_5_22 = pkgs.postgresql_9_5.overrideAttrs(attrs: {
  version = "9.5.22";
  src = pkgs.fetchurl {
    url = "https://ftp.postgresql.org/pub/source/v9.5.22/postgresql-9.5.22.tar.gz";
    sha256 = "0vapi50g6i0qh5bd8svf29aj6j1l5x3brk1kn1rnkski6lvfsxz9";
  };
});
#{{ end }}

#{{ define "packages" }}
loon-postgresql-9_5_22
#{{ end }}
