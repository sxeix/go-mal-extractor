# go-mal-extractor

Golang script to extract a user's [myanimelist](https://myanimelist.net/).


### Prerequisites
[Golang](https://golang.org/)

### Installation

`git clone https://github.com/sxeix/go-mal-extractor.git && cd go-mal-extractor/cmd/myanime && go install`

Then use from terminal with command

`myanime -user=YOUR_USERNAME`

### Flags
*Mandatory flags*

`-user=YOUR_USERNAME`

*Optional flags*

`-sort=BOOLEAN `
`-score=BOOLEAN`
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

Find reference to the docs for Jikan API which this script uses [here](https://jikan.docs.apiary.io/)