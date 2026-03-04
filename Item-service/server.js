const express = require("express");
const cors = require("cors");

const app = express();
const itemRoutes = require("./routes/item-routes");

app.use(cors());
app.use(express.json());

app.use("/", itemRoutes);

app.get("/", (req, res) => {
  res.json({ message: "Item Service Running" });
});

const PORT = 8081;

app.listen(PORT, () => {
  console.log(`Item Service running on port ${PORT}`);
});