# Hipchat Message Parser

Parse mentions, emoticons and links from chat messages

# Deployed on Heroku

- [getting started with heroku](https://devcenter.heroku.com/articles/getting-started-with-go#introduction)
 - Deployed on free tier of Heroku, so first request may timeout.

# Example

`curl -XPOST https://chat-parser.herokuapp.com/parse -d '{"msg":"@bob @john (megusta) (Kappa) best website ever http://www.nbcolympics.com"}'`
