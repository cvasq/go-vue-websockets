<template>
  <div class="container">
        <WebsocketClient/>
  </div>
</template>

<script>
import WebsocketClient from './WebsocketClient'
export default {
  name: 'DefaultContainer',
  components: {
    WebsocketClient
  },
  data () {
    return {
    }
  }
}
</script>

<style scoped>
.container {
  font-family: 'Roboto', sans-serif;
  max-width: 700px;
}
</style>
