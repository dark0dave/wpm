# Weidu package manager
A Fast and Flexible Package Manager, designed to help wiedu modders share code.

Usage:
  wpm [flags]
  wpm [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  install     Install all the dependencies in your project
  version     Version info

Flags:
  -h, --help   help for wpm

Use "wpm [command] --help" for more information about a command.

## Example wpm.yaml file

```yaml
name: "mymods"
version: "v1.0.0"
dependancies:
  git:
    - name: "BaeBG2"
      path: "https://github.com/dark0dave/BaeBG2"
      version_type: "branch"
      version: "main"
  url:
    - name: "faldorn-bg2ee"
      location: "https://downloads.weaselmods.net/download/faldorn-bg2ee/?wpdmdl=480&refresh=66ccff1d9457d1724710685"

```

## Run

Create your wpm.yaml file, see example above.

Then run:
```sh
$ wpm(.exe) install
```
the package manager will add the files to the weidu_modules folder
```sh
$ ls weidu_modules
git url
$ ls weidu_modules/url
faldorn-bg2ee.zip
$ ls weidu_modules/git
BaeBG2
```
