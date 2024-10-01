# Inventory Management App

**Inventory Management App** an app which i am trying to make. It will be a full-fledged inventory management system. Making it for my path to learn go and next.js more!


## TODO

- [x] **Login/Register Email based Authentication**
- [ ] **Dashboard Admin/Other staff**
- [ ] **Implement RBAC**
- [ ] **Add Database for Item/Inventory/Seller**
- [ ] **QR Code labeling system for products**
- [ ] **Stock Management**
- [ ] **Notifications & Alerts**

## Technologies Used

### Frontend
- **Next.js**: A JavaScript library for building user interfaces.

### Backend
- **Fiber**: A Golang package for server-side logic.
- **PGX**: A PostgreSQL driver for Go.
- **Go-JWT**: For authentication and authorization.
- **Go-Redis**: For caching and verification management.

### Other Stuff
- **Resend**: For managing and sending mails
- **Docker**: For containerization and deployment.
- **PostgreSQL**: For the database.
- **Redis**: For caching and verification management.

## Installation

To set up this project locally, follow the steps below:

1. **Clone the Repository**
    ```bash
    git clone https://github.com/AriaFantom/Inventory-Management.git
    cd Inventory-Management
    ```


4. **For development the App**

    - **Backend (Golang 1.23)**:
      ```bash
      cd backend
      ```
      ```bash
      go install github.com/air-verse/air@latest
      ```
      ```bash
      go install github.com/pressly/goose/v3/cmd/goose@latest
      ```
      ```bash
      go mod tidy
      ```
      ```bash
      docker-compose up -d
      ```
      ```bash
      air
      ```

    - **Frontend (Node.js 21.0.0)**:
      ```bash
      cd frontend
      ```
      ```bash
      npm install
      ```
      ```bash
      npm run dev
      ```

    The backend will run by default on `http://localhost:8080` and the frontend on `http://localhost:3000`.


## Contributing

We welcome contributions! Please follow these steps to contribute:

1. Fork the repository.
2. Create a new branch.
3. Make your changes
4. Open a pull request and explain what changes you’ve made.

## License

This project is licensed under the MIT License. See the `LICENSE` file for more details.