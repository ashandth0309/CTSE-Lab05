const express = require("express");
const { createProxyMiddleware } = require("http-proxy-middleware");

const app = express();

// Route to Item Service
app.use(
  "/items",
  createProxyMiddleware({
    target: "http://item-service:8081",
    changeOrigin: true,
    pathRewrite: { "^/items": "" }, 
  })
);

// Route to Order Service
app.use(
  "/orders",
  createProxyMiddleware({
    target: "http://order-service:8082",
    changeOrigin: true,
    pathRewrite: { "^/orders": "" }, 
  })
);

// Route to Payment Service
app.use(
  "/payments",
  createProxyMiddleware({
    target: "http://payment-service:8083",
    changeOrigin: true,
     pathRewrite: { "^/payments": "" },
  })
);

app.listen(8080, () => {
  console.log("API Gateway running on port 8080");
});