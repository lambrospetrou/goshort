## GoShort

GoShort is the companion command line utility tool to create short links for long URLs or for text you want to share. The service being used is my own **Spi.to** which I strongly suggest you to check out and give me feedback.
Visit Spi.to at [http://spi.to](http://spi.to "Spi.to online service")

_Spi.to_ supports expiring links, therefore this tool also allows you to specify an expiration time. Specify the lifetime of your Spit link in seconds.

### How to use

To use `GoShort`:

* Clone the repository

  ```
  git clone https://github.com/lambrospetrou/goshort.git
  ```

* Run it directly without installing

  ```
  cd goshort
  go run goshort.go -c http://mylongurl.com 
  ```

  OR, you can install it in your $PATH with the following command (you need GOBIN and GOROOT env variables to be set) to be able to call it from anywhere!

  ```
  go install goshort.go
  ```
  
  Call it with the following command

  ```
  goshort -c http://mylongurl.com
  ```

By default GoShort will create short links for URLs that expire in 1 day so please remember to specify your preferred lifespan for the link.

The lifespan of the Spit link can be specified with the `-e` flag (measured in seconds from now), as can be seen in the following command that creates a short link for the website `http://spi.to` that **never** expires.
  
  ```
  goshort -c http://spi.to -e 0
  ```

### Options

You have the option to create short links for plain text (or source code) apart from URLs. 
The option `-t` defines the _type_ of the Spit you want to create. If you supply the word "text" you do not create a Spit for a URL link but for plain text.

  ```
  goshort -c "Hello world! This is not a URL link but it can still be shortened..." -t "text" -e 100
  ```

The above command creates a short link for the specified text which will expire in 100 seconds.

#### Summary for Flags

* `-c` - Content
  After this flag you should put the content of your spit, either the URL link or the text.

* `-e` - Expiration
  Specify the lifespan of your link in seconds (0 stands for never expiry - live forever).

* `-t` - Spit Type
  It can be either _url_ or _text_.

## Copyright

Copyright (c) 2015 _Lambros Petrou_. See LICENSE for further details.
