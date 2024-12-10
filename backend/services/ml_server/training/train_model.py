"""
Model Training and Evaluation Module

This module provides functions for training and evaluating a Multinomial Naive Bayes
classifier for email classification. It includes functionality for model training
and comprehensive performance evaluation using standard metrics.

Functions:
    train_model: Trains a Multinomial Naive Bayes classifier
    evaluate_model: Evaluates model performance using classification metrics
"""

from sklearn.naive_bayes import MultinomialNB
from sklearn.metrics import classification_report, accuracy_score


def train_model(x_train, y_train):
    """
    Train a Multinomial Naive Bayes model on the provided training data.

    Args:
        x_train: array-like of shape (n_samples, n_features)
            Training data features, typically TF-IDF or similar text representations
        y_train: array-like of shape (n_samples,)
            Target values (class labels) for training data

    Returns:
        MultinomialNB: Trained Naive Bayes classifier

    Note:
        Uses default parameters for MultinomialNB:
        - alpha=1.0 (Laplace smoothing parameter)
        - fit_prior=True (learn class prior probabilities)
        - class_prior=None (automatic prior calculation)
    """
    print("Training Multinomial Naive Bayes model...")
    model = MultinomialNB()
    model.fit(x_train, y_train)
    return model


def evaluate_model(model, x_test, y_test):
    """
    Evaluate the trained model using multiple performance metrics.

    Args:
        model: Trained MultinomialNB classifier
        x_test: array-like of shape (n_samples, n_features)
            Test data features to evaluate the model on
        y_test: array-like of shape (n_samples,)
            True class labels for test data

    Prints:
        - Detailed classification report including:
            - Precision
            - Recall
            - F1-score
            - Support for each class
        - Overall model accuracy

    Note:
        The classification report includes per-class metrics and weighted averages,
        providing a comprehensive view of model performance across all categories.
    """
    print("Evaluating model...")
    # Generate predictions for test data
    y_pred = model.predict(x_test)

    # Print detailed classification metrics
    print("Classification Report:")
    print(classification_report(y_test, y_pred))

    # Calculate and print overall accuracy
    accuracy = accuracy_score(y_test, y_pred)
    print(f"Accuracy: {accuracy:.4f}")
