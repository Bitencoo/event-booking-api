# Event Booking API

The Event Booking API is a RESTful service built in Go for managing events, users, and event registrations. It allows you to create events, register users, and allocate users to specific events. The API handles common event management operations such as creating, updating, and deleting events, as well as registering and unregistering users for those events.

## Features

- **User Management:**
  - Register new users
  - Validate user credentials
  - Secure password storage and validation

- **Event Management:**
  - Create, update, and delete events
  - Retrieve all events or a specific event by ID
  - Manage event registrations

- **Event Registration:**
  - Register users for specific events
  - Cancel user registrations for events

## Prerequisites

- **Go 1.x or higher**: Ensure you have Go installed. You can download it from [here](https://golang.org/dl/).

## Setup and Installation

### Step 1: Clone the Repository

```bash
git clone https://github.com/yourusername/event-booking-api.git
cd event-booking-api
```

### Step 2: Install Dependencies
Use Go modules to install the required dependencies:
```bash
go mod tidy
```

### Step 3: Initialize and Migrate the Database
The database is automatically initialized and migrated using SQLite when you run the API. Ensure that db.go is correctly set up to create the necessary tables upon startup.

### Step 4: Start the API Server
Run the API server:
```bash
go run main.go
```
The server will start running on http://localhost:8080.

## API Endpoints

> **Note:** The application uses JWT (JSON Web Tokens) for authentication, and all user passwords are securely hashed before being stored in the database. Some endpoints require authentication while other don't, represented by the tags - Not Authenticated (Doesn't need authentication) or - Authenticated (Needs authentication).

### User
Endpoints responsible for registering and loging into the application.

<details>
  <summary>User Endpoints</summary>
  
#### Register a New User - Not Authenticated

- **Endpoint**: `POST /signup`
- **Request Body**:
  ```json
  {
    "email": "user@example.com",
    "password": "securepassword"
  }

#### Login - Not Authenticated

- **Endpoint**: `POST /login`
- **Request Body**:
  ```json
  {
    "email": "user@example.com",
    "password": "securepassword"
  }
</details>

### Events
Endpoints responsible for Events. An Event might be a show, movie, party or others. In here you will find option to visualize all events, single event, register for an event, cancel registration, create, delete and update events.

<details>
  <summary><strong>Events Endpoints</strong></summary>
  <br>
<details>
  <summary>Get All Events - Not Authenticated</summary>
  
- **Endpoint**: `GET /events`
- **Response Body**:
  ```json
  [
    {
      "ID": 1,
      "Name": "Sample Event",
      "Description": "This is a sample description of an event.",
      "Location": "1234 Sample Street, Example City, 90210",
      "DateTime": "2023-04-25T15:04:05Z",
      "UserID": 1
    },
    {
      "ID": 2,
      "Name": "Sample Event",
      "Description": "This is a sample description of an event.",
      "Location": "1234 Sample Street, Example City, 90210",
      "DateTime": "2023-04-25T15:04:05Z",
      "UserID": 1
    }
  ]

</details>

<details>
  <summary>Get Single Event - Not Authenticated</summary>
  
- **Endpoint**: `GET /events/:userId`
- **Response Body**:
  ```json
  {
    "ID": 2,
    "Name": "Sample Event",
    "Description": "This is a sample description of an event.",
    "Location": "1234 Sample Street, Example City, 90210",
    "DateTime": "2023-04-25T15:04:05Z",
    "UserID": 1
  }

</details>

<details>
  <summary>Register Logged User to an Event - Authenticated</summary>
  
- **Endpoint**: `POST /events/:eventId/register`
- **Headers**:
  - content-type: application/json
  - Authorization: jwt
- **Response Body**:
  ```json
  {
    "message": "Registration created succesfully",
    "registrationID": 3
  }

</details>

<details>
  <summary>Update Event - Authenticated</summary>
  
- **Endpoint**: `PUT /events/:eventId`
- **Headers**:
  - content-type: application/json
  - Authorization: jwt
- **Request Body**:
  ```json
  {
    "Name": "Sample Event",
    "Description": "This is a sample description of an event.",
    "Location": "1234 Sample Street, Example City, 90210",
    "DateTime": "2023-04-25T15:04:05Z"
  } 
- **Response Body**:
  ```json
  {
    "message": "Event updated succesfully"
  }

</details>

<details>
  <summary>Create Event - Authenticated</summary>
  
- **Endpoint**: `POST /events`
- **Headers**:
  - content-type: application/json
  - Authorization: jwt
- **Request Body**:
  ```json
  {
    "Name": "Sample Event",
    "Description": "This is a sample description of an event.",
    "Location": "1234 Sample Street, Example City, 90210",
    "DateTime": "2023-04-25T15:04:05Z"
  }
- **Response Body**:
  ```json
  {
    "event": {
      "ID": 3,
      "Name": "Sample Event",
      "Description": "This is a sample description of an event.",
      "Location": "1234 Sample Street, Example City, 90210",
      "DateTime": "2023-04-25T15:04:05Z",
      "UserID": 2
    },
    "message": "Event Created!"
  }

</details>

<details>
  <summary>Delete Event - Authenticated</summary>
  
- **Endpoint**: `DELETE /events/:eventId`
- **Headers**:
  - content-type: application/json
  - Authorization: jwt
- **Response Body**:
  ```json
  {
    "message": "Deleted the Event succesfully!"
  }

</details>

<details>
  <summary>Cancel registration of logged User on an event- Authenticated</summary>
  
- **Endpoint**: `DELETE /events/:eventId/register`
- **Headers**:
  - content-type: application/json
  - Authorization: jwt
- **Response Body**:
  ```json
  {
    "message": "Registration Canceled!"
  }

</details>
</details>
