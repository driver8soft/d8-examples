import grpc
import hello_pb2
import hello_pb2_grpc

def run(inputname):
    with grpc.insecure_channel('localhost:50051') as channel:
        stub = hello_pb2_grpc.D8grpcStub(channel)
        r = stub.Hello(hello_pb2.MsgReq(hello_name=inputname))
    print(f"Result: {r.response}")

if __name__ == '__main__':
    # Get user Input 
    inputname = input("Please enter name: ")
    run(inputname)