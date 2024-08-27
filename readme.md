# Weidu package manager
![](https://img.shields.io/badge/go-65A2BE2?logo=go&style=for-the-badge&logoColor=grey)
[![](https://img.shields.io/badge/Linux-FCC624?style=for-the-badge&logo=linux&logoColor=black)](https://github.com/dark0dave/wpm/releases/latest)
[![](https://img.shields.io/badge/Windows-0078D6?&style=for-the-badge&logoColor=white&logo=git-for-windows)](https://github.com/dark0dave/wpm/releases/latest)
[![](https://img.shields.io/badge/mac%20os-grey?style=for-the-badge&logo=apple&logoColor=white)](https://github.com/dark0dave/wpm/releases/latest)
[![](https://img.shields.io/github/actions/workflow/status/dark0dave/wpm/main.yaml?style=for-the-badge)](https://github.com/dark0dave/wpm/actions/workflows/main.yaml)
[![](https://img.shields.io/github/license/dark0dave/wpm?style=for-the-badge)](./LICENSE)
```sh
A Fast and Flexible Package Manager, designed to help wiedu modders share code.

Usage:
  wpm [flags]
  wpm [command]

Available Commands:
  add         Add dependencies
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  install     Install all the dependencies from your project file (wpm.yaml)
  remove      Removes dependencies
  version     Version info

Flags:
  -h, --help   help for wpm

Use "wpm [command] --help" for more information about a command.
```

## Example wpm.yaml file

```yaml
dependencies:
  git:
    - name: "BaeBG2"
      path: "https://github.com/dark0dave/BaeBG2"
      version_type: "branch"
      version: "main"
  url:
    - name: "faldorn-bg2ee"
      location: "https://downloads.weaselmods.net/download/faldorn-bg2ee/?wpdmdl=480&refresh=66ccff1d9457d1724710685"
name: "mymods"
version: "v1.0.0"
```

## Install command

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

## Add Command

```sh
Add dependencies to a manifest file

Usage:
  wpm add [flags]
  wpm add [command]

Aliases:
  add, a

Available Commands:
  git         Add git dependencies
  url         Add url dependencies

Flags:
  -h, --help   help for add

Use "wpm add [command] --help" for more information about a command.
```

### Url Dependency Example
#### Long Form

```sh
$ wpm(.exe) add url --name "test" --path "testPath.com" --version "1.0.0"
```

#### Short Form

```sh
$ wpm(.exe) a u -n "test" -p "testPath.com" -v "1.0.0"
```
#### Result

```sh
$ cat wpm.yaml
```

```yaml
dependencies:
    git:
        - name: BaeBG2
          path: https://github.com/dark0dave/BaeBG2
          version: main
          version_type: ""
    url:
        - name: faldorn-bg2ee
          path: https://downloads.weaselmods.net/download/faldorn-bg2ee/?wpdmdl=480&refresh=66ccff1d9457d1724710685
          version: ""
        - name: test
          path: testPath.com
          version: 1.0.0
name: mymods
version: v1.0.0
```

### Git Dependency Example
#### Long Form

```sh
$ wpm(.exe) add git --name "test" --path "github.com/dark0dave/BaeBG2" --version "main"
```

#### Short Form

```sh
$ wpm(.exe) a g -n "test" -p "github.com/dark0dave/BaeBG2" -v "main"
```

#### Result

```sh
$ cat wpm.yaml
```
```yaml
dependencies:
    git:
        - name: BaeBG2
          path: https://github.com/dark0dave/BaeBG2
          version_type: ""
          version: main
        - name: test
          path: github.com/dark0dave/BaeBG2
          version_type: ""
          version: main
    url:
        - name: faldorn-bg2ee
          path: https://downloads.weaselmods.net/download/faldorn-bg2ee/?wpdmdl=480&refresh=66ccff1d9457d1724710685
          version: ""
name: mymods
version: v1.0.0
```

## Rm Command

This command removes an entry to your wpm.yaml file

### Url Dependency Example
#### Long Form

```sh
$ wpm(.exe) rm url --name "test" --path "testPath.com" --version "1.0.0"
```

#### Short Form

```sh
$ wpm(.exe) r u -n "test" -p "testPath.com" -v "1.0.0"
```

### Git Dependency Example
#### Long Form

```sh
$ wpm(.exe) rm git --name "test" --path "testPath.com" --version "1.0.0"
```

#### Short Form

```sh
$ wpm(.exe) r g -n "test" -p "github.com/dark0dave/BaeBG2" -v "main"
```
