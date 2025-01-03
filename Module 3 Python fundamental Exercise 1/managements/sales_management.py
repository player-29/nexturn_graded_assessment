
from .customer_management import Customer
from .book_management import books

class Transaction(Customer):
    def __init__(self, name, email, phone, book_title, quantity_sold):
        super().__init__(name, email, phone)
        self.book_title = book_title
        self.quantity_sold = quantity_sold

    def display_details(self):
        return f"Customer: {self.name}, Book: {self.book_title}, Quantity Sold: {self.quantity_sold}"

sales = []

def sell_book(customer_name, email, phone, book_title, quantity_sold):
    for book in books:
        if book.title.lower() == book_title.lower():
            if book.quantity >= quantity_sold:
                book.quantity -= quantity_sold
                sales.append(Transaction(customer_name, email, phone, book_title, quantity_sold))
                return f"Sale Successful. Remaining Books:{book.quantity}"
            else:
                return "Error: Not enough stock available."
    return "Error: Book not found."

def view_sales():
    return [sale.display_details() for sale in sales] if sales else "No sales records available."
