<template>
  <div class="footer">
    <small><button class="btn btn-link"><a href="https://cvasquez.dev/">Return to cvasquez.dev</a></button></small>
</div>
</template>

<script>
export default {
}
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css?family=Roboto|VT323');

.footer {
    position: fixed;
    left: 0;
    bottom: 0;
    height: 40px;
    width: 100%;
    background-color: black;
    color: white;
    text-align: left;
}

/* unvisited link */
a:link {
  color: white;
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
