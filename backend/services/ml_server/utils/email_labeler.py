"""
EmailLabeler Class

A sophisticated email classification system that uses multiple signals to categorize emails
into predefined categories. It combines natural language processing, pattern matching,
and scoring mechanisms to determine the most likely category for each email.

Features:
- NLP-based text processing using spaCy
- Pattern matching for sender domains, keywords, and subjects
- Confidence scoring system
- Support for batch processing
- Detailed classification reasoning
"""

import re
import spacy
import numpy as np
from sklearn.feature_extraction.text import TfidfVectorizer
from collections import defaultdict
from category_data import category_patterns


class EmailLabeler:
    """
    A class for automatically labeling emails into predefined categories using
    multiple classification signals and NLP techniques.

    The labeler uses a combination of:
    - Sender domain matching
    - Keyword detection
    - Subject line pattern matching
    - Header analysis
    - TF-IDF vectorization for content analysis

    Attributes:
        nlp: spaCy language model for text processing
        category_patterns: Dictionary of patterns for each category
        vectorizer: TF-IDF vectorizer for text feature extraction
    """

    def __init__(self) -> None:
        """
        Initialize the EmailLabeler with required models and patterns.
        Loads spaCy model, category patterns, and initializes TF-IDF vectorizer.
        """
        # Load spaCy model for NLP processing
        self.nlp = spacy.load("en_core_web_sm")

        # Load category patterns from imported data
        self.category_patterns = category_patterns

        # Initialize TF-IDF vectorizer with specified parameters
        self.vectorizer = TfidfVectorizer(
            max_features=1000,  # Limit features to top 1000
            stop_words="english",  # Remove English stop words
            ngram_range=(1, 2),  # Use both unigrams and bigrams
        )

    def preprocess_text(self, text):
        """
        Clean and process text using spaCy NLP.

        Args:
            text (str): Raw text to process

        Returns:
            str: Processed text containing only relevant lemmatized tokens

        Note:
            Focuses on content-bearing parts of speech (nouns, verbs, adjectives)
            and removes stop words and punctuation
        """
        doc = self.nlp(text)

        # Extract relevant tokens and lemmatize
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

    def calculate_category_score(self, email, category, patterns):
        """
        Calculate a confidence score for how well an email matches a category.

        Args:
            email (dict): Email data containing subject, body, sender, and headers
            category (str): Category name being evaluated
            patterns (dict): Pattern rules for the category

        Returns:
            tuple: (score, confidence_factors)
                - score (float): Numerical score indicating category match
                - confidence_factors (list): Reasons for the score

        Note:
            Scoring weights:
            - Sender domain match: 10 points
            - Keyword match: 3 points each
            - Subject pattern match: 5 points
            - Header indicator match: 4 points
        """
        score = 0
        confidence_factors = []

        # Preprocess email content for analysis
        text = self.preprocess_text(f"{email['subject']} {email['body']}").lower()
        sender = email["sender"].lower()

        # Check sender domain (highest weight)
        for domain in patterns["senders"]:
            if domain in sender:
                score += 10
                confidence_factors.append(f"Sender matches {domain}")

        # Analyze keyword presence
        keyword_matches = 0
        for keyword in patterns["strong_keywords"]:
            if keyword in text:
                keyword_matches += 1
        score += keyword_matches * 3
        if keyword_matches > 0:
            confidence_factors.append(f"Found {keyword_matches} keywords")

        # Check subject line patterns
        for pattern in patterns["subject_patterns"]:
            if re.search(pattern, email["subject"].lower()):
                score += 5
                confidence_factors.append(f"Subject matches pattern {pattern}")

        # Analyze email headers
        headers = email["headers"]
        for indicator in patterns["header_indicators"]:
            if any(indicator in str(value).lower() for value in headers.values()):
                score += 4
                confidence_factors.append(f"Header contains {indicator}")

        return score, confidence_factors

    def label_email(self, email):
        """
        Assign most likely category to an email with confidence score.

        Args:
            email (dict): Email data containing subject, body, sender, and headers

        Returns:
            dict: Classification results containing:
                - category: Best matching category
                - confidence: Confidence score (0-1)
                - score: Raw classification score
                - factors: List of reasons for classification
        """
        scores = {}
        all_factors = {}

        # Calculate scores for all categories
        for category, patterns in self.category_patterns.items():
            score, factors = self.calculate_category_score(email, category, patterns)
            scores[category] = score
            all_factors[category] = factors

        # Determine best category
        max_score = max(scores.values())
        best_category = max(scores.items(), key=lambda x: x[1])[0]

        # Calculate confidence level (normalized to 0-1)
        confidence_level = min(max_score / 20, 1.0)

        return {
            "category": best_category,
            "confidence": confidence_level,
            "score": max_score,
            "factors": all_factors[best_category],
        }

    def batch_label_emails(self, emails):
        """
        Process and label multiple emails in batch.

        Args:
            emails (list): List of email dictionaries to classify

        Returns:
            list: List of dictionaries containing classification results for each email:
                - message_id: Original email ID
                - category: Assigned category
                - confidence: Classification confidence
                - factors: Reasoning for classification
        """
        results = []
        for email in emails:
            label_info = self.label_email(email)
            results.append(
                {
                    "message_id": email["message_id"],
                    "category": label_info["category"],
                    "confidence": label_info["confidence"],
                    "factors": label_info["factors"],
                }
            )
        return results
