# Short URL Service

This project is a URL shortening service built with Go and MongoDB. It provides APIs for shortening URLs and retrieving them.

## Setup Instructions

Follow these steps to set up and run the project.

### 1. Clone the Repository

git clone https://github.com/your-username/your-repo.git
cd your-repo

### 2. Install Go Dependencies

Run the following command to install all necessary Go modules:

go mod tidy

### 3. Set Up Environment Variables

Copy the `.env.example` file to create a `.env` file, then edit it with your specific values:

cp .env.example .env
# Edit the .env file with your specific values

Make sure to set valid environment variables in the `.env` file, such as `DATABASE_URL`, `SHORT_URL_OPTIONS`, and `SHORT_URL_LENGTH`.

Example `.env` file:

DATABASE_URL=mongodb://localhost:27017  
PORT=8080  
SHORT_URL_OPTIONS=abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789  
SHORT_URL_LENGTH=6

### 4. Run the Project with Docker Compose

To build and run the project using Docker Compose, use the following command:

docker-compose up --build

### 5. Access the Application

Once the application is running, you can access it by navigating to:

http://localhost:8080

