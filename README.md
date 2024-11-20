# 🎓 Student Registration API - Golang 

## 📚 About The Project
A simple yet powerful RESTful API built with Go (Golang) for managing student registration data. This project allows you to store and retrieve student information through a clean and efficient API interface.

## 🛠️ Built With
- 🐹 Golang
- 🗄️ MySQL Database
- 📡 RESTful API Architecture

## 📁 Project Structure
```
├── 📂 controllers
│   └── student_controller.go
├── 📂 models
│   └── student_model.go
├── 📂 routes
│   └── routes.go
├── 📂 utils
│   └── respons.go
├── 📂 view
│   └── 📂 static
│       ├── style.css
│       └── index.html
├── 📄 .env
├── 📄 go.mod
├── 📄 go.sum
└── 📄 main.go
```

## 🚀 Features
- ✨ Student Registration
- 📋 View Student List
- 🔒 Secure Data Storage
- 🎯 RESTful Endpoints

## 🔌 API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/students` | Register new student |
| `GET`  | `/students` | Get all students |

## 🏃‍♂️ How to Run

1. Clone the repository
```bash
git clone <your-repo-url>
```

2. Set up environment variables
```bash
# Copy example env file
cp .env.example .env
# Configure your database credentials
```

3. Install dependencies
```bash
go mod tidy
```

4. Run the application
```bash
go run main.go
```

## 💡 Usage Example

### Register New Student
```bash
curl -X POST http://localhost:8080/students \
-H "Content-Type: application/json" \
-d '{
    "name": "Daniels",
    "nim": "12345",
    "Fakultas": "Computer Science",
    "Program_study": "Rekayasa Perangkat lunak"
}'
```

### Get All Students
```bash
curl http://localhost:8080/students
```

## 📝 Database Structure

### Students Table
| Field | Type | Description |
|-------|------|-------------|
| id | int | Primary Key |
| name | varchar | Student's name |
| nim | varchar | Student ID number |
| faculty | varchar | Faculty name |
| study_program | varchar | Study program name |


## 📫 Contact
Fiqri Agustriawan - [@fkriq4](https://www.instagram.com/fkriq4/)

## 🙏 Acknowledgments
* [Golang](https://golang.org/)
* [MySQL](https://www.mysql.com/)
* [Your other dependencies]

## 📜 License
This project is licensed under the MIT License - see the LICENSE file for details

---
⭐️ From [Fiqri Agustriawan](https://github.com/faradaydayday)
