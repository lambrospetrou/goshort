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

* Run it directly without installing (XXX is any positive number)

  ```
  cd goshort
  go run goshort.go -l http://mylongurl.com -exp XXX
  ```

  OR, you can install it in your $PATH with the following command (you need GOBIN and GOROOT env variables to be set)

  ```
  go install goshort.go
  ```
  
  and then call it with the following command

  ```
  goshort -l http://mylongurl.com
  ```

By default GoShort will create links that expire in 1 day so please remember to specify your preferred lifespan for the link.

The lifespan of the Spit link can be specified with the `-exp` flag (measured in seconds from now), as can be seen in the following command that creates a short link for the website `http://spi.to` that **never** expires.
  
  ```
  goshort -l http://spi.to -exp 0
  ```

## Copyright

Copyright (c) 2015 _Lambros Petrou_. See LICENSE for further details.