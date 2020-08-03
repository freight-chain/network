## Görli Testnet faucet

This Faucet is based on https://github.com/poanetwork/poa-faucet

### Building from source

1. Clone repository
  ```
  git clone https://github.com/slockit/goerli-faucet
  ```
2. Copy `config.json.example` to `config.json`
  ```
  cp config.json.example config.json
  ```
2. Update config.json `./config.json` (see config.json with placeholders below)
3. Update `./public/index.html`: Find `<div class="g-recaptcha" data-sitekey="type your reCaptcha plugin secret here"></div>` line and type your reCaptcha plugin secret in `data-sitekey` attribute. For more info, [see](https://developers.google.com/recaptcha/docs/verify?hl=ru)
4. Install dependencies `npm install` from the project's root
5. Run faucet with `npm start`. the faucet will be launched at `http://localhost:5001`

### Server config.json (`./config.json`) with placeholders
```
{
  "environment": "switcher between configurations: 'prod' or 'dev'",
  "debug": "switch on/off server logs: true or false",
  "Captcha": {
    "secret": "reCaptcha plugin secret"
  },
  "Ethereum": {
    "etherToTransfer": "The number of milliEther to be sent from the faucet. For example, 500",
    "gasLimit": "Transaction gas limit, for example, 21000",
    "prod": {
      "rpc": "JSON RPC endpoint. For example, https://rpc.slock.it/goerli",
      "account": "The address from which the funds will be drained",
      "privateKey": "Private key of the account"
    },
    "dev": {
      ...
    }
  }
}
```

### Configure using environment variables

You can also configure things by using environment variables:

* `ACCOUNT` - Ethereum account to send fund from
* `KEY` - Private key for that account
* `RPC` - RPC endpoint to use 
