<template>
  <div id="app">
    <DefaultContainer/>
    <Footer/>
  </div>
</template>

<script>
import DefaultContainer from './components/DefaultContainer'
import Footer from './components/Footer'

export default {
  name: 'App',
  components: {
    DefaultContainer,
    Footer
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
  padding-top: 1.5rem;
  padding-bottom: 1.5rem;
}
</style>
