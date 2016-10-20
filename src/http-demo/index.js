var http = require('http');
var server = http.createServer(function (request, response) {
    response.writeHead(200, { "Content-Type": "text/plain" });
    response.end("Hello World\n");
    console.log(request.connection.remoteAddress, request.url, request.method)
});
var port = 3000;
console.log("starting demo application on", port);
server.listen(port);
