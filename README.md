1- Init frontend on docker
 
	cd frontend/
	docker-compose up

2- Initialize backend api on docker

	cd backendAPI/
	docker-compose up

3- Open a browser and go to the http://localhost:8080

4- Upload test_file.csv* to see the result 


*I created a small cvs file because the one i got from zip file it's all mess up.



What i did:

	-Create a fronted app with vue.js, really minimalist design, form to upload a csv file and running on port 8080;	
	-BackedAPI with golang running on port 7000;
	-Function that handles call from /api;
	-Download(make a copy) of uploaded file;
	-Open csv file and read each line;
	-fill user struct with email, phone, parcel and country which is determined by phone number through regular expression match;
	-turn slice of user into json;
	-write to the page all the user in json;
	
	 
	
