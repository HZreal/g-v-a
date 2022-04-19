<template>
  <div>
    ---------------测试创建新页面--------------
  </div>
</template>

<script>

import { testView } from '../../api/testView'

export default {
  name: 'index',
  data() {
    return {
      postData: {
        "name": "huang",
        "age": 22
      }
    }
  },
  methods: {
    testPost(){
      testView(this.postData)
    }
  },
  created() {
    this.testPost()
  }
}

</script>

<style scoped>

</style>
