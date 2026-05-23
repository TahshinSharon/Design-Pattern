#Open Closed Principle
from abc import ABC,abstractmethod

class PaymentMethod(ABC):
    @abstractmethod
    def pay(self,amount:int):
        pass
class BkashPayment(PaymentMethod):
    def pay(self, amount:int):
        print(f"Paring through Bkash of Tk: {amount}")

class DebitCardPayment(PaymentMethod):
    def pay(self,amount:int):
        print(f"Paring through DebitCard of Tk: {amount}")
class PayPalPayment(PaymentMethod):
    def pay(self,amount:int):
        print(f"Paring through PayPal of Tk: {amount}")
class PaymentProcessor:
    def process_payment(self,payment_method:PaymentMethod,amount:int):
        payment_method.pay(amount)
bKash=BkashPayment()
debit=DebitCardPayment()
paypal=PayPalPayment()

pay_process=PaymentProcessor()
pay_process.process_payment(bKash,500)
pay_process.process_payment(debit,1000)
pay_process.process_payment(paypal,700)
