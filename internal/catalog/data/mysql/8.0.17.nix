#{{ define "derivation" }}
loon-mysql-8_0_17 = pkgs.mysql80.overrideAttrs(attrs: {
  version = "8.0.17";
  src = pkgs.fetchurl {
    url = "https://cdn.mysql.com//Downloads/MySQL-8.0/mysql-8.0.17.tar.gz";
    sha256 = "1mjrlxn8vigi69r0r674j2dibdnkaar01ji5965gsyx7k60z7qy6";
  };
});
#{{ end }}

#{{ define "packages" }}
loon-mysql-8_0_17
#{{ end }}{
