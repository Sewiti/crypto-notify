# Crypto notify

CLI program that notifies when a cryptocurrency triggers any rule from a given file.\
Couple rule files examples could be found in [data folder](data).

*This project is made as a recruitment task.*

## Running

There are multiple ways of running the program. Easiest would be either:

    make run

or

    go run ./cmd/cryptonotify

You can also specify your custom set of rules file as a first cli argument, example:

    ./bin/cryptonotify ./data/rules-set-2.json

## Building

You can build using `make` command or by yourself for your own needs, it's your choice. The default output is in `bin` folder.

## Testing

There are couple unit and integration tests written, all could be ran using `make test`