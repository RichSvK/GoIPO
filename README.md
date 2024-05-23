## Indonesia IPO Stock with Go

### Stock Market: Indonesia
### Sample Data Source: E-IPO Stock Prospectus
### Link: https://e-ipo.co.id/en

### System Requirements
- Software used in developing this program:
  - Go 1.22
  - MySQL 8.3.0
  - Text Editor: Visual Studio Code

### Program Description
This program allows users to export a CSV file containing the information of IPO data with value and underwriter filter. The result will be saved in `UW` or `value` folder inside the `output` folder based on the type of filter applied by the user. Additionally, it provides functionality to insert data, with the file format in `data` folder into a MySQL database using GORM.

### Program Preparation
1. Install all the required software such as Go and MySQL.
2. Get Gorm and MySQL Driver using command `go get -u gorm.io/gorm` and `go get -u gorm.io/driver/mysql`.
3. User can comment the logger info in `database.go` inside `configs` folder if the user doesn't want to see the database query log
   ![image](https://github.com/RichSvK/IPO_Stock/assets/87809864/abc34ad5-05be-4b65-b186-99738a4e6b34)

### Database Setup
1. Create database using the commands in `StokDatabase.sql`.
2. Check the `.env` file and set the `DB_SOURCE` based on user setting with format `user:password@tcp(host:port)/databaseName?parseTime=true`.
   ![image](https://github.com/RichSvK/IPO_Stock/assets/87809864/4e3df8b3-d905-4d16-a95d-114407bd042c)

### Program Flow
1. Run the go program using command `go run .`
2. Program will enter the main menu
   <br>
   ![image](https://github.com/RichSvK/IPO_Stock/assets/87809864/61307dd8-2e89-4770-a94a-abd19cd4cd55)
   
4. If the user input invalid input, program will warned the user with **Invalid Input** message 
   <br>
   ![image](https://github.com/RichSvK/IPO_Stock/assets/87809864/12edd24c-e619-4fb7-bb87-5ff053f75908)
   
5. If this is the user's first time using the program, user should select the database menu by entering option 3
   <br>
   ![image](https://github.com/RichSvK/IPO_Stock/assets/87809864/d7f304ee-d137-443e-865a-2a2ce099f69f)
   * Create a database table using automigrate from GORM by entering option 1
   * Clear table data using `TRUNCATE` by entering option 2
   * Delete table using `DROP` by entering option 3
   * Return to main menu by entering option 4

6. Insert data by entering option 2 in the main menu
   <br>
   ![image](https://github.com/RichSvK/IPO_Stock/assets/87809864/e03ffd72-6ac1-4d43-9355-3c2b81053fb7)
   * Insert the with following step to prevent foreign key error:
     - Insert Data IPO Stock by entering option 1
       ![image](https://github.com/RichSvK/IPO_Stock/assets/87809864/913654f2-9231-4301-b7ef-336ab4b324dd)
       
     - Insert Data underwriter by entering option 2
       ![image](https://github.com/RichSvK/IPO_Stock/assets/87809864/453ef1fa-ccde-4bd8-a1d3-fdb76163bf2d)

     - Insert Data IPO Detail by entering option 3
       ![image](https://github.com/RichSvK/IPO_Stock/assets/87809864/4305110d-18eb-4add-9b20-486d95dd3d0d)
   * Return to main menu by entering option 4

7. Export the data by entering option 1 from the main menu
   <br>
   ![image](https://github.com/RichSvK/IPO_Stock/assets/87809864/b8c12066-fd36-47a2-8e63-3b97f899b062)
   * Export by Underwriter filter by entering option 1
     ![image](https://github.com/RichSvK/IPO_Stock/assets/87809864/775fb028-d24a-4512-a65d-0173cb1ad0af)

   * Export by Value filter by entering option 2
     - User can add more filter `ALL` underwriter
       ![image](https://github.com/RichSvK/IPO_Stock/assets/87809864/38953243-78fa-4f8a-8f87-604aca8ae1a6)

     - User can add more filter `XX` underwriter
       ![image](https://github.com/RichSvK/IPO_Stock/assets/87809864/eda1eae9-779a-4c07-bea6-0a3731d5e174)     
       
8. Open the exported CSV in Excel by importing in the `data` tab and choose `From Text/CSV`
   <br>
   ![image](https://github.com/RichSvK/IPO_Stock/assets/87809864/eeb83b15-8f46-494c-871a-8a2987ac8342)
   <br>
   **Notes: Some cell of excel may convert the data to a general format makes some data seems different so user can check the csv in some text editor**
   <br><br>
   For example: <br>
   ![image](https://github.com/RichSvK/IPO_Stock/assets/87809864/4b95d1ad-6d22-4e6e-ac2c-083af5146e5f)
   ![image](https://github.com/RichSvK/IPO_Stock/assets/87809864/8fc166ed-9c62-4708-85cc-1111b04a5883)
   In the above image, we can see the `warrant` column in the Excel seems to be different from the original in CSV because the excel auto format it to time.

9. Exit program by entering option 4 from the main menu
   <br>
   ![image](https://github.com/RichSvK/IPO_Stock/assets/87809864/bc57e31c-2917-4273-88be-c2f63cd8aa4c)
