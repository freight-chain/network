const WebSocket = require('ws'),
      { Client } = require('pg'),
      EventEmitter = require('events'),
      { migrate } = require("postgres-migrations"),
      { resolve } = require('path');

const emitter = new EventEmitter();

(async () => {
  await migrate({
    user: process.env.DATABASE_USER,
    password: process.env.DATABASE_PASSWORD,
    host: process.env.DATABASE_HOST,
    database: process.env.DATABASE_DATABASE,
    port: parseInt(process.env.DATABASE_PORT),
  }, resolve(__dirname, './migrations/'), undefined);
})();

const client = new Client({
    user: process.env.DATABASE_USER,
    password: process.env.DATABASE_PASSWORD,
    host: process.env.DATABASE_HOST,
    database: process.env.DATABASE_DATABASE,
    port: parseInt(process.env.DATABASE_PORT),
});
(async () => {
  await client.connect()
})();

client.on('notification', function (msg) {
  const payload = msg.payload
  console.log(payload)

  // Send payload into a queue etc...
  emitter.emit('event', payload);
});

// Listen for NOTIFY calls
(async () => {
  var res = await client.query('LISTEN db_notifications')
})();

console.log(process.env.HOST, parseInt(process.env.PORT))
const wss = new WebSocket.Server({ host: process.env.HOST, port: parseInt(process.env.PORT) });

wss.on('connection', function connection(ws) {
  console.log('new connection');

  let id = '';
  
  ws.on('message', function incoming(message) {
    console.log('received: %s', message);

    switch (message['action']) {
      case 'subscribe':
        id = message['id']

        emitter.on('event', function listener(payload) {
          if (payload['input_id'] === id) {
            ws.send(payload);
          }
        })

        break


      case 'update':
        pgClient.query('UPDATE inputs SET value = ? WHERE input_id = ?', [message['id'], message['value']], (err, res) => {
          console.log(err, res)
        })

        break

      default:
        console.log('unknown action');
    }
  });
});
