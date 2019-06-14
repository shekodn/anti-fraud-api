# Anti Fraud Service

## What Is It?

According to the article [Fraud Detection With GPS and Airplanes](https://inversegravity.net/2019/fraud-detection-with-GPS-and-airplanes/), It's possible to calculate the distance between two given cities and
make sure it’s realistic to travel the distance in the given time via airplane.

For example: If Alice makes a transaction in Berlin DE at ``` 05 Aug 12 02:47 UTC ```
and 10 hours later she makes one in Toronto CA at ``` 05 Aug 12 14:30 UTC ``` we
cannot flag this transaction as fraudulent since it is mathematically possible.

However, if there is an attempt to do a transaction in Cancun MX
at ```05 Aug 12 15:03 UTC```, this time it would be flagged as
fraudulent. This is because is not possible for Alice to travel from Toronto to
Cancun in that amount of time.

## How It Works?

![current_app](/docs/images/anti-fraud-api.png)

## How To Run

1. Clone the project
2. Run ``` glide install ```
3. Set up an .env file in the project's root directory
```
POSTGRES_DB=anti-fraud-api
POSTGRES_PASSWORD=mysecretdbpassword
POSTGRES_USER=postgres
DB_TYPE=postgres
DB_HOST=localhost
DB_HOST_DOCKER=host.docker.internal
DB_PORT=5432
CACHE_PORT=6379
CACHE_PASSWORD=mysecretcachepassword
PORT=8000

```
4. Simulate an environment with Redis (cache), Postgres (db), and the
Application by running ``` docker-compose up --build ```


## How To Use
You can use postman or curl in order to post a new transaction in the following
endpoint: ```localhost:8000/new```

### First Transaction
This transaction should be a legit one since it is the first one and it can
happen anywhere in the world.

#### Sample input
```
{
	  "user_id" : 2,
	  "city_name" : "Berlin",
	  "country_code" : "de",
	  "time" : "05 Aug 12 02:47 UTC"
}
```
#### Sample Output:
```
{"message":"success","status":true,"tx":{"ID":1,"CreatedAt":"2019-06-17T02:54:11.6864658Z","UpdatedAt":"2019-06-17T02:54:11.6864658Z","DeletedAt":null,"city_name":"berlin","country_code":"de","time":"05 Aug 12 02:47 UTC","user_id":1}}

```

### Second Transaction
This one should also be legit. It is possible to travel from Berlin
to Toronto within the time difference between the previous transaction (Berlin)
and the current transaction (Toronto).

#### Sample input
```
{
	  "user_id" : 2,
	  "city_name" : "toronto",
	  "country_code" : "ca",
	  "time" : "05 Aug 12 14:30 UTC"
}
```

#### Sample Output:
```
{"message":"success","status":true,"tx":{"ID":2,"CreatedAt":"2019-06-17T04:17:18.5402872Z","UpdatedAt":"2019-06-17T04:17:18.5402872Z","DeletedAt":null,"city_name":"toronto","country_code":"ca","time":"05 Aug 12 13:30 UTC","user_id":2}}
```


### Third Transaction
This one should be flagged as fraud since it is impossible to travel from
Toronto CA to Cancun MX in just 33 minutes.
#### Sample input
```
{
	  "user_id" : 2,
	  "city_name" : "cancun",
	  "country_code" : "mx",
	  "time" : "05 Aug 12 15:03 UTC"

}
```
#### Sample Output:
```
{"message":"Transaction is Fraudalent","status":false}
```

## References
[0] https://inversegravity.net/2019/fraud-detection-with-GPS-and-airplanes/
