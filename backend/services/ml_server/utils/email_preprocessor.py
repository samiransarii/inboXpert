"""
EmailPreProcessor Class

A specialized class for preprocessing email content for natural language processing tasks.
Handles cleaning, normalization, and feature extraction from email text data through
a combination of regex cleaning, spaCy NLP processing, and TF-IDF vectorization.

Features:
- Text cleaning and normalization
- Email-specific content handling (removing emails, URLs)
- NLP-based text processing using spaCy
- Combined subject and body text processing
- TF-IDF vectorization capability
"""

import re
import spacy
import pandas as pd
from sklearn.feature_extraction.text import TfidfVectorizer


class EmailPreProcessor:
    """
    A preprocessing class specifically designed for email content preparation.

    This class combines multiple text processing techniques to convert raw email
    content into a format suitable for machine learning and analysis tasks.

    Attributes:
        nlp: spaCy language model for advanced NLP processing
        vectorizer: TF-IDF vectorizer for text feature extraction
    """

    def __init__(self) -> None:
        """
        Initialize the EmailPreProcessor with required NLP models and vectorizers.
        Sets up:
        - spaCy language model for text processing
        - TF-IDF vectorizer with specific parameters for email content
        """
        # Initialize spaCy model for NLP processing
        self.nlp = spacy.load("en_core_web_sm")

        # Configure TF-IDF vectorizer with email-appropriate parameters
        self.vectorizer = TfidfVectorizer(
            max_features=1000,  # Limit to top 1000 features
            stop_words="english",  # Remove English stop words
            ngram_range=(1, 2),  # Use both single words and pairs
        )

    def preprocess_text(self, text):
        """
        Process text using spaCy NLP for advanced linguistic analysis.
        Args:
            text (str): Raw text input
        Returns:
            str: Processed text containing only relevant lemmatized tokens

        Note:
            Focuses on content-bearing parts of speech:
            - NOUN: Nouns
            - VERB: Verbs
            - ADJ: Adjectives
            - PROPN: Proper nouns

            Removes stop words and punctuation for cleaner output.
        """
        doc = self.nlp(text)
        tokens = [
            token.lemma_.lower()
            for token in doc
            if (
                token.pos_ in ["NOUN", "VERB", "ADJ", "PROPN"]
                and not token.is_stop
                and not token.is_punct
            )
        ]

        return " ".join(tokens)

    def clean_text(self, text):
        """
        Clean and normalize text content by removing noise and standardizing format.

        Args:
            text (str): Raw text to be cleaned
        Returns:
            str: Cleaned and normalized text

        Note:
            Performs the following cleaning operations:
            1. Converts to lowercase
            2. Removes email addresses
            3. Removes URLs (both https and www formats)
            4. Removes special characters and numbers
            5. Applies NLP preprocessing
        """
        # Handle non-string input
        if not isinstance(text, str):
            return ""

        # Convert to lowercase
        text = text.lower()

        # Remove email addresses using regex
        text = re.sub(r"\S*@\S*\s?", "", text)

        # Remove URLs (both https and www formats)
        text = re.sub(r"https\S+|www.\S+", "", text)

        # Remove special characters and numbers
        text = re.sub(r"[^a-zA-Z\s]", "", text)

        # Apply NLP preprocessing
        text = self.preprocess_text(text)
        return text

    def preprocess_email(self, email_data):
        """
        Process complete email data by combining and cleaning subject and body text.

        Args:
            email_data (dict): Dictionary containing email data with 'subject' and 'body' fields
        Returns:
            str: Processed and combined email text

        Note:
            - Handles missing/null values in subject or body
            - Combines subject and body text (subject implicitly weighted by position)
            - Applies full text cleaning pipeline
        """
        # Handle potentially missing subject or body
        subject = email_data["subject"] if pd.notnull(email_data["subject"]) else ""
        body = email_data["body"] if pd.notnull(email_data["body"]) else ""

        # Combine subject and body (subject comes first for implicit weighting)
        combined_text = subject + " " + body

        # Apply full cleaning pipeline
        cleaned_text = self.clean_text(combined_text)
        return cleaned_text
