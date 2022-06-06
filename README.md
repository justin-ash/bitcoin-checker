Instructions
1. Clone application to the local machine from git repo
    git clone https://github.com/
2. Edit .env file (change TO_EMAIL and mail trap credentials), change max, min price if needed
3. Run the command *docker-compose up*
4. Containerization will begin shortly and automatically start the application with port :8080
5. Trigger the api from your postman
    http://localhost:8080/api/prices/btc?date=06-06-2022&limit=100&offset=0