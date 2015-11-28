# Websocket Chat Demo

Heavily borrowed from the
[Heroku Example](https://github.com/heroku-examples/go-websocket-chat-demo)
to validate running and deploying a websocket server on Heroku.

Check out the [live demo](http://go-websocket-chat-demo.herokuapp.com) or [read the docs](https://devcenter.heroku.com/articles/go-websockets).

## Deploy

Create a Heroku app if you haven't already.

    heroku create -b heroku/go

Add the redis addon.

    heroku addons:create heroku-redis

## Docs

- https://devcenter.heroku.com/articles/go-support
- https://devcenter.heroku.com/articles/go-dependencies-via-godep
