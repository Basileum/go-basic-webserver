# Basic webserver in Go (go-basic-webserver)

Thanks to https://github.com/AkhilSharma90

Code written following the training course of freeCodeCamp.org : Learn Go Programming by Building 11 Projects – Full Course - https://www.youtube.com/watch?v=jFfo23yIWac

## Purpose of the project
This project goal is to serve three URL paths : 
- / : that will display the home page (index.html)
- /form : that will manage form submitted
- /hello : that will display a hello message


## Specific
As I was working o a specific workspace on my computer, and not the one that is set by default for GO, I add to change the env settings. 
First, I check the env variables of Go with this command :
```
go env
```
Then I set the GO111MODULE to "Auto" with this command : 
```
go env -w GO111MODULE=auto
```

