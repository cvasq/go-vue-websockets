<template>
<div>
   <div class="header clearfix">
      <nav>
         <ul class="nav nav-pill float-right">
            <li class="nav-item">
               Websocket Status:
               <button @click="disconnect" v-if="status === 'connected'" class="btn btn-success btn-sm">Connected</button>
               <button @click="connect" v-if="status === 'disconnected'" class="btn btn-secondary btn-sm">Disconnected</button>
            </li>
         </ul>
      </nav>
      <div align="left">
         <div style="display: inline-block">
            <h2 class="text-muted" v-text="title"><a href="/"></a></h2>
         </div>
      </div>
   </div>
   <div class="jumbotron">
      Type a message to send to the public WebSocket Echo server at <b><a href="https://www.websocket.org/echo.html">echo.websocket.org</a></b>
      <div v-if="status === 'connected'">
         <br>
         <input @click="sendMessage" type="submit" value="Send Message" style="float: right" />
         <div style="overflow: hidden; padding-right: .5em;">
            <form @submit.prevent="sendMessage" action="#">
               <input v-model="message" type="text" style="width: 100%;" />
            </form>
         </div>
         <br>
         <table class="table table-striped">
            <thead>
               <tr>
                  <th>
                     Event
                  </th>
                  <th>
                     Data
                  </th>
               </tr>
            </thead>
            <tbody>
               <tr v-for="log in logs" v-bind:key="log.event">
                  <td >
                     <a v-text="log.event"></a>
                  </td>
                  <td >
                     <a v-text="log.data"></a>
                  </td>
               </tr>
            </tbody>
         </table>
      </div>
   </div>
</div>
</template>

<script>
export default {
  name: 'WebsocketClient',
  data () {
    return {
      title: 'Websocket Echo Client',
      message: '',
      logs: [],
      status: 'disconnected'

    }
  },
  created () {
    this.connect()
  },
  methods: {
    connect () {
      this.socket = new WebSocket('wss://echo.websocket.org')
      this.socket.onopen = () => {
        this.status = 'connected'
        console.log('WebSocket connected to:', this.socket.url)
        this.logs.push({event: 'WebSocket Connected', data: this.socket.url})

        this.socket.onmessage = ({data}) => {
          this.logs.push({event: 'Recieved message', data})
          console.log('Received:', data)
        }
      }
    },
    disconnect () {
      this.socket.close()
      this.status = 'disconnected'
      this.logs = []
    },
    sendMessage (e) {
      this.socket.send(this.message)
      this.logs.push({ event: 'Sent message', data: this.message })
      console.log('Sent:', this.message)
      this.message = ''
    }
  }
}
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css?family=Roboto|VT323');
@import url('https://fonts.googleapis.com/css?family=Montserrat');
h5 {
    font-family: 'Montserrat', sans-serif;
    font-size: large;
    font-weight: bold;

}
body {
  padding-top: 1.5rem;
  padding-bottom: 1.5rem;
  font-family: 'Roboto', sans-serif;
}
a {
    color: rgb(27, 27, 27);
}
.status {
    color:  white;
}
.header {
  padding-bottom: 1rem;
  border-bottom: .05rem solid #e5e5e5;
  font-family: 'VT323', monospace;
}
</style>
