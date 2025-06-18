
# 🛠️ cut-tool

A minimal reimplementation of the Unix `cut` command written in Go. This tool extracts specific fields from lines of a CSV or TSV file using user-defined delimiters and field numbers.

## 🚀 Features

- Extract one or more fields from structured text input.
- Support for custom delimiters via `-d` flag.
- Defaults to tab-delimited (`.tsv`) files if no delimiter is provided.
- Compatible with UNIX-style pipelines (e.g., `| head`, `| wc`, etc.).
- Lightweight and dependency-free.

## 📦 Installation

```bash
git clone https://github.com/hiabhi-cpu/cut-tool.git
cd cut-tool
go build -o cut-tool
```

## 🔧 Usage

```bash
./cut-tool -f <field_number> [-d <delimiter>] [file]
```

## Examples
Extract the first field (default tab-delimited):

```bash
./cut-tool -f1 sample.tsv
```

Extract the second field from a CSV (comma-delimited):
```bash
./cut-tool -f2 -d, sample.csv
```

Chain with other Unix tools:

```bash
./cut-tool -f1 -d, data.csv | head -n 10
```

## ⚠️ Notes
- Field indexing starts at 1 (same as Unix cut).

- -f is required; if not provided, the program will print an error.

- -d is optional; defaults to tab (\t).
