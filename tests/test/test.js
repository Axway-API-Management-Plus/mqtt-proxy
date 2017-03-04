var mqtt = require('mqtt')

function expect(a,b) {
    if (a!=b) {
        throw new Error("Expecting "+a+" == "+ b)
    }
}

describe('User module', () => {

    it('normal operation', (done) => {
      const client = mqtt.connect('mqtt://mqtt-proxy', { username: "guest", password: "goodpass"})
      let ok=false
      client.on('connect', function () {
        client.subscribe('presence')
        client.publish('presence', 'Hello mqtt')
      })
      client.on('message', function (topic, message) {
        try {
            expect(topic, 'presence')
            expect(message.toString(), 'Hello mqtt')
            ok=true
        } catch(e) {
            done(new Error("error"+e))
        }
        if(ok) done()
      })
      client.on('error', function (err) {
          if (!ok) done(err)
      })
      client.on('close', function (err) {
          if (!ok) done(err)
      })
    })

    it('bad user/password', (done) => {
        const client= mqtt.connect('mqtt://mqtt-proxy', { username: "guest", password: "badpass"})
        client.on('connect', function () {
            done(new Error("Should not connect"))
        })
        client.on('close', function (err) {
            done()
        })
    })

    it('bad topic on subscribe', (done) => {
        const client= mqtt.connect('mqtt://mqtt-proxy', { username: "bad_topic_on_subscribe", password: "goodpass"})
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
        const client= mqtt.connect('mqtt://mqtt-proxy', { username: "bad_topic_on_publish", password: "goodpass"})
        client.on('connect', function () {
            client.subscribe('bad_topic_on_publish')
            client.publish('bad_topic_on_publish', 'Hello mqtt')
        })
        client.on('message', function (topic, message) {
            done(new Error('Should not receive bad_topic_on_publish messages'))
        })
        client.on('close', function (err) {
            done()
        })
    })

    it('bad topic on receive', (done) => {
        const client= mqtt.connect('mqtt://mqtt-proxy', { username: "bad_topic_on_receive", password: "goodpass"})
        client.on('connect', function () {
            client.subscribe('bad_topic_on_receive')
            client.publish('bad_topic_on_receive', 'Hello mqtt')
        })
        client.on('message', function (topic, message) {
            let ok=false
            done(new Error('Should not receive bad_topic_on_receive messages'))
        })
        client.on('close', function (err) {
            done()
        })
    })

    it('rename topic on subscribe', (done) => {
        const client= mqtt.connect('mqtt://mqtt-proxy', { username: "rename_topic_on_subscribe", password: "goodpass"})
        let ok=false
        client.on('connect', function () {
            client.subscribe('rename_topic_on_subscribe')
            client.publish('renamed_topic_on_subscribe', 'Hello mqtt')
        })
        client.on('message', function (topic, message) {
            if (topic!="renamed_topic_on_subscribe") {
                return done(new Error('Unexpected topic :' + topic))
            }
            ok=true
            done()
            client.close()
        })
        client.on('close', function (err) {
            if (!ok) done(new Error("premature close"))
        })
    })

    it('rename topic on publish', (done) => {
        const client= mqtt.connect('mqtt://mqtt-proxy', { username: "rename_topic_on_publish", password: "goodpass"})
        let ok=false
        client.on('connect', function () {
            client.subscribe('renamed_topic_on_publish')
            client.publish('rename_topic_on_publish', 'Hello mqtt')
        })
        client.on('message', function (topic, message) {
            if (topic!="renamed_topic_on_publish") {
                return done(new Error('Unexpected topic :' + topic))
            }
            ok=true
            done()
            client.close()
        })
        client.on('close', function (err) {
            if (!ok) done(new Error("premature close"))
        })
    })

    it('rename topic on receive', (done) => {
        const client= mqtt.connect('mqtt://mqtt-proxy', { username: "rename_topic_on_receive", password: "goodpass"})
        let ok=false
        client.on('connect', function () {
            client.subscribe('rename_topic_on_receive')
            client.publish('rename_topic_on_receive', 'Hello mqtt')
        })
        client.on('message', function (topic, message) {
            if (topic!="renamed_topic_on_receive") {
                return done(new Error('Unexpected topic :' + topic))
            }
            ok=true
            done()
            client.close()
        })
        client.on('close', function (err) {
            if (!ok) done(new Error("premature close"))
        })
    })

    it('alter message on publish', (done) => {
        const client= mqtt.connect('mqtt://mqtt-proxy', { username: "alter_message_on_publish", password: "goodpass"})
        let ok=false
        client.on('connect', function () {
            client.subscribe('alter_message_on_publish')
            client.publish('alter_message_on_publish', 'alter_message_on_publish')
        })
        client.on('message', function (topic, message) {
            if (message!="altered_message_on_publish") {
                return done(new Error('Unexpected message :' + topic))
            }
            ok=true
            done()
            client.close()
        })
        client.on('close', function (err) {
            if (!ok) done(new Error("premature close"))
        })
    })
    it('alter message on receive', (done) => {
        const client= mqtt.connect('mqtt://mqtt-proxy', { username: "alter_message_on_receive", password: "goodpass"})
        let ok=false
        client.on('connect', function () {
            client.subscribe('alter_message_on_receive')
            client.publish('alter_message_on_receive', 'alter_message_on_receive')
        })
        client.on('message', function (topic, message) {
            if (message!="altered_message_on_receive") {
                return done(new Error('Unexpected message :' + topic))
            }
            ok=true
            done()
            client.close()
        })
        client.on('close', function (err) {
            if (!ok) done(new Error("premature close"))
        })
    })
})
