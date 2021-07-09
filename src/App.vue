<template>
  <div @mousedown="resetCount">
    <p>{{count}}</p>
    <button
    @click="turn"
    >
    Turn {{toggle}}
    </button>
  </div>
</template>

<script>
const { remote } = require('electron');

export default {
  data () {
    return {
      count: 0,
      enabled: false,
      toggle: 'ON',
    };
  },
  created: function() {
    setInterval(()=>{
      this.countUp();
    }, 1000);
    window.addEventListener("mousedown", this.resetCount);
    window.addEventListener('mouseup', 
        () => console.log('mouse up captured outside the window'));
  },
  methods: {
    turn (){
      if (this.enabled) {
        this.enabled = false
        this.toggle = 'ON '
      }else{
        this.enabled = true
        this.toggle = 'OFF'
      }
    },
    countUp () {
      if (this.enabled){
        if (this.count < 30){
          this.count++;
        }else{
          this.count = 1;
        }
      }
    },
    resetCount (e) {
      console.log(e.buttons);
      if (e.buttons == 24) {
        this.count = 0;
      }
    },
    onMouseEnter () {
      // v-on:mouseenter="onMouseEnter"
      // v-on:mouseleave="onMouseLeave"
      remote.getCurrentWindow().setIgnoreMouseEvents(false);
    },
    onMouseLeave () {
      remote.getCurrentWindow().setIgnoreMouseEvents(true, { forward: true })
    }
  },
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
