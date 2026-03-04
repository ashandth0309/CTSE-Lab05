from fastapi import FastAPI
from routes.payment_routes import router as payment_router

app=FastAPI(title="Payment service")

app.include_router(payment_router)

@app.get("/")
def read_root():
    return {"message": "Payment Service Running "}