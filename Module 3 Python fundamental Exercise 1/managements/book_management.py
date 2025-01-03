
class Book:
    def __init__(self, title, author, price, quantity):
        self.title = title
        self.author = author
        self.price = price
        self.quantity = quantity

    def display_details(self):
        return f"Title: {self.title}, Author: {self.author}, Price: ${self.price}, Quantity: {self.quantity}"

books = []

def add_book(title, author, price, quantity):
    try:
        price = float(price)
        quantity = int(quantity)
        if price <= 0 or quantity <= 0:
            raise ValueError("Price and quantity must be positive numbers.")
        books.append(Book(title, author, price, quantity))
    except ValueError as e:
        return f"Error: {e}"

def view_books():
    return [book.display_details() for book in books] if books else "No books available."

def search_book(query):
    results = [book.display_details() for book in books if query.lower() in book.title.lower() or query.lower() in book.author.lower()]
    return results if results else "No matching books found."
