# PizzaHub

PizzaHub is a scalable and efficient pizza ordering system built in Go. Users can place orders for different types of pizzas, and the system processes these orders based on the selected pizza type. The application is designed to handle multiple orders simultaneously, utilizing a pool of chefs to ensure quick processing.

## Table of Contents

- [Features](#features)
- [Technologies Used](#technologies-used)
- [API Endpoints](#api-endpoints)
- [Installation](#installation)
- [Usage](#usage)
- [Testing](#testing)
- [OpenAPI Documentation](#openapi-documentation)
- [Contributing](#contributing)

## Features

- **Order Processing**: Users can place orders for pizzas, which are processed based on the type of pizza ordered.
- **Scalability**: The system can handle multiple orders concurrently, with each chef processing only one order at a time.
- **RESTful API**: The application exposes a RESTful API for interaction.
- **Clean Code**: The code is structured to follow best practices in object-oriented programming.
- **Testing**: Comprehensive unit tests are included to ensure reliability.
- **OpenAPI Specification**: The API is documented using OpenAPI for better understanding and integration.

## Technologies Used

- Go (Golang)
- OpenAPI for API documentation

## API Endpoints

### Chefs

- **POST /chefs**
  - Description: Add a new chef.
  - Request Body: None
  - Response: `201 Created` on success.

- **GET /chefs**
  - Description: Retrieve the list of chefs.
  - Response: `200 OK` with a JSON array of chefs.

### Menus

- **POST /menus**
  - Description: Add a new menu item.
  - Request Body: JSON object with `Name` and `Duration`.
  - Response: `201 Created` on success.

- **GET /menus**
  - Description: List available pizzas.
  - Response: `200 OK` with a JSON object of menu items.

### Orders

- **POST /orders**
  - Description: Add a new order.
  - Request Body: JSON object with `ID` and `Pizza`.
  - Response: `201 Created` on success or `400 Bad Request` for invalid pizza type.

- **GET /orders**
  - Description: Retrieve the list of orders.
  - Response: `200 OK` with a JSON array of orders.

## Installation

1. **Clone the repository**:
   ```bash
   git clone git@github.com:balqisgautama/pizza-hub.git
   cd pizza-hub
   ```

2. **Build the application**:
   ```bash
   make build
   ```

3. **Run the application**:
   ```bash
   make run
   ```

4. **Rebuild the application**:
   ```bash
   make rebuild
   ```

5. **Stop the application**:
   ```bash
   make stop
   ```

6. **Clean the application**:
   ```bash
   make clean
   ```

## Usage

You can use tools like `curl` or Postman to interact with the API. Here are some example commands:

### Adding a Chef
```bash
curl -X POST http://localhost:8080/chefs -H "Content-Type: application/json" -d '{"name": "gautama"}'
```

### Getting Chefs
```bash
curl http://localhost:8080/chefs
```

### Adding a Menu Item
```bash
curl -X POST http://localhost:8080/menus -H "Content-Type: application/json" -d '{"name": "Pepperoni", "duration": 4}'
```

### Listing Menus
```bash
curl http://localhost:8080/menus
```

### Adding an Order
```bash
curl -X POST http://localhost:8080/orders -H "Content-Type: application/json" -d '{"menu_id": 1}'
```

### Getting Orders
```bash
curl http://localhost:8080/orders
```

## Testing

To run the tests, use the following command:
```bash
make test
```

## OpenAPI Documentation

The API is documented using OpenAPI. You can find the documentation in the `oas.yaml` file. This file provides a detailed description of the API endpoints, request/response formats, and other relevant information.

## Contributing

Contributions are welcome! If you have suggestions for improvements or new features, please open an issue or submit a pull request.