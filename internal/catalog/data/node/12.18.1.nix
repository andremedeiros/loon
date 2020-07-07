#{{ define "derivation" }}
loon-nodejs-12_18_1 = pkgs.nodejs-12_x.overrideAttrs(attrs: {
  name = "nodejs-12.18.1";
  version = "12.18.1";
  src = pkgs.fetchurl {
    url = "https://nodejs.org/dist/v12.18.1/node-v12.18.1.tar.gz";
    sha256 = "be532894112d06e671cb3d9388e2a3fa2c690d40bdfa9e432921a9e078b3f241";
  };
});
#{{ end }}

#{{ define "packages" }}
loon-nodejs-12_18_1
#{{ end }}
