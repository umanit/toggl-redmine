# toggl-redmine

Application desktop (Wails v2, Go + React) qui synchronise les temps enregistrés sur **toggl track** vers **Redmine**.

## Fonctionnement

- Chargement des temps enregistrés sur toggl track pour une période donnée.
- Rapprochement avec les entrées de temps déjà présentes sur Redmine, pour éviter les doublons.
- Blocage de la synchronisation d'une tâche toggl track en cours d'enregistrement.
- Blocage de la synchronisation d'une tâche dont le ticket Redmine associé est fermé depuis plus de 15 jours.
- Gestion du dépassement de quota de l'API toggl track (affichage du délai avant de pouvoir réessayer).
- Envoi des entrées sélectionnées vers Redmine en un clic, depuis l'écran « Synchroniser ».

## Configuration

Les clés API et URLs de toggl track et Redmine se saisissent dans l'écran « Configurer » de l'application.

Elles sont stockées dans `~/.toggl-redmine/config.json` (fichier créé automatiquement au premier lancement). Valeurs
par défaut des URLs :

- toggl track : `https://api.track.toggl.com/api/v9`
- Redmine : `https://suivi.umanit.fr`

## Développement

L'environnement de dev est fourni par Nix + direnv (`flake.nix` / `.envrc`) : Go, la CLI Wails, Node.js et les
dépendances système GTK/WebKit nécessaires sont chargés automatiquement.

```sh
direnv allow
```

Puis, pour lancer l'application en mode développement (hot reload du frontend via Vite, devtools accessibles sur
`http://localhost:34115`) :

```sh
wails dev
```

> Le build tag Go `webkit2_41` (requis pour lier contre webkitgtk 4.1) est déjà fixé via `build:tags` dans
> `wails.json` : pas besoin de le passer à la main.

### Configurer son IDE (GOROOT)

Le shellHook du flake affiche le `GOROOT` à utiliser à chaque entrée dans l'environnement (`direnv allow` ou
`nix develop`). À défaut, `go env GOROOT` donne la même information.

## Build

```sh
wails build
```

Le binaire est généré dans `build/bin/`.

## Stack

- Go 1.25, [Wails v2.13](https://wails.io/)
- Frontend : React 18 + Vite 3 (`frontend/`)

Voir [`CHANGELOG.md`](CHANGELOG.md) pour l'historique des versions.

## Release

Les releases sont construites automatiquement par la CI GitHub Actions (`.github/workflows/main.yaml`) sur chaque
tag Git poussé : builds `linux/amd64` et `darwin/universal`, publiés en release GitHub.

Le binaire `linux/amd64` est compilé sous Ubuntu et lié dynamiquement contre les libs système à leur emplacement
standard (FHS) : il fonctionne tel quel sur Ubuntu/Debian et distributions similaires, mais **pas sur NixOS**
(erreur `error while loading shared libraries: libglib-2.0.so.0: cannot open shared object file`, NixOS n'ayant
pas cette arborescence FHS). Sur NixOS, utiliser plutôt le flake du repo :

```sh
nix run github:umanit/toggl-redmine
```

ou, depuis un clone local :

```sh
nix build .
nix run .
```
