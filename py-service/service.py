from concurrent import futures
import grpc
import common_pb2
import common_pb2_grpc

class PyService(common_pb2_grpc.DemoServiceServicer):
    def Process(self, request, context):
        return common_pb2.Response(output=f"Python processed: {request.input}")

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    common_pb2_grpc.add_DemoServiceServicer_to_server(PyService(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    server.wait_for_termination()

if __name__ == '__main__':
    serve()