# ImagerNamer-Go Documentation

## Command Line Interface

## Rename

### Usage

`imagenamer-go rename [arguments] [flags]`

### Descriptipon

Use this command to rename images in directories after your configuration in your config file.
Provide the paths as an argument.

### Options

```
  -e, --exclude strings   Exclude folders in given paths
  -h, --help              help for rename
  -s, --safe              Disable or enable safe rename function (default: true)
  -y, --yes               Confirm all question with yes (default: false)
```

## Global Options

```
      --config string   ocation of the configuration file (default: $HOME/.imagenamer/config.yaml)
      --debug           Show debug messages
  -h, --help            help for imagenamer-go
  -v, --version         version for imagenamer-go
```
