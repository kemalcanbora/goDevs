![alt text](https://media.kommunity.com/communities/goturkiye/19786/goturkey.png)

## Sheet Method

- Open the Google Sheet and share the spreadsheet (button in the top right corner) so anyone with the link can view it without logging in.
- Copy the spreadsheet ID, which is the long random string in the URL of the spreadsheet. Make sure to copy the entire random part of the URL between two slashes.
- We’ll be using this URL to get the spreadsheet’s data, replacing spreadsheet_id with your spreadsheet ID from the previous step:
Example:
`https://docs.google.com/spreadsheets/d/spreadsheet_id/gviz/tq?tqx=out:json`
  
##  How to run it
 -  Add the entry `127.0.0.1 kafka1` to your /etc/hosts file
 -  `docker-compose up` for ES, Kafka, Logstash and Zookeeper
 -  go run app.go 