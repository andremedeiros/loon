#{{ define "derivation" }}
loon-mysql-8_0_21 = pkgs.mysql80.overrideAttrs(attrs: {
  version = "8.0.21";
  src = pkgs.fetchurl {
    url = "https://cdn.mysql.com//Downloads/MySQL-8.0/mysql-8.0.21.tar.gz";
    sha256 = "0d00k55rkzdgn5wj32vxankjk5x3ywfqw62zxzg3m503xrg56mmd";
  };
  patches = [];
  buildInputs = [pkgs.boost172] ++ attrs.buildInputs;
});
#{{ end }}

#{{ define "packages" }}
loon-mysql-8_0_21
#{{ end }}
