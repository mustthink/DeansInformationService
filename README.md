# DeansInformationService 
## Service for storing, editing and releasing information

>### Setup:
>>1) Install Docker with Docker Compose (v. 3.1+)
>>2) Run `docker-compose.yml` in Docker
>### OR
>>1) Install mysql server
>>2) Run mysql server
>>3) Add to server all tables from /mysql/db
>>4) Install NodeJS
>>5) Install all node modules
>>6) Run Vue build
>>7) Run GO build
>>8) HOPE THAT THE PROJECT WILL WORK 
>### If there are any errors
>>clear the images and volumes of the service and try to run it again:
>> 1. Stop all DIS containers 
>> 2. ` docker rmi $(docker images -q)`
>> 3. `docker volume rm $(docker volume ls -q)`
