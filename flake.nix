{
  description = "kube-no-trouble";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/2fbfb1d73d239d2402a8fe03963e37aab15abe8b";
    flake-utils.url = "github:numtide/flake-utils/";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in
      {
        devShells.default = pkgs.mkShell {
          name = "golang-env";
          # Libs
          buildInputs = with pkgs; [
            go
            openssl
          ];
          # Tools
          nativeBuildInputs = with pkgs; [
            delve
            git
            golangci-lint
            gopls
            gotools
            pre-commit
          ];
          shellHook = ''
            export CGO_ENABLED=0
            export GOPATH="$PWD/.go"
            export PATH="$GOPATH/bin:$PATH"
            mkdir -p .go/bin
          '';
        };
        formatter = pkgs.nixpkgs-fmt;
      }
    );
}
