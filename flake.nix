{
  description = "Environnement de développement pour toggl-redmine (Wails v2 + React)";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in
      {
        devShells.default = pkgs.mkShell {
          packages = [
            pkgs.go
            pkgs.wails
            pkgs.nodejs_24
            pkgs.pkg-config
            pkgs.gtk3
            pkgs.webkitgtk_4_1
          ];

          GOFLAGS = "-tags=webkit2_41";

          shellHook = ''
            echo "toggl-redmine dev shell"
            echo "  go:    $(go version)"
            echo "  GOROOT: $(go env GOROOT)"
            echo "  wails: $(wails version)"
            echo "  node:  $(node --version)"
            echo "  npm:   $(npm --version)"
          '';
        };
      });
}
