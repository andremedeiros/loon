#{{ define "derivation" }}
loon-nodejs-14_4_0 = pkgs.nodejs-14_x.overrideAttrs(attrs: {
  name = "nodejs-14.4.0";
  version = "14.4.0";
  src = pkgs.fetchurl {
    url = "https://nodejs.org/dist/v14.4.0/node-v14.4.0.tar.gz";
    sha256 = "1xxpdpf2w7hpmil2bggmn2w6a9kj0jrflr4xmf6z7qqryrncwsap";
  };
});
#{{ end }}

#{{ define "packages" }}
loon-nodejs-14_4_0
#{{ end }}
