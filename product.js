
const productsData = `
[
  { "id": 1, "name": "Laptop", "category": "Electronics", "price": 1200, "available": true },
  { "id": 2, "name": "Smartphone", "category": "Electronics", "price": 800, "available": false },
  { "id": 3, "name": "Desk", "category": "Furniture", "price": 300, "available": true }
]
`;

function parseJSONData(data) {// function to parse JSON data
  return JSON.parse(data);
}

let products = parseJSONData(productsData);// parsing the JSON data and store it in the products array


function addProduct(newProduct) {
  products.push(newProduct);
}

function updateProductPrice(productId, newPrice) {
  const product = products.find((p) => p.id === productId);
  if (product) {
    product.price = newPrice;
    return `Product ID ${productId} price updated to ${newPrice}`;
  } else {
    return `Error: Product with ID ${productId} not found.`;
  }
}


function filterAvailableProducts() {
  return products.filter((product) => product.available);
}


function filterProductsByCategory(category) {
  return products.filter((product) => product.category === category);
}



addProduct({ id: 4, name: "Chair", category: "Furniture", price: 100, available: true });

console.log(updateProductPrice(1, 1300));  
console.log(filterAvailableProducts());
console.log(filterProductsByCategory("Electronics"));
