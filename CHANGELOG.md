# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Ajout d'un environnement de développement Nix (`flake.nix` + `.envrc`) fournissant Go, la CLI Wails, Node.js et les
  dépendances système GTK/WebKit nécessaires

### Changed

- Mise à jour de wails/v2 vers la v2.13.0
- Mise à jour de viper vers la v1.21.0
- Mise à jour de la directive `go` du module vers 1.25 et suppression du verrou `toolchain` (désormais géré par
  l'environnement Nix)
- Mise à jour des dépendances frontend : bootstrap 5.3.8, react/react-dom 18.3.1, react-bootstrap 2.10.10,
  react-router-dom 6.30.4, @types/react 18.3.31, @types/react-dom 18.3.7, @vitejs/plugin-react 2.2.0, vite 3.2.11
- Suppression du pin Volta dans `frontend/package.json` (Node désormais fourni par l'environnement Nix)
- Fixation du build tag Go `webkit2_41` via `build:tags` dans `wails.json`, requis pour lier contre webkitgtk 4.1

### Removed

- Suppression d'une directive `replace` obsolète et inactive dans `go.mod`

## [2.1.1] - 2024-10-18

### Fixed

- Le préfixe des numéros de ticket peut comporter des caractères unicodes

## [2.1.0] - 2024-10-11

### Added

- Permet de matcher le n° de ticket même s’il y a du texte avant

## [2.0.2] - 2024-09-30

### Fixed

- Désactivation de l’accélération matérielle

## [2.0.1] - 2024-09-30

### Fixed

- Mise à jour du CLI de wails
- Mise à jour de l’action GitHub pour build les releases

## [2.0.0] - 2024-06-02

Première version pour la 2.0 🚀

[Unreleased]: https://github.com/umanit/toggl-redmine/compare/2.1.1...HEAD

[2.1.1]: https://github.com/umanit/toggl-redmine/compare/2.1.0...2.1.1

[2.1.0]: https://github.com/umanit/toggl-redmine/compare/2.0.2...2.1.0

[2.0.2]: https://github.com/umanit/toggl-redmine/compare/2.0.1...2.0.2

[2.0.1]: https://github.com/umanit/toggl-redmine/compare/2.0.0...2.0.1

[2.0.0]: https://github.com/umanit/toggl-redmine/releases/tag/2.0.0
