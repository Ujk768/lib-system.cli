# 📚 Library Management System (Go)

A simple **console-based Library Management System** written in **Go**.  
This project demonstrates the use of **structs, slices, maps, error handling, and user input** in Go.  
It also includes a minimal **console UI** with colors and centered headings.

---

## ✨ Features

✅ **Add Book**  
- Add a new book by providing **ID, Title, and Author**.  
- Prevents duplicate IDs using a `map[int]*Book`.  

✅ **Delete Book**  
- Delete a book by ID.  
- Updates both the **slice** (for ordered storage) and **map** (for quick lookups).  

✅ **Find Book by ID**  
- Retrieve book details using the unique ID.  
- Returns helpful error messages if not found.  

✅ **View All Books**  
- Lists all books currently in the library.  
- Uses a slice to maintain insertion order.  

✅ **Colorful Console UI**  
- Menu and messages styled with ANSI color codes.  
- Titles and options are **centered** for better readability.  

✅ **Graceful Exit**  
- Option to exit the program with a goodbye message.  

---

## 🛠️ Tech Stack

- **Language**: Go (Golang)  
- **Libraries Used**:  
  - [`containerd/console`](https://github.com/containerd/console) → for console width detection  
  - Standard Go libraries (`fmt`, `os`, `bufio`, `strconv`, `strings`)  

---

## 📂 Project Structure

```
.
├── main.go       # Entry point and core logic
├── README.md     # Project documentation
```

---

## ▶️ How to Run

1. **Clone this repository**
   ```bash
   git clone https://github.com/your-username/library-management-system-go.git
   cd library-management-system-go
   ```

2. **Run the program**
   ```bash
   go run main.go
   ```

---

## 📸 Demo (Console UI)

```
           📚 Library Management System 📚

       Select one of the options...
          1. Add Book
          2. Delete Book
          3. Find Book By Id
          4. View All Books
          5. Exit
```

---

## 🚀 Next Planned Features

- [ ] **Search by Title / Author**  
- [ ] **Update Book Details**  
- [ ] **Persistent Storage (save to JSON or DB)**  
- [ ] **Export Books to File (CSV/JSON)**  

---

## 🧠 Learning Goals

This project is part of a **Go learning plan** to practice:
- Structs (`Book`)
- Slices (`library`)
- Maps (`libMap`)
- Functions with error handling
- User input and validation
- Console formatting with ANSI codes

---

## 👨‍💻 Author
Made with ❤️ in Go by *Utkarsh Kanade*  
