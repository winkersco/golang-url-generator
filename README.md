# Golang Url Generator (GOUG)

`GOUG` get list of urls and extensions then generate huge list of them (use power of concurrency in golang)

## Usage
```
   ____    ___    _   _    ____
  / ___|  / _ \  | | | |  / ___|
 | |  _  | | | | | | | | | |  _
 | |_| | | |_| | | |_| | | |_| |
  \____|  \___/   \___/   \____|

Usage:
  main [OPTIONS]

Application Options:
  -u, --urls=       Path to file that containing the urls
  -o, --output=     Path to the output file (default: output.txt)
  -w, --workers=    Number of background workers (default: 5)
  -e, --extensions= Path to file that containing extensions (default: extensions.txt)

Help Options:
  -h, --help        Show this help message

```

## Example
You can run `main.go` by passing a sample urls file and a extensions file, like below:

```bash
go run main.go -e extensions.txt -u urls.txt -o output.txt -w 100
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[GNU](https://choosealicense.com/licenses/gpl-3.0/)