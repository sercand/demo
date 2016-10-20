from concurrent import futures
import time
import grpc
import feedprovider_pb2
import color_pb2
import common_pb2
import os

_ONE_DAY_IN_SECONDS = 60 * 60 * 24


class FeedProvider(feedprovider_pb2.FeedProviderServicer):
    def __init__(self, color):
        self.name = os.environ['HOSTNAME']
        self.color = color

    def Get(self, request, context):
        print("get request")
        return common_pb2.FeedGetResponse(
            items=[common_pb2.FeedItem(color=self.color, provider_name=self.name, title="python", score=500,
                                       text="python")])


def serve():
    colorServerEndpoint = os.environ["COLOR_SERVER"]
    channel = grpc.insecure_channel(colorServerEndpoint)
    stub = color_pb2.ColorGeneratorStub(channel)
    color = stub.Next(color_pb2.NewColorRequest())

    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    feedprovider_pb2.add_FeedProviderServicer_to_server(FeedProvider(color.hex), server)
    server.add_insecure_port('[::]:3040')
    print("staring server on :3040")
    server.start()
    try:
        while True:
            time.sleep(_ONE_DAY_IN_SECONDS)
    except KeyboardInterrupt:
        server.stop(0)


if __name__ == '__main__':
    serve()
