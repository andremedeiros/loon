#{{ define "derivation" }}
loon-nodejs-12_18_2 = pkgs.nodejs-12_x.overrideAttrs(attrs: {
  name = "nodejs-12.18.2";
  version = "12.18.2";
  src = pkgs.fetchurl {
    url = "https://nodejs.org/dist/v12.18.2/node-v12.18.2.tar.gz";
    sha256 = "052sv29v32wcws48pya8mwghmss6l9g1rwpnndg1m7lg0vb6dqrb";
  };
});
#{{ end }}

#{{ define "packages" }}
loon-nodejs-12_18_2
#{{ end }}
