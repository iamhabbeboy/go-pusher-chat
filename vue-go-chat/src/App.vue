<script>
import {onMounted, ref} from 'vue';
import Pusher from 'pusher-js'

export default {
  setup() {
    const username = ref('')
    const messages = ref([])
    const message = ref('')

    const apiKey = import.meta.env.VITE_PUSHER_APP_KEY
    const apiUrl = import.meta.env.VITE_API_URL_ENDPOINT

    onMounted(() => {

      if(!apiKey || !apiUrl) {
        return alert("API keys are not yet set.")
      }
      Pusher.logToConsole = true;

      const pusher = new Pusher(apiKey, {
        cluster: 'eu'
      });

      const channel = pusher.subscribe('chat');
      channel.bind('message', data => {
        messages.value.push(data);
      });
    })


    const submit = async () => {

      const payload = {
        username: username.value,
        message: message.value
      }

      await fetch(apiUrl.toString(), {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(payload)
      })
      message.value = '';
    }

    return {
      username,
      messages,
      message,
      submit
    }
  }
}
</script>

<template>
  <div class="w-4/12 mx-auto my-2">
    <div class="border p-5">
      <ul class="mb-10">
        <li v-for="message in messages" class="border-b p-3">
          <img src="https://gravatar.com/avatar/aa43d8ed527bc33e6077d120632e835a?s=400&d=identicon&r=x"
               class="w-5 rounded-full mt-1" align="left"/>
          <span class="pl-2 name">{{ message.username}}</span>
          <p style="font-size: 12px">{{message.message}} </p>
        </li>
      </ul>

      <input type="text" v-model="username" class="border p-2 w-full" placeholder="Username"/> <br/>
      <textarea v-model="message" class="border p-2 mt-3 w-full" placeholder="Message"></textarea>
      <button @click="submit" class="mt-2 bg-purple-500 py-3 px-8 rounded-md text-white hover:bg-purple-600">Send</button>
    </div>
  </div>
</template>
<style>
body {
  font-family: 'Robot', sans-serif;
}

.name {
  font-family: 'Noto Sans', sans-serif;
}
</style>