# vido
[![Go Report Card](https://goreportcard.com/badge/github.com/tcm5343/vido)](https://goreportcard.com/report/github.com/tcm5343/vido)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## What?
vido (video do) is a CLI tool, written in Go, to help manage YouTube videos downloaded using [yt-dlp](https://github.com/yt-dlp/yt-dlp)/youtube-dl/youtube-dlc. This project will be pedagogical in nature and attempt to follow these [guidelines](https://clig.dev/).

# Features (WIP):
* reading archive files of locally downloaded videos to see current status of them on YouTube (exist, deleted, private, age restricted, etc)
* updating the generated archive file to reflect videos which may be removed from the local library
* produce a CSV file with as much data as possible about the library for further analysis by a user
* produce a `video_id.log` based on the `archive.log` to use as input for historical processing
* index the videos with an ascending prefix
