
# Cloud Foundry CLI manifest validator

This project can be used to validate Pivotal Cloud foundry deploy
manifests locally or as part of continuous integration.

Binaries are cross-compiled and checked in for the following platforms:
    - windows/amd64
    - linux/amd64
    
# Build from source

First obtain gvt as it is used as a package manager to
include the latest clf cli version from github

```sh
go get -u github.com/FiloSottile/gvt
```

```sh
gvt fetch
```

```sh
./cross_compile_linux_amd64.sh
```
# Usage

Windows:

````commandline
manifest_validator.exe fixtures\inherited-manifest-valid.yml
````

Linux:

````commandline
./manifest_validator fixtures/inherited-manifest-valid.yml
````

The details of a successfully parsed manifest will be printed
and the program exits normally.

The cli takes one argument which is the absolute or relative
path of the manifest. Inheritance is deprecated by cf cli, yet
still supported. If your manifest make use of inheritance, specify
the child since base manifests will be loaded automatically.

The program will return with a non-zero exit code if cf cli
is unable to process the manifest files. Moreover, the error
details will be printed.
