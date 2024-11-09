# Add-PathEntry

A command-line tool for adding an entry to the user PATH environment variable on Windows. The utility can take a relative or absolute directory path as a command-line argument or default to the current working directory if no argument is provided.

## Requirements

[Go](https://go.dev/) is required to build and install this application.

## Installation

```
go install github.com/chtozamm/Add-PathEntry@latest
```

## Usage

```
Usage: Add-PathEntry [path]
If no path is provided, the current working directory will be added to the PATH.
```

## Examples

Add a directory to the PATH using absolute filepath:

```
> Add-PathEntry C:\MyFolder
Successfully added "C:\MyFolder" to user PATH.
```

1. Add current working directory to the PATH 
2. Try to add it the second time 
3. Add a parent directory using a relative filepath

```
my-project > Add-PathEntry 
Successfully added "C:\Users\Home\my-project" to user PATH.

my-project > Add-PathEntry
"C:\Users\Home\my-project" is already in the PATH.

my-project > Add-PathEntry ../
Successfully added "C:\Users\Home" to user PATH.
```
