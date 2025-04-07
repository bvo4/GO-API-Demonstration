# GO-API-Demonstration
A demonstration of a Web API being used in the GO Programming language.  Will plan to be used for a demonstration for a coding interview on 4/8/2025

## Function

- When given a directory, will scan and parse all contents from all .csv files to load as Order Details.
- When having a valid set of Order Details, will go through each individual order and use their Order GUID to query to an unspecified API for another company
- Upon receiving a valid HTTP Response and receiving details about that particular Order GUID, parse the .Json response and convert it into a map of that particular type struct array
- After parsing all Order GUIDs, the program will format the data received  and upload the data contents to a local MySQL server through a bulk transaction
- After the SQL upload has been performed, the program will demonstrate Web API Listening capabilities by opening an HTTP Server and processing API Requests by acting as a form of middle-ware to process and forward API Requests from the user to the end-party.

https://stackoverflow.com/questions/71055589/building-a-docker-image-to-run-go-applications