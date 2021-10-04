## How to run

Standard Golang code (installation and so);
For the code to work download dataset from https://alphamev.ai and extract it to `./dataset` folder.

It has to be structured as so:

- Polygon:

```
block,transaction,trace
...
```

- Ethereum:

```
txHash,txData,txTrace
...
```

Then you can run the code with `go run .`

## How decoder works (currently, not in the future)

It tried to decode each transaction on its own. It parses the csv line by line (to avoid hude memory usage issue).

For each transaction it loads the trace as a recursive struct and checks every bridge on its own. This is the part where the actual decoder kicks in.

Decoder tries to recursively match a certain fingerprint trace on every level. For example I know how the Polygon Bridge Deposit transaction looks like on a trace level. It is unique because the same bridge smart contract code always creates the same trace. If it finds the correct trace fingerprint of any bridge method on any level it then further processes the transaction and checks the status and amount and so on. If everything goes through it then saves this transaction to the database with `gorm`.

## Adding a new bridge

Should be as straight forward as duplicating the PolygonPOS Goland file and replacing the trace checker with the correct fingerprint.

### TODO

Make this step a little more automatic -> make helper functions which take correct trace fingerprint as input and all other params as well (contract addresses, method and so on).

<!-- ## How Matching Engine works

NOT YET IMPLEMENTED. -->
