// Array of books (each book is represented as an object)
let books = [
  { title: 'The Catcher in the Rye', author: 'J.D. Salinger', year: 1951 },
  { title: 'To Kill a Mockingbird', author: 'Harper Lee', year: 1960 },
  { title: '1984', author: 'George Orwell', year: 1949 },
];

// Function to display information about a book
function displayBookInfo(book) {
  console.log(`Title: ${book.title}`);
  console.log(`Author: ${book.author}`);
  console.log(`Year: ${book.year}`);
  console.log('------------------------');
}

// Display information about each book in the array
console.log('List of Books:');
books.forEach(displayBookInfo);

// Add a new book to the array
let newBook = { title: 'The Great Gatsby', author: 'F. Scott Fitzgerald', year: 1925 };
books.push(newBook);

// Display information about the new book
console.log('List of Books (after adding a new book):');
books.forEach(displayBookInfo);

