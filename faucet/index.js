const express = require('express')
const fs = require('fs')
const bodyParser = require('body-parser')
let app = express()

require('./src/helpers/blockchain-helper')(app)

process.on('SIGINT', function () {
    console.log('Exiting.')
    process.exit()
})

let config
const configPath = './config.json'
const configExists = fs.existsSync(configPath, fs.F_OK)
if (configExists) {
    config = JSON.parse(fs.readFileSync(configPath, 'utf8'))
} else {
    return console.log('There is no config.json file')
}
app.config = config

const envAccount = process.env.ACCOUNT
if (envAccount) {
    config.Ethereum.live.account = envAccount
    console.log('Using account from environment: ' + envAccount)
}

const envPk = process.env.KEY
if (envPk) {
    config.Ethereum.live.privateKey = envPk
    console.log('Using key from environment')
}

const envRpc = process.env.RPC
if (envRpc) {
    config.Ethereum.live.rpc = envRpc
    console.log('Using rpc from environment: ' + envRpc)
}

app.configureWeb3(config)
    .then((web3) => {
        app.web3 = web3
        app.use(express.static(__dirname + '/public'))
        app.use(
            bodyParser.json({
                limit: '50mb',
            })
        )
        app.use(
            bodyParser.urlencoded({
                limit: '50mb',
                extended: true,
            })
        )

        require('./src/controllers/index')(app)

        app.get('/', function (request, response) {
            response.send('Goerli Network faucet')
        })

        app.set('port', process.env.PORT || 5001)

        app.listen(app.get('port'), function () {
            console.log(
                'Goerli Network faucet is running on port',
                app.get('port')
            )
        })
    })
    .catch((error) => {
        return console.log(error)
    })
