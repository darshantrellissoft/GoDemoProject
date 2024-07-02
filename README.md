# Transact APP

Transact APP is a demo application built using Go with the Gin framework, PostgreSQL for the database, and various utility packages for email sending, logging, and more. This application includes features for user authentication, payment processing, and daily report generation.

## Features

### User Authentication

- **Register User**: Users can register with their details and a company profile.
- **Login**: Authenticate users and return a JWT token.
- **Logout**: Invalidate the current user's token.
- **Update User Details**: API to update user information.

### Payment Processing

- **Create Payment**: API to create a new payment.
- **Confirm Payment**: API to confirm a payment using the transaction ID.
- **Get Transaction Details**: Retrieve transaction details by transaction ID.

### Background Tasks

- **Daily Report**: A cron job to generate and send a daily transaction report via email.

## Project Structure

.
├── controllers
│ ├── auth_controller.go
│ └── payment_controllers.go
├── cron
│ └── cron.go
├── models
│ ├── company_profile.go
│ ├── transaction.go
│ └── user.go
├── utils
│ ├── auth_utils.go
│ ├── base_utils.go
│ ├── email_utils.go
│ ├── excel_utils.go
│ └── json_response.go
├── templates
│ ├── confirmation.html
│ └── welcome.html
├── config
│ ├── config.go
│ └── logger.go
├── main.go
├── Makefile
└── README.md


## Setup Instructions

1. **Clone the repository**:
    ```sh
    git clone <repository_url>
    cd TransactAPP
    ```

2. **Set up environment variables**: Create a `.env` file with the following variables and also configure db details in Makefile:
    ```sh
    PASETO_SECRET_KEY=your_paseto_secret_key
    EMAIL_FROM=your_email@example.com
    SMTP_HOST=smtp.example.com
    SMTP_USER=smtp_user
    SMTP_PASS=smtp_password
    ```
3. **Create database**:
    ```sh
    make create_db 
    ```

4. **Install dependencies**:
    ```sh
    go mod tidy
    ```

5. **create database migrations files**:
    ```sh
    make create_migration name=<name>
    ```

6. **Run database migrations**:
    ```sh
    make migrate_up
    ```

7. **Run database migrations reverse**:
    ```sh
    make migrate_down
    ```

8. **Start the application**:
    ```sh
    make run
    ```

## Running the Application

- **Register a user**:
    ```sh
    POST /auth/register
    {
      "firstName": "John",
      "lastName": "Doe",
      "email": "john.doe@example.com",
      "password": "password123",
      "confirmPassword": "password123",
      "companyName": "ExampleCorp"
    }
    ```

- **Login**:
    ```sh
    POST /auth/login
    {
      "email": "john.doe@example.com",
      "password": "password123"
    }
    ```

- **Create a payment**:
    ```sh
    POST /payments
    {
      "cardNumber": "1234567812345678",
      "expiryDate": "12/25",
      "cvv": "123",
      "amount": 100.00
    }
    ```

- **Confirm a payment**:
    ```sh
    GET /payments/confirm/{transactionID}
    ```

- **Get transaction details**:
    ```sh
    GET /payments/{transactionID}
    ```

## Logging

Logs are generated using the `logrus` package and are stored in a log file located in the `logs` directory. Each API handler logs relevant information such as payloads, errors, and other significant events to aid in debugging and monitoring the application.
