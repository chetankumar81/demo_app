# demo_app

## API DETAILS
---
1. **Start Game**
----
  Returns json data about whether the game started or not.

* **URL**

  /start_game

* **Method:**

  `POST`
  
*  **URL Params**

    None

* **Data Params**

    `{
        "user1": "x1",
        "user2": "x2"
    }`

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `{"code":200,"msg":"","model":{"gameDetails":{"Id":3,"User1":{"Id":1,"UserName":"x1","EmailId":"x1@gmail.com","State":"1"},"User2":{"Id":2,"UserName":"x2","EmailId":"x2@gmail.com","State":"1"},"Status":1,"Timer":"30","Result":0,"Started":"2020-07-31T08:40:58.44140131+05:30","Ended":"0001-01-01T00:00:00Z"}}}`
 
* **Error Response:**

  * **Code:** 200 <br />
    **Content:** `{"code":455,"msg":"","model":"User1 / user2 already in game"}`

  OR

    * **Code:** 200 <br />
    **Content:** `{"code":454,"msg":"","model":"Invalid Params for starting game"}`

  OR

  * **Code:** 401 UNAUTHORIZED <br />
    **Content:** `{ error : "You are unauthorized to make this request." }`

* **Sample Call:**

  ```http
    POST http://127.0.0.1:3000/start_game

    {
        "user1": "x1",
        "user2": "x2"
    }
  ```
----

2. **Pick a Card**
----
  Returns json data about response when a player picks a card for the game.

* **URL**

  /pick_card

* **Method:**

  `POST`
  
*  **URL Params**

    None

* **Data Params**

    `{
        "gameId": "2",
        "user": "x1",
        "card": "26",
        "pickedTime": "2020-07-30 14:50:28"
    }`

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `{"code":200,"msg":"success","model":{}}`

    OR

  * **Code:** 200 <br />
    **Content:** `{"code":200,"msg":"success","model":{"winner":"x1"}}`
 
* **Error Response:**

  * **Code:** 200 <br />
    **Content:** `{"code":400,"msg":"","model":"Invalid game / Invalid user / Invalid card"}`

  OR

  * **Code:** 401 UNAUTHORIZED <br />
    **Content:** `{ error : "You are unauthorized to make this request." }`

* **Sample Call:**

  ```http
    POST http://127.0.0.1:3000/pick_card

    {
        "gameId": "2",
        "user": "x1",
        "card": "26",
        "pickedTime": "2020-07-30 14:50:28"
    }
  ```

----
3. **Show Game Details**
----
  Returns json data about the game.

* **URL**

  /game_details/:gameId

* **Method:**

  `GET`
  
*  **URL Params**

   **Required:**
 
   `gmaeId=[integer]`

* **Data Params**

  None

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `{"code":200,"msg":"Success","model":{"cardDetails":[{"Id":17,"UserId":1,"Card":26,"PickedTime":"2020-07-30T20:20:28+05:30"},{"Id":16,"UserId":2,"Card":26,"PickedTime":"2020-07-30T20:20:28+05:30"},{"Id":14,"UserId":1,"Card":25,"PickedTime":"2020-07-30T20:20:28+05:30"},{"Id":13,"UserId":1,"Card":23,"PickedTime":"2020-07-30T20:20:28+05:30"},{"Id":11,"UserId":2,"Card":22,"PickedTime":"2020-07-30T20:20:28+05:30"},{"Id":10,"UserId":1,"Card":20,"PickedTime":"0001-01-01T00:00:00Z"}],"gameDetails":{"Id":2,"User1":{"Id":1,"UserName":"x1","EmailId":"x1@gmail.com","State":"1"},"User2":{"Id":2,"UserName":"x2","EmailId":"x2@gmail.com","State":"1"},"Status":0,"Timer":"30","Result":1,"Started":"2020-07-30T21:44:38+05:30","Ended":"2020-07-31T08:12:20+05:30"}}}`
 
* **Error Response:**

  * **Code:** 404 NOT FOUND <br />
    **Content:** `{"code":404,"msg":"","model":"Error Game Not Found"}`

  OR

  * **Code:** 401 UNAUTHORIZED <br />
    **Content:** `{ error : "You are unauthorized to make this request." }`

* **Sample Call:**

  ```http
    GET http://127.0.0.1:3000/game_details?gameId=2
  ```