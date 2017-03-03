var mqtt = require('mqtt')

describe('User module', () => {
    var client
    it('connect to mqtt-proxy', (done) => {
      client = mqtt.connect('mqtt://mqtt-proxy'/*, { username: "guest", password: "guestpass"}*/)
      client.on('connect', function () {
          console.log("connected")
          done()
      })
      client.on('error', function (err) {
          console.log("error", err)
          done(err)
      })
    })

    it('pub/sub to mqtt-proxy', (done) => {

      client.on('message', function (topic, message) {
        // message is Buffer
        console.log(topic, message.toString())
        done()
      })
      client.subscribe('presence')
      client.publish('presence', 'Hello mqtt')
    })
})
