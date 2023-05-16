# Jekyll2Pelican

This is a fork that fix the project to work with the latest versions of Golang.

    $ git clone https://github.com/Juev/jekyll2pelican.git
    $ cd jekyll2pelican
    $ go mod init
    $ go mod tidy
    $ go build -mod=readonly
    $ mkdir source
    $ cp <jekyll_src>/_posts/* source/
    $ ./jekyll2pelican

Result files will be located in new directory "result".
