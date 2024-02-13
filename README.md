# vido
Vido (video do) is a CLI tool, written in Go, to help manage YouTube videos downloaded using [yt-dlp](https://github.com/yt-dlp/yt-dlp)/youtube-dl/youtube-dlp. 

Tasks such as:
* reading archive files of locally downloaded videos to see current status of them YouTube (exist, deleted, private, age restricted, etc)
* updating the generated archive file to reflect videos which may be removed from the local library
* produce a CSV file with as much data as possible about the library for further analysis by a user

This project will be pedagogical in nature and will attempt to follow these [guidelines](https://clig.dev/). 
