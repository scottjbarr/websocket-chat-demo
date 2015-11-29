# Websocket Chat Demo

Totally pinched from the
[Heroku Example](https://github.com/heroku-examples/go-websocket-chat-demo)
to validate running and deploying a websocket server on Heroku.

Check out the [live demo](http://go-websocket-chat-demo.herokuapp.com) or [read the docs](https://devcenter.heroku.com/articles/go-websockets).

Godeps kind of suck but I got them there after a little luck.

End result was a websocket server that didn't really work on Heroku. Messages
could not be read by server. The same code deployed to an instance on
DigitalOcean worked well. I suspect interference from a proxy in Heroku world.

Bottom line is it doesn't work on Heroku, but it works well elsewhere.

See the play app run nicely on a [DigitalOcean instance](http://wsdemo.scottjbarr.com:5000/)
*(only available during the Golang Meetup Talk)*

### Todo

- More investigation into Heroku fail needed.
- Put a proxy (e.g. HAProxy) in from of the websocket server on the
  DigitalOcean instance. I've seen this work in the past.
- Don't use Heroku for Websocket server? Sometimes free is only free if your
  time is worth nothing.

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

## To Really Deploy

Push this to a server that isn't Heroku :)

## Docs

- https://devcenter.heroku.com/articles/go-support
- https://devcenter.heroku.com/articles/go-dependencies-via-godep
