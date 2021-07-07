# go-mal-extractor

Golang command line tool to extract a user's [myanimelist](https://myanimelist.net/).

## Usage
`go run main.go ...`

On windows if built

`./mal.exe ... `

### Flags
`-user=YOUR_USERNAME`

`-status=STATUS`

  **List of status options**
* all
* watching
* completed
* onhold
* dropped
* plantowatch
* ptw (same as plan to watch)

Run with  `-h` flag for further help

Find reference to the docs for Jikan API which this tool uses [here](https://jikan.docs.apiary.io/)