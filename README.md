Go and React Movie App

This is mostly a learning tool for me, using react v6 hooks and routing to pull Go Rest Api Server supplied json data.

A few prereqs to install:
  REACT:
    - npm install react-router-dom | React-Router https://reactrouter.com/
    - npm install bootstrap        | Bootstrap https://getbootstrap.com/
    
  GOLANG:
    - go get -u github.com/julienschmidt/httprouter
    - go get -u github.com/lib/pq

You will want a postgres db for data models. Use postgres/go_movies.sql to populate your db.
