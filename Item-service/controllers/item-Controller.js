let items = ["Book", "Laptop", "Phone"];
let idCounter = 1;

const getItems = (req, res) => {
  res.json(items);
};

const addItem = (req, res) => {
  const { name } = req.body;

  if (!name) {
    return res.status(400).json({ message: "Item name is required" });
  }

  items.push(name);

  res.status(201).json({
    id: idCounter++,
    name,
  });
};

const getItemById = (req, res) => {
  const id = parseInt(req.params.id);

  if (id < 0 || id >= items.length) {
    return res.status(404).json({ message: "Item not found" });
  }

  res.json({ id, name: items[id] });
};

module.exports = {
  getItems,
  addItem,
  getItemById,
};