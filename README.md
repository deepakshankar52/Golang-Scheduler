# Project Setup Instructions

## Backend (Go Server)

1. Navigate to the `Mail-Backend` directory:
    ```bash
    cd Mail-Backend
    ```

2. Start the Go server:
    ```bash
    go run main.go
    ```

## Frontend (Vue.js)

1. Navigate to the `file-frontend` directory:
    ```bash
    cd file-frontend
    ```

2. Run the Vue frontend:
    ```bash
    npm run serve
    ```

---

By following these steps, your Go server should be running on `localhost:8080`, and your Vue frontend should be running on `localhost:8081`.

---
# Project Overview

## Functionality

This project allows users to upload an attendance Excel file through the Vue.js frontend. The Excel file should contain the following columns:

- **Student's Name**
- **Leave Count**
- **Email ID**

## Process Flow

1. **File Upload (Frontend):**
    - The user uploads the Excel file via the Vue.js frontend.

2. **File Processing (Backend):**
    - The Go server receives the uploaded Excel file.
    - It processes the file to extract the student's name, leave count, and email ID.

3. **Warning Mail Notification:**
    - The Go server checks each student's leave count.
    - If a student exceeds the specified number of leaves, the server sends a personalized warning email to the respective student.

## Technical Details

- **Frontend:** Vue.js handles the file upload interface.
- **Backend:** The Go server processes the Excel file and sends emails.

## Key Features

- **Automated Email Alerts:** Automatically sends warning emails based on the leave count.
- **Seamless Integration:** Smooth interaction between the frontend and backend to handle file uploads and processing.

