# main.py
from managements.book_management import add_book, view_books, search_book
from managements.customer_management import add_customer, view_customers
from managements.sales_management import sell_book, view_sales

def main():
    while True:
        print("\n=== BookMart Management System ===")
        print("1. Book Management")
        print("2. Customer Management")
        print("3. Sales Management")
        print("4. Exit")
        choice = input("Enter your choice: ")

        if choice == "1":
            print("\n--- Book Management ---")
            print("1. Add Book")
            print("2. View Books")
            print("3. Search Book")
            sub_choice = input("Enter your choice: ")
            if sub_choice == "1":
                title = input("Enter book title: ")
                author = input("Enter book author: ")
                price = input("Enter book price: ")
                quantity = input("Enter book quantity: ")
                print(add_book(title, author, price, quantity))
            elif sub_choice == "2":
                print("\n".join(view_books()))
            elif sub_choice == "3":
                query = input("Enter title or author to search: ")
                print("\n".join(search_book(query)))

        elif choice == "2":
            print("\n--- Customer Management ---")
            print("1. Add Customer")
            print("2. View Customers")
            sub_choice = input("Enter your choice: ")
            if sub_choice == "1":
                name = input("Enter customer name: ")
                email = input("Enter customer email: ")
                phone = input("Enter customer phone: ")
                print(add_customer(name, email, phone))
            elif sub_choice == "2":
                print("\n".join(view_customers()))

        elif choice == "3":
            print("\n--- Sales Management ---")
            print("1. Sell Book")
            print("2. View Sales")
            sub_choice = input("Enter your choice: ")
            if sub_choice == "1":
                customer_name = input("Enter customer name: ")
                email = input("Enter customer email: ")
                phone = input("Enter customer phone: ")
                book_title = input("Enter book title: ")
                quantity_sold = int(input("Enter quantity sold: "))
                print(sell_book(customer_name, email, phone, book_title, quantity_sold))
            elif sub_choice == "2":
                print("\n".join(view_sales()))

        elif choice == "4":
            print("Exiting the system. Goodbye!")
            break
        else:
            print("Invalid choice. Try again.")

if __name__ == "__main__":
    main()
