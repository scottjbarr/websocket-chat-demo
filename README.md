# Websocket Chat Demo

Totally pinched from the
[Heroku Example](https://github.com/heroku-examples/go-websocket-chat-demo)
to validate running and deploying a websocket server on Heroku.

Check out the [live demo](http://go-websocket-chat-demo.herokuapp.com) or [read the docs](https://devcenter.heroku.com/articles/go-websockets).

Godeps kind of suck but I got them there after a little luck.

## Run Locally

    make deps
    make run

## Deploy

Create a Heroku app if you haven't already.

    heroku create -b heroku/go

Create the Redis addon.

    heroku addons:create heroku-redis

Push the code to Heroku.

    git push origin heroku

Open the app up in a browser.

    heroku open

## Docs

- https://devcenter.heroku.com/articles/go-support
- https://devcenter.heroku.com/articles/go-dependencies-via-godep
