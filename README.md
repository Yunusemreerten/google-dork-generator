# Google Dork Generator

## Project Description
This project is a web-based Google Dork Generator developed using Go (Golang).  
It generates Google search queries (dorks) based on user input such as a domain, keyword, or company name.

## Features
- Generates 47 Google dork queries
- Supports domain, keyword, and company targets
- Category-based filtering
- User-friendly interface
- Dynamic query generation

## Technologies Used
- Go (Golang)
- HTML
- CSS
- JSON

## How It Works
1. User enters a target (domain or keyword)
2. Selects target type and category
3. System generates relevant Google dork queries
4. Results are displayed on a separate page


Screenshots
<img width="1441" height="603" alt="image" src="https://github.com/user-attachments/assets/cd3badf4-7a55-4157-b963-da6c2602eafb" />
<img width="1273" height="634" alt="image" src="https://github.com/user-attachments/assets/597f3416-26d4-450a-9393-5e42f0e0f56c" />

## How to Run
```bash
git clone https://github.com/Yunusemreerten/google-dork-generator.git
cd google-dork-generator
go run main.go
Then open:
http://localhost:8080

Project Structure
main.go → backend logic
data/dorks.json → dork dataset
templates/ → HTML pages
static/ → CSS and JS files

