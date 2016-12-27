* Jekyll2Pelican

    $ git clone https://github.com/Juev/jekyll2pelican.git
    $ cd jekyll2pelican
    $ govendor sync
    $ govendor build
    $ mkdir source
    $ cp <jekyll_src>/_posts/* source/
    $ ./jekyll2pelican

Result files will be located in new directory "result".
