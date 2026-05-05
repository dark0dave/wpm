{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable-small";
  };
  outputs =
    { self, nixpkgs }:
    let
      systems = [
        "x86_64-linux"
        "aarch64-linux"
        "aarch64-darwin"
      ];
      forEachSystem = f: nixpkgs.lib.genAttrs systems (system: f system);
    in
    {
      devShells = forEachSystem (
        system:
        let
          pkgs = import nixpkgs { inherit system; };
          remote = pkgs.fetchFromGitHub {
            owner = "jdx";
            repo = "hk";
            rev = "refs/tags/v1.44.2";
            hash = "sha256-PJ8RaUeHfOVWl9wwQo5sYbuo8kap8DhtcunI6XosBCg=";
          };
          hk = pkgs.callPackage (remote + "/default.nix") { };
        in
        {
          default =
            with pkgs;
            mkShell rec {
              nativeBuildInputs = with pkgs; [
                codespell
                delve
                git
                golangci-lint
                gopls
                gotools
                hk
                nixfmt
                pre-commit
                yamlfmt
              ];
              buildInputs = with pkgs; [
                go
                openssl
              ];
              shellHook = ''
                export GOPATH="$PWD/.go"
                export PATH="$GOPATH/bin:$PATH"
                mkdir -p .go/bin
              '';
              env.HK_PKL_BACKEND = "pklr";
              env.CGO_ENABLED = 0;
            };
        }
      );
    };
}
