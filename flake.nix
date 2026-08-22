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
        in
        {
          default =
            with pkgs;
            mkShell {
              nativeBuildInputs = [
                codespell
                delve
                git
                golangci-lint
                gopls
                gotools
                hk
                musl
                nil
                nixfmt
                pre-commit
                yamlfmt
              ];
              buildInputs = [
                go
                openssl
              ];
              shellHook = ''
                export GOPATH="$PWD/.go"
                export PATH="$GOPATH/bin:$PATH"
                mkdir -p .go/bin
              '';
              env.CGO_ENABLED = 0;
              ldflags = [
                "-linkmode external"
                "-extldflags '-static -L${musl}/lib'"
              ];
            };
        }
      );
      formatter = forEachSystem (system: nixpkgs.legacyPackages.${system}.nixfmt);
    };
}
