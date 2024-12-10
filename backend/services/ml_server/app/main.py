import grpc
import os

# import numpy as np
import protos.ml_message_pb2 as msgpb
import protos.ml_service_pb2 as svcpb
import protos.ml_service_pb2_grpc as svcpb_grpc

from concurrent import futures
from grpc_reflection.v1alpha import reflection
from .predict_category import EmailCategoryPredictor


class EmailPredictionServicer(svcpb_grpc.EmailPredictionServicer):
    def __init__(self):
        # Initialize the correct model path
        current_file_path = os.path.abspath(__file__)
        current_dir = os.path.dirname(current_file_path)
        model_file_path = os.path.join(
            current_dir, "..", "models", "email_classifier.pkl"
        )

        self.model_path = model_file_path
        self.category_predictor = EmailCategoryPredictor(self.model_path)

    def CategorizeEmail(self, request, context):
        try:
            # Extract email body and subject and make prediction
            email_text = request.subject + " " + request.body
            ml_category, ml_confidence = self.category_predictor.predict_email(
                email_text
            )

            return msgpb.CategoryResponse(
                id=request.id,
                category=ml_category,
                confidence=ml_confidence,
                keywords=[],
                alternatives=[msgpb.AlternativeCategory()],
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
