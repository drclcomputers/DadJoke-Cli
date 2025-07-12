# 😂 DadJoke CLI

A simple and fun command-line tool to fetch and display random dad jokes! Powered by Go, the awesome [Cobra](https://github.com/spf13/cobra) library and the [Official-Joke-API](https://official-joke-api.appspot.com)

## Features

- 🤖 Fetches random dad jokes from the internet
- 👨‍💻 Programming jokes and general jokes
- 🎨 Output formatting: plain text or JSON

## Usage

```sh
joke [flags]
```

### Flags

- `-p`, `--programming`   Show programming jokes
- `-f`, `--format`        Output format: `text` (default) or `json`
- `-v`, `--version`       Show app version
- `-h`, `--help`          Show help message

### Examples

```sh
joke -p                # Get a programming joke
joke --format json     # Get a joke in JSON format
joke --version         # Show version info
```

## Output Formats

- **Text:**
  > Why do programmers prefer dark mode? 
  > - Because light attracts bugs!

- **JSON:**
  ```json
  {
    "setup": "Why do programmers prefer dark mode?",
    "punchline": "Because light attracts bugs!"
  }
  ```

## Contributing

Pull requests welcome! Feel free to add more joke categories or improve the CLI. 😄

## License

MIT © 2025 @drclcomputers
