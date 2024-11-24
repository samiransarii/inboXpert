import grpc

# import numpy as np
import protos.ml_message_pb2 as msgpb
import protos.ml_service_pb2 as svcpb
import protos.ml_service_pb2_grpc as svcpb_grpc

from concurrent import futures
from grpc_reflection.v1alpha import reflection


class EmailPredictionServicer(svcpb_grpc.EmailPredictionServicer):
    def __init__(self):
        # Initialize the model here
        self.model_version = "v1.0"

    def CategorizeEmail(self, request, context):
        try:
            # ML logic goes here
            # For now, returns a dummy response
            return msgpb.CategoryResponse(
                id="1",
                category="INVOICE",
                confidence=0.8,
                keywords=["invoice", "pay", "due"],
                alternatives=[
                    msgpb.AlternativeCategory(category="FINANCE", confindence=0.75)
                ],
            )
        except Exception as e:
            context.set_code(grpc.StatusCode.INTERNAL)
            context.set_details(f"Error during prediction: {str(e)}")
            return msgpb.CategoryResponse()


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    svcpb_grpc.add_EmailPredictionServicer_to_server(EmailPredictionServicer(), server)

    # Enable reflection
    SERVICE_NAMES = (
        svcpb.DESCRIPTOR.services_by_name["EmailPrediction"].full_name,
        reflection.SERVICE_NAME,
    )
    reflection.enable_server_reflection(SERVICE_NAMES, server)

    server.add_insecure_port("[::]:50055")
    server.start()
    print("ML Prediction Server started on port 50055")
    server.wait_for_termination()


if __name__ == "__main__":
    serve()
