new Vue({
  el: "#app",
  data: {
    message: "",
    logs: [],
    status: "disconnected"
  },
  methods: {
    connect() {
      this.socket = new WebSocket("wss://echo.websocket.org");
      this.socket.onopen = () => {
        this.status = "connected";
        this.logs.push({ event: "WebSocket Connect", data: this.socket.url})
        

        this.socket.onmessage = ({data}) => {
          this.logs.push({ event: "Recieved message", data });
        };
      };
    },
    disconnect() {
      this.socket.close();
      this.status = "disconnected";
      this.logs = [];
    },
    sendMessage(e) {
      this.socket.send(this.message);
      this.logs.push({ event: "Sent message", data: this.message });
      this.message = "";
    }
  }
});
