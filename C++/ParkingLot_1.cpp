#include<bits/stdc++.h>
#include <chrono>
#include<ctime>

using namespace std;

/*
	PARKING LOT

	High Level Entites

	ParkingSpot
		- id uuid
		- Level/floor int
		- isAvailable/functional
		- occupiedBy string
		- reservedFor string

		checkAvailabilityForSpot()
		MarkSpotOccupied(Vehicle)
		MarkTheSpotForReservation(Vehicle)
	    UpdateReservationDetails(Vehicle)
	    RemoveReservationStatusForSpot()

	Vehicle
	  - vehicleType ENUM TWO_WHEELER, FOUR_WHEELER, HEAVY_VEHICLE
	  - RegistrationNo
	  - DriverDetails
	  - EntryTimestamp
	  - isBlackListed


	ParkingLotAssignment
		- floorsToAllowedLevelsMap map<vehicleType,List<int>>

		GetLevelForVehicle(vehicleTypeEnum vehicleType, string carNo)

	Ticket
	 - registrationNo
	 - EntryTimeStamp
	 - ExitTimeStamp
	 - Remarks // charges details
	 - AmountPaid // to be populated at exit


	- ParkingLotService
		ParkingLotService* instance;

		map<Level, List<ParkingSpots>> spotsData


		GetSpotforVehicle(vehicle)
		IssueSpotTicketForVehicle(vehicle)
		GetChargesForUser(vehicle)
		MarkVehicleBlacklisted(vehicle)
		UpdateBoothDetails();
*/

// Vehicle
//   - vehicleType ENUM TWO_WHEELER, FOUR_WHEELER, HEAVY_VEHICLE
//   - RegistrationNo
//   - EntryTimestamp


enum VehicleType {
	TWO_WHEELER,
	FOUR_WHEELER,
	HEAVY_VEHICLE
};

class Vehicle {
public:
	VehicleType vehicleType;
	string registrationNo;
	string entryTimestamp;
};


class ParkingSpot {
private:
	string id;
	int level;
	bool isAvaiable;
	string occupiedBy; // stores the registration no of the vehicle
	string reservedFor;
public:

	bool checkAvailabilityForSpot() {
		return this->isAvaiable;
	}

	void MarkSpotOccupied(string id) {
		this->occupiedBy = id;
	}

	void UpdateReservationDetails(Vehicle* vehicle) {
		this->reservedFor = vehicle->registrationNo;
	}

	void RemoveReservationStatusForSpot() {
		this->reservedFor = "";
	}

};

class Ticket {
public:
	string registrationNo;
	chrono::system_clock entryTimestamp;
	chrono::system_clock exitTimestamp; // initially empty
	string remarks; // charges details
	string amountPaid; // initial amount paid
	VehicleType vehicleType;

	virtual Ticket* GetTicketForVehicle(Vehicle* vehicle) = 0;
	virtual Ticket* GetFinalChargesForTicket(Ticket* ticket) = 0;
};

class TwoWheelerTicket: public Ticket {
	const int BASE_PRICE = 50;
	const int INCREAMENTAL_PRICE = 100;
public:

	Ticket* GetTicketForVehicle(Vehicle* vehicle) override {
		Ticket* ticket = new TwoWheelerTicket();
		ticket->registrationNo = vehicle->registrationNo;
		//	ticket->entryTimestamp = chrono::system_clock::(now);
		vehicleType = TWO_WHEELER;
	}

	Ticket* GetFinalChargesForTicket(Ticket* ticket) {
		//	ticket->exitTimestamp = chrono::system_clock::now();
		double seconds = 0; // difftime(ticket->exitTimestamp, ticket.entryTimestamp);
		int hours = (seconds + 3599) / 3600;
		ticket->amountPaid = this->BASE_PRICE + (hours - 1) * (this->INCREAMENTAL_PRICE);
		return ticket;
	}

};

class FourWheelerTicket: public Ticket {
	const int BASE_PRICE = 100;
	const int INCREAMENTAL_PRICE = 100;
public:

	Ticket* GetTicketForVehicle(Vehicle* vehicle) override {
		Ticket* ticket = new TwoWheelerTicket();
		ticket->registrationNo = vehicle->registrationNo;
		//	ticket->entryTimestamp = chrono::system_clock::now();
		vehicleType = FOUR_WHEELER;
	}

	Ticket* GetFinalChargesForTicket(Ticket* ticket) {
		//	ticket->exitTimestamp = chrono::system_clock::now();
		double seconds = 0; //difftime(ticket->exitTimestamp, ticket.entryTimestamp);
		int hours = (seconds + 3599) / 3600;
		ticket->amountPaid = this->BASE_PRICE + (hours - 1) * (this->INCREAMENTAL_PRICE);
		return ticket;
	}
};

class HeavyVehicleTicket: public Ticket {
	const int BASE_PRICE = 150;
	const int INCREAMENTAL_PRICE = 100;
public:

	Ticket* GetTicketForVehicle(Vehicle* vehicle) override {
		Ticket* ticket = new TwoWheelerTicket();
		ticket->registrationNo = vehicle->registrationNo;
		//	ticket->entryTimestamp = chrono::system_clock::now();
		vehicleType = HEAVY_VEHICLE;
	}

	Ticket* GetFinalChargesForTicket(Ticket* ticket) {
		//	ticket->exitTimestamp = chrono::system_clock::now();
		double seconds = 0; //difftime(ticket->exitTimestamp, ticket.entryTimestamp);
		int hours = (seconds + 3599) / 3600;
		ticket->amountPaid = this->BASE_PRICE + (hours - 1) * (this->INCREAMENTAL_PRICE);
		return ticket;
	}
};


class ParkingLot {
	ParkingLot() {} // making constructor private
	map<int, vector<ParkingLot*>> spotsData;
	map<ParkingLot*, Vehicle*> currentParkingLotStatus;
	map<string, Ticket*> ticketsIssuanceData;
	Ticket* ticketStrategy;
public:
	static ParkingLot* parkingLotInstance;

	static ParkingLot* getParkingLotInstance() {
		if (parkingLotInstance == NULL) {
			cout << "Instance initiated\n";
			parkingLotInstance = new ParkingLot();
		}
		return parkingLotInstance;
	}

	void setTicketStrategy(Ticket* ticketStrategy) {
		this->ticketStrategy = ticketStrategy;
	}

	Ticket* executeStrategy(string ticketId, Vehicle* vehicle) {
		if (ticketId != "") return ticketStrategy->GetFinalChargesForTicket(ticketsIssuanceData[ticketId]);
		return ticketStrategy->GetTicketForVehicle(vehicle);
	}


};

// global/static variables are assinged segment memory
// local variables and functions are assigned stack memory
// dynamically allocated variables are allocated in heap

ParkingLot* ParkingLot::parkingLotInstance = NULL; // assigning memory globally

int main() {
	ParkingLot* parkingLotService = ParkingLot::getParkingLotInstance();
	cout << (parkingLotService != NULL);
	parkingLotService = ParkingLot::getParkingLotInstance();
	parkingLotService = ParkingLot::getParkingLotInstance();

}


