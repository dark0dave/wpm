{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable-small";
    hk = {
      url = "github:jdx/hk/v1.44.2";
    };
  };
  outputs =
    {
      self,
      nixpkgs,
      hk,
    }:
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
                hk.packages.${system}.default
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
      formatter = forEachSystem (system: nixpkgs.${system}.nixfmt);
    };
}
