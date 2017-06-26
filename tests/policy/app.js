const express = require('express')
const fs = require('fs')
const http = require('http')
const https = require('https')
const privateKey  = fs.readFileSync('./certs/server.key', 'utf8')
const certificate = fs.readFileSync('./certs/server.pem', 'utf8')

const app = express()
const bodyParser = require('body-parser')
app.use(bodyParser.json())

process.on('SIGTERM', function () {
    process.exit(0);
});

app.post('/mqtt/connect', function (req, res) {
  console.log("/mqtt/connect", req.body)
  if (req.body.Password!="goodpass") {
    return res.status(403).send({ message: "bad password"})
  }
  res.send({ })
})

app.post('/mqtt/subscribe', function (req, res) {
  const resp={}
  console.log("/mqtt/subscribe", req.body)
  if (req.body.Topic=="bad_topic_on_subscribe") {
    return res.status(400).send({ message: "bad topic on subscribe"})
  }
  if (req.body.Topic=="rename_topic_on_subscribe") {
    resp.Topic="renamed_topic_on_subscribe"
  }
  res.send(resp)
})

app.post('/mqtt/publish', function (req, res) {
  const resp={}
  console.log("/mqtt/publish", req.body)
  if (req.body.Topic=="bad_topic_on_publish") {
    return res.status(400).send({ message: "bad topic on publish"})
  }
  if (req.body.Payload=="bad_payload_on_publish") {
    return res.status(400).send({ message: "bad payload on publish"})
  }
  if (req.body.Topic=="rename_topic_on_publish") {
    resp.Topic="renamed_topic_on_publish"
  }
  if (req.body.Payload=="alter_message_on_publish") {
    resp.Payload="altered_message_on_publish"
  }
  res.send(resp)
})

app.post('/mqtt/receive', function (req, res) {
  const resp={}
  console.log("/mqtt/receive", req.body)
  if (req.body.Topic=="bad_topic_on_receive") {
    return res.status(400).send({ message: "bad topic on receive"})
  }
  if (req.body.Payload=="bad_payload_on_receive") {
    return res.status(400).send({ message: "bad payload on receive"})
  }
  if (req.body.Topic=="rename_topic_on_receive") {
    resp.Topic="renamed_topic_on_receive"
  }
  if (req.body.Payload=="alter_message_on_receive") {
    resp.Payload="altered_message_on_receive"
  }
  res.send(resp)
})

const httpServer = http.createServer(app)
const httpsServer = https.createServer({key: privateKey, cert: certificate}, app)

httpServer.listen(3000, () => {
  console.log('Example app listening on port 3000 for http!')
})

httpsServer.listen(3001, () => {
  console.log('Example app listening on port 3001 for https!')
})
