# Real-time Messaging with Ember.js and Go

![Real time messaging logo](https://imgs.search.brave.com/YwAb3-7k9JHOybiGWVb7t-N77xMDf_zPNj7eWURnAAc/rs:fit:500:0:0/g:ce/aHR0cHM6Ly9pay5p/bWFnZWtpdC5pby9h/Ymx5L2dob3N0L3By/b2QvMjAyMy8wMS9i/dWlsZC1hLXJlYWx0/aW1lLWNoYXQtYXBw/LWZyb20tc2NyYXRj/aC0tMS0ucG5nP3Ry/PXctMTcyOCxxLTUw)

## 📌 Overview

This project provides a comprehensive solution for real-time messaging using **Ember.js** on the frontend and **Go** (Golang) on the backend. The synergy of Ember.js's reactivity and Go's concurrent capabilities ensures swift messaging experiences, suitable for applications demanding real-time interactions.

### 🚀 What this project achieves:

1. **Real-time Messaging**: Messages are delivered in real-time, ensuring immediate communication between users.
2. **Scalability**: Built on Go's lightweight, concurrent foundation, this messaging system can handle thousands of simultaneous users.
3. **Reactive UI**: With Ember.js, any new message or notification updates the UI in real-time, offering users an engaging experience.
4. **Secure Authentication**: A robust authentication system protects user data and ensures only authenticated users access the platform.
5. **Cross-platform**: Designed to be responsive, the Ember.js frontend can adapt to various devices, from desktops to mobiles.

## 🛠️ Technologies Used

- ![Ember.js Icon](https://imgs.search.brave.com/DHToW1v6msszbLQ5ax9JU-yMebRY1UPovvLzlH9T1N0/rs:fit:860:0:0/g:ce/aHR0cHM6Ly9paDEu/cmVkYnViYmxlLm5l/dC9pbWFnZS4zOTI4/NjQ3ODUuNzM0NS9y/YWYsMzYweDM2MCww/NzUsdCxmYWZhZmE6/Y2E0NDNmNDc4Ni5q/cGc)
- **Ember.js**: A productive, battle-tested JavaScript framework for building modern web applications.
- ![Go Icon](https://imgs.search.brave.com/WbsQe2Yih9zzjmP3z645d_mnmx_jIOqHI5tCrwCiNCc/rs:fit:860:0:0/g:ce/aHR0cHM6Ly9paDEu/cmVkYnViYmxlLm5l/dC9pbWFnZS44MzE2/ODcyMzEuMjUzNC9w/b3N0ZXIsNTA0eDQ5/OCxmOGY4ZjgtcGFk/LDYwMHg2MDAsZjhm/OGY4LnU3LmpwZw)
- **Go (Golang)**: An open-source programming language that makes it easy to build efficient and reliable software.
- ![PostgreSQL Icon](https://imgs.search.brave.com/yMz2d3lR3SJ1a84hOzYyM89xt7x_LiW3MFKWNFvS4ac/rs:fit:860:0:0/g:ce/aHR0cHM6Ly9jZG4u/d29ybGR2ZWN0b3Js/b2dvLmNvbS9sb2dv/cy9wb3N0Z3Jlc3Fs/LnN2Zw.svg)
- **PostgreSQL**: A powerful, open-source object-relational database system.

## 🚢 Getting Started

### Prerequisites

The version and cli running below 

![embercli](./frontend/img/embercli.png)

- Node.js and npm (for Ember.js)
- Go (for the backend server)
- PostgreSQL (for the database)

### Setting Up

1. **Backend (Go)**:
    ```bash
    # Navigate to the backend directory
    cd path-to-your-go-backend

    # Set up the environment variables for database
    export DB_HOST=localhost
    export DB_PORT=5432
    export DB_USER=postgres
    export DB_PASSWORD=your_password
    export DB_NAME=your_database_name

    # Run the Go application
    go run *.go
    ```

2. **Frontend (Ember.js)**:
    ```bash
    # Navigate to the frontend directory
    cd path-to-your-ember-frontend

    # Install the dependencies
    npm install

    # Start the Ember.js server
    ember serve
    ```

Visit `http://localhost:4200/` on your browser to view the Ember.js frontend, while the Go backend runs on `http://localhost:8080/`.

## 📘 Documentation

For detailed documentation, including API endpoints, database schemas, and frontend components, refer to [DOCUMENTATION.md](./DOCUMENTATION.md).

## 🙌 Contributing

We welcome contributions! Please read our [CONTRIBUTING.md](./CONTRIBUTING.md) for guidelines on how to proceed.

## 📜 License

This project is licensed under the MIT License. See [LICENSE.md](./LICENSE.md) for more details.

## 🤝 Acknowledgements

- Ember.js community and developers
- The Go team and contributors
- Everyone who provided feedback and contributed to this project

---

Made with ❤️ by [Adesoji Alu or Team Name](https://your-link.com).
