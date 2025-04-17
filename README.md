# file
Usage Instructions
Enter the required ClickHouse connection details.
Click the Choose file button to select the flat file for ingestion.
Specify the delimiter used in the file.
Click Connect to establish a connection to the ClickHouse database.
Click Load Columns to view the file columns.
Click Start Ingestion to ingest data into the specified database.
# Data Ingestion Tool

## Setup Instructions
1. **Clone the Repository**
   ```bash
   git clone https://github.com/your-repo/data-ingestion-tool.git
   Run the Application Start the application 
go run main.go
Configuration Details
ClickHouse Connection

Host: localhost
Port: 8080
Database: sql
User: Your ClickHouse username
JWT Token: Your JWT token for authentication
Flat File Selection

File Input: Choose a flat file (e.g., sodapdf-converted.pdf) to process.
Delimiter: Set the delimiter (default is 30).
Run Instructions
Open your web browser and navigate to http://localhost:8081/main.go.
Fill in the ClickHouse connection details and select the flat file.
Click the Connect button to establish the connection to ClickHouse.
Click the Load Columns button to load the columns from the selected flat file.
Click the Start Ingestion button to begin the data ingestion process.
