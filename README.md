# myhttp tool
A tool to make parallel http requests to multiple urls and print corresponding MD5 hash of the reponse body. Use <-parallel n> to indicate the parallel request limit. Default limit is 10. 
## Compile
$ go build
## Run
$ ./myhttp [-parallel n] url1 url2 url3 ...

Or

$ go run main.go [-parallel n] url1 url2 url3 ...
### Example
```
./myhttp -parallel 3 adjust.com google.com facebook.com yahoo.com yandex.com twitter.com reddit.com/r/funny reddit.com/r/notfunny baroquemusiclibrary.com
[-parallel 3 adjust.com google.com facebook.com yahoo.com yandex.com twitter.com reddit.com/r/funny reddit.com/r/notfunny baroquemusiclibrary.com]
http://google.com 4f1c969a838cdf9bc6ca702a7d938674
http://adjust.com 21aaee7adb366770ee0999ee903db700
http://yandex.com fd7d0256f58f1a4516b2a15087c847d3
http://facebook.com 0c8a06ee51a77e519d0eaedda5a905da
http://twitter.com 511982089739c61679313823309afccf
http://reddit.com/r/funny efadb7bcbd146695f2c277ac9c04742c
http://yahoo.com e5a27ac41d6a642cca935cccc7a22891
http://reddit.com/r/notfunny ff3b2b7dcd9e716ca0adcbd208061c9a
http://baroquemusiclibrary.com c1da7598e4e627c0e32b5206600e127c
```
## Run unit test
$ go test -v