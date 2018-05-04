============================================
The Chamber of Secrets: A Distributed Paxos Diary
============================================

Created for:
CPSC 416 Distributed Systems at the University of British Columbia (UBC)

Authors: Graham L. Brown, Alex Budkina, Larissa Feng, Harryson Hu, Sharon Yang.

The application can be run both on a local machine and on Azure with different VMs.

Let PORT be some valid free port number.

To run the app locally:
— Start server from go/src: go run distributeddiaryserver/server.go 12345 --local
- Start app from go/src: go run distributeddiaryapp/app.go 127.0.0.1:12345 PORT --local

This will run server and apps at 127.0.0.1:PORT

Let ADDRESS be the outgoing IP for the server VM.

To run the app in production:
— Start server from src: go run distributeddiaryserver/server.go 12345
- Start app from src: go run distributeddiaryapp/app.go ADDRESS:12345 PORT

This will run apps on machine's outbound IP on port PORT

NOTE:   Running the app in production mode may not work if not on an Azure VM.
        The reason for this is that the code to retrieve the outbound IP address has been written specifically
        to fetch the outbound IP *for an Azure machine*. Individual instances of Azure VMs are all behind the
        Azure load balancer and have no idea what the original, public-facing IP address is.
        Therefore, in production, the app must make a `curl` call out to Azure in order to retrieve the public-facing IP.

The performance logs are stored under src/logs
To view the performance at real time add “--debug” in the end of the command that runs the app.

