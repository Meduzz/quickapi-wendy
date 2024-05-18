# quickapi-wendy
Wendy take on the quickapi stuff.

Each entity gets its own wendy module. Routing is done as per normal wendy.

## Fork in the road

As with quickapi, there's 2 ways to get the thing going.

1. `Run` which you feed your entities. It will start a wendy-rpc-server and expose it over nats according to nuts.
2.  `For` which feed one of your entities and returns a wendy-module that you then have to deal with your self.

## Differences

Obviously there are differencenes. Quickapi-wendy have it's own sub-api. Ie `Create` expects a `api.Create`-struct and so on. How that struct is created and sent, is your problem :).