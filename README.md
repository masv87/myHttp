#Myhttp

Command line tool which makes http requests and prints the address of the request along with the
MD5 hash of the response.

Build:  ``` go build -o myhttp```

Usage: ```./myhttp [-parallel int] URLs...``` 
```
  -parallel int
        parallel workers count (default 10)
  URLs 
        space separated list of url addresses to retrieve the content  
```
