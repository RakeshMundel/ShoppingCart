const axios = require("axios");
const all_product = require("../Assets/all_product.js");

all_product.forEach(async (product) => {
  try {
    await axios.post("http://localhost:4000/products", {
      name: product.name,
      description: "Seeded from script",
      image: product.image,
      price: product.new_price,
      category: product.category,
    });
    console.log(`✅ Inserted: ${product.name}`);
  } catch (err) {
    console.error(`❌ Error inserting ${product.name}:`, err.message);
  }
});
