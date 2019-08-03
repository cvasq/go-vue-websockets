<template>
  <div class="footer">
    <small><button class="btn btn-link"><a href="https://cvasquez.dev/">Return to cvasquez.dev</a></button></small>
    <small><button class="btn btn-link metrics-link"><a v-bind:href="metrics_link">Metrics</a></button></small>
    <small><button class="btn btn-link metrics-link"><a v-bind:href="source_link">Source Code</a></button></small>
</div>
</template>

<script>
export default {
  data () {
    return { 
      metrics_link: 'https://cvasquez.dev/demo/vue-websocket-echo/metrics',
      source_link: 'https://github.com/cvasq/go-vue-websockets',
    }
  }
}
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css?family=Roboto|VT323');

.footer {
    position: fixed;
    left: 0;
    bottom: 0;
    height: 42px;
    width: 100%;
    background-color: black;
    color: white;
    text-align: left;
}

.metrics-link {
    float: right;
}

/* unvisited link */
a:link {
  color: #F0F0F0;
}

/* visited link */
a:visited {
  color: gray;
}

/* mouse over link */
a:hover {
  color: gray;
}

</style>
