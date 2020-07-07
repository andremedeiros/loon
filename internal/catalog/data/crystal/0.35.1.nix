#{{ define "derivation" }}
loon-crystal-0_35_1-archs = {
  x86_64-darwin = "darwin-x86_64";
  x86_64-linux  = "linux-x86_64";
};

loon-crystal-0_35_1-sha256s = {
  x86_64-darwin = "7d75f70650900fa9f1ef932779bc23f79a199427c4219204fa9e221c330a1ab6";
  x86_64-linux = "6c3fd36073b32907301b0a9aeafd7c8d3e9b9ba6e424ae91ba0c5106dc23f7f9";
};


loon-crystal-0_35_1 = stdenv.mkDerivation rec {
  arch = loon-crystal-0_35_1-archs.${pkgs.stdenv.system};

  pname = "crystal";
  version = "0.35.1";
  src = pkgs.fetchurl {
    url = "https://github.com/crystal-lang/crystal/releases/download/0.35.1/crystal-${version}-1-${arch}.tar.gz";
    sha256 = loon-crystal-0_35_1-sha256s.${pkgs.stdenv.system};
  };

  buildCommand = ''
    mkdir -p $out
    tar --strip-components=1 -C $out -xf ${src}
  '';
};
#{{ end }}

#{{ define "packages" }}
loon-crystal-0_35_1
#{{ end }}
