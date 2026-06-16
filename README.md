# Load Balancing IP HASH method
This is a really simple load balancing proof of concept using docker in the final version and with no docker in 1.0.0 version
## Dependencies
go version 1.26+
docker
## How to install
```bash
  git clone https://github.com/fsbitbybit/Load_Balancing.git
  cd Load_Balancing
  
  #If docker is not running
  sudo systemctl docker

  #The -d flag helps it run in the background
  docker compose up --build -d
```
## To test it
Run this command to make sure it is working or enter localhost:8080 in your browser of choice
```bash
  curl "http://localhost:8080"
```
You should see any of these three appear as an answer:
```
  Backend server 1
  Backend server 2
  Backend server 3
```

Also you can test it to see if this working for sure running this:
```bash
  for i in {1..30}; do
    docker run --rm --network load\_balancing\_app-network alpine sh -c "wget -q -O- http://loadbalancer:8080/" &
  done
  wait
```
