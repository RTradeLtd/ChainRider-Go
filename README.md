# ChainRider-Go

Basic wrapper around the ChainRider API. Eventually aims to support all calls (events included).

## Notes

For some API calls, the authentication token is included in the body, for others it is a url param. When using a payload within the body,we format it like so:

* `"{\n \"token\": \"%s\"\n}"`

The url pattern for chain rider is as follows:

* `https://api.chainrider.io/<API_VERSION>/<DIGITAL_CURRENCY>/<BLOCKCHAIN>`

Examples:

* `https://api-dot-chainrider.io/v1/dash/main`
* `https://api-dot-chainrider.io/v1/dash/testnet`

## Supported Calls - DASH

Supported calls for the DASH ChainRider API

### Address

* Balance For Address

### Block

* Block By Hash

* Last Block Hash

### Blockchain

* Blockchain Sync Status

### Client

* Rate Limit

* Information

### Payments

* Create Payment Forward
  * Supported parameters:
    * destination address (required)
    * callback url (optional)

* Get Payment Forward By ID

* Delete Payment Forward

### Transactions

* Transaction By Hash

* Transactions For Address