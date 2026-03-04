from pydantic import BaseModel

class PaymentRequest(BaseModel):
    orderId: int
    amount: float
    method: str