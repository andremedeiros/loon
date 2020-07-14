#{{ define "derivation" }}
loon-nodejs-14_5_0 = pkgs.nodejs-14_x.overrideAttrs(attrs: {
  name = "nodejs-14.5.0";
  version = "14.5.0";
  src = pkgs.fetchurl {
    url = "https://nodejs.org/dist/v14.5.0/node-v14.5.0.tar.gz";
    sha256 = "1bn3m1z79l1adndr3v25c7cz4lmb16y6rviczln8asrw7ycjzrvd";
  };
});
#{{ end }}

#{{ define "packages" }}
loon-nodejs-14_5_0
#{{ end }}
