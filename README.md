
** backend-coding-test**<br/>
** Clone this project and run <strong>go run main.go</strong> command**
----
 Get List Github Repositories by user ID

* **URL**

  http://localhost:8080/api/v1/:userId/repositories

* **Method:**
  
  `GET`
  
*  **URL Params**

  userId

   **Required:**
 
   `userId=[integer]`

   **Optional:**
 
   `orderBy=[alphanumeric]`
   `orderDirection=[asc/desc]`

* **Data Params**

  <_If making a post request, what should the body payload look like? URL Params rules apply here too._>

* **Success Response:**
  
 

  * **Code:** 200 <br />
    **Content:** `[{"id":397481783,"name":"backend-coding-test","description":"","stargazersCount":0,"ownerId":40974995,"ownerLogin":"trunglv2018"}]`
 
* **Error Response:**

  <_Most endpoints will have many ways they can fail. From unauthorized access, to wrongful parameters etc. All of those should be liste d here. It might seem repetitive, but it helps prevent assumptions from being made where they should be._>

  * **Code:** 400 <br />
    **Content:** `{ error : "Not Found" }`

  OR


* **Sample Call:**

  http://localhost:8080/api/v1/trunglv2018/repositories

* **Notes:**

  