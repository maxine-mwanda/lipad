Lipad Loan Request API

A Go backend API to manage users and loan requests, simulate a credit score API callback, and track loan approvals. Designed for fintech applications with traceable incoming and outgoing API requests.

Repository

https://github.com/maxine-mwanda/lipad

Project Description

This API allows:

Creating and listing users

Submitting loan requests

Validating loan amounts and preventing duplicate pending requests

Simulating calls to a credit scoring service

Updating loan request statuses based on credit score results

All API requests and outgoing calls are logged for traceability.

Prerequisites

Go 1.25+

MySQL 8+

Postman
 (optional, for testing API endpoints)

1. Clone the repository
git clone https://github.com/maxine-mwanda/lipad.git
cd lipad

2. Environment Variables

Create a .env file in the project root:

PORT=8080
DB_USER=root
DB_PASSWORD=your_mysql_password
DB_HOST=localhost
DB_NAME=lipad_test

3. Database Setup

Log in to MySQL and create the database:

CREATE DATABASE lipad_test;
USE lipad_test;


Create the users table:

CREATE TABLE user (
    user_id INT(15) PRIMARY KEY AUTO_INCREMENT,
    email VARCHAR(30) NOT NULL,
    phone_number VARCHAR(20) NOT NULL UNIQUE,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP
);


Create the loan_requests table:

CREATE TABLE loan_requests (
    loan_id VARCHAR(36) PRIMARY KEY,
    user INT NOT NULL,
    amount DECIMAL(12, 2) NOT NULL DEFAULT 0.00,
    reason VARCHAR(255),
    status VARCHAR(30),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user) REFERENCES user(user_id)
);

4. Install Dependencies
go mod tidy

5. Run the Server
go run main.go


Server listens on the port specified in .env (default 8080)

Logs all incoming requests and outgoing API calls

API Endpoints
Create a User

POST /users

Request:

{
  "email": "john@example.com",
  "phone_number": "0712345678"
}


Response:

{
  "status": "success",
  "data": { "user_id": 1, "email": "john@example.com", "phone_number": "0712345678" }
}

Create a Loan Request

POST /loan-requests

Request:

{
  "user_id": 1,
  "amount": 5000,
  "reason": "Emergency medical loan"
}


Validates amount > 0 and < 1,000,000

Rejects duplicate pending requests

Automatically triggers simulated credit score callback

List Loan Requests

GET /loan-requests/{id}

Replace {id} with the user ID

Returns all loan requests for that user

Webhook: Credit Score Callback

POST /webhook/credit-score

Request:

{
  "loan_id": "uuid-string",
  "score": 720,
  "status": "APPROVED",
  "reason": "Good credit history"
}


status must be "APPROVED" or "REJECTED"

Updates loan request in the database

Traceability

Incoming requests and outgoing API calls are logged for audit purposes

Use the logs to trace requests during Postman testing or production runs