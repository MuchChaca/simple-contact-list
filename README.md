# Simple-Contact-List
## The app
* **About:** A simple app with a backend written in go and a frontend written in vuejs.
* **Purpose:** Experimenting how go interacts with vuejs

## List of experiments
- [x] Make go send JSON data to vuejs
- [x] Fetch JSON data from go with vuejs
- [x] Send JSON data from vuejs to go
- [x] Have CRUD working
- [x] Practice some markdown :+1:
- [ ] Make it work with JWT (+middlewares)
- [ ] Replace the maker.sh by docker (container + images ready for production and development)
- [x] SSH connection (kinda out of context here but still)
- [ ] Deploy a go + vuejs app without Docker
- [ ] Deploy a go + vuejs app with Docker
- [ ] Test a go + vuejs app on heroku  

## Status
* **Contribution:** No contribution :no_entry_sign:  
* **Activity:** Active :large_blue_circle:  
* **Evolution:** Not finished :construction:  

## Dependencies
* **GO:** (https://golang.org/)
* **MySQL for go:** (https://github.com/go-sql-driver/mysql)
* **Echo:** (https://echo.labstack.com/)
* **Echo - Middlewares:** (https://echo.labstack.com/middleware)
* **NPM:** (https://www.npmjs.com/)
* **Vue.Js** (https://vuejs.org/)
* **Axios:** (https://www.npmjs.com/package/axios)
* **JWT:** (https://github.com/dgrijalva/jwt-go)

## Build
> ``git clone https://github.com/MuchChaca/simple-contact-list.git``  
> ``cd simple-contact-list``  
> Install the MySQL database from the db directory (create database included) and adapt the alpha.go according to your databse setup  
:warning: Note that you will need to remove the "Training/" in the imports paths of the go files!  
> ``docker/maker.sh``  
> *The contact list will be started on http://localhost:8080*  

## Screenshot
![simple-contact-list](https://image.ibb.co/cbnhNb/simple_contact_list_screen.png)  

## Emoji's commit legend
* :art: when improving the format/structure of the code
* :rocket: when improving performance
* :sweat_drops: plugging memory leaks
* :memo: when writing docs
* :penguin: when fixing something on Linux
* :apple: when fixing something on Mac OS
* :computer: when fixing something on Windows
* :name_badge: when fixing a bug
* :exclamation: when a bug is spotted
* :fire: `:fire:` when removing code or files
* :black_medium_small_square: when fixing the CI build
* :white_check_mark: `:white_check_mark:` when adding tests
* :lock: `:lock:` when dealing with security
* :arrow_up: `:arrow_up:` when upgrading dependencies
* :arrow_down: `:arrow_down:` when downgrading dependencies
* :shirt: `:shirt:` when removing linter warnings
* :construction: when rebuilding something
* :octocat: when changing something about git/github


-----------------------------

<img src="https://upload.wikimedia.org/wikipedia/commons/4/44/Gophercolor.jpg" alt="Golang" width="300px"/><img src="https://upload.wikimedia.org/wikipedia/commons/thumb/5/53/Vue.js_Logo.svg/1000px-Vue.js_Logo.svg.png" alt="Vuejs" width="300px"/>
