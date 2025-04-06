from concurrent import futures
import grpc
from common_pb2 import Response
from common_pb2_grpc import DemoServiceServicer, add_DemoServiceServicer_to_server

class DemoService(DemoServiceServicer):
    def Process(self, request, context):
        return Response(output=f"Python processed: {request.input}")

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    add_DemoServiceServicer_to_server(DemoService(), server)
    server.add_insecure_port("[::]:50051")
    server.start()
    server.wait_for_termination()

if __name__ == "__main__":
    serve()