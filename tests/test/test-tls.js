const mqtt = require('mqtt')

function expect(a, b) {
    if (a != b) {
        throw new Error("Expecting " + a + " == " + b)
    }
}

class Once {
  constructor(done) {
    this.done = done
    this.flag = false
  }
  done(err, res) {
    if (!this.flag) {
      this.flag = true
      this.done(err, res)
    }
  }
}


process.env.NODE_TLS_REJECT_UNAUTHORIZED = "0"

const MQTT_BROKER = "mqtts://mqtt-proxy:1884"

describe('mqtt-proxy (TLS)', () => {

    it('normal operation (TLS)', (done) => {
      const client = mqtt.connect(MQTT_BROKER, { username: "guest", password: "goodpass", rejectUnauthorized: false })
      const once = new Once(done)
      client.on('connect', function () {
        client.subscribe('presence')
        client.publish('presence', 'Hello mqtt')
      })
      client.on('message', function (topic, message) {
        try {
            expect(topic, 'presence')
            expect(message.toString(), 'Hello mqtt')
        } catch (e) {
            return once.done(new Error("error" + e))
        }
        once.done()
        client.close()
      })
      client.on('error', function (err) {
          once.done(err)
      })
      client.on('close', function (err) {
          once.done(err)
      })
    })

    it('bad user/password', (done) => {
        const client = mqtt.connect(MQTT_BROKER, { username: "guest", password: "badpass", rejectUnauthorized: false })
        client.on('connect', function () {
            done(new Error("Should not connect"))
        })
        client.on('close', function (err) {
            done()
        })
    })

    it('bad topic on subscribe', (done) => {
        const client = mqtt.connect(MQTT_BROKER, { username: "bad_topic_on_subscribe", password: "goodpass", rejectUnauthorized: false })
        client.on('connect', function () {
            client.subscribe('bad_topic_on_subscribe')
            client.publish('bad_topic_on_subscribe', 'Hello mqtt')
        })
        client.on('message', function (topic, message) {
            done(new Error('Should not receive bad_topic_on_subscribe messages'))
        })
        client.on('close', function (err) {
            done()
        })
    })

    it('bad topic on publish', (done) => {
        const client = mqtt.connect(MQTT_BROKER, { username: "bad_topic_on_publish", password: "goodpass", rejectUnauthorized: false })
        let once = new Once(done)
        client.on('connect', function () {
            client.subscribe('bad_topic_on_publish')
            client.publish('bad_topic_on_publish', 'Hello mqtt')
        })
        client.on('message', function (topic, message) {
            once.done(new Error('Should not receive bad_topic_on_publish messages'))
        })
        client.on('close', function (err) {
            once.done()
        })
    })

    it('bad topic on receive', (done) => {
        const client = mqtt.connect(MQTT_BROKER, { username: "bad_topic_on_receive", password: "goodpass", rejectUnauthorized: false })
        client.on('connect', function () {
            client.subscribe('bad_topic_on_receive')
            client.publish('bad_topic_on_receive', 'Hello mqtt')
        })
        client.on('message', function (topic, message) {
            done(new Error('Should not receive bad_topic_on_receive messages'))
        })
        client.on('close', function (err) {
            done()
        })
    })

    it('rename topic on subscribe', (done) => {
        const client = mqtt.connect(MQTT_BROKER, { username: "rename_topic_on_subscribe", password: "goodpass", rejectUnauthorized: false })
        let once = new Once(done)
        client.on('connect', function () {
            client.subscribe('rename_topic_on_subscribe')
            client.publish('renamed_topic_on_subscribe', 'Hello mqtt')
        })
        client.on('message', function (topic, message) {
            if (topic != "renamed_topic_on_subscribe") {
                return done(new Error('Unexpected topic :' + topic))
            }
            once.done()
            client.close()
        })
        client.on('close', function (err) {
            once.done(new Error("premature close"))
        })
    })

    it('rename topic on publish', (done) => {
        const client = mqtt.connect(MQTT_BROKER, { username: "rename_topic_on_publish", password: "goodpass", rejectUnauthorized: false })
        const once = new Once(done)
        client.on('connect', function () {
            client.subscribe('renamed_topic_on_publish')
            client.publish('rename_topic_on_publish', 'Hello mqtt')
        })
        client.on('message', function (topic, message) {
            if (topic != "renamed_topic_on_publish") {
                return once.done(new Error('Unexpected topic :' + topic))
            }
            once.done()
            client.close()
        })
        client.on('close', function (err) {
            once.done(new Error("premature close"))
        })
    })

    it('rename topic on receive', (done) => {
        const client = mqtt.connect(MQTT_BROKER, { username: "rename_topic_on_receive", password: "goodpass", rejectUnauthorized: false })
        const once = new Once(done)
        client.on('connect', function () {
            client.subscribe('rename_topic_on_receive')
            client.publish('rename_topic_on_receive', 'Hello mqtt')
        })
        client.on('message', function (topic, message) {
            if (topic != "renamed_topic_on_receive") {
                return once.done(new Error('Unexpected topic :' + topic))
            }
            once.done()
            client.close()
        })
        client.on('close', function (err) {
            once.done(new Error("premature close"))
        })
    })

    it('alter message on publish', (done) => {
        const client = mqtt.connect(MQTT_BROKER, { username: "alter_message_on_publish", password: "goodpass", rejectUnauthorized: false })
        const once = new Once(done)
        client.on('connect', function () {
            client.subscribe('alter_message_on_publish')
            client.publish('alter_message_on_publish', 'alter_message_on_publish')
        })
        client.on('message', function (topic, message) {
            if (message != "altered_message_on_publish") {
                return done(new Error('Unexpected message :' + topic))
            }
            once.done()
            client.close()
        })
        client.on('close', function (err) {
            once.done(new Error("premature close"))
        })
    })
    it('alter message on receive', (done) => {
        const client = mqtt.connect(MQTT_BROKER, { username: "alter_message_on_receive", password: "goodpass", rejectUnauthorized: false })
        const once = new Once(done)
        client.on('connect', function () {
            client.subscribe('alter_message_on_receive')
            client.publish('alter_message_on_receive', 'alter_message_on_receive')
        })
        client.on('message', function (topic, message) {
            if (message != "altered_message_on_receive") {
                return done(new Error('Unexpected message :' + topic))
            }
            once.done()
            client.close()
        })
        client.on('close', function (err) {
            once.done(new Error("premature close"))
        })
    })
})
