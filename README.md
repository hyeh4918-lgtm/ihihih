# go-reloaded

A command-line tool written in Go that reads a text file, applies a set of
formatting and modification rules based on inline tags, and writes the
result to a new file.

## Usage

```bash
go run . <input-file> <output-file>
```

Example:

```bash
$ cat sample.txt
it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom
$ go run . sample.txt result.txt
$ cat result.txt
It was the best of times, it was the worst of TIMES, it was the age of wisdom
```

## Supported tags

| Tag | Effect | Example |
|---|---|---|
| `(hex)` | Converts the preceding hexadecimal word to decimal | `1E (hex)` → `30` |
| `(bin)` | Converts the preceding binary word to decimal | `10 (bin)` → `2` |
| `(up)` | Uppercases the preceding word | `go (up)` → `GO` |
| `(low)` | Lowercases the preceding word | `LOUD (low)` → `loud` |
| `(cap)` | Capitalizes the preceding word | `bridge (cap)` → `Bridge` |
| `(up, N)` / `(low, N)` / `(cap, N)` | Applies the transform to the previous `N` words instead of just one | `so exciting (up, 2)` → `SO EXCITING` |

In addition to tags, the tool automatically:

- Fixes spacing around punctuation (`,` `.` `!` `?` `:` `;`), including runs
  like `...` or `!?`, so they stay glued to the previous word with a single
  space before the next one.
- Tightens single-quote pairs (`' word '` → `'word'`), including
  multi-word spans.
- Converts `a` to `an` (and back) depending on whether the next word starts
  with a vowel or `h`.

## Input validation

Before any tag processing happens, the whole file is checked for:

- **Empty input** — a file with no content (or only whitespace) is rejected.
- **Non-ASCII characters** — any byte above the standard ASCII range causes
  a rejection.
- **Length** — files longer than 2000 characters are rejected.

While processing, any text that looks like an attempted tag but isn't
well-formed (wrong keyword, wrong casing, missing/invalid count, not enough
preceding words, an invalid hex/binary value, etc.) causes the program to
stop and report an error instead of guessing at what was intended.

On any of the above failures, the output file contains a single line
describing the error instead of transformed text.

## Line handling

The input is processed line by line. Each line is transformed
independently, and blank lines are preserved, so a 3-line input file always
produces a 3-line output file.

## Project structure

| File | Responsibility |
|---|---|
| `main.go` | Reads the input file, calls `ProcessText`, writes the output file. |
| `checker.go` | Whole-file validation: emptiness, ASCII-only, character limit. |
| `dbugs.go` | Tag detection: wraps valid tags in spaces for splitting, and flags malformed tag-like text. |
| `hundler.go` | Orchestrates the full pipeline: validation, line splitting, tag resolution, and final formatting. |
| `Hex.go` | Converts a hexadecimal string to decimal. |
| `Bin.go` | Converts a binary string to decimal. |
| `upper.go` | Uppercases the last N words in a slice. |
| `lower.go` | Lowercases the last N words in a slice. |
| `cap.go` | Capitalizes the last N words in a slice. |
| `punctuation.go` | Fixes spacing around punctuation marks. |
| `quotes.go` | Fixes spacing around single-quote pairs. |
| `vowel.go` | Converts `a` to `an` (and back) based on the following word. |

## Building

```bash
go build ./...
```