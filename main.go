package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/containerd/console"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

/**

"%[1]0*s"

Let's quickly break it down:

%: Starts the format verb.

[1]: Refers to the first argument.

0: The flag to use '0' for padding.

*: The dynamic width specifier.

s: The format verb for a string.

*/

type Book struct {
	ID     int
	Title  string
	Author string
}

var library = []Book{{ID: 1, Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"}}

var libMap = make(map[int]Book)

func AddBook(bookId int, title string, author string) {
	if _, exists := libMap[bookId]; exists {
		fmt.Println(Red + "Error: Book with this ID already exists!" + Reset)
		return
	}
	book := Book{ID: bookId, Title: title, Author: author}
	library = append(library, book)
	libMap[bookId] = book
	fmt.Println(Green + "Book added successfully!" + Reset)

}

func FindBookByID(id int) (*Book, error) {
	if book, exists := libMap[id]; exists {
		return &book, nil
	}
	return nil, fmt.Errorf("Book with ID %d not found", id)
}

func DeleteBook(id int) (*Book, error) {
	if book, exists := libMap[id]; exists {
		delete(libMap, id)
		for i, b := range library {
			if b.ID == id {
				library = append(library[:i], library[i+1:]...)
				break
			}
		}
		return &book, nil
	}
	return nil, fmt.Errorf("Book with ID %d not found", id)

}

func SearchByTitle(title string) []Book {
	matchingBooks := []Book{}
	for _, book := range library {
		if strings.Contains(strings.ToLower(book.Title), strings.ToLower(title)) {
			matchingBooks = append(matchingBooks, book)
		}
	}
	return matchingBooks
}

func center(s string, w int) string {

	return fmt.Sprintf("%[1]*s", -w, fmt.Sprintf("%[1]*s", (w+len(s))/2, s))

}

func main() {
	winSize, err := console.Current().Size()

	if err != nil {
		fmt.Println("Error geting console width")
	}

	width := int(winSize.Width)
	libMap[1] = Book{ID: 1, Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"}

	fmt.Println(center(Cyan+"Library Management System"+Reset, width))

	for {
		fmt.Println()
		fmt.Println(center(Magenta+"Select one of the options..."+Reset, width))
		fmt.Println(center(Yellow+"1. Add Book "+Reset, width))
		fmt.Println(center(Yellow+"2. Delete Book "+Reset, width))
		fmt.Println(center(Yellow+"3. Find Book By Id "+Reset, width))
		fmt.Println(center(Yellow+"4. View All Books "+Reset, width))
		fmt.Println(center(Yellow+"5. Search Book By Title "+Reset, width))
		fmt.Println(center(Yellow+"6. Exit "+Reset, width))

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')

		input = strings.TrimSpace(input)
		choice, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println(Red + "Error: Invalid input. Please enter a number." + Reset)
			continue // Skip to the next loop iteration
		}

		switch choice {
		case 1:
			fmt.Println(Magenta + "You chose Add Book." + Reset)
			fmt.Printf(Blue + "Enter Book Id: " + Reset)
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)
			bookId, err := strconv.Atoi(id)
			if err != nil {
				fmt.Println(Red + "Error: Invalid Book ID. Please enter a valid number." + Reset)
				continue
			}
			fmt.Printf(Blue + "Enter Book Title: " + Reset)
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)

			fmt.Printf(Blue + "Enter Book Author: " + Reset)
			author, _ := reader.ReadString('\n')
			author = strings.TrimSpace(author)

			if title == "" || author == "" {
				fmt.Println(Red + "Error: Title and Author cannot be empty." + Reset)
				continue
			}

			AddBook(bookId, title, author)

		case 2:
			fmt.Println("You chose Delete Book.")
			// Call your DeleteBook function here
			fmt.Printf(Blue + "Enter Book Id to delete: " + Reset)
			id, _ := reader.ReadString(('\n'))
			id = strings.TrimSpace(id)
			bookId, err := strconv.Atoi(id)
			if err != nil {
				fmt.Println(Red + "Error: Invalid Book ID. Please enter a valid number." + Reset)
				continue
			}
			book, err := DeleteBook(bookId)
			if err != nil {
				fmt.Println(Red + err.Error() + Reset)
			} else {
				fmt.Println(Green+"Deleted Book Successfully :", "ID:", book.ID, "Title:", book.Title, "Author:", book.Author+Reset)

			}
		case 3:
			fmt.Println("You chose Find Book By Id.")
			// Call your FindBookByID function here
			fmt.Printf(Blue + "Enter Book Id to find: " + Reset)
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)
			bookId, err := strconv.Atoi(id)
			if err != nil {
				fmt.Println(Red + "Error: Invalid Book ID. Please enter a valid number." + Reset)
				continue
			}
			book, err := FindBookByID(bookId)
			if err != nil {
				fmt.Println(Red + err.Error() + Reset)
			} else {
				fmt.Println(Green+"Book found:", "ID:", book.ID, "Title:", book.Title, "Author:", book.Author, Reset)
			}
		case 4:
			fmt.Println("You chose View All Books.")
			// Call your ViewAllBooks function here
			for _, book := range library {
				fmt.Println("ID:", book.ID, "Title:", book.Title, "Author:", book.Author)
			}
		case 5:
			fmt.Println("You chose Search Book By Title.")
			// Call your SearchByTitle function here
			fmt.Printf(Blue + "Enter Book Title to search: " + Reset)
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)
			if title == "" {
				fmt.Println(Red + "Error: Title cannot be empty." + Reset)
				continue
			}
			matchingBooks := SearchByTitle(title)
			if len(matchingBooks) == 0 {
				fmt.Println(Red + "No books found matching the title." + Reset)
			} else {
				fmt.Println(Green + "Matching Books:" + Reset)
				for _, book := range matchingBooks {
					matchInd := strings.Index(strings.ToLower(book.Title), strings.ToLower(title))
					fmt.Println("ID:", book.ID, "Title:", book.Title[:matchInd]+Yellow+book.Title[matchInd:matchInd+len(title)]+Reset+book.Title[matchInd+len(title):], "Author:", book.Author)
				}
			}
		case 6:
			fmt.Println(center(Magenta+"Exiting Application. GoodBye!!!"+Reset, width))
			return // Use return to exit the main function and terminate the program.
		default:
			fmt.Println(Red + "Error: Invalid choice. Please select an option from 1 to 5." + Reset)
		}

	}

}
