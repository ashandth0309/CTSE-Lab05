payments = []
id_counter = 1

async def get_payments():
    return payments

async def process_payment(payment):
    global id_counter
    
    new_payment = payment.dict()
    new_payment["id"] = id_counter
    new_payment["status"] = "SUCCESS"
    
    id_counter += 1
    payments.append(new_payment)

    return new_payment

async def get_payment_by_id(payment_id: int):
    for p in payments:
        if p["id"] == payment_id:
            return p
    return {"error": "Payment not found"}