Todo: validation to check if url is of correct type
support different languages
scraper subcommand

go build

format:
./ao3-scraper <subcommand> <flags>

subcommands:
download:
flags:
  -end int
        last page to download
        only applies to type 1.
        default: all
  -format int
        format of download file
        0: pdf
        1: html
        2: mobi
        3: epub
        4: azw3
        default: 0
  -type int
        type of url
        0: download single fic from fic url
        1: download fics from search/bookmarks/works url
        default: 0
  -url string
        url to download from

./ao3-scraper download -url <url> -format 1 -end 4 -type 1