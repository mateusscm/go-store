# go-store

This project is a simple CRUD (Create, Read, Update, Delete) application developed using Go and PostgreSQL. It serves as an initial attempt to build a web application with these technologies, following the "Go: Create a Web Application" course provided by Alura.

## Features

- **Product Management**: Add, view, update, and delete products with attributes such as description, price, and quantity.

## Technologies Used

- **Go**: The primary programming language for the application.
- **PostgreSQL**: Database system used to store product information.

## Project Structure

The project is organized into the following directories:

- `controllers/`: Contains the handlers for HTTP requests.
- `models/`: Defines the data structures and interfaces for database interactions.
- `routes/`: Manages the routing of HTTP requests.
- `templates/`: Holds the HTML templates for rendering web pages.
- `db/`: Includes database connection and migration files.

## Setup and Installation

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/mateusscm/go-store.git
   ```

2. **Navigate to the Project Directory**:

   ```bash
   cd go-store
   ```

3. **Install Dependencies**:

   Ensure you have Go installed. Initialize and download the required Go modules:

   ```bash
   go mod tidy
   ```

4. **Set Up the Database**:

   - Install PostgreSQL.
   - Create a new database named `go_store`.
   - Update the database connection settings in the project to match your PostgreSQL configuration.

5. **Run Database Migrations**:

   Apply any necessary database migrations (if applicable) to set up the required tables.

6. **Start the Application**:

   ```bash
   go run main.go
   ```

   The application should now be running, and you can access it via your web browser.

## Acknowledgments

This project is based on the "Go: Create a Web Application" course provided by [Alura](https://www.alura.com.br/).

---

For more details, please refer to the original repository: [https://github.com/mateusscm/go-store](https://github.com/mateusscm/go-store).

