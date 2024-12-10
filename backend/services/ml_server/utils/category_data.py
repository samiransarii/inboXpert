"""
Email Classification Pattern Dictionary

This dictionary defines comprehensive patterns and rules for classifying emails into different categories.
Each category contains four types of classification signals:
- senders: Domain patterns that typically send emails of this category
- strong_keywords: Words and phrases strongly associated with the category
- subject_patterns: Regular expressions to match subject lines
- header_indicators: Special header tags that indicate the category
"""

category_patterns = {
    "PERSONAL": {
        "senders": ["@gmail.com", "@yahoo.com", "@hotmail.com", "@outlook.com"],
        # Keywords for identifying personal and social communications
        "strong_keywords": [
            "family",
            "personal",
            "friend",
            "dinner",
            "lunch",
            "weekend",
            "vacation",
            "holiday",
            "birthday",
            "social",
            "meetup",
        ],
        "subject_patterns": [
            r"\b(hey|hi|hello)\b",  # Casual greetings
            r"personal matter",  # Direct personal indicators
            r"family|friend",  # Relationship indicators
        ],
        "header_indicators": ["personal", "private", "confidential"],
    },
    "CAREER": {
        "senders": [
            "@linkedin.com",
            "@indeed.com",
            "@careers.",
            "@recruit.",
            "@talent.",
            "@hr.",
            "@jobs.",
        ],
        # Keywords for identifying job and career-related communications
        "strong_keywords": [
            "job",
            "career",
            "position",
            "opportunity",
            "resume",
            "cv",
            "interview",
            "recruitment",
            "hiring",
            "salary",
            "application",
            "qualification",
            "experience",
            "skill",
        ],
        "subject_patterns": [
            r"job (opportunity|opening|position)",
            r"career|position|vacancy",
            r"interview|recruitment",
        ],
        "header_indicators": ["job", "career", "recruitment", "employment"],
    },
    "FINANCE": {
        "senders": [
            "@bank.",
            "@paypal.",
            "@billing.",
            "@finance.",
            "@accounting.",
            "@invoice.",
            "@tax.",
        ],
        # Keywords for identifying financial and banking communications
        "strong_keywords": [
            "transaction",
            "payment",
            "invoice",
            "bank",
            "financial",
            "account",
            "balance",
            "credit",
            "debit",
            "statement",
            "tax",
            "investment",
            "money",
            "fund",
        ],
        "subject_patterns": [
            r"payment|invoice|transaction",
            r"financial|statement|balance",
            r"tax|investment",
        ],
        "header_indicators": ["finance", "banking", "payment", "invoice"],
    },
    "SHOPPING": {
        "senders": [
            "@amazon.",
            "@ebay.",
            "@walmart.",
            "@shop.",
            "@store.",
            "@retail.",
            "@order.",
        ],
        # Keywords for identifying shopping and order-related communications
        "strong_keywords": [
            "order",
            "purchase",
            "delivery",
            "shipping",
            "tracking",
            "item",
            "product",
            "cart",
            "buy",
            "price",
            "discount",
            "deal",
            "sale",
            "shop",
        ],
        "subject_patterns": [
            r"order (confirmation|status|shipped)",
            r"delivery|tracking",
            r"purchase|shopping",
        ],
        "header_indicators": ["order", "purchase", "shipping", "delivery"],
    },
    "HEALTH": {
        "senders": [
            "@hospital.",
            "@clinic.",
            "@healthcare.",
            "@medical.",
            "@health.",
            "@doctor.",
        ],
        # Keywords for identifying health and medical communications
        "strong_keywords": [
            "appointment",
            "medical",
            "health",
            "doctor",
            "clinic",
            "prescription",
            "medication",
            "treatment",
            "insurance",
            "patient",
            "healthcare",
            "wellness",
        ],
        "subject_patterns": [
            r"medical|health|appointment",
            r"doctor|clinic|hospital",
            r"prescription|medication",
        ],
        "header_indicators": ["medical", "health", "healthcare", "appointment"],
    },
    "SUBSCRIPTIONS": {
        "senders": [
            "@netflix.",
            "@spotify.",
            "@subscription.",
            "@service.",
            "@membership.",
        ],
        # Keywords for identifying subscription and membership communications
        "strong_keywords": [
            "subscription",
            "membership",
            "renewal",
            "plan",
            "service",
            "account",
            "streaming",
            "monthly",
            "annual",
            "auto-renewal",
            "premium",
        ],
        "subject_patterns": [
            r"subscription|membership",
            r"renewal|plan|service",
            r"account (status|update)",
        ],
        "header_indicators": [
            "subscription",
            "membership",
            "service",
            "account",
        ],
    },
    "BUSINESS COMMUNICATION": {
        "senders": [
            "@company.",
            "@corp.",
            "@business.",
            "@team.",
            "@hr.",
            "@consulting.",
        ],
        # Keywords for identifying business and professional communications
        "strong_keywords": [
            "meeting",
            "agenda",
            "project",
            "business",
            "collaboration",
            "team",
            "report",
            "proposal",
            "task",
            "update",
            "deadline",
            "strategy",
            "client",
            "conference",
            "presentation",
        ],
        "subject_patterns": [
            r"project|report|business",
            r"collaboration|meeting|update",
            r"proposal|deadline|strategy",
        ],
        "header_indicators": [
            "business",
            "project",
            "team",
            "update",
            "meeting",
        ],
    },
}
