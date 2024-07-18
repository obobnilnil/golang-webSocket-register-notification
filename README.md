<div align="center">
    <img src="https://miro.medium.com/v2/resize:fit:563/1*XcBOdmY4KHMVzcWux1mekw.png" alt="Go" width="300"/>
    <img src="https://miro.medium.com/v2/resize:fit:1072/1*LPLP3nZxRX7p36HUhk0P0A.png" alt="Go" width="400"/>
</div>
<H1>ReadME</H1>
<p>This example Golang-based code demonstrates how to write Golang code to create a session between a client and server for real-time communication using the WebSocket protocol, specifically leveraging the socket.io library.</p>
<h1>Here's a brief overview of the process:
</h1>
In this example, we use a POST method as an API to create an account. Once the account is created, 
a notification is sent to an admin who is connected to the session. The admin is identified by a specific role, and the notification is sent to the company matching the applicant.
<div align="center">
    <h2>The socket.io session has been connected. Once the account is successfully registered, all the data will be transferred to the socket.io session.</h2>
    <img src="https://github.com/user-attachments/assets/cc19e337-81ba-4724-9241-876ed693abee" alt="Socket.io" width="1000"/>
    <img src="https://github.com/user-attachments/assets/157b7a7a-f490-449f-b623-7d070eb7da50" alt="Socket.io" width="1000"/>
    <h2>All the data that we sent to the parameter named Broadcast in our code is stored in the variable msg within the HandlerMessage function.</h2>
    <img src="https://github.com/user-attachments/assets/49bc389e-3d83-4252-b986-f93846cc1285" alt="Socket.io" width="1000"/>
</div>
Golang/Gin-framework Database/Postgresql+Pgadmin WebSocket/Socket.io
