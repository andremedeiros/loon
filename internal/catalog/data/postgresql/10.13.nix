#{{ define "derivation" }}
loon-postgresql-10_13 = pkgs.postgresql_10.overrideAttrs(attrs: {
  version = "10.13";
  src = pkgs.fetchurl {
    url = "https://ftp.postgresql.org/pub/source/v10.13/postgresql-10.13.tar.gz";
    sha256 = "0gagakw8iy4rz8lf405vyw83lmg6gj39rad83yxybfryzgmch40i";
  };
});
#{{ end }}

#{{ define "packages" }}
loon-postgresql-10_13
#{{ end }}
