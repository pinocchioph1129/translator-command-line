# Translator Command Line Tool

A simple command-line language translator built in Go, using Google Cloud's Translation API. This tool allows users to translate sentences between various languages dynamically.

## Features

- Translate sentences from one language to another.
- Interactive menu to dynamically set source and target languages.
- Supports human-readable language names (e.g., "English" or "Chinese") and maps them to BCP 47 language tags.
- User-friendly error handling for invalid inputs.
- Powered by Google Cloud Translation API.

## Prerequisites

- [Go](https://golang.org/) installed on your system.
- A valid Google Cloud project with the **Translation API** enabled.
- A service account JSON key for authentication.

## Installation

1. Clone the repository:
```bash
git clone https://github.com/your-username/translator-command-line.git
cd translator-command-line
```
2.	Initialize and tidy up dependencies:
```bash
go mod init translator-command-line
go mod tidy
```
3.	Set up your Google Cloud credentials:
```
export GOOGLE_APPLICATION_CREDENTIALS="/path/to/your-service-account-key.json"
```

## Usage

Running the Tool

1.	Compile and run the program:
 ```
go run main.go --from English --to Chinese
```
2.	Follow the interactive menu to:
  
  -	Change the source language.

  -	Change the target language.
  
  -	Translate sentences dynamically.

## Example

Translate a sentence from English to Chinese:

```
./translator-command-line --from English --to Chinese
Enter the sentence to translate:
What is your name?
Translating 'What is your name?' from English (en) to Chinese (zh)...
Translated text: 你的名字是什么？
```

## File Structure
```
translator-command-line/
├── src/
│   └── bcp/
│       └── bcp.go   # Contains the language map and helper functions
├── main.go          # Main program file
├── .gitignore       # Git ignore rules
└── README.md        # Documentation
```
## Supported Languages

This project supports a wide range of languages. For a full list, see src/bcp/bcp.go.

## Contributing

Contributions are welcome! Here’s how you can contribute:
1.	Fork the repository.
2.	Create a new branch:

```
git checkout -b feature-your-feature-name
```

3.	Commit your changes:
```
git commit -m "Add feature: your-feature-name"
```

4.	Push to your fork:
```
git push origin feature-your-feature-name
```

5.	Open a pull request.

## License

This project is licensed under the MIT License. See the LICENSE file for details.
