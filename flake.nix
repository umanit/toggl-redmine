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

        frontendDist = pkgs.buildNpmPackage {
          pname = "toggl-redmine-frontend";
          version = "0.0.0";

          src = ./frontend;

          npmDepsHash = "sha256-iW4qAO2nmER73VQtwQ1U8er86xY0NtqcowH/Q9AJwa0=";

          npmBuildScript = "build";

          installPhase = ''
            mkdir -p $out
            cp -r dist/. $out/
          '';
        };
      in
      {
        packages.default = pkgs.buildGoModule {
          pname = "toggl-redmine";
          version = "2.2.0";

          src = ./.;

          vendorHash = "sha256-ggOBBYw3eL/nYm1mvlZkCSNvExCvayDFG2jYXAk7QOM=";

          tags = [ "production" "webkit2_41" ];

          env.CGO_ENABLED = "1";

          nativeBuildInputs = [ pkgs.pkg-config ];
          buildInputs = [ pkgs.gtk3 pkgs.webkitgtk_4_1 ];

          postPatch = ''
            rm -rf frontend/dist
            cp -r ${frontendDist} frontend/dist
          '';

          postInstall = ''
            install -Dm444 build/linux/usr/local/share/pixmaps/toggl-redmine.png \
              $out/share/pixmaps/toggl-redmine.png
            install -Dm444 build/linux/usr/local/share/applications/toggl-redmine.desktop \
              $out/share/applications/toggl-redmine.desktop
            substituteInPlace $out/share/applications/toggl-redmine.desktop \
              --replace-fail 'Exec=' 'Exec=toggl-redmine'
          '';

          meta = {
            description = "Synchronise les entrées toggl track vers Redmine";
            mainProgram = "toggl-redmine";
          };
        };

        apps.default = {
          type = "app";
          program = "${self.packages.${system}.default}/bin/toggl-redmine";
        };

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
