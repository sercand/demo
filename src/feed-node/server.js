const PROTO_PATH = __dirname + '/../../protos/feedprovider.proto';
const grpc = require('grpc');
const feed_proto = grpc.load(PROTO_PATH).feed.v1

function newItem() {

}

function getItems(call, callback) {
    console.log("GET", call.request);
    callback(null, {items: [{title:"ok"},{title:"ok1"}]});
}

function main() {
    var server = new grpc.Server();
    server.addProtoService(feed_proto.FeedProvider.service, {get: getItems});
    server.bind('0.0.0.0:3020', grpc.ServerCredentials.createInsecure());
    server.start();
}

main();