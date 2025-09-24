# 📬 InboXpert — Smart Email Management Extension

A lightweight browser extension that helps you tame your inbox by
**auto-categorizing emails, filtering spam, and surfacing what matters**—all with a clean, simple UI.

---

## 📌 Introduction
**InboXpert** makes everyday email easier. It watches your inbox (on supported web clients),
sorts messages into useful categories, flags likely spam, and highlights important emails
so you can focus on what actually needs your attention.

Under the hood, it blends **machine learning** with a few sensible **rules**,
and connects to a small set of **modular backend services** to stay fast and scalable.

---

## 🎯 Key Features
- **Automated Categorization**  
  Emails are sorted into buckets like **Work**, **Promotions**, and **Social** using an ML model plus simple rules (e.g., domain hints such as `@company.com` for Work).

- **Spam Filtering**  
  Obvious junk is flagged and separated (keywords like “win big” are caught by rules; an ML spam model is planned).

- **Priority Highlighting**  
  Messages from frequent contacts or with urgent words (e.g., “deadline”) are marked as **High Priority** so they don’t get buried.

- **Simple, Fast UI**  
  A small popup provides tabs for categories, quick actions like **Unsubscribe** and **Mark All Read**, and a clean way to move messages.

---

## 🛠️ Tech Stack

**Frontend**
- React + Vite + TypeScript  
- Tailwind CSS  
- Chrome Extension APIs (`chrome.runtime`, `chrome.storage`)  
- `MutationObserver` to handle dynamic DOM changes in webmail UIs

**Backend**
- Go (Gin)  
- Microservices for categorization, spam filtering, and priority detection  
- **API Gateway** for routing, retries, and rate limiting  
- PostgreSQL (uses **JSONB** to store flexible email metadata)

**Machine Learning**
- Python ML server (gRPC/REST)  
- Multinomial Naive Bayes with **TF-IDF** for categorization  
- (Planned) ML spam model and NLP summarization

---

## 🧩 How It Works (at a glance)
1. **Content/UI**  
   The extension injects a small UI and listens for inbox changes with `MutationObserver`.

2. **Classification Request**  
   When new or selected emails are detected, the extension sends a request to the **API Gateway**.

3. **Microservices**  
   The gateway forwards the request to the right service (categorization / spam / priority).  
   The categorization service talks to the **ML server** for real-time predictions.

4. **Store & Return**  
   Results (category, confidence, flags) are saved in **PostgreSQL** and returned to the extension.

5. **Update the UI**  
   The popup and inline indicators update so you can act quickly (move, mark read, unsubscribe, etc.).

---

## 📂 Project Structure
```bash
inboXpert/
├── extension/           # Frontend Chrome extension
│   ├── src/             # React + TypeScript components
│   └── public/          # Manifest, icons, assets
├── backend/             # Golang microservices
│   ├── api-gateway/     # API Gateway (Gin)
│   ├── categorization/  # ML-based categorization service
│   ├── spam/            # Rule-based spam filtering service
│   └── priority/        # Priority highlighting service
├── ml-server/           # Python ML server
│   ├── models/          # Trained ML models
│   └── api/             # gRPC + REST endpoints
├── db/                  # Database migrations & schema
└── docs/                # Presentations, reports
