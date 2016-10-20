const PROTO_PATH = __dirname + '/../../protos/feedprovider.proto';
const COLOR_PROTO_PATH = __dirname + '/../../protos/color.proto';

const grpc = require('grpc');
const feedpb = grpc.load(PROTO_PATH).feed.v1;
const colorpb = grpc.load(COLOR_PROTO_PATH).feed.v1;

var pod_color = '#000000';

function newItem(title, text, score) {
    return {
        title: title,
        text: text,
        color: pod_color,
        provider_name: process.env.HOSTNAME,
        score: Math.floor(score),
    }
}

function getItems(call, callback) {
    console.log('GET', call.request);
    let items = [];
    let score = 400;
    for (let i = 0; i < call.request.request.limit; i++) {
        items.push(newItem('node-title-' + i, 'text-' + i, score / (i + 1)))
    }
    callback(null, {items: items});
}

function nextColor(done) {
    var client = new colorpb.ColorGenerator(process.env.COLOR_SERVER || 'localhost:4000', grpc.credentials.createInsecure());
    client.next({}, (err, response) => {
        if (err) {
            return console.error(err);
        }
        pod_color = response.hex;
        done()
    });
}

function main() {
    let port = process.env.SERVER_PORT || '3020';
    let server = new grpc.Server();
    server.addProtoService(feedpb.FeedProvider.service, {get: getItems});
    server.bind(`0.0.0.0:${port}`, grpc.ServerCredentials.createInsecure());
    console.log('binding to', port);
    server.start();
}

nextColor(main);
