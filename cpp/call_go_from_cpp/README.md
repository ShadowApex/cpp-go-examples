# Build
`go build -buildmode=c-shared -o libgohello.a gohello.go`    
`scons`    

# References
* http://blog.ralch.com/tutorial/golang-sharing-libraries/

### Other stuff
Go plugins (won't work with C/C++)    
`go build -buildmode=plugin -o libgohello.so gohello.go`    

Built as an archive doesn't seem to work.    
`go build -buildmode=c-archive -o libgohello.a gohello.go`    
