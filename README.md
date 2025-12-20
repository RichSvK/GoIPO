## Indonesia IPO Stock with Go

### Sample Data Source: E-IPO Stock Prospectus
### Link: https://e-ipo.co.id/en

## Related Repositories
- **Frontend**: https://github.com/RichSvK/StockWeb
- **Stock Services**: https://github.com/RichSvK/Stock_Backend
- **IPO console application**: https://github.com/RichSvK/GoIPO
  
### System Requirements
- Software used in developing this program:
  - Go 1.22
  - MySQL 8.3.0
  - GORM

### Program Description
This program is a console-based Go application for processing IPO data with several features:
- Import CSV IPO data to database
- Export CSV IPO data with filtering by value or underwriter. The filtered results are saved in the `UW` or `value` folder inside the `output` folder based on the selected filter.

### Program Preparation
1. Install all the required software such as Go and MySQL.
2. Get Gorm and MySQL Driver using command `go get -u gorm.io/gorm` and `go get -u gorm.io/driver/mysql`.
3. User can comment the logger info in `database.go` inside `configs` folder if the user doesn't want to see the database query log

### Database Setup
1. Create database using the commands in `StokDatabase.sql`.
2. Create a `.env` file based on the `.env.example` file and configure the environment variables accordingly.

### Program Flow
1. Run the go program using command `go run .`
2. Program will enter the main menu <br>
   <img width="250" height="185" alt="image" src="https://github.com/user-attachments/assets/e9b90fe8-eb52-4742-8e0c-60100cf9190c" />

4. If the user input invalid input, program will warned the user with **Invalid Input** message  <br>
   <img width="250" height="185" alt="image" src="https://github.com/user-attachments/assets/f2aed364-6b00-4546-afbf-95a43fa152b9" />
   
5. If this is the user's first time using the program, user should select the database menu by entering option 3 <br>
   <img width="250" height="185" alt="image" src="https://github.com/user-attachments/assets/475fad41-6ed1-4d60-b9d0-4861a051f7f7" />
   * Create a database table using automigrate from GORM by entering option 1
   * Clear table data using `TRUNCATE` by entering option 2
   * Delete table using `DROP` by entering option 3
   * Return to main menu by entering option 4

6. Insert data by entering option 2 in the main menu <br>
   <img width="250" height="185" alt="image" src="https://github.com/user-attachments/assets/0cd93cc9-445f-4be3-8174-689f4ee00109" />
   * Insert the with following step to prevent foreign key error:
     - Insert IPO Stock data by entering option 1
     - Insert Underwriter data by entering option 2
     - Insert IPO Detail IPO by entering option 3
   * Return to main menu by entering option 4

7. Export the data by entering option 1 from the main menu <br>
   <img width="250" height="185" alt="image" src="https://github.com/user-attachments/assets/0a9c3f7c-e197-4854-86f1-03d3f2a1c774" />
   * Export by Underwriter filter by entering option 1 <br>
     <img width="250" height="185" alt="image" src="https://github.com/user-attachments/assets/07696358-f35f-4ea1-b282-a67884e75ee8" />

   * Export by Value filter by entering option 2<br>
     <img width="250" height="185" alt="image" src="https://github.com/user-attachments/assets/3252c967-6bc6-46a8-9084-4c00827bb11b" />
     - User can add more filter `ALL` or `XX` underwriter<br>
       <img width="250" height="185" alt="image" src="https://github.com/user-attachments/assets/1680d3a0-a184-4762-939e-84f58797ca0d" />  
       
8. Open the exported CSV in Excel by importing in the `Data` tab and import the CSV data <br>
   <img width="600" height="200" alt="image" src="https://github.com/user-attachments/assets/8f14d7b7-9c1d-4131-bc3d-9737a05036dd" /> <br>
   **Notes: Some cell of excel may convert the data to a general format makes some data seems different so user can check the csv in some text editor**
   
9. Exit program by entering option 4 from the main menu <br>
   <img width="250" height="185" alt="image" src="https://github.com/user-attachments/assets/4b5ff79e-302f-44d1-ad67-476e5871ba2c" />

