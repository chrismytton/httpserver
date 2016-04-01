# Static http server

Ridiculously simple http server for static files.

## Install

On OS X you can install with Homebrew:

    $ brew install chrismytton/formula/httpserver

Or you can install from source using Go:

    $ go get github.com/chrismytton/httpserver

## Usage

To serve the current directory on port 8080:

    $ httpserver

To use a different port specify with the `-port` flag:

    $ httpserver -port 5000

To serve a different directory use the `-root` flag:

    $ httpserver -root public

## Options

`-port` Defines the TCP port to listen on. (Defaults to 8080).

`-root` Defines the directory to serve. (Defaults to the current directory).

## Credits

Inspired by
[nodeapps/http-server](https://github.com/nodeapps/http-server).

## MIT License

Copyright (c) Chris Mytton

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject
to the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR
ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF
CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
