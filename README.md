# 😂 DadJoke CLI

A simple and fun command-line tool to fetch and display random dad jokes! Powered by Go, the awesome [Cobra](https://github.com/spf13/cobra) library and the [Official-Joke-API](https://official-joke-api.appspot.com).

## Features

- 🤖 Fetches random dad jokes from the internet
- 👨‍💻 Programming jokes and general jokes
- 🎨 Save output to file: plain text or JSON

## Usage

```sh
joke [flags]
```

### Flags

- `-p`, `--programming`   Show programming jokes
- `-s`, `--save`          Save the joke to a file
- `-v`, `--version`       Show app version
- `-h`, `--help`          Show help message

### Examples

```sh
joke -p                   # Get a programming joke
joke --save dadjoke.json  # Get a joke in JSON format
joke --version            # Show version info
```

## Contributing

Pull requests welcome! Feel free to add more joke categories and APIs or improve the CLI. 😄

## License

MIT © 2025 @drclcomputers
