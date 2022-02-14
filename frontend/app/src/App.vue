<template>
  <div id='app'>
    <section class="w-50 mx-auto">
      <span class='title-text'>gRPC Test Game</span>
      <div class='row justify-content-center mt-4'>
        <input v-model='userId' placeholder='名無し'>
        <button v-if="!ready" @click="initialize" class='mt-4 btn btn-primary'>
          ランダムマッチ
        </button>
        <button v-else v-bind:disabled="count !== 0 || endCount === 0" @click="countNum" class='mt-4 btn btn-primary'>
          クリック！
        </button>
      </div>
    </section>
    <section class="mt-4 w-70 mx-auto">
      <h2 v-if="!start && ready">参加者を待っています。。。</h2>
      <h2 v-if="start && (count > 0)">ゲームスタートまで{{ count }}秒</h2>
      <h2 v-else-if="count == 0">スタート！！ 残り時間{{ endCount }}秒</h2>
    </section>
    <section v-if="count == 0">
      <h2>あなたのスコア: {{ totalnumsList.filter(el => el.player)[0].total }}</h2>
      <h2>{{ totalnumsList.filter(el => !el.player)[0].userid }}のスコア: {{ totalnumsList.filter(el => !el.player)[0].total }}</h2>
    </section>
  </div>
</template>

<script>
import { Null, joinRoomParams, addRoomParams, countNumparams, getRoomTotalNumParams } from './count_pb'
import { addNumServiceClient } from './count_grpc_web_pb'

export default {
  name: 'App',
  components: {},
  data: function () {
    return {
      num: 0,
      roomId: '',
      userId: `名無し${Math.floor(Math.random() * 1000)}`,
      ready: false,
      start: false,
      stream: null,
      totalnumsList: [],
      count: 10,
      endCount: 60
    }
  },
  created: function () {
    // eslint-disable-next-line
    this.client = new addNumServiceClient(`${location.protocol}//${location.hostname}:9000`, null, null)
  },
  watch: {
    start: function(nv) {
      if (nv) {
        this.timer(0, 10, () => {this.count -= 1})
      }
    },
    count: function(nv) {
      if (nv === 0) {
        this.timer(0, 60, () => {this.endCount -= 1})
      }
    }
  },
  methods: {
    initialize: function () {
      this.getRooms().then(() => this.fetchRoomTotalNums())
    },
    timer: function (i, m, f) {
      if (i < m) {
        f()
        setTimeout(function(){this.timer(i+1, m, f)}.bind(this), 1000);
      }
    },
    getRooms: async function () {
      let getRoomsParam = new Null()
      await new Promise(resolve => {
        this.client.getRooms(getRoomsParam, {}, (err, response) => {
          const rooms = response.toObject()
          if (rooms.roomsList.find(el => el.isactive) !== undefined) {
            this.roomId = rooms.roomsList.find(el => el.isactive).roomid
            this.joinRoom().then(() => resolve())
          } else {
            this.addRoom().then(() => this.joinRoom()).then(() => resolve())
          }
        })
      })
    },
    addRoom: async function () {
      let addRoomParam = new addRoomParams()
      addRoomParam.setRoomid(this.generateUuid())
      await new Promise(resolve => {
        this.client.addRoom(addRoomParam, {}, (err, response) => {
          const result = response.toObject()
          this.roomId = result.roomid
          resolve()
        })
      })
    },
    joinRoom: async function () {
      let getJoinRoomReq = new joinRoomParams()
      getJoinRoomReq.setUserid(this.userId)
      getJoinRoomReq.setRoomid(this.roomId)
      await new Promise(resolve => {
        this.client.joinRoom(getJoinRoomReq, {}, (err, response) => {
          const result = response.toObject()
          this.ready = result.issuccess
          resolve()
        })
      })
    },
    fetchRoomTotalNums: function () {
      let getRoomTotalNum = new getRoomTotalNumParams()
      getRoomTotalNum.setRoomid(this.roomId)
      this.stream = this.client.getRoomTotalNum(getRoomTotalNum)
      this.stream.on('data', (response) => {
        let totalnumsList = response.toObject().totalnumsList
        let userNums = []
        for (let v in totalnumsList.filter(el => el.roomid === this.roomId)) {
          userNums.push({'userid': totalnumsList[v].userid, 'player': this.userId == totalnumsList[v].userid,'total': totalnumsList[v].total})
        }
        this.start = this.ready && (userNums.length > 1)
        this.totalnumsList = userNums
      })
    },
    countNum: async function () {
      let countNumParam = new countNumparams()
      countNumParam.setUserid(this.userId)
      countNumParam.setRoomid(this.roomId)
      await this.client.countNum(countNumParam, {}, (err, response) => {
        const result = response.toObject()
        console.log(result)
      })
    },
    generateUuid: function () {
      let chars = "xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx".split("");
      for (let i = 0, len = chars.length; i < len; i++) {
          switch (chars[i]) {
              case "x":
                  chars[i] = Math.floor(Math.random() * 16).toString(16);
                  break;
              case "y":
                  chars[i] = (Math.floor(Math.random() * 4) + 8).toString(16);
                  break;
          }
      }
      return chars.join("");
    }
  }
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
