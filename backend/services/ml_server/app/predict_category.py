"""
EmailCategoryPredictor Class

A production-ready class for predicting email categories using a pre-trained model.
Handles both single email and batch prediction requests, providing category
predictions with confidence scores.

Features:
- Model and vectorizer loading from saved files
- Text preprocessing integration
- Single email prediction
- Batch prediction support
- Confidence score calculation
"""

import joblib
from utils.email_preprocessor import EmailPreProcessor


class EmailCategoryPredictor:
    """
    A class for predicting email categories using a pre-trained classification model.

    This class loads a trained model and vectorizer, processes incoming emails,
    and provides category predictions with confidence scores.

    Attributes:
        model_path (str): Path to the saved model and vectorizer file
        model: Loaded classification model
        vectorizer: Loaded TF-IDF vectorizer
        preprocessor (EmailPreProcessor): Instance of email preprocessing class
    """

    def __init__(self, model_path) -> None:
        """
        Initialize the predictor with a pre-trained model.

        Args:
            model_path (str): Path to the joblib file containing the trained model
                            and vectorizer

        Note:
            The joblib file should contain a dictionary with:
            - 'model': trained classification model
            - 'vectorizer': fitted TF-IDF vectorizer
        """
        self.model_path = model_path
        self.model = None
        self.vectorizer = None
        self.load_model()
        self.preprocessor = EmailPreProcessor()

    def load_model(self):
        """
        Load the trained model and vectorizer from the specified path.

        Loads:
            - Classification model for category prediction
            - TF-IDF vectorizer for text feature extraction

        Raises:
            FileNotFoundError: If model file is not found
            KeyError: If model file doesn't contain required components
        """
        print(f"Loading model from {self.model_path}")
        loaded_model = joblib.load(self.model_path)
        self.model = loaded_model["model"]
        self.vectorizer = loaded_model["vectorizer"]

    def predict_email(self, email_text):
        """
        Predict category and confidence score for a single email.

        Args:
            email_text (str): Combined email subject and body text

        Returns:
            tuple: (predicted_category, confidence_score)
                - predicted_category (str): Most likely email category
                - confidence_score (float): Probability score for the prediction (0-1)

        Process:
            1. Clean and preprocess the email text
            2. Transform text to TF-IDF features
            3. Predict category using the model
            4. Calculate confidence score from prediction probabilities
        """
        # Process the email text
        processed_text = self.preprocessor.clean_text(email_text)

        # Transform text to feature vector
        email_features = self.vectorizer.transform([processed_text])

        # Predict category
        predicted_category = self.model.predict(email_features)[0]

        # Calculate confidence score for the predicted category
        confidence_score = self.model.predict_proba(email_features)[
            0, self.model.classes_.tolist().index(predicted_category)
        ]

        return predicted_category, confidence_score

    def batch_predict_emails(self, email_data):
        """
        Predict categories and confidence scores for multiple emails.

        Args:
            email_data (list): List of email dictionaries, each containing:
                - email_id: Unique identifier for the email
                - subject: Email subject
                - body: Email body text

        Returns:
            list: List of dictionaries containing prediction results:
                - email_id: Original email identifier
                - predicted_category: Predicted email category
                - confidence_score: Prediction confidence (0-1)

        Note:
            Combines subject and body text for prediction, giving equal weight
            to both components.
        """
        predictions = []
        for email in email_data:
            # Combine subject and body text
            email_text = email["subject"] + " " + email["body"]

            # Get prediction and confidence
            predicted_category, confidence_score = self.predict_email(email_text)

            # Store results with email ID
            predictions.append(
                {
                    "email_id": email["email_id"],
                    "predicted_category": predicted_category,
                    "confidence_score": confidence_score,
                }
            )

        return predictions
