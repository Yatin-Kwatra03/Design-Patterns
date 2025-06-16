#include<bits/stdc++.h>
#include<mutex>

using namespace std;

/*

	Rate Limiter -

	is used to limit the no of calls to an API/some system.

	ApiTier
	  TIER1
	  TIER2
	  TIER3

	Api
	 ApiTier apiTier;


	ClientApi:Api
		virtual callServerApi() = 0;


	// our critical api schema which needs to be rate limited
	ServerApi:Api
		performOperation()

		// timestamp at which api was hit to check if rate limiting is working for the server api
		// it is temporary and could be removed once server side logic validated
		List<int> logsOfHit;


	Tier1Api: ClientApi
		static map<ServerApi, List<int>> apiCallData; // stores the successful call made to the api by tier1 clients, time here is represented by int
		callServerApi(serverApi)

	Tier2Api: ClientApi
		static map<ServerApi, List<int>> apiCallData; // stores the successful call made to the api by tier1 clients, time here is represented by int
		callServerApi(serverApi)

	Tier3Api: ClientApi
		static map<ServerApi, List<int>> apiCallData; // stores the successful call made to the api by tier1 clients, time here is represented by int
		callServerApi(serverApi)


map<ServerApi, List<int>> Tier1Api::apiCallData;
map<ServerApi, List<int>> Tier2Api::apiCallData;
map<ServerApi, List<int>> Tier3Api::apiCallData;


	Enum TimeUnit
		HOUR
		MINUTE
		SECOND

	Limit -
	  NoOfCalls int
	  unit TimeUnit

	RateLimiter -
		map<api,Limit> apiThresholds;
		ClientApi* clientApiStrategy;

		RegisterApiForRateLimiting(ServerApi, limit) // client APIs would have a separate allowance to call this Api.
		// Tier1 service would get 50% quota
		// Tier2 service would get 30% quota
		// Tier3 and below would get 20% quota

		UpdateRateLimits(api, limit)
		SetStrategy(ClientApi* newStrategy)
		ExecuteStrategy(serverApi)

*/

mutex mtx;
int tme; // to represent for our LLD purpose

enum ApiTier {
	TIER1,
	TIER2,
	TIER3
};

enum TimeUnit {
	HOUR,
	MINUTE,
	SECOND
};

class Limit {
	int noOfCallsAllowed;
	TimeUnit unit;
public:
	Limit(int noOfCallsAllowed, TimeUnit unit) {
		this->noOfCallsAllowed = noOfCallsAllowed;
		this->unit = unit;
	}
	int getNoOfCallsAllowed() {
		return this->noOfCallsAllowed;
	}
	TimeUnit getTimeUnit() {
		return this->unit;
	}
};

class API {
	ApiTier apiTier;
public:
	ApiTier getApiTier() {
		return this->apiTier;
	}
};

class ServerApi: public API {
	string apiName;
public:
	ServerApi(string apiName) {
		this->apiName = apiName;
	}

	void performOperation() {
		cout << "Sent notification to the user\n";
	}
	string getApiName() {
		return this->apiName;
	}
};

class ClientApi: public API {
public:
	virtual void callServerApi(ServerApi* serverApi, Limit* serverApiLimit) = 0;
};

// Tier1Api: Api
// 	static map<ServerApi, List<int>> apiCallData; // stores the successful call made to the api by tier1 clients, time here is represented by int
// 	callServerApi(serverApi)

class Tier1Api: public ClientApi {
	const double ALLOTED_QUOTA = 0.5;
	// stores the successful call made to the api by tier1 clients, time here is represented by int
	static map<ServerApi*, queue<int>> apiCallData;

public:

	int getTimeUnitInSeconds(TimeUnit timeUnit) {
		switch (timeUnit) {
		case TimeUnit::HOUR:
			return 3600;
		case TimeUnit::MINUTE:
			return 60;
		default:
			return 1;
		}
	}

	int getRemainingQuota(ServerApi* serverApi, Limit* serverApiLimit) {
		queue<int> &callData = apiCallData[serverApi];
		int callsAllowedForTheTier = 0.5 * serverApiLimit->getNoOfCallsAllowed(); // 50% of total quota
		int timeWindowInSeconds = getTimeUnitInSeconds(serverApiLimit->getTimeUnit());
		if (apiCallData.empty()) return serverApiLimit->getNoOfCallsAllowed(); // full quota is remaining

		while (!callData.empty() and tme - callData.front() > timeWindowInSeconds) {
			callData.pop();
		}

		return callsAllowedForTheTier - (int)(callData.size());
	}

	void callServerApi(ServerApi* serverApi, Limit* serverApiLimit) {
		int remainingQuota = getRemainingQuota(serverApi, serverApiLimit);
		if (remainingQuota == 0) {
			cout << "CALL FAILED AT " << tme << " TIME\n";
			return;
		}
		serverApi->performOperation();
		apiCallData[serverApi].push(tme);
	}
};

class Tier2Api: public ClientApi {
	const double ALLOTED_QUOTA = 0.3;
	// stores the successful call made to the api by tier1 clients, time here is represented by int
	static map<ServerApi*, queue<int>> apiCallData;

public:
	int getTimeUnitInSeconds(TimeUnit timeUnit) {
		switch (timeUnit) {
		case TimeUnit::HOUR:
			return 3600;
		case TimeUnit::MINUTE:
			return 60;
		default:
			return 1;
		}
	}

	int getRemainingQuota(ServerApi* serverApi, Limit* serverApiLimit) {
		queue<int> &callData = apiCallData[serverApi];
		int callsAllowedForTheTier = 0.5 * serverApiLimit->getNoOfCallsAllowed(); // 50% of total quota
		int timeWindowInSeconds = getTimeUnitInSeconds(serverApiLimit->getTimeUnit());
		if (apiCallData.empty()) return serverApiLimit->getNoOfCallsAllowed(); // full quota is remaining

		while (!callData.empty() and tme - callData.front() > timeWindowInSeconds) {
			callData.pop();
		}

		return callsAllowedForTheTier - (int)(callData.size());
	}

	void callServerApi(ServerApi* serverApi, Limit* serverApiLimit) {
		int remainingQuota = getRemainingQuota(serverApi, serverApiLimit);
		if (remainingQuota == 0) {
			cout << "CALL FAILED AT " << tme << " TIME\n";
			return;
		}
		serverApi->performOperation();
		apiCallData[serverApi].push(tme);
	}
};


class Tier3Api: public ClientApi {
	const double ALLOTED_QUOTA = 0.2;
	// stores the successful call made to the api by tier1 clients, time here is represented by int
	static map<ServerApi*, queue<int>> apiCallData;

public:
	int getTimeUnitInSeconds(TimeUnit timeUnit) {
		switch (timeUnit) {
		case TimeUnit::HOUR:
			return 3600;
		case TimeUnit::MINUTE:
			return 60;
		default:
			return 1;
		}
	}

	int getRemainingQuota(ServerApi* serverApi, Limit* serverApiLimit) {
		queue<int> &callData = apiCallData[serverApi];
		int callsAllowedForTheTier = 0.5 * serverApiLimit->getNoOfCallsAllowed(); // 50% of total quota
		int timeWindowInSeconds = getTimeUnitInSeconds(serverApiLimit->getTimeUnit());
		if (apiCallData.empty()) return serverApiLimit->getNoOfCallsAllowed(); // full quota is remaining

		while (!callData.empty() and tme - callData.front() > timeWindowInSeconds) {
			callData.pop();
		}

		return callsAllowedForTheTier - (int)(callData.size());
	}

	void callServerApi(ServerApi* serverApi, Limit* serverApiLimit) {
		int remainingQuota = getRemainingQuota(serverApi, serverApiLimit);
		if (remainingQuota == 0) {
			cout << "CALL FAILED AT TIME " << tme << endl;
			return;
		}
		serverApi->performOperation();
		apiCallData[serverApi].push(tme);
	}
};

class ClientApiFactory {
public:

	ClientApi* getClientApi(ApiTier apiTier) {
		switch (apiTier) {
		case ApiTier::TIER1:
			return new Tier1Api();
		case ApiTier::TIER2:
			return new Tier2Api();
		case ApiTier::TIER3:
			return new Tier3Api();
		default:
			// do something
			cerr << "unknown api tier\n";
			return NULL;
		}
		return NULL;
	}
};


class RateLimiter {
private:
	map<ServerApi*, Limit*>apiThresholds;
	ClientApi* clientApiStrategy;
	RateLimiter() {} // private constructor
	static RateLimiter* rateLimiterInstance;
public:

	static RateLimiter* getRateLimiterInstance() {
		// take lock here
		lock_guard<mutex> lock(mtx);

		if (rateLimiterInstance == NULL) {
			cout << "rate limiter instance created\n";
			rateLimiterInstance = new RateLimiter();
		}
		return rateLimiterInstance;
	}

	void RegisterApiForRateLimiting(ServerApi* serverApi, Limit* limit) {
		this->apiThresholds[serverApi] = limit;
	}

	void setStrategy(ClientApi* newStrategy) {
		this->clientApiStrategy = newStrategy;
	}

	void executeStrategy(ServerApi* serverApi) {
		clientApiStrategy->callServerApi(serverApi, apiThresholds[serverApi]);
	}
};

// defining static variables
RateLimiter* RateLimiter::rateLimiterInstance = NULL;
map<ServerApi*, queue<int>> Tier1Api::apiCallData;
map<ServerApi*, queue<int>> Tier2Api::apiCallData;
map<ServerApi*, queue<int>> Tier3Api::apiCallData;

int main() {
	RateLimiter* instance = RateLimiter::getRateLimiterInstance();
	if (instance == NULL) {
		cout << "FAILED to initialize rate limiter\n";
		return 0;
	}

	ServerApi* serverApi = new ServerApi("accountsService");
	Limit* limit = new Limit(10, TimeUnit::MINUTE); // quota will be 5,3,2 requests for tier1, tier2 and tier3 in 10 seconds

	instance->RegisterApiForRateLimiting(serverApi, limit);

	ClientApiFactory* clientApiFactory = new ClientApiFactory();
	ClientApi* clientApiStrategy = clientApiFactory->getClientApi(ApiTier::TIER1);
	instance->setStrategy(clientApiStrategy);


	tme = 0;
	instance->executeStrategy(serverApi);
	tme++;
	instance->executeStrategy(serverApi);
	tme++;
	instance->executeStrategy(serverApi);
	tme++;
	instance->executeStrategy(serverApi);
	tme++;
	instance->executeStrategy(serverApi);
	// last 2 calls should fail
	tme++;
	instance->executeStrategy(serverApi);
	tme++;
	instance->executeStrategy(serverApi);

	// try creating new instance and still it should fail
	tme++;
	clientApiStrategy = clientApiFactory->getClientApi(ApiTier::TIER1);
	instance->setStrategy(clientApiStrategy);
	instance->executeStrategy(serverApi);

	tme = 100;

	// after one minute they should pass through
	instance->executeStrategy(serverApi);
	instance->executeStrategy(serverApi);

}
