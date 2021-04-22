<h1>Instruction</h1>

1- Init frontend on docker
 
	cd frontend/
	docker-compose up

2- Initialize backend api on docker

	cd backendAPI/
	docker-compose up

3- Open a browser and go to the http://localhost:8080

4- Upload test_file.csv* to see the result 
<br>

<h1>What i did:</h1><br>
	<li>Create a fronted app with vue.js, really minimalist design, form to upload a csv file and running on port 8080;	
	<li>BackedAPI with golang running on port 7000;
	<li>Function that handles call from /api;
	<li>Download(make a copy) of uploaded file;
	<li>Open csv file and read each line;
	<li>Fill user struct with email, phone, parcel and country which is determined by phone number through regular expression match;
	<li>Turn slice of user into Json;
	<li>Write to the page all the user on Json notation;
	
	 
	
