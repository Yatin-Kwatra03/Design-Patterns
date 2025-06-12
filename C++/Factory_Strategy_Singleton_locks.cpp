#include<iostream>
#include<memory>
#include <utility>
#include<mutex>

using namespace std;


class PaymentStrategy {
public:
	virtual void payMoney() = 0;
};

class UpiPayment: public PaymentStrategy {
public:
	void payMoney() override {
		cout << "Money paid using UPI\n";
	}
};


class CreditCardPayment: public PaymentStrategy {
public:
	void payMoney() override {
		cout << "Money paid using CreditCard\n";
	}
};


class DebitCardPayment: public PaymentStrategy {
public:
	void payMoney() override {
		cout << "Money paid using DebitCard\n";
	}
};

class AmazonPayPayment: public PaymentStrategy {
public:
	void payMoney() override {
		cout << "Money paid using AmazonPay\n";
	}
};

class PaymentService {
private:
	unique_ptr<PaymentStrategy> paymentStrategy;
	static unique_ptr<PaymentService> instance;
	static mutex mtx;
	PaymentService() {}
public:

	static PaymentService* getInstance() {
	    unique_lock<mutex> lock(mtx);
	    
		if (instance == NULL) {
			cout << "instance created\n";
			instance = unique_ptr<PaymentService>(new PaymentService());
		}
		return instance.get();
	}

	void setPaymentStrategy(unique_ptr<PaymentStrategy> newStrategy) {
		paymentStrategy = move(newStrategy);
	}

	void executeStrategy() {
		if (paymentStrategy) {
			paymentStrategy->payMoney();
		}
		else {
			cout << "No Payment Strategy chosen";
		}
	}
};


unique_ptr<PaymentService> PaymentService::instance = NULL; // initiating static variable
// this declared inside class but storage is given outside the class
// otherwise linker will throw an error

mutex PaymentService::mtx;

class PaymentFactory {
public:
	static unique_ptr<PaymentStrategy> getUserDesiredStrategy(string paymentMood) {
		if (paymentMood == "Upi") return make_unique<UpiPayment>();
		else if (paymentMood == "CreditCard") return make_unique<CreditCardPayment>();
		else if (paymentMood == "DebitCard") return make_unique<DebitCardPayment>();
		else if (paymentMood == "AmazonPay") return make_unique<AmazonPayPayment>();
		return NULL;
	}
};




int main() {

	PaymentService* paymentServiceInstance = PaymentService::getInstance();

	if (!paymentServiceInstance) {
		cout << "Failed allocation\n";
		return 0;
	}

	string paymentMood = "Upi";
	unique_ptr<PaymentStrategy> paymentStrategy = PaymentFactory::getUserDesiredStrategy(paymentMood);
	paymentServiceInstance->setPaymentStrategy(move(paymentStrategy));
	paymentServiceInstance->executeStrategy();


	paymentServiceInstance = PaymentService::getInstance();
	paymentMood = "CreditCard";
	paymentStrategy = PaymentFactory::getUserDesiredStrategy(paymentMood);
	paymentServiceInstance->setPaymentStrategy(move(paymentStrategy));
	paymentServiceInstance->executeStrategy();

	paymentServiceInstance = PaymentService::getInstance();
	paymentMood = "DebitCard";
	paymentStrategy = PaymentFactory::getUserDesiredStrategy(paymentMood);
	paymentServiceInstance->setPaymentStrategy(move(paymentStrategy));
	paymentServiceInstance->executeStrategy();
}


/*

	So basically
	unique_ptr and make_unique both are used for memory allocations
	which handles the deallocation on its own

	it's just that to initiate using unique_ptr<T>(pass the new T() object)
	for make_unique<T>() pass nothing

	to transfer assign/ownership use move(ptr)
	After moving, the original pointer becomes nullptr.


	to pass raw ptr for reference pass ptr_which_is_unique.get()
	
	
	
	We can unique_lock which helps us acquire lock2
	we can unlock and lock at our will, and it would
	also automattically unlock once it goes out of scope

    unique_lock<mutex>lock(mtx);
    
    
    If multiple locks involved, use 
    lock(mtx1, mtx2)
    and then take ownership of lock
    unique_lock<mutex> lock(mtx1); // would ty to lock again
    so use unique_lock<mutex> lock(mtx1, adopt_lock);
    
*/
