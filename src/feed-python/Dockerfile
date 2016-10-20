FROM sercand/alpine-grpc-python:v1.0

WORKDIR "/demo/feed"

ADD color_pb2.py /demo/feed/color_pb2.py
ADD common_pb2.py /demo/feed/common_pb2.py
ADD feedprovider_pb2.py /demo/feed/feedprovider_pb2.py
ADD server.py /demo/feed/server.py

EXPOSE 3040
CMD ["python","server.py"]
