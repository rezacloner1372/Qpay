# Koanf Package

koanf is a library for reading configuration from different sources in different formats in Go applications. It is a cleaner, lighter alternative to spf13/viper with better abstractions and extensibility and far fewer dependencies.

## Installation

To install the Koanf package, use the following command:

```bash
# Install the core.
go get -u github.com/knadh/koanf/v2

# Install the necessary Provider(s).
# Available: file, env, posflag, basicflag, confmap, rawbytes,
#            structs, fs, s3, appconfig/v2, consul/v2, etcd/v2, vault/v2, parameterstore/v2
# eg: go get -u github.com/knadh/koanf/providers/s3
# eg: go get -u github.com/knadh/koanf/providers/consul/v2

go get -u github.com/knadh/koanf/providers/file


# Install the necessary Parser(s).
# Available: toml, json, yaml, dotenv, hcl, hjson, nestedtext
# go get -u github.com/knadh/koanf/parsers/$parser

go get -u github.com/knadh/koanf/parsers/toml
```

### Usage

To use the Koanf package in your Go project, import it as follows:

go
Copy code

```go
import (
    "github.com/knadh/koanf/v2"
    "github.com/knadh/koanf/parsers/json"
    "github.com/knadh/koanf/parsers/yaml"
    "github.com/knadh/koanf/providers/file"
)
```

## Alternative to viper

koanf is a lightweight alternative to the popular spf13/viper. It was written as a result of multiple stumbling blocks encountered with some of viper's fundamental flaws.

* viper breaks JSON, YAML, TOML, HCL language specs by forcibly lowercasing keys.
* Significantly bloats build sizes.
* Tightly couples config parsing with file extensions.
* Has poor semantics and abstractions. Commandline, env, file etc. and various parses are hardcoded in the core. There are no primitives that can be extended.
* Pulls a large number of third party dependencies into the core package. For instance, even if you do not use YAML or flags, the dependencies are still pulled as a result of the coupling.
* Imposes arbitrary ordering conventions (eg: flag -> env -> config etc.)
* Get() returns references to slices and maps. Mutations made outside change the underlying values inside the conf map.
* Does non-idiomatic things such as throwing away O(1) on flat maps.
* Viper treats keys that contain an empty map (eg: my_key: {}) as if they were not set (ie: IsSet("my_key") == false).
* There are a large number of open issues.

## Why we use Koanf package

We incorporate the Koanf package into our project to streamline and enhance configuration management. Koanf offers a user-friendly and versatile approach to loading and retrieving configuration settings, enabling us to configure our application seamlessly through diverse sources like files, environment variables, and other options. Leveraging Koanf's capabilities, we can effortlessly fetch configuration values and efficiently handle distinct configurations tailored for various environments. This not only enhances the adaptability of our application but also contributes to its overall maintainability.
