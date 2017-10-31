<template>
  <div class="row">
    <div class="col-md-12">
      <h2>Contacts</h2><br>
      <div class="col-md-5">
        <ul class="list-group">
          <li class="list-group-item" 
            v-for="(contact, index) in contacts"
            v-bind:key="contact.ID">
            <strong>{{contact.FName}} {{contact.LName}}</strong><br>
            <strong>E-mail: </strong>{{contact.Email}}<br>
            <strong>Portable: </strong>{{contact.Portable}}<br>
            <strong>Phone: </strong>{{contact.Phone}}
            <span class="pull-right">
              <button class="btn btn-xs btn-info" v-on:click="setModContact(index)">
                <i class="fa fa-pencil-square-o" aria-hidden="true"></i>
              </button>
              <br>
              <button class="btn btn-xs btn-danger" v-on:click="delContact(index)">
                <i class="fa fa-trash-o" aria-hidden="true"></i>
              </button>
            </span>
          </li>
        </ul>
      </div>
      <div class="col-md-7">
        <div class="col-md-6">
          <div class="addContact">
            <div class="input-group">
                <h3>Add a contact</h3>
              <div class="input-list">
                <input type="text" class="form-control" placeholder="First Name" required v-model="newContact.FName">
                <input type="text" class="form-control" placeholder="Last Name" required v-model="newContact.LName">
                <input type="text" class="form-control" placeholder="E-mail" required v-model="newContact.Email">
                <input type="text" class="form-control" placeholder="Portable" required v-model="newContact.Portable">
                <input type="text" class="form-control" placeholder="Phone" required v-model="newContact.Phone">
              </div>
              <div class="col-md-1">
                <span class="input-group-btn">
                  <button class="btn btn-primary" type="button" v-on:click="addContact">Create</button>
                </span>
              </div>
            </div> <!-- //INPUT-GROUP -->
          </div><!-- //addContact -->

          <div v-if="modContact.index" class="modContact">
            <div class="input-group">
                <h3>Edit a contact <button class="btn btn-xs btn-default" v-on:click="unsetModContact">
                    <i class="fa fa-times" aria-hidden="true"></i>
                  </button></h3>
              <div class="input-list">
                <input type="text" class="form-control" value="{{modContact.contact.LName}}" required v-model="modContact.contact.FName">
                <input type="text" class="form-control" value="{{modContact.contact.FName}}" required v-model="modContact.contact.LName">
                <input type="text" class="form-control" value="{{modContact.contact.Email}}" required v-model="modContact.contact.Email">
                <input type="text" class="form-control" value="{{modContact.contact.Portable}}" required v-model="modContact.contact.Portable">
                <input type="text" class="form-control" value="{{modContact.contact.Phone}}" required v-model="modContact.contact.Phone">
              </div>
              <div class="col-md-1">
                <span class="input-group-btn">
                  <button class="btn btn-primary" type="button" v-on:click="modifContact">Save</button>
                </span>
              </div>
            </div><!-- //INPUT-GROUP -->
          </div><!-- //V-IF -->
        </div><!-- //COL-MD -->
      </div>
    </div>

    <!-- ERRORS -->
    <div class="row">
      <ul class="list-group">
        <li class="list-group-item"
        v-for="error in errors"
        v-bind:key="error.id">
          <p>{{error}}</p>
        </li>
      </ul>
    </div>
    <!-- //ERRORS -->
  </div>

</template>

<script>
import axios from 'axios'

export default {
  data () {
    return {
      errors: [],
      contacts: [],
      modContact: {},
      newContact: {}
    }
  },
  // This is run whenever the page is loaded to make sure we have a current contact list
  created: function() {
    this.fetchContacts()
  },
  methods: {
    fetchContacts: function () {
      // Use the axios to fetch data from the / route
      axios.get(`/index`)
      .then(response => {
        // JSON response are automatically parsed.
        this.contacts = response.data.items
      })
      .catch(e => {
        this.errors.push(e)
      })
    },

    addContact: function(){
      if (!$.trim(this.newContact.FName)){
        this.newContact = {}
        return
      }
      
      // Post the new contact to the /add route using axios
      axios.post('/add', this.newContact)
      .then(response => {
        this.newContact.ID = response.created
        this.contacts.push(this.newContact)
        // console.log("Contact created!")
        this.newContact = {}
      })
      .catch(e => {
        this.errors.push(error)
      })
    },

    //* Modification
    // set the variable
    setModContact: function(index){
      this.modContact = {
        index: index,
        contact: this.contacts[index]
      }
    },
    // empty the variable
    unsetModContact: function(){
      this.modContact = {}
    },
    // make the modifications
    modifContact: function() {
      axios.put('/mod', this.modContact.contact)
      .then(response => {
        this.modContact = {}
      })
      .catch(e => {
        this.errors.push(e)
      })
    },

    //* Careful the "documentation" version doesn't work very well !!
    delContact: function(index) {
      // Use axios to send a delete task by id
      /* var index = this.contacts.indexOf(contact)
      console.log(this.contacts)
      console.log(index) */
      // return
      axios.delete('/del/' + this.contacts[index].ID)
      .then(response => {
        this.contacts.splice(index, 1)
        // console.log("Contact deleted!")
      })
      .catch(e => {
        this.errors.push(error)
      })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: block;
  margin: 0 10px;
}

a {
  color: #42b983;
}

/* .list-group{
  width: 60%;
}
*/
.list-group-item{
  text-align: left;
}
/* 
.input-group{
  margin-bottom: 5%;
} */

.addContact{
  position: fixed;
  width: 20%;
}
.modContact {
  position: fixed;
  width: 20%;
  margin-top: 240px;
}

.input-group-btn:last-child>.btn, .input-group-btn:last-child>.btn-group{
  height: 170px;
  margin-left: -14px;
  margin-top: -136px;
  width: 70px;
}

.input-list{
  width: 81%;
}

button.btn.btn-xs.btn-info {
    /* height: 100px; */
    margin-top: -110px;
    font-size: 20px;
    vertical-align: middle;
    width: 30px;
}

.btn.btn-xs.btn-danger{
  margin-top: -60px;
  font-size: 20px;
  width: 30px;
  vertical-align: middle;
}

.btn.btn-default{
  width: 70px;
  font-size: 20px;
  height: 45px;
  margin-top: -10px;
  vertical-align: middle;
  color: red;
  position: absolute;
  margin-left: 9.5%;
}

h3 {
  margin-bottom: 10px;
}

.row{
  margin-left: 20%;
}

h2 {
  margin-left: -220px;
}
</style>
