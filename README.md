Go and React Movie App

This is a learning tool for me, using react v6 hooks and routing to pull Go Rest Api Server supplied json data.

A few prereqs to install.
  REACT:
  
    - npm install react-router-dom | React-Router https://reactrouter.com/
    - npm install bootstrap        | Bootstrap https://getbootstrap.com/
    
  GOLANG:
  
    - go get -u github.com/julienschmidt/httprouter
    - go get -u github.com/lib/pq

You will want a postgres db for data models. Use postgresql/go_movies.sql to populate your db.

I use docker for database along with an adminer docker image to have a gui. 

NOTE: 
Though this is obvious once you catch it, I had issues at first connecting adminer to postgresql instance from adminer interface.

You will not be connecting localhost:[exposed_port] from browser but rather eth_addr:[exposed_port]. So for instance my 192.168 address.
  Server: 192.168.12.24:5432
  Hope this helps someone and saves the 20min or so I lost troubleshooting.
