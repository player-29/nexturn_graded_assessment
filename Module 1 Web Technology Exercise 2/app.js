// JavaScript for Expense Tracker

// Select DOM elements
const expenseForm = document.getElementById('expense-form');
const expenseTableBody = document.getElementById('expense-table-body');
const categorySummary = document.getElementById('category-summary');
const expenseChart = document.getElementById('expense-chart');

let expenses = JSON.parse(localStorage.getItem('expenses')) || [];

// Function to update localStorage
function updateLocalStorage() {
    localStorage.setItem('expenses', JSON.stringify(expenses));
}

// Function to render the expense table
function renderExpenseTable() {
    expenseTableBody.innerHTML = '';
    expenses.forEach((expense, index) => {
        const row = document.createElement('tr');

        row.innerHTML = `
            <td>${expense.amount}</td>
            <td>${expense.description}</td>
            <td>${expense.category}</td>
            <td><button class="delete-btn" data-index="${index}">Delete</button></td>
        `;

        expenseTableBody.appendChild(row);
    });
}

// Function to render the category summary
function renderCategorySummary() {
    const categoryTotals = expenses.reduce((totals, expense) => {
        totals[expense.category] = (totals[expense.category] || 0) + parseFloat(expense.amount);
        return totals;
    }, {});

    categorySummary.innerHTML = '';
    for (const category in categoryTotals) {
        const listItem = document.createElement('li');
        listItem.innerHTML = `${category}: $${categoryTotals[category].toFixed(2)}`;
        categorySummary.appendChild(listItem);
    }
}

// Function to add a new expense
function addExpense(event) {
    event.preventDefault();

    const amount = parseFloat(document.getElementById('amount').value);
    const description = document.getElementById('description').value;
    const category = document.getElementById('category').value;

    if (!amount || amount <= 0) {
        alert('Please enter a valid amount.');
        return;
    }

    const newExpense = { amount, description, category };
    expenses.push(newExpense);
    updateLocalStorage();
    renderExpenseTable();
    renderCategorySummary();
    renderChart();

    expenseForm.reset();
}

// Function to delete an expense
function deleteExpense(event) {
    if (event.target.classList.contains('delete-btn')) {
        const index = event.target.dataset.index;
        expenses.splice(index, 1);
        updateLocalStorage();
        renderExpenseTable();
        renderCategorySummary();
        renderChart();
    }
}

// Function to render the chart
function renderChart() {
    const categoryTotals = expenses.reduce((totals, expense) => {
        totals[expense.category] = (totals[expense.category] || 0) + parseFloat(expense.amount);
        return totals;
    }, {});

    const categories = Object.keys(categoryTotals);
    const totals = Object.values(categoryTotals);

    const ctx = expenseChart.getContext('2d');
    new Chart(ctx, {
        type: 'pie',
        data: {
            labels: categories,
            datasets: [{
                data: totals,
                backgroundColor: ['#FF6384', '#36A2EB', '#FFCE56', '#4BC0C0'],
            }],
        },
    });
}

// Event Listeners
expenseForm.addEventListener('submit', addExpense);
expenseTableBody.addEventListener('click', deleteExpense);

// Initial Render
renderExpenseTable();
renderCategorySummary();
renderChart();
